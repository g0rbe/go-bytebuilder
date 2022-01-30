package bytebuilder

import "io"

type Buffer struct {
	b []byte
}

func NewBuffer(bytes []byte) Buffer {
	return Buffer{b: bytes}
}

// NewEmpty creates a Buffer with a zero length byte slice.
func NewEmpty() Buffer {
	return Buffer{b: make([]byte, 0)}
}

// ReadAll reads from r until an error or EOF and returns a Buffer from the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func ReadAll(r io.Reader) (Buffer, error) {

	buf := Buffer{b: make([]byte, 0, 512)}

	for {
		if len(buf.b) == cap(buf.b) {
			buf.b = append(buf.b, 0)[:len(buf.b)]
		}

		n, err := r.Read(buf.b[len(buf.b):cap(buf.b)])

		buf.b = buf.b[:len(buf.b)+n]

		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return buf, err
		}
	}

}

// Empty returns whether b is empty.
func (b *Buffer) Empty() bool {
	return len(b.b) == 0
}

// Size returns the size of the underlying byte slice
func (b *Buffer) Size() int {
	return len(b.b)
}

// Bytes returns the underlying byte slice.
func (b *Buffer) Bytes() []byte {
	return b.b
}

// BytesPointer returns a pointer to the underlying byte slice.
func (b *Buffer) BytesPointer() *[]byte {
	return &b.b
}
