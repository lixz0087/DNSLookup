package server

import (
	"strings"
	"testing"
)

// TestValidDigInput test function validDigInput
func TestValidDigInput(t *testing.T) {
	tests := []struct {
		name             string
		digInput         *DigInput
		expectedErrorStr string
	}{
		{
			name: "Successful case: no error",
			digInput: &DigInput{
				Domain:    "test_domain",
				DnsServer: "test_dnsServer",
				DnsType:   "test_type",
			},
		},
		{
			name: "Error case: no domain",
			digInput: &DigInput{
				DnsServer: "test_dnsServer",
				DnsType:   "test_type",
			},
			expectedErrorStr: "please provide domain",
		},
		{
			name: "Error case: no dns_server",
			digInput: &DigInput{
				Domain:  "test_domain",
				DnsType: "test_type",
			},
			expectedErrorStr: "please provide dns server",
		},
		{
			name: "Error case: no dns_type",
			digInput: &DigInput{
				Domain:    "test_domain",
				DnsServer: "test_dnsServer",
			},
			expectedErrorStr: "please provide dns type",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validDigInput(tt.digInput)
			if tt.expectedErrorStr == "" {
				if err != nil {
					t.Errorf("Expected get no error, but got error: %v", err)
				}
			} else if !strings.Contains(err.Error(), tt.expectedErrorStr) {
				t.Errorf("Expected error: %v, got error: %v", tt.expectedErrorStr, err)
			}
		})
	}
}
