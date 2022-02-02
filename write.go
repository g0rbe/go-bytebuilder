package bytebuilder

import (
	"crypto/rand"
	"time"
)

// WriteBytes appends bytes at the end of b.
func (b *Buffer) WriteBytes(bytes ...byte) {
	b.b = append(b.b, bytes...)
}

// WriteUint8 appends v at the end of b.
func (b *Buffer) WriteUint8(v uint8) {
	b.WriteBytes(byte(v))
}

// WriteUint16 appends v at the end of b.
func (b *Buffer) WriteUint16(v uint16) {
	b.WriteBytes(byte(v>>8), byte(v))
}

// WriteUint24 appends v at the end of b.
func (b *Buffer) WriteUint24(v uint32) {
	b.WriteBytes(byte(v>>16), byte(v>>8), byte(v))
}

// WriteUint32 appends v at the end of b.
func (b *Buffer) WriteUint32(v uint32) {
	b.WriteBytes(byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteGMTUnixTime32 appends time to b in 32 bit unix format.
func (b *Buffer) WriteGMTUnixTime32(time time.Time) {
	b.WriteUint32(uint32(time.Unix()))
}

// WriteRandom appends n random bytes to b.
// If failed to read random, this function panic.
func (b *Buffer) WriteRandom(n int) {
	v := make([]byte, n)

	_, err := rand.Read(v)
	if err != nil {
		panic("failed to read rand: " + err.Error())
	}
}
