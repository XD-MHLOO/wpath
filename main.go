package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

var Version = "dev"

func main() {
	showVersion := flag.Bool("v", false, "Print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println("Version:", Version)
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	seen := make(map[string]struct{})

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		u, err := url.Parse(line)
		if err != nil {
			continue
		}
		if idx := strings.LastIndex(u.Path, "."); idx > 0 {
			u.Path = u.Path[:idx]
		}
		segments := strings.SplitSeq(u.Path, "/")
		for s := range segments {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			if _, ok := seen[s]; !ok {
				fmt.Println(s)
				seen[s] = struct{}{}
			}

		}
	}
}
