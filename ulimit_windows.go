// +build windows

package ulimit

func getRlimit() (uint64, uint64, error) {
	return 0, 0, nil
}

func setRlimit(soft uint64, max uint64) error {
	return nil
}
