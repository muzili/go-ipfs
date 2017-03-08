package merkledag

import (
	"fmt"
)

// CidVersion is a a cid version number that is guaranteed to be valid
type CidVersion struct {
	version uint64
}

var (
	CID0 = CidVersion{0}
	CID1 = CidVersion{1}
)

func NewCidVersion(version int) (CidVersion, error) {
	switch version {
	case 0, 1:
		return CidVersion{uint64(version)}, nil
	default:
		return CidVersion{}, fmt.Errorf("unknown CID version: %d", version)
	}
}

func (v CidVersion) Version() uint64 {
	return v.version
}
