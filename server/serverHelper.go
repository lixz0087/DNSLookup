package server

import (
	"DNSLookup/utils"
	"errors"
	"github.com/miekg/dns"
	"net"
	"strings"
)

// validDigInput valid all fields in DigInput are non-empty
func validDigInput(digInput *DigInput) error {
	if digInput.Domain == "" {
		return errors.New("please provide domain")
	}

	if digInput.DnsServer == "" {
		return errors.New("please provide dns server")
	}

	if digInput.DnsType == "" {
		return errors.New("please provide dns type")
	}

	return nil
}

// executeDig utilized Go library miekg/dns to do the DNS lookup
func executeDig(c *dns.Client, domain string, dnsType uint16, dnsServer string) (interface{}, string, error) {
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(domain), dnsType)
	msg.RecursionDesired = true

	r, _, err := c.Exchange(msg, net.JoinHostPort(dnsServer, "53"))
	if r == nil {
		return nil, "", err
	}

	var resSlice []string

	for _, ansStr := range r.Answer {
		resSlice = append(resSlice, ansStr.String())
	}

	ansStr := strings.Join(resSlice, "\n")

	return r.Answer, ansStr, nil
}

// getDNSType bases on the user input return a DNS type
func getDNSType(inputType string) (uint16, error) {
	inputType = strings.ToLower(inputType)
	if dnsType, exist := utils.DnsTypeMap[inputType]; exist {
		return dnsType, nil
	} else {
		return 0, errors.New("ERROR: can't find matched DNS type")
	}
}
