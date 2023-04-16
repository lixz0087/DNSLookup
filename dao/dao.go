package dao

import (
	"DNSLookup/utils"
	"time"
)

// DNSLookupDAO is the data access object for the project, it stores DNS lookup histories.
type DNSLookupDAO struct {
	LookupHistory []*DNSLookupRecord
}

type DNSLookupRecord struct {
	Domain     string      `json:"domain"`
	DnsType    string      `json:"dns_type"`
	DnsServer  string      `json:"dns_server"`
	DnsRecords interface{} `json:"dns_records"`
	Timestamp  time.Time   `json:"time_stamp"`
}

// CreateDNSLookupDAO creates a DNSLookupDAO
func CreateDNSLookupDAO() *DNSLookupDAO {
	return &DNSLookupDAO{
		LookupHistory: []*DNSLookupRecord{},
	}
}

// AddDNSHistoryItem add a DNS history item to the history slice
func (dao *DNSLookupDAO) AddDNSHistoryItem(recordItem *DNSLookupRecord) {
	dao.LookupHistory = append(dao.LookupHistory, recordItem)

	if len(dao.LookupHistory) > utils.DNSHistoryLen {
		dao.LookupHistory = dao.LookupHistory[1:]
	}
}

// GetAllDNSHistory gets DNS histories
func (dao *DNSLookupDAO) GetAllDNSHistory() []*DNSLookupRecord {
	return dao.LookupHistory
}
