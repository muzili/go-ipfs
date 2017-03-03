package merkledag

import (
	"fmt"
)

// CidVersion is a a cid version number that is guaranteed to be valid
type CidVersion struct {
	version uint64
}

func NewCidVersion(version int) (CidVersion, error) {
	switch version {
	case 0:
		return CidVersion{0}, nil
	case 1:
		return CidVersion{1}, nil
	default:
		return CidVersion{}, fmt.Errorf("unknown CID version: %d", version)
	}
}

func (v CidVersion) Version() uint64 {
	return v.version
}

