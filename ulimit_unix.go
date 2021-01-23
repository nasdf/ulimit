// +build darwin linux netbsd openbsd freebsd

package ulimit

import (
	"golang.org/x/sys/unix"
)

func getRlimit() (uint64, uint64, error) {
	var limit unix.Rlimit
	if err := unix.Getrlimit(unix.RLIMIT_NOFILE, &limit); err != nil {
		return 0, 0, err
	}

	return limit.Cur, limit.Max, nil
}

func setRlimit(soft uint64, max uint64) error {
	limit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}

	return unix.Setrlimit(unix.RLIMIT_NOFILE, &limit)
}
