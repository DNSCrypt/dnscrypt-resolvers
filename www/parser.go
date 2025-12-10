package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	stamps "github.com/jedisct1/go-dnsstamps"
)

type ResolverType string

const (
	TypeResolver     ResolverType = "resolver"
	TypeRelay        ResolverType = "relay"
	TypeODoHServer   ResolverType = "odoh_server"
	TypeODoHRelay    ResolverType = "odoh_relay"
	TypeOnionService ResolverType = "onion"
)

type Resolver struct {
	Name        string
	Description string
	Stamps      []string
	Type        ResolverType
	SourceFile  string
}

type Parser struct {
	dir string
}

func NewParser(dir string) *Parser {
	return &Parser{dir: dir}
}

func (p *Parser) ParseAll() ([]Resolver, error) {
	var allResolvers []Resolver

	files := []struct {
		name       string
		typ        ResolverType
		stampTypes []stamps.StampProtoType
	}{
		{"public-resolvers.md", TypeResolver, []stamps.StampProtoType{
			stamps.StampProtoTypeDNSCrypt,
			stamps.StampProtoTypeDoH,
			stamps.StampProtoTypeTLS,
			stamps.StampProtoTypeDoQ,
		}},
		{"relays.md", TypeRelay, []stamps.StampProtoType{
			stamps.StampProtoTypeDNSCryptRelay,
		}},
		{"odoh-servers.md", TypeODoHServer, []stamps.StampProtoType{
			stamps.StampProtoTypeODoHTarget,
		}},
		{"odoh-relays.md", TypeODoHRelay, []stamps.StampProtoType{
			stamps.StampProtoTypeODoHRelay,
		}},
		{"onion-services.md", TypeOnionService, nil},
	}

	for _, f := range files {
		path := filepath.Join(p.dir, f.name)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		resolvers, err := p.parseFile(path, f.typ)
		if err != nil {
			return nil, err
		}
		allResolvers = append(allResolvers, resolvers...)
	}

	return allResolvers, nil
}

func (p *Parser) parseFile(path string, typ ResolverType) ([]Resolver, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var resolvers []Resolver
	var current *Resolver

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "## ") {
			if current != nil && len(current.Stamps) > 0 {
				resolvers = append(resolvers, *current)
			}
			name := strings.TrimPrefix(line, "## ")
			current = &Resolver{
				Name:       name,
				Type:       typ,
				SourceFile: filepath.Base(path),
			}
		} else if current != nil {
			if strings.HasPrefix(line, "sdns://") {
				current.Stamps = append(current.Stamps, line)
			} else if line != "" && !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "--") {
				if current.Description == "" {
					current.Description = line
				} else {
					current.Description += " " + line
				}
			}
		}
	}

	if current != nil && len(current.Stamps) > 0 {
		resolvers = append(resolvers, *current)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return resolvers, nil
}

func GetStampProto(stampStr string) (stamps.StampProtoType, error) {
	stamp, err := stamps.NewServerStampFromString(stampStr)
	if err != nil {
		return 0, err
	}
	return stamp.Proto, nil
}

func ParseStamp(stampStr string) (stamps.ServerStamp, error) {
	return stamps.NewServerStampFromString(stampStr)
}
