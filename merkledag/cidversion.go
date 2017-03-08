package merkledag

import (
	"fmt"
)

// CidVersion is a a cid version number.  Do not set directly, instead
// use NewCidVersion to create.  It is safe to panic if it is invalid.
type CidVersion uint64

const (
	CID0 CidVersion = 0
	CID1 CidVersion = 1
)

func NewCidVersion(version int) (CidVersion, error) {
	switch version {
	case 0, 1:
		return CidVersion(version), nil
	default:
		return CidVersion(0), fmt.Errorf("unknown CID version: %d", version)
	}
}
