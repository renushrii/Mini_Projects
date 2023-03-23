package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	p := newParser()

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		p.lines++

		parsed, err := parse(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		domain, visits := parsed.domain, parsed.visits

		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}

		p.total += visits
		p.sum[domain] = result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}

	}

	fmt.Printf("%-30s %10s\n", "Domain", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("\n%-30s %10d\n", "Total", p.total)

}
