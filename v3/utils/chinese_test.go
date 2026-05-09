package utils

import (
	"testing"
)

func TestProcessNumberForChinese(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		wantToContains  string
		wantRegion      string
		shouldErr       bool
	}{
		{
			name:           "11-digit Chinese number (17601205205 case)",
			input:          "17601205205",
			wantToContains: "+8617601205205",
			wantRegion:     "CN",
			shouldErr:      false,
		},
		{
			name:           "11-digit Chinese number (13800138000)",
			input:          "13800138000",
			wantToContains: "+8613800138000",
			wantRegion:     "CN",
			shouldErr:      false,
		},
		{
			name:           "Chinese number with +86",
			input:          "+8613800138000",
			wantToContains: "+8613800138000",
			wantRegion:     "CN",
			shouldErr:      false,
		},
		{
			name:           "Chinese number with country code",
			input:          "8613800138000",
			wantToContains: "+8613800138000",
			wantRegion:     "CN",
			shouldErr:      false,
		},
		{
			name:           "US number with +1",
			input:          "+12025551234",
			wantToContains: "+12025551234",
			wantRegion:     "US",
			shouldErr:      false,
		},
		{
			name:           "US number without +",
			input:          "12025551234",
			wantToContains: "+12025551234",
			wantRegion:     "US",
			shouldErr:      false,
		},
		{
			name:           "UK number",
			input:          "+447911123456",
			wantToContains: "+447911123456",
			wantRegion:     "GG", // phonenumbers library returns GG for this number
			shouldErr:      false,
		},
		{
			name:           "Invalid number",
			input:          "invalid",
			wantToContains: "",
			wantRegion:     "",
			shouldErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processed, _, _, region, err := ProcessNumberForChinese(tt.input)

			if tt.shouldErr && err == nil {
				t.Errorf("ProcessNumberForChinese() expected error, got nil")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("ProcessNumberForChinese() unexpected error: %v", err)
			}

			if !tt.shouldErr {
				if processed != tt.wantToContains {
					t.Errorf("ProcessNumberForChinese() processed = %s, want %s", processed, tt.wantToContains)
				}
				if region != tt.wantRegion {
					t.Errorf("ProcessNumberForChinese() region = %s, want %s", region, tt.wantRegion)
				}
			}
		})
	}
}

func TestProcessNumberForChineseEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{"Empty string", "", true},
		{"Non-numeric", "abcdefghijk", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, _, _, err := ProcessNumberForChinese(tt.input)
			if tt.wantError && err == nil {
				t.Errorf("ProcessNumberForChinese(%s) expected error, got nil", tt.input)
			}
			if !tt.wantError && err != nil {
				t.Errorf("ProcessNumberForChinese(%s) unexpected error: %v", tt.input, err)
			}
		})
	}
}

func TestProcessNumberForChineseDifferentFormats(t *testing.T) {
	// Test that different input formats for the same number produce consistent processed output
	formats := []string{
		"13800138000",
		"+8613800138000",
		"8613800138000",
	}

	var firstProcessed string
	var firstRegion string

	for i, format := range formats {
		processed, _, _, region, err := ProcessNumberForChinese(format)
		if err != nil {
			t.Fatalf("ProcessNumberForChinese(%s) unexpected error: %v", format, err)
		}

		if i == 0 {
			firstProcessed = processed
			firstRegion = region
		} else {
			if processed != firstProcessed {
				t.Errorf("Format %s: Inconsistent processed number: got %s, want %s", format, processed, firstProcessed)
			}
			if region != firstRegion {
				t.Errorf("Format %s: Inconsistent region: got %s, want %s", format, region, firstRegion)
			}
		}
	}
}









