package bech32

import (
	"strings"
	"testing"
)

func TestSegwitAddrDecodeAcceptsUppercaseHRP(t *testing.T) {
	program := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	addr, err := SegwitAddrEncode("BGL", 0, program)
	if err != nil {
		t.Fatalf("SegwitAddrEncode() error = %v", err)
	}
	if addr != strings.ToUpper(addr) {
		t.Fatalf("SegwitAddrEncode() with uppercase HRP returned non-uppercase address %q", addr)
	}

	version, decoded, err := SegwitAddrDecode("BGL", addr)
	if err != nil {
		t.Fatalf("SegwitAddrDecode() with uppercase HRP/address error = %v", err)
	}
	if version != 0 {
		t.Fatalf("SegwitAddrDecode() version = %d, want 0", version)
	}
	if len(decoded) != len(program) {
		t.Fatalf("decoded program length = %d, want %d", len(decoded), len(program))
	}
	for i := range program {
		if decoded[i] != program[i] {
			t.Fatalf("decoded[%d] = %d, want %d", i, decoded[i], program[i])
		}
	}
}
