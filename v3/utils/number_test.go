package utils

import (
	"testing"
)

func TestParseNumber(t *testing.T) {
	tests := []struct {
		name          string
		number        string
		wantCountry   int64
		wantNational  int64
		wantRegion    string
		shouldErr     bool
	}{
		{
			name:          "Valid Chinese number with +86",
			number:        "+8613800138000",
			wantCountry:   86,
			wantNational:  13800138000,
			wantRegion:    "CN",
			shouldErr:     false,
		},
		{
			name:          "Valid Chinese number without +",
			number:        "8613800138000",
			wantCountry:   86,
			wantNational:  13800138000,
			wantRegion:    "CN",
			shouldErr:     false,
		},
		{
			name:          "Valid US number",
			number:        "+12025551234",
			wantCountry:   1,
			wantNational:  2025551234,
			wantRegion:    "US",
			shouldErr:     false,
		},
		{
			name:          "Valid US number without +",
			number:        "12025551234",
			wantCountry:   1,
			wantNational:  2025551234,
			wantRegion:    "US",
			shouldErr:     false,
		},
		{
			name:          "Invalid number format",
			number:        "invalid",
			wantCountry:   0,
			wantNational:  0,
			wantRegion:    "",
			shouldErr:     true,
		},
		{
			name:          "Empty number",
			number:        "",
			wantCountry:   0,
			wantNational:  0,
			wantRegion:    "",
			shouldErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			country, national, region, err := ParseNumber(tt.number)

			if tt.shouldErr && err == nil {
				t.Errorf("ParseNumber() expected error, got nil")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("ParseNumber() unexpected error: %v", err)
			}

			if !tt.shouldErr {
				if country != tt.wantCountry {
					t.Errorf("ParseNumber() country code = %d, want %d", country, tt.wantCountry)
				}
				if national != tt.wantNational {
					t.Errorf("ParseNumber() national number = %d, want %d", national, tt.wantNational)
				}
				if region != tt.wantRegion {
					t.Errorf("ParseNumber() region = %s, want %s", region, tt.wantRegion)
				}
			}
		})
	}
}

func TestParseNumberWithVariousCountries(t *testing.T) {
	tests := []struct {
		number      string
		expectError bool
	}{
		{"+8613800138000", false},    // China
		{"+12025551234", false},      // USA
		{"+447911123456", false},     // UK
		{"+81312345678", false},      // Japan
		{"+33123456789", false},      // France
		{"+49301234567", false},      // Germany
		{"+invalid", true},           // Invalid
		{"+1", true},                 // Too short
	}

	for _, tt := range tests {
		t.Run(tt.number, func(t *testing.T) {
			_, _, _, err := ParseNumber(tt.number)
			if tt.expectError && err == nil {
				t.Errorf("ParseNumber(%s) expected error, got nil", tt.number)
			}
			if !tt.expectError && err != nil {
				t.Errorf("ParseNumber(%s) unexpected error: %v", tt.number, err)
			}
		})
	}
}

