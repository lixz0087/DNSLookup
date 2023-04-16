package dao

import (
	"DNSLookup/utils"
	"testing"
)

// TestAddDNSHistoryItem test function AddDNSHistoryItem
func TestAddDNSHistoryItem(t *testing.T) {
	tests := []struct {
		name      string
		itemCount int
	}{
		{
			name:      "Successfully add 3 record items",
			itemCount: 3,
		},
		{
			name:      "LookupHistory will not longer than utils.DNSHistoryLen",
			itemCount: utils.DNSHistoryLen + 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dnsLookupDAO := CreateDNSLookupDAO()

			for i := 0; i < tt.itemCount; i++ {
				dnsLookupDAO.AddDNSHistoryItem(&DNSLookupRecord{})
			}

			if tt.itemCount <= utils.DNSHistoryLen {
				if len(dnsLookupDAO.LookupHistory) != tt.itemCount {
					t.Errorf("dnsLookupDAO.LookupHistory should have the expected length. "+
						"dnsLookupDAO.LookupHistory: %v, expected length: %v", dnsLookupDAO.LookupHistory, tt.itemCount)
				}
			} else {
				if len(dnsLookupDAO.LookupHistory) > utils.DNSHistoryLen {
					t.Errorf("dnsLookupDAO.LookupHistory length should bigger than utils.DNSHistoryLen"+
						"dnsLookupDAO.LookupHistory: %v, expected length: %v", dnsLookupDAO.LookupHistory, utils.DNSHistoryLen)
				}
			}
		})
	}
}
