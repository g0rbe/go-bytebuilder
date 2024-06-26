package bytebuilder

import "fmt"

// WriteByte appends v at the end of b.
func WriteByte(b *[]byte, v byte) {
	*b = append(*b, v)
}

// WriteBytes appends bytes at the end of b.
func WriteBytes(b *[]byte, v ...byte) {
	*b = append(*b, v...)
}

// WriteUint8 appends v at the end of b.
func WriteUint8(b *[]byte, v uint8) {
	WriteByte(b, byte(v))
}

// WriteInt8 appends v at the end of b.
func WriteInt8(b *[]byte, v int8) {
	WriteByte(b, byte(v))
}

// WriteLittleUint16 appends v at the end of b in little-endian order.
func WriteLittleUint16(b *[]byte, v uint16) {
	WriteBytes(b, byte(v), byte(v>>8))
}

// WriteBigUint16 appends v at the end of b in big-endian order.
func WriteBigUint16(b *[]byte, v uint16) {
	WriteBytes(b, byte(v>>8), byte(v))
}

// WriteUint16 appends v at the end of b in native-endian order.
func WriteUint16(b *[]byte, v uint16) {
	switch NativeEndian {
	case LittleEndian:
		WriteLittleUint16(b, v)
	case BigEndian:
		WriteBigUint16(b, v)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// WriteLittleInt16 appends v at the end of b in little-endian order.
func WriteLittleInt16(b *[]byte, v int16) {
	WriteBytes(b, byte(v), byte(v>>8))
}

// WriteBigInt16 appends v at the end of b in big-endian order.
func WriteBigInt16(b *[]byte, v int16) {
	WriteBytes(b, byte(v>>8), byte(v))
}

// WriteInt16 appends v at the end of b in native-endian order.
func WriteInt16(b *[]byte, v int16) {
	switch NativeEndian {
	case LittleEndian:
		WriteLittleInt16(b, v)
	case BigEndian:
		WriteBigInt16(b, v)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// WriteLittleUint24 appends v at the end of b in little-endian order.
func WriteLittleUint24(b *[]byte, v uint32) {
	WriteBytes(b, byte(v), byte(v>>8), byte(v>>16))
}

// WriteBigUint24 appends v at the end of b in big-endian order.
func WriteBigUint24(b *[]byte, v uint32) {
	WriteBytes(b, byte(v>>16), byte(v>>8), byte(v))
}

// WriteUint24 appends v at the end of b in native-endian order.
func WriteUint24(b *[]byte, v uint32) {
	switch NativeEndian {
	case LittleEndian:
		WriteLittleUint24(b, v)
	case BigEndian:
		WriteBigUint24(b, v)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// WriteInt24 appends v at the end of b in little-endian order.
func WriteLittleInt24(b *[]byte, v int32) {
	WriteBytes(b, byte(v), byte(v>>8), byte(v>>16))
}

// WriteInt24 appends v at the end of b in big-endian order.
func WriteBigInt24(b *[]byte, v int32) {
	WriteBytes(b, byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt24 appends v at the end of b in native-endian order.
func WriteInt24(b *[]byte, v int32) {
	switch NativeEndian {
	case LittleEndian:
		WriteLittleInt24(b, v)
	case BigEndian:
		WriteBigInt24(b, v)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// WriteUint32 appends v at the end of b in little-endian order.
func WriteLittleUint32(b *[]byte, v uint32) {
	WriteBytes(b, byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}

// WriteUint32 appends v at the end of b in big-endian order.
func WriteBigUint32(b *[]byte, v uint32) {
	WriteBytes(b, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteUint32 appends v at the end of b in native-endian order.
func WriteUint32(b *[]byte, v uint32) {
	switch NativeEndian {
	case LittleEndian:
		WriteLittleUint32(b, v)
	case BigEndian:
		WriteBigUint32(b, v)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// WriteInt32 appends v at the end of b in little-endian order.
func WriteLittleInt32(b *[]byte, v int32) {
	WriteBytes(b, byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}

// WriteInt32 appends v at the end of b in big-endian order.
func WriteBigInt32(b *[]byte, v int32) {
	WriteBytes(b, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt32 appends v at the end of b in native-endian order.
func WriteInt32(b *[]byte, v int32) {
	switch NativeEndian {
	case LittleEndian:
		WriteLittleInt32(b, v)
	case BigEndian:
		WriteBigInt32(b, v)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// WriteUint64 appends v at the end of b in little-endian order.
func WriteLittleUint64(b *[]byte, v uint64) {
	WriteBytes(b, byte(v), byte(v>>8), byte(v>>16), byte(v>>24), byte(v>>32), byte(v>>40), byte(v>>48), byte(v>>56))
}

// WriteUint64 appends v at the end of b in big-endian order.
func WriteBigUint64(b *[]byte, v uint64) {
	WriteBytes(b, byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteUint64 appends v at the end of b in native-endian order.
func WriteUint64(b *[]byte, v uint64) {
	switch NativeEndian {
	case LittleEndian:
		WriteLittleUint64(b, v)
	case BigEndian:
		WriteBigUint64(b, v)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// WriteInt64 appends v at the end of b in little-endian order.
func WriteLittleInt64(b *[]byte, v int64) {
	WriteBytes(b, byte(v), byte(v>>8), byte(v>>16), byte(v>>24), byte(v>>32), byte(v>>40), byte(v>>48), byte(v>>56))
}

// WriteInt64 appends v at the end of b in big-endian order.
func WriteBigInt64(b *[]byte, v int64) {
	WriteBytes(b, byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt64 appends v at the end of b in native-endian order.
func WriteInt64(b *[]byte, v int64) {
	switch NativeEndian {
	case LittleEndian:
		WriteLittleInt64(b, v)
	case BigEndian:
		WriteBigInt64(b, v)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}
