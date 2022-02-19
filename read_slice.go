package bytebuilder

// ReadBytes removes the first n bytes from b and returns it.
// If the read failed, returns nil.
func ReadBytes(b *[]byte, n int) []byte {

	if len(*b) < n || n < 0 {
		return nil
	}

	v := (*b)[:n]
	*b = (*b)[n:]

	return v
}

// Skip removes the first n bytes from b.
// Returns whether it was successful.
func Skip(b *[]byte, n int) bool {
	return ReadBytes(b, n) != nil
}

// ReadUint8 removes the first byte from b and returns it as an uint8.
// The bool indicates whether the read was successful.
func ReadUint8(b *[]byte) (uint8, bool) {

	v := ReadBytes(b, 1)
	if v == nil {
		return 0, false
	}

	return uint8(v[0]), true
}

// ReadUint16 removes the first bytes from b and returns it as an uint16.
// The bool indicates whether the read was successful.
func ReadUint16(b *[]byte) (uint16, bool) {

	v := ReadBytes(b, 2)
	if v == nil {
		return 0, false
	}

	return uint16(v[0])<<8 | uint16(v[1]), true
}

// ReadUint24 removes the first bytes from b and returns it as a uint32.
// The bool indicates whether the read was successful.
func ReadUint24(b *[]byte) (uint32, bool) {

	v := ReadBytes(b, 3)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<16 | uint32(v[1])<<8 | uint32(v[2]), true
}

// ReadUint32 removes the first bytes from b and returns it as an uint32.
// The bool indicates whether the read was successful.
func ReadUint32(b *[]byte) (uint32, bool) {

	v := ReadBytes(b, 4)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<24 | uint32(v[1])<<16 | uint32(v[2])<<8 | uint32(v[3]), true
}

// ReadInt removes the first bytes (depends on bitSize) from b and returns it as an int.
func ReadInt(b *[]byte, bitSize int) (int, bool) {

	switch bitSize {
	case 8:
		l, ok := ReadUint8(b)
		return int(l), ok
	case 16:
		l, ok := ReadUint16(b)
		return int(l), ok
	case 24:
		l, ok := ReadUint24(b)
		return int(l), ok
	case 32:
		l, ok := ReadUint32(b)
		return int(l), ok
	default:
		panic("invalid bitSize value")
	}
}
