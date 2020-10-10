package cidrgrep

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"regexp"
)

// IPAddressRegexp is regexp for matching IP address
var IPAddressRegexp = regexp.MustCompile(`[0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}`)

// Filter reads from its input reader, looking for lines containing IP addresses
// which match the supplied pattern, which is a CIDR network address expression
// as accepted by net.ParseCIDR. Any matching lines are written to the output
// writer. All other input is ignored. The input reader is read to completion,
// or the first read error. A prefix is prefixed to the output for when parsing
// multiple files
func Filter(r io.Reader, w io.Writer, pattern string, prefix string) {
	scanner := bufio.NewScanner(r)

	// IP is not needed here, just the *IPNet
	_, ipv4Net, err := net.ParseCIDR(pattern)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		for _, match := range IPAddressRegexp.FindAllString(line, -1) {
			if ipv4Net.Contains(net.ParseIP(match)) {
				fmt.Fprintf(w, "%s%s\n", prefix, line)
				break
			}
		}
	}
}
