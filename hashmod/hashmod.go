/*
 The package hashmod is a utility for taking lists of unique identifiers
 and placing them in a designated set of buckets with (hopefully)
 uniform distribution. The intial use case is for returning a boolean
 value for a given hostname or IP address to enable or disable a feature
 without requiring a global directory.
*/
package hashmod

import (
	"encoding/binary"
	"hash"
)

type Hashmod struct {
	buckets uint64
	hasher  hash.Hash
}

func New(buckets uint64, hasher hash.Hash) *Hashmod {
	instance := &Hashmod{
		buckets: buckets,
		hasher:  hasher,
	}

	return instance
}

func (h *Hashmod) IsHostEnabled(value string, buckets uint64) bool {
	hashedValAsBytes := h.compute(value)

	hashedValAsInt := byteArrayToUint64(hashedValAsBytes)

	modulo := hashedValAsInt % buckets

	var enabled bool

	if modulo == 0 {
		enabled = true
	} else {
		enabled = false
	}

	return enabled
}

func (h *Hashmod) compute(value string) []byte {
	h.hasher.Write([]byte(value))

	hashedValAsBytes := h.hasher.Sum(nil)

	h.hasher.Reset()

	return hashedValAsBytes
}

func byteArrayToUint64(arr []byte) uint64 {
	return binary.BigEndian.Uint64(arr)
}
