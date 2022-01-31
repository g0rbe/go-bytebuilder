package bytebuilder

import (
	"io"
	"math/rand"
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
func (b *Buffer) WriteRandom(n int) error {
	v := make([]byte, n)

	_, err := rand.Read(v)

	return err
}

// WriteBytes appends bytes at the end of in.
func WriteBytes(in *[]byte, bytes ...byte) {
	*in = append(*in, bytes...)
}

// WriteUint8 appends v at the end of in.
func WriteUint8(in *[]byte, v uint8) {
	WriteBytes(in, byte(v))
}

// WriteUint16 appends v at the end of in.
func WriteUint16(in *[]byte, v uint16) {
	WriteBytes(in, byte(v>>8), byte(v))
}

// WriteUint24 appends v at the end of in.
func WriteUint24(in *[]byte, v uint32) {
	WriteBytes(in, byte(v>>16), byte(v>>8), byte(v))
}

// WriteUint32 appends v at the end of in.
func WriteUint32(in *[]byte, v uint32) {
	WriteBytes(in, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteWriterBytes appends v at the end of in.
func WriteWriterBytes(in io.Writer, bytes ...byte) error {

	_, err := in.Write(bytes)

	return err
}

// WriteWriterUint8 appends v at the end of in.
func WriteWriterUint8(in io.Writer, v uint8) error {
	return WriteWriterBytes(in, byte(v))
}

// WriteWriterUint16 appends v at the end of in.
func WriteWriterUint16(in io.Writer, v uint16) error {
	return WriteWriterBytes(in, byte(v>>8), byte(v))
}

// WriteWriterUint24 appends v at the end of in.
func WriteWriterUint24(in io.Writer, v uint32) error {
	return WriteWriterBytes(in, byte(v>>16), byte(v>>8), byte(v))
}

// WriteWriterUint32 appends v at the end of in.
func WriteWriterUint32(in io.Writer, v uint32) error {
	return WriteWriterBytes(in, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}
