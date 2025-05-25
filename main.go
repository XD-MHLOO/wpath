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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
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

		segments := strings.Split(u.Path, "/")
		for i, s := range segments {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			if i == len(segments)-1 {
				if idx := strings.LastIndex(s, "."); idx > 0 {
					s = s[:idx]
				}
			}

			if _, ok := seen[s]; !ok {
				writer.WriteString(s + "\n")
				seen[s] = struct{}{}
			}

		}
	}
}
