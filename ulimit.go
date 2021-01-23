// Package ulimit contains methods for setting resource limits.
package ulimit

// GetRlimit returns the soft and hard resource limits.
func GetRlimit() (uint64, uint64, error) {
	return getRlimit()
}

// SetRlimit attempts to set the soft and hard resource limits.
// If the hard limit fails to update only the soft limit is set.
func SetRlimit(target uint64) error {
	soft, hard, err := getRlimit()
	if err != nil {
		return err
	}

	if target <= soft {
		return nil
	}

	if err = setRlimit(target, target); err == nil {
		return nil
	}

	if target > hard {
		target = hard
	}

	return setRlimit(target, hard)
}
