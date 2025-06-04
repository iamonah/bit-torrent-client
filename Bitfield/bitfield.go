package bitfield

type Bitfield []byte

// tells if a bit field has a particular index set
func (bf Bitfield) HasPiece(index int) bool {
	byteIndex := index / 8 //byte postion
	offset := index % 8    //bit position
	if byteIndex < 0 || byteIndex >= len(bf) {
		return false
	}
	bitPos := 7 - offset
	return bf[byteIndex]&1<<(bitPos) != 0
}

// set a bit in the field
func (bf Bitfield) SetPiece(index int) {
	byteIndex := index / 8
	offset := index % 8

	//silently discard invalid bounded index
	if byteIndex < 0 || byteIndex >= len(bf) {
		return
	}
	bitPos := 7 - offset
	bf[byteIndex] |= 1 << (bitPos)
}
