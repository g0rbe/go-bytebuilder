package bytebuilder

import "fmt"

// ReadByte removes the first byte from b and returns it as a byte.
// The bool indicates whether the read was successful.
func ReadByte(b *[]byte) (byte, bool) {

	if len(*b) < 1 {
		return 0, false
	}

	v := (*b)[0]
	*b = (*b)[1:]

	return v, true
}

// ReadBytes removes the first n bytes from b and returns it.
// If the read failed, returns nil.
func ReadBytes(b *[]byte, n int) []byte {

	if len(*b) < n || n < 1 {
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

// ReadInt8 removes the first byte from b and returns it as an int8.
// The bool indicates whether the read was successful.
func ReadInt8(b *[]byte) (int8, bool) {

	v := ReadBytes(b, 1)
	if v == nil {
		return 0, false
	}

	return int8(v[0]), true
}

// ReadLittleUint16 removes the first bytes from b and returns it as an uint16 in little-endian order.
// The bool indicates whether the read was successful.
func ReadLittleUint16(b *[]byte) (uint16, bool) {

	v := ReadBytes(b, 2)
	if v == nil {
		return 0, false
	}

	return uint16(v[1])<<8 | uint16(v[0]), true
}

// Uint16BE removes the first bytes from b and returns it as an uint16 in big-endian order.
// The bool indicates whether the read was successful.
func ReadBigUint16(b *[]byte) (uint16, bool) {

	v := ReadBytes(b, 2)
	if v == nil {
		return 0, false
	}

	return uint16(v[0])<<8 | uint16(v[1]), true
}

// ReadUint16 removes the first bytes from b and returns it as an uint16 in native endian order.
// The bool indicates whether the read was successful.
func ReadUint16(b *[]byte) (uint16, bool) {

	switch NativeEndian {
	case LittleEndian:
		return ReadLittleUint16(b)
	case BigEndian:
		return ReadBigUint16(b)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// ReadLittleInt16 removes the first bytes from b and returns it as an int16 in little-endian order.
// The bool indicates whether the read was successful.
func ReadLittleInt16(b *[]byte) (int16, bool) {

	v := ReadBytes(b, 2)
	if v == nil {
		return 0, false
	}

	return int16(v[1])<<8 | int16(v[0]), true
}

// ReadBigInt16 removes the first bytes from b and returns it as an int16 in big-endian order.
// The bool indicates whether the read was successful.
func ReadBigInt16(b *[]byte) (int16, bool) {

	v := ReadBytes(b, 2)
	if v == nil {
		return 0, false
	}

	return int16(v[0])<<8 | int16(v[1]), true
}

// ReadInt16 removes the first bytes from b and returns it as an int16 in native endian order.
// The bool indicates whether the read was successful.
func ReadInt16(b *[]byte) (int16, bool) {

	switch NativeEndian {
	case LittleEndian:
		return ReadLittleInt16(b)
	case BigEndian:
		return ReadBigInt16(b)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// ReadLittleUint24 removes the first bytes from b and returns it as a uint32 in little-endian order.
// The bool indicates whether the read was successful.
func ReadLittleUint24(b *[]byte) (uint32, bool) {

	v := ReadBytes(b, 3)
	if v == nil {
		return 0, false
	}

	return uint32(v[2])<<16 | uint32(v[1])<<8 | uint32(v[0]), true
}

// ReadBigUint24 removes the first bytes from b and returns it as a uint32 in big-endian order.
// The bool indicates whether the read was successful.
func ReadBigUint24(b *[]byte) (uint32, bool) {

	v := ReadBytes(b, 3)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<16 | uint32(v[1])<<8 | uint32(v[2]), true
}

// ReadUint24 removes the first bytes from b and returns it as a uint32 in native endian order.
// The bool indicates whether the read was successful.
func ReadUint24(b *[]byte) (uint32, bool) {

	switch NativeEndian {
	case LittleEndian:
		return ReadLittleUint24(b)
	case BigEndian:
		return ReadBigUint24(b)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// ReadInt24 removes the first bytes from b and returns it as a int32 in little-endian order.
// The bool indicates whether the read was successful.
func ReadLittleInt24(b *[]byte) (int32, bool) {

	v := ReadBytes(b, 3)
	if v == nil {
		return 0, false
	}

	return int32(v[2])<<16 | int32(v[1])<<8 | int32(v[0]), true
}

// ReadInt24 removes the first bytes from b and returns it as a int32 in big-endian order.
// The bool indicates whether the read was successful.
func ReadBigInt24(b *[]byte) (int32, bool) {

	v := ReadBytes(b, 3)
	if v == nil {
		return 0, false
	}

	return int32(v[0])<<16 | int32(v[1])<<8 | int32(v[2]), true
}

// ReadInt24 removes the first bytes from b and returns it as a int32 in native endian order.
// The bool indicates whether the read was successful.
func ReadInt24(b *[]byte) (int32, bool) {

	switch NativeEndian {
	case LittleEndian:
		return ReadLittleInt24(b)
	case BigEndian:
		return ReadBigInt24(b)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// ReadLittleUint32 removes the first bytes from b and returns it as an uint32 in little-endian order.
// The bool indicates whether the read was successful.
func ReadLittleUint32(b *[]byte) (uint32, bool) {

	v := ReadBytes(b, 4)
	if v == nil {
		return 0, false
	}

	return uint32(v[3])<<24 | uint32(v[2])<<16 | uint32(v[1])<<8 | uint32(v[0]), true
}

// ReadBigUint32 removes the first bytes from b and returns it as an uint32 in big-endian order.
// The bool indicates whether the read was successful.
func ReadBigUint32(b *[]byte) (uint32, bool) {

	v := ReadBytes(b, 4)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<24 | uint32(v[1])<<16 | uint32(v[2])<<8 | uint32(v[3]), true
}

// ReadUint32 removes the first bytes from b and returns it as an uint32 in native-endian order.
// The bool indicates whether the read was successful.
func ReadUint32(b *[]byte) (uint32, bool) {

	switch NativeEndian {
	case LittleEndian:
		return ReadLittleUint32(b)
	case BigEndian:
		return ReadBigUint32(b)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// ReadLittleInt32 removes the first bytes from b and returns it as an int32 in little-endian order.
// The bool indicates whether the read was successful.
func ReadLittleInt32(b *[]byte) (int32, bool) {

	v := ReadBytes(b, 4)
	if v == nil {
		return 0, false
	}

	return int32(v[3])<<24 | int32(v[2])<<16 | int32(v[1])<<8 | int32(v[0]), true
}

// ReadBigInt32 removes the first bytes from b and returns it as an int32 in big-endian order.
// The bool indicates whether the read was successful.
func ReadBigInt32(b *[]byte) (int32, bool) {

	v := ReadBytes(b, 4)
	if v == nil {
		return 0, false
	}

	return int32(v[0])<<24 | int32(v[1])<<16 | int32(v[2])<<8 | int32(v[3]), true
}

// ReadInt32 removes the first bytes from b and returns it as an int32 in native-endian order.
// The bool indicates whether the read was successful.
func ReadInt32(b *[]byte) (int32, bool) {

	switch NativeEndian {
	case LittleEndian:
		return ReadLittleInt32(b)
	case BigEndian:
		return ReadBigInt32(b)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// ReadUint64 removes the first bytes from b and returns it as an uint64 in little-endian order.
// The bool indicates whether the read was successful.
func ReadLittleUint64(b *[]byte) (uint64, bool) {

	v := ReadBytes(b, 8)
	if v == nil {
		return 0, false
	}

	return uint64(v[7])<<56 | uint64(v[6])<<48 | uint64(v[5])<<40 | uint64(v[4])<<32 | uint64(v[3])<<24 | uint64(v[2])<<16 | uint64(v[1])<<8 | uint64(v[0]), true
}

// ReadUint64 removes the first bytes from b and returns it as an uint64 in big-endian order.
// The bool indicates whether the read was successful.
func ReadBigUint64(b *[]byte) (uint64, bool) {

	v := ReadBytes(b, 8)
	if v == nil {
		return 0, false
	}

	return uint64(v[0])<<56 | uint64(v[1])<<48 | uint64(v[2])<<40 | uint64(v[3])<<32 | uint64(v[4])<<24 | uint64(v[5])<<16 | uint64(v[6])<<8 | uint64(v[7]), true
}

// ReadUint64 removes the first bytes from b and returns it as an uint64 in native-endian order.
// The bool indicates whether the read was successful.
func ReadUint64(b *[]byte) (uint64, bool) {

	switch NativeEndian {
	case LittleEndian:
		return ReadLittleUint64(b)
	case BigEndian:
		return ReadBigUint64(b)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}

// ReadInt64 removes the first bytes from b and returns it as an int64 in little-endian order.
// The bool indicates whether the read was successful.
func ReadLittleInt64(b *[]byte) (int64, bool) {

	v := ReadBytes(b, 8)
	if v == nil {
		return 0, false
	}

	return int64(v[7])<<56 | int64(v[6])<<48 | int64(v[5])<<40 | int64(v[4])<<32 | int64(v[3])<<24 | int64(v[2])<<16 | int64(v[1])<<8 | int64(v[0]), true
}

// ReadInt64 removes the first bytes from b and returns it as an int64 in big-endian order.
// The bool indicates whether the read was successful.
func ReadBigInt64(b *[]byte) (int64, bool) {

	v := ReadBytes(b, 8)
	if v == nil {
		return 0, false
	}

	return int64(v[0])<<56 | int64(v[1])<<48 | int64(v[2])<<40 | int64(v[3])<<32 | int64(v[4])<<24 | int64(v[5])<<16 | int64(v[6])<<8 | int64(v[7]), true
}

// ReadInt64 removes the first bytes from b and returns it as an int64 in native-endian order.
// The bool indicates whether the read was successful.
func ReadInt64(b *[]byte) (int64, bool) {

	switch NativeEndian {
	case LittleEndian:
		return ReadLittleInt64(b)
	case BigEndian:
		return ReadBigInt64(b)
	default:
		panic(fmt.Sprintf("Invalid NativeEndian: %d", NativeEndian))
	}
}
