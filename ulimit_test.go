package ulimit

import (
	"testing"
)

func TestSetRlimit(t *testing.T) {
	const target = uint64(2048)
	if err := SetRlimit(target); err != nil {
		t.Fatal("failed to set rlimit")
	}

	soft, hard, err := GetRlimit()
	if err != nil {
		t.Fatal("failed to get rlimit")
	}

	if soft != target {
		t.Error("unexpected soft target")
	}

	if hard != target {
		t.Error("unexpected hard target")
	}
}

func TestGetRlimit(t *testing.T) {
	soft, hard, err := GetRlimit()
	if err != nil {
		t.Fatal("failed to get rlimit")
	}

	if soft == 0 {
		t.Error("unexpected soft limit")
	}

	if hard == 0 {
		t.Error("unexpected hard limit")
	}
}
