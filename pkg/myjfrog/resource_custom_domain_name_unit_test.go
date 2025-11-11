package myjfrog

import (
	"testing"
)

func TestNormalizeCertificateData(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Unix line endings (LF) - no change",
			input:    "-----BEGIN CERTIFICATE-----\nMIID\ntest\n-----END CERTIFICATE-----\n",
			expected: "-----BEGIN CERTIFICATE-----\nMIID\ntest\n-----END CERTIFICATE-----\n",
		},
		{
			name:     "Windows line endings (CRLF) - normalized to LF",
			input:    "-----BEGIN CERTIFICATE-----\r\nMIID\r\ntest\r\n-----END CERTIFICATE-----\r\n",
			expected: "-----BEGIN CERTIFICATE-----\nMIID\ntest\n-----END CERTIFICATE-----\n",
		},
		{
			name:     "Mixed line endings - all CRLF converted to LF",
			input:    "-----BEGIN CERTIFICATE-----\r\nMIID\ntest\r\n-----END CERTIFICATE-----\n",
			expected: "-----BEGIN CERTIFICATE-----\nMIID\ntest\n-----END CERTIFICATE-----\n",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "No line endings",
			input:    "test data without line endings",
			expected: "test data without line endings",
		},
		{
			name:     "Multiple consecutive CRLF",
			input:    "line1\r\n\r\nline2\r\n",
			expected: "line1\n\nline2\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeCertificateData(tt.input)
			if result != tt.expected {
				t.Errorf("normalizeCertificateData() = %q, want %q", result, tt.expected)
			}
		})
	}
}
