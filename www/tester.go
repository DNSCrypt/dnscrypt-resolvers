package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	stamps "github.com/jedisct1/go-dnsstamps"
	"github.com/miekg/dns"
)

type Tester struct {
	db          *DB
	concurrency int
	timeout     time.Duration
	httpClient  *http.Client
}

type TestJob struct {
	Resolver   Resolver
	ResolverID int64
	Stamp      string
}

type TestJobResult struct {
	ResolverID int64
	Stamp      string
	Success    bool
	RTT        time.Duration
	Error      string
}

func NewTester(db *DB, concurrency int, timeout time.Duration) *Tester {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
			MinVersion:         tls.VersionTLS12,
		},
		DisableKeepAlives:     true,
		ForceAttemptHTTP2:     true, // Enable HTTP/2 for DoH
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       30 * time.Second,
		ResponseHeaderTimeout: timeout,
	}

	return &Tester{
		db:          db,
		concurrency: concurrency,
		timeout:     timeout,
		httpClient: &http.Client{
			Transport: transport,
			Timeout:   timeout,
		},
	}
}

func (t *Tester) TestAll(resolvers []Resolver) {
	jobs := make(chan TestJob, len(resolvers)*2)
	results := make(chan TestJobResult, len(resolvers)*2)

	var wg sync.WaitGroup
	for i := 0; i < t.concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				result := t.testOne(job)
				results <- result
			}
		}()
	}

	go func() {
		for _, r := range resolvers {
			resolverID, err := t.db.UpsertResolver(r.Name, string(r.Type), r.Description, r.SourceFile)
			if err != nil {
				log.Printf("Failed to upsert resolver %s: %v", r.Name, err)
				continue
			}
			for _, stamp := range r.Stamps {
				jobs <- TestJob{
					Resolver: Resolver{
						Name:        r.Name,
						Description: r.Description,
						Type:        r.Type,
						SourceFile:  r.SourceFile,
					},
					ResolverID: resolverID,
					Stamp:      stamp,
				}
			}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	successCount, failCount := 0, 0

	for result := range results {
		if result.ResolverID == 0 {
			log.Printf("Warning: resolver ID is 0 for stamp %s", result.Stamp[:min(50, len(result.Stamp))])
			continue
		}

		errMsg := ""
		if result.Error != "" {
			errMsg = result.Error
		}

		err := t.db.RecordTest(result.ResolverID, result.Stamp, result.Success, result.RTT.Milliseconds(), errMsg)
		if err != nil {
			log.Printf("Failed to record test result: %v", err)
		}

		if result.Success {
			successCount++
		} else {
			failCount++
		}
	}

	log.Printf("Test completed: %d success, %d failed", successCount, failCount)
}

func (t *Tester) testOne(job TestJob) TestJobResult {
	resolverID := job.ResolverID

	stamp, err := stamps.NewServerStampFromString(job.Stamp)
	if err != nil {
		return TestJobResult{
			ResolverID: resolverID,
			Stamp:      job.Stamp,
			Success:    false,
			Error:      fmt.Sprintf("invalid stamp: %v", err),
		}
	}

	start := time.Now()
	var testErr error

	switch stamp.Proto {
	case stamps.StampProtoTypeDoH:
		testErr = t.testDoH(stamp)
	case stamps.StampProtoTypeDNSCrypt:
		testErr = t.testDNSCrypt(stamp)
	case stamps.StampProtoTypeTLS:
		testErr = t.testDoT(stamp)
	case stamps.StampProtoTypeDoQ:
		testErr = t.testDoQ(stamp)
	case stamps.StampProtoTypeDNSCryptRelay:
		testErr = t.testRelay(stamp)
	case stamps.StampProtoTypeODoHTarget:
		testErr = t.testODoHTarget(stamp)
	case stamps.StampProtoTypeODoHRelay:
		testErr = t.testODoHRelay(stamp)
	default:
		testErr = fmt.Errorf("unsupported protocol: %s", stamp.Proto.String())
	}

	rtt := time.Since(start)

	result := TestJobResult{
		ResolverID: resolverID,
		Stamp:      job.Stamp,
		RTT:        rtt,
	}

	if testErr != nil {
		result.Success = false
		result.Error = testErr.Error()
	} else {
		result.Success = true
	}

	status := "OK"
	if !result.Success {
		status = "FAIL"
	}
	log.Printf("[%s] %s (%s) - %v - %s", status, job.Resolver.Name, stamp.Proto.String(), rtt.Round(time.Millisecond), result.Error)

	return result
}

func (t *Tester) testDoH(stamp stamps.ServerStamp) error {
	host := stamp.ProviderName
	path := stamp.Path
	if path == "" {
		path = "/dns-query"
	}

	// Build URL with hostname for TLS SNI
	u := &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}

	// Create DNS query
	query := new(dns.Msg)
	query.SetQuestion("example.com.", dns.TypeA)
	query.Id = dns.Id()
	query.RecursionDesired = true

	wireFormat, err := query.Pack()
	if err != nil {
		return fmt.Errorf("failed to pack DNS query: %v", err)
	}

	// Create custom transport if we have an IP address in the stamp
	client := t.httpClient
	if stamp.ServerAddrStr != "" {
		client = t.createDoHClient(stamp.ServerAddrStr, host)
	}

	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)
	defer cancel()

	// Try POST first
	resp, err := t.doDoHRequest(ctx, client, u.String(), wireFormat, "POST")
	if err != nil || resp.StatusCode != http.StatusOK {
		if resp != nil {
			resp.Body.Close()
		}
		// Fall back to GET
		resp, err = t.doDoHRequest(ctx, client, u.String(), wireFormat, "GET")
		if err != nil {
			return fmt.Errorf("DoH request failed: %v", err)
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("DoH returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 65535))
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	response := new(dns.Msg)
	if err := response.Unpack(body); err != nil {
		return fmt.Errorf("failed to parse DNS response: %v", err)
	}

	// Verify we got a valid response
	if response.Id != query.Id {
		return fmt.Errorf("response ID mismatch")
	}

	return nil
}

func (t *Tester) createDoHClient(serverAddr, hostname string) *http.Client {
	// Extract IP and port from serverAddr
	ipStr := serverAddr
	port := "443"
	if h, p, err := net.SplitHostPort(serverAddr); err == nil {
		ipStr = h
		port = p
	}

	// Strip port from hostname for TLS ServerName (if present)
	tlsServerName := hostname
	if h, _, err := net.SplitHostPort(hostname); err == nil {
		tlsServerName = h
	}

	// Create a dialer that connects to the IP from stamp
	dialer := &net.Dialer{Timeout: t.timeout}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			ServerName:         tlsServerName,
			InsecureSkipVerify: false,
			MinVersion:         tls.VersionTLS12,
		},
		DisableKeepAlives: true,
		ForceAttemptHTTP2: true, // Enable HTTP/2
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// Connect to the IP from stamp instead of resolving hostname
			return dialer.DialContext(ctx, network, net.JoinHostPort(ipStr, port))
		},
	}

	return &http.Client{
		Transport: transport,
		Timeout:   t.timeout,
	}
}

func (t *Tester) doDoHRequest(ctx context.Context, client *http.Client, urlStr string, wireFormat []byte, method string) (*http.Response, error) {
	var req *http.Request
	var err error

	if method == "GET" {
		// For GET, encode the query in the URL
		encoded := base64.RawURLEncoding.EncodeToString(wireFormat)
		if strings.Contains(urlStr, "?") {
			urlStr = urlStr + "&dns=" + encoded
		} else {
			urlStr = urlStr + "?dns=" + encoded
		}
		req, err = http.NewRequest("GET", urlStr, nil)
	} else {
		req, err = http.NewRequest("POST", urlStr, bytes.NewReader(wireFormat))
		if err == nil {
			req.Header.Set("Content-Type", "application/dns-message")
		}
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/dns-message")
	req = req.WithContext(ctx)

	return client.Do(req)
}

func (t *Tester) testDNSCrypt(stamp stamps.ServerStamp) error {
	addr := stamp.ServerAddrStr
	if addr == "" {
		return fmt.Errorf("no server address in stamp")
	}

	providerName := stamp.ProviderName
	if providerName == "" {
		return fmt.Errorf("no provider name in stamp")
	}

	// DNSCrypt certificate is retrieved by querying TXT record for the provider name
	// The provider name in the stamp is already the full name to query (e.g., "2.dnscrypt.default.ns1.adguard.com")
	certName := providerName

	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)
	defer cancel()

	dialer := net.Dialer{Timeout: t.timeout}

	// Try UDP first, then TCP
	var conn net.Conn
	var err error
	useTCP := false

	conn, err = dialer.DialContext(ctx, "udp", addr)
	if err != nil {
		conn, err = dialer.DialContext(ctx, "tcp", addr)
		if err != nil {
			return fmt.Errorf("connection failed: %v", err)
		}
		useTCP = true
	}
	defer conn.Close()

	// Create TXT query for the certificate
	query := new(dns.Msg)
	query.SetQuestion(dns.Fqdn(certName), dns.TypeTXT)
	query.Id = dns.Id()

	wireFormat, err := query.Pack()
	if err != nil {
		return fmt.Errorf("failed to pack query: %v", err)
	}

	conn.SetDeadline(time.Now().Add(t.timeout))

	if useTCP {
		// TCP requires 2-byte length prefix
		length := make([]byte, 2)
		length[0] = byte(len(wireFormat) >> 8)
		length[1] = byte(len(wireFormat))
		if _, err := conn.Write(length); err != nil {
			return fmt.Errorf("failed to write length: %v", err)
		}
	}

	if _, err := conn.Write(wireFormat); err != nil {
		return fmt.Errorf("failed to write query: %v", err)
	}

	buffer := make([]byte, 4096)
	var n int

	if useTCP {
		lengthBuf := make([]byte, 2)
		if _, err := io.ReadFull(conn, lengthBuf); err != nil {
			return fmt.Errorf("failed to read length: %v", err)
		}
		respLen := int(lengthBuf[0])<<8 | int(lengthBuf[1])
		if respLen > len(buffer) {
			return fmt.Errorf("response too large: %d bytes", respLen)
		}
		n, err = io.ReadFull(conn, buffer[:respLen])
	} else {
		n, err = conn.Read(buffer)
	}

	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	response := new(dns.Msg)
	if err := response.Unpack(buffer[:n]); err != nil {
		return fmt.Errorf("failed to parse DNS response: %v", err)
	}

	if response.Rcode != dns.RcodeSuccess {
		return fmt.Errorf("DNS query failed with rcode: %s", dns.RcodeToString[response.Rcode])
	}

	// Extract and validate DNSCrypt certificates from TXT records
	var validCerts int
	var lastCertErr error
	now := time.Now()

	for _, rr := range response.Answer {
		txt, ok := rr.(*dns.TXT)
		if !ok {
			continue
		}

		// Join TXT strings and unescape (matching dnscrypt-proxy's approach)
		binCert := packTXTRR(strings.Join(txt.Txt, ""))

		if err := validateDNSCryptCert(binCert, now, stamp.ServerPk); err != nil {
			lastCertErr = err
			continue
		}
		validCerts++
	}

	if validCerts == 0 {
		if lastCertErr != nil {
			return fmt.Errorf("no valid certificates: %v", lastCertErr)
		}
		return fmt.Errorf("no certificates found in response")
	}

	return nil
}

// packTXTRR unescapes a TXT record string from the dns library
// Copied from dnscrypt-proxy's dnsutils.go
func packTXTRR(s string) []byte {
	bs := []byte(s)
	msg := make([]byte, 0, len(bs))
	for i := 0; i < len(bs); i++ {
		if bs[i] == '\\' {
			i++
			if i == len(bs) {
				break
			}
			if i+2 < len(bs) && isDigit(bs[i]) && isDigit(bs[i+1]) && isDigit(bs[i+2]) {
				msg = append(msg, dddToByte(bs[i:]))
				i += 2
			} else if bs[i] == 't' {
				msg = append(msg, '\t')
			} else if bs[i] == 'r' {
				msg = append(msg, '\r')
			} else if bs[i] == 'n' {
				msg = append(msg, '\n')
			} else {
				msg = append(msg, bs[i])
			}
		} else {
			msg = append(msg, bs[i])
		}
	}
	return msg
}

func dddToByte(s []byte) byte {
	return (s[0]-'0')*100 + (s[1]-'0')*10 + (s[2] - '0')
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

// CertMagic is the magic bytes at the start of a DNSCrypt certificate
var CertMagic = [4]byte{0x44, 0x4e, 0x53, 0x43} // "DNSC"

// validateDNSCryptCert validates a DNSCrypt certificate
// Certificate format (124 bytes minimum):
// - Magic: 4 bytes ("DNSC")
// - ES Version: 2 bytes (0x0001 = X25519-XSalsa20Poly1305, 0x0002 = X25519-XChaCha20Poly1305)
// - Protocol minor version: 2 bytes
// - Signature: 64 bytes (Ed25519)
// - Resolver public key: 32 bytes
// - Client magic: 8 bytes
// - Serial: 4 bytes
// - TS Start: 4 bytes (Unix timestamp)
// - TS End: 4 bytes (Unix timestamp)
// - Extensions: variable
func validateDNSCryptCert(binCert []byte, now time.Time, serverPk []byte) error {
	if len(binCert) < 124 {
		return fmt.Errorf("certificate too short: %d bytes", len(binCert))
	}

	// Check magic bytes
	if !bytes.Equal(binCert[:4], CertMagic[:]) {
		return fmt.Errorf("invalid cert magic: %x", binCert[:4])
	}

	// Check ES version
	esVersion := binary.BigEndian.Uint16(binCert[4:6])
	if esVersion != 0x0001 && esVersion != 0x0002 {
		return fmt.Errorf("unsupported ES version: 0x%04x", esVersion)
	}

	// Parse timestamps
	tsBegin := binary.BigEndian.Uint32(binCert[116:120])
	tsEnd := binary.BigEndian.Uint32(binCert[120:124])

	if tsBegin >= tsEnd {
		return fmt.Errorf("certificate ends before it starts (%d >= %d)", tsBegin, tsEnd)
	}

	nowUnix := uint32(now.Unix())
	if nowUnix > tsEnd {
		return fmt.Errorf("certificate expired (ended %s)", time.Unix(int64(tsEnd), 0).Format(time.RFC3339))
	}
	if nowUnix < tsBegin {
		return fmt.Errorf("certificate not yet valid (starts %s)", time.Unix(int64(tsBegin), 0).Format(time.RFC3339))
	}

	// TODO: Optionally verify Ed25519 signature using serverPk from stamp
	// signature := binCert[8:72]
	// signed := binCert[72:]
	// if !ed25519.Verify(pk, signed, signature) { ... }

	return nil
}

func (t *Tester) testDoT(stamp stamps.ServerStamp) error {
	host := stamp.ProviderName
	addr := stamp.ServerAddrStr

	if addr == "" {
		addr = host + ":853"
	}

	serverHost := host
	if serverHost == "" {
		serverHost, _, _ = net.SplitHostPort(addr)
	}

	dialer := &net.Dialer{Timeout: t.timeout}
	conn, err := tls.DialWithDialer(dialer, "tcp", addr, &tls.Config{
		ServerName:         serverHost,
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
	})
	if err != nil {
		return fmt.Errorf("TLS connection failed: %v", err)
	}
	defer conn.Close()

	query := new(dns.Msg)
	query.SetQuestion(".", dns.TypeNS)
	query.Id = dns.Id()
	query.RecursionDesired = true

	conn.SetDeadline(time.Now().Add(t.timeout))

	dnsConn := &dns.Conn{Conn: conn}
	if err := dnsConn.WriteMsg(query); err != nil {
		return fmt.Errorf("failed to write query: %v", err)
	}

	_, err = dnsConn.ReadMsg()
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	return nil
}

func (t *Tester) testDoQ(stamp stamps.ServerStamp) error {
	return fmt.Errorf("DoQ testing not implemented")
}

func (t *Tester) testRelay(stamp stamps.ServerStamp) error {
	addr := stamp.ServerAddrStr
	if addr == "" {
		return fmt.Errorf("no server address in stamp")
	}

	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)
	defer cancel()

	dialer := net.Dialer{Timeout: t.timeout}

	// DNSCrypt relays primarily use UDP
	conn, err := dialer.DialContext(ctx, "udp", addr)
	if err != nil {
		// Fall back to TCP test
		conn, err = dialer.DialContext(ctx, "tcp", addr)
		if err != nil {
			return fmt.Errorf("connection failed (UDP and TCP): %v", err)
		}
		conn.Close()
		return nil
	}
	defer conn.Close()

	// Send a minimal probe packet to check if the relay is responsive
	// DNSCrypt relays expect DNSCrypt traffic, but we can at least verify
	// the port is open and not firewalled by sending a small packet
	conn.SetDeadline(time.Now().Add(t.timeout))

	// Send a small probe (this won't be valid DNSCrypt, but helps detect
	// firewalled ports that would return ICMP unreachable)
	probe := []byte{0x00}
	if _, err := conn.Write(probe); err != nil {
		return fmt.Errorf("UDP write failed: %v", err)
	}

	// For UDP, we won't get a response to an invalid packet, but if the
	// port is closed/filtered, the Write or a subsequent Read would fail
	// with an ICMP error on some systems. Give it a brief moment.
	conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
	buf := make([]byte, 1)
	_, err = conn.Read(buf)
	// Timeout is expected (no response to invalid probe), but connection
	// refused or network unreachable indicates the relay is down
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			// Timeout is OK - relay didn't respond to invalid probe but port is open
			return nil
		}
		// Check for ICMP errors that indicate port is closed
		if isConnectionRefused(err) {
			return fmt.Errorf("UDP port closed: %v", err)
		}
	}

	return nil
}

func isConnectionRefused(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "network is unreachable") ||
		strings.Contains(errStr, "no route to host")
}

func (t *Tester) testODoHTarget(stamp stamps.ServerStamp) error {
	host := stamp.ProviderName

	// ODoH targets must serve their public key configs at /.well-known/odohconfigs
	// The stamp.Path is for DNS queries, not for config discovery
	u := &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   "/.well-known/odohconfigs",
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("ODoH config fetch failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ODoH config returned status %d", resp.StatusCode)
	}

	// Read and validate the ODoHConfigs response
	// ODoHConfigs is a sequence of ODoHConfig structures:
	// - 2 bytes: length of config
	// - N bytes: ODoHConfig (HPKE key configuration)
	body, err := io.ReadAll(io.LimitReader(resp.Body, 65535))
	if err != nil {
		return fmt.Errorf("failed to read ODoH config: %v", err)
	}

	if len(body) < 4 {
		return fmt.Errorf("ODoH config too short: %d bytes", len(body))
	}

	// Validate it looks like HPKE key config (first 2 bytes are length, should be reasonable)
	configLen := int(body[0])<<8 | int(body[1])
	if configLen < 8 || configLen > len(body)-2 {
		return fmt.Errorf("invalid ODoH config length: %d (body: %d bytes)", configLen, len(body))
	}

	return nil
}

func (t *Tester) testODoHRelay(stamp stamps.ServerStamp) error {
	host := stamp.ProviderName
	path := stamp.Path
	if path == "" {
		path = "/"
	}

	u := &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}

	// ODoH relays expect POST requests with OHTTP messages (message/ohttp-req content type)
	// We send a minimal probe to verify the relay endpoint is alive
	// A proper OHTTP request would require HPKE encryption, so we just verify connectivity
	// The relay should respond with some HTTP status (even 400/415 is fine - it means relay is running)
	req, err := http.NewRequest("POST", u.String(), bytes.NewReader([]byte{}))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "message/ohttp-req")

	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("ODoH relay request failed: %v", err)
	}
	defer resp.Body.Close()

	// Status 400/415/422 etc are acceptable - they mean the relay is running
	// but rejected our malformed request (expected). Only 5xx errors or
	// connection failures indicate a problem.
	if resp.StatusCode >= 500 {
		return fmt.Errorf("ODoH relay returned server error: %d", resp.StatusCode)
	}

	return nil
}
