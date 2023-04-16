package server

import (
	"DNSLookup/dao"
	"errors"
	"fmt"
	"github.com/miekg/dns"
	"time"
)

// DigInput is the input for /dig API
type DigInput struct {
	Domain    string `json:"domain"`
	DnsType   string `json:"dns_type"`
	DnsServer string `json:"dns_server"`
}

// DigHandler handlers request to /dig API
func DigHandler(digInput *DigInput, dnsLookupDao *dao.DNSLookupDAO) (string, error) {
	err := validDigInput(digInput)
	if err != nil {
		return "", errors.New(fmt.Sprintf("ERROR: invald input: %v", err.Error()))
	}

	c := new(dns.Client)
	dnsType, err := getDNSType(digInput.DnsType)
	if err != nil {
		return "", errors.New(fmt.Sprintf("ERROR: Got error get DNS type. Input type: %v, error msg: %v", digInput.DnsType, err.Error()))
	}

	ans, ansStr, err := executeDig(c, digInput.Domain, dnsType, digInput.DnsServer)
	if err != nil {
		return "", errors.New(fmt.Sprintf("ERROR: Got error when dig domain. Domain name: %v, error msg: %v", digInput.Domain, err.Error()))
	}

	dnsLookupDao.AddDNSHistoryItem(&dao.DNSLookupRecord{
		Domain:     digInput.Domain,
		DnsType:    digInput.DnsType,
		DnsServer:  digInput.DnsServer,
		DnsRecords: ans,
		Timestamp:  time.Now(),
	})

	return ansStr, nil
}

// DigHistoryHandler handlers request to /dig/history API
func DigHistoryHandler(dnsLookupDao *dao.DNSLookupDAO) []*dao.DNSLookupRecord {
	return dnsLookupDao.GetAllDNSHistory()
}
