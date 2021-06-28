package oksvg

import (
	"testing"
)

func TestParseSVGColorNum(t *testing.T) {
	for _, tt := range []struct {
		desc    string
		hex     string
		wantR   uint8
		wantG   uint8
		wantB   uint8
		wantA   uint8
		wantErr bool
	}{
		{"3-char hex", "CCC", 0xCC, 0xCC, 0xCC, 0xFF, false},
		{"4-char hex", "AAAA", 0xAA, 0xAA, 0xAA, 0xAA, false},
		{"6-char hex", "ABCDEF", 0xAB, 0xCD, 0xEF, 0xFF, false},
		{"8-char hex", "ABCDEF56", 0xAB, 0xCD, 0xEF, 0x56, false},
	} {
		t.Run(tt.desc, func(t *testing.T) {
			gotR, gotG, gotB, gotA, gotErr := ParseSVGColorNum(tt.hex)
			if gotR != tt.wantR {
				t.Errorf("unexpected color R: got %d, want %d", gotR, tt.wantR)
			}
			if gotG != tt.wantG {
				t.Errorf("unexpected color G: got %d, want %d", gotG, tt.wantG)
			}
			if gotB != tt.wantB {
				t.Errorf("unexpected color B: got %d, want %d", gotB, tt.wantB)
			}
			if gotA != tt.wantA {
				t.Errorf("unexpected color A: got %d, want %d", gotA, tt.wantA)
			}
			if tt.wantErr != (gotErr != nil) {
				t.Errorf("unexpected error: got %v, want %v", gotErr != nil, tt.wantErr)
			}
		})
	}
}
