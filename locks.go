package appconfig

import "fmt"

// Lock represents a key/value pair that is currently
// locked on App Config
type Lock struct {
	key   string
	label string
}

// LockKV locks the key/value pair at key, optionally with a
// label. If you don't want to pass a label, pass an empty
// string.
//
// Returns an error if the lock couldn't be acquired,
// and returns the lock itself if the lock could be acquired
func (cl *Client) LockKV(key, label string) (*Lock, error) {
	lck := &Lock{
		key:   key,
		label: label,
	}
	path := fmt.Sprintf(
		"/locks/%s?api-version=%s",
		key,
		cl.cfg.APIVersion(),
	)
	if label != "" {
		path += fmt.Sprintf("&label=%s", label)
	}
	// TODO: PUT HTTP request to path
	return lck, nil
}

// UnlockKV unlocks the key/value pair at key, optionally
// with a label.
//
// Returns nil if the key was locked and unlock succeeded,
// and a non-nil error if either of those cases are not true
func (cl *Client) UnlockKV(lck *Lock) error {
	path := fmt.Sprintf(
		"/locks/%s?api-version=%s",
		lck.key,
		cl.cfg.APIVersion(),
	)
	if lck.label == "" {
		path += fmt.Sprintf("?label=%s", lck.label)
	}
	// TODO: DELETE HTTP request to path
	return nil
}
