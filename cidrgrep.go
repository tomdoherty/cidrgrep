package cidrgrep

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"regexp"
)

// Filter filters input for lines containing an IP within cidr
func Filter(input io.Reader, output io.Writer, cidr string) {
	scanner := bufio.NewScanner(input)

	_, ipv4Net, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`[0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}`)
	for scanner.Scan() {
		line := scanner.Text()
		for _, match := range re.FindAllString(line, -1) {
			if ipv4Net.Contains(net.ParseIP(match)) {
				fmt.Fprintf(output, "%s", line)
				break
			}
		}
	}
}
