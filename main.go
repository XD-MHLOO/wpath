package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
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
		segments := strings.SplitSeq(u.Path, "/")
		for s := range segments {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}

			if idx := strings.LastIndex(s, "."); idx != -1 {
				if _, ok := seen[s[:idx]]; !ok {
					fmt.Println(s[:idx])
					seen[s[:idx]] = struct{}{}
				}
			} else {
				if _, ok := seen[s]; !ok {
					fmt.Println(s)
					seen[s] = struct{}{}
				}
			}
		}
	}
}
