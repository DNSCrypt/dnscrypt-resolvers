package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type ResolverType string

const (
	TypeResolver   ResolverType = "resolver"
	TypeRelay      ResolverType = "relay"
	TypeODoHServer ResolverType = "odoh_server"
	TypeODoHRelay  ResolverType = "odoh_relay"
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
		name string
		typ  ResolverType
	}{
		{"public-resolvers.md", TypeResolver},
		{"relays.md", TypeRelay},
		{"odoh-servers.md", TypeODoHServer},
		{"odoh-relays.md", TypeODoHRelay},
		// Onion services are skipped - testing requires Tor
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
