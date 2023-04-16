package utils

import "github.com/miekg/dns"

const (
	ApiPrefix     = "/api/v1"
	DNSHistoryLen = 30
)

var (
	DnsTypeMap = map[string]uint16{
		"a":       dns.TypeA,
		"aaaa":    dns.TypeAAAA,
		"cname":   dns.TypeCNAME,
		"mx":      dns.TypeMX,
		"ns":      dns.TypeNS,
		"ptr":     dns.TypePTR,
		"soa":     dns.TypeSOA,
		"srv":     dns.TypeSRV,
		"txt":     dns.TypeTXT,
		"dnskey":  dns.TypeDNSKEY,
		"ds":      dns.TypeDS,
		"nsec":    dns.TypeNSEC,
		"nsec3":   dns.TypeNSEC3,
		"rrsig":   dns.TypeRRSIG,
		"afsdb":   dns.TypeAFSDB,
		"atma":    dns.TypeATMA,
		"caa":     dns.TypeCAA,
		"cert":    dns.TypeCERT,
		"dhcid":   dns.TypeDHCID,
		"dname":   dns.TypeDNAME,
		"hinfo":   dns.TypeHINFO,
		"isdn":    dns.TypeISDN,
		"loc":     dns.TypeLOC,
		"mb":      dns.TypeMB,
		"mg":      dns.TypeMG,
		"minfo":   dns.TypeMINFO,
		"mr":      dns.TypeMR,
		"naptr":   dns.TypeNAPTR,
		"nsapptr": dns.TypeNSAPPTR,
		"rp":      dns.TypeRP,
		"rt":      dns.TypeRT,
		"tlsa":    dns.TypeTLSA,
		"x25":     dns.TypeX25,
	}
)
