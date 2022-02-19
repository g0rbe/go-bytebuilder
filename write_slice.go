package bytebuilder

// WriteBytes appends bytes at the end of b.
func WriteBytes(b *[]byte, bytes ...byte) {
	*b = append(*b, bytes...)
}

// WriteUint8 appends v at the end of b.
func WriteUint8(b *[]byte, v uint8) {
	WriteBytes(b, byte(v))
}

// WriteUint16 appends v at the end of b.
func WriteUint16(b *[]byte, v uint16) {
	WriteBytes(b, byte(v>>8), byte(v))
}

// WriteUint24 appends v at the end of b.
func WriteUint24(b *[]byte, v uint32) {
	WriteBytes(b, byte(v>>16), byte(v>>8), byte(v))
}

// WriteUint32 appends v at the end of b.
func WriteUint32(b *[]byte, v uint32) {
	WriteBytes(b, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt append v at the end of b with size bitSize.
func WriteInt(b *[]byte, v int, bitSize int) {
	switch bitSize {
	case 8:
		WriteUint8(b, uint8(v))
	case 16:
		WriteUint16(b, uint16(v))
	case 24:
		WriteUint24(b, uint32(v))
	case 32:
		WriteUint32(b, uint32(v))
	default:
		panic("invalid bitSize value")
	}
}
