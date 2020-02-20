package appconfig

import "fmt"

type Lock struct{}

// LockKV locks the key/value pair at key, optionally with a
// label. If you don't want to pass a label, pass an empty
// string.
//
// Returns an error if the lock couldn't be acquired,
// and returns the lock itself if the lock could be acquired
func (a *Client) LockKV(key, label string) (*Lock, error) {
	lck := &Lock{}
	urlStr := fmt.Sprintf(
		"/locks/%s?api-version=%s",
		key,
		a.cfg.APIVersion(),
	)
	if label != "" {
		urlStr += fmt.Sprintf("&label=%s", label)
	}
	return lck, nil
}
