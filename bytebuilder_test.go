package bytebuilder

import (
	"io"
	"os"
	"testing"
)

func TestUint8(t *testing.T) {

	buf := NewEmpty()

	buf.WriteUint8(234)
	buf.WriteUint8(123)

	if num, ok := buf.ReadUint8(); ok != true {
		t.Fatalf("Failed to read the first uint8 in TestUint8")
	} else if num != 234 {
		t.Fatalf("The first num is not 234")
	}

	b := buf.BytesPointer()

	if num, ok := ReadUint8(b); ok != true {
		t.Fatalf("Failed to read second byte")
	} else if num != 123 {
		t.Fatalf("The second num is not 123")
	}

	if !buf.Empty() {
		t.Fatalf("buf is not empty")
	}

}

func TestUint16(t *testing.T) {

	buf := NewEmpty()

	buf.WriteUint16(65234)
	buf.WriteUint16(56123)

	if num, ok := buf.ReadUint16(); ok != true {
		t.Fatalf("Failed to read the first uint16 in TestUint16")
	} else if num != 65234 {
		t.Fatalf("The first num is not 65234")
	}

	b := buf.BytesPointer()

	if num, ok := ReadUint16(b); ok != true {
		t.Fatalf("Failed to read second byte")
	} else if num != 56123 {
		t.Fatalf("The second num is not 56123")
	}

	if !buf.Empty() {
		t.Fatalf("buf is not empty")
	}

}

func TestUint24(t *testing.T) {

	buf := NewEmpty()

	buf.WriteUint24(12365234)
	buf.WriteUint24(13256123)

	if num, ok := buf.ReadUint24(); ok != true {
		t.Fatalf("Failed to read the first uint16 in TestUint16")
	} else if num != 12365234 {
		t.Fatalf("The first num is not 12365234")
	}

	b := buf.BytesPointer()

	if num, ok := ReadUint24(b); ok != true {
		t.Fatalf("Failed to read second byte")
	} else if num != 13256123 {
		t.Fatalf("The second num is not 13256123")
	}

	if !buf.Empty() {
		t.Fatalf("buf is not empty")
	}

}

func TestUint32(t *testing.T) {

	buf := NewEmpty()

	buf.WriteUint32(4294967295)
	buf.WriteUint32(4294967294)

	if num1, ok := buf.ReadUint32(); ok != true {
		t.Fatalf("Failed to read the first num")
	} else if num1 != 4294967295 {
		t.Fatalf("The first num is not 4294967295")
	}

	b := buf.BytesPointer()

	if num2, ok := ReadUint32(b); ok != true {
		t.Fatalf("Failed to read second num")
	} else if num2 != 4294967294 {
		t.Fatalf("The second num is not 4294967294: %d", num2)
	}

	if !buf.Empty() {
		t.Fatalf("buf is not empty")
	}

}

func TestOthers(t *testing.T) {

	buf := NewBuffer([]byte{0x00})

	if !buf.Skip(1) {
		t.Fatalf("Failed to skip")
	}

	if buf.Size() != 0 {
		t.Fatalf("buf is not empty")
	}
}

func TestFileUint8(t *testing.T) {

	file, err := os.OpenFile("testfile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("Failed to create testfile: %s", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	if err = WriteWriterUint8(file, 123); err != nil {
		t.Fatalf("Filed to write uint8: %T", err)
	}

	if err = WriteWriterUint8(file, 213); err != nil {
		t.Fatalf("Filed to write uint8: %T", err)
	}

	file.Seek(0, 0)

	if num, err := ReadReaderUint8(file); err != nil && err != io.EOF {
		t.Fatalf("Failed to read first uint8: %s", err)
	} else if num != 123 {
		t.Fatalf("first num mismatch: 123 vs %x", num)
	}

	if num, err := ReadReaderUint8(file); err != nil && err != io.EOF {
		t.Fatalf("Failed to read second uint8: %s", err)
	} else if num != 213 {
		t.Fatalf("second num mismatch: 213 vs %x", num)
	}
}

func TestFileUint16(t *testing.T) {

	file, err := os.OpenFile("testfile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("Failed to create testfile: %s", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	if err = WriteWriterUint16(file, 65234); err != nil {
		t.Fatalf("Filed to write first uint16: %T", err)
	}

	if err = WriteWriterUint16(file, 56123); err != nil {
		t.Fatalf("Filed to write second uint16: %T", err)
	}

	file.Seek(0, 0)

	if num, err := ReadReaderUint16(file); err != nil && err != io.EOF {
		t.Fatalf("Failed to read first uint16: %s", err)
	} else if num != 65234 {
		t.Fatalf("first num mismatch: 65234 vs %d", num)
	}

	if num, err := ReadReaderUint16(file); err != nil && err != io.EOF {
		t.Fatalf("Failed to read second uint16: %s", err)
	} else if num != 56123 {
		t.Fatalf("second num mismatch: 56123 vs %d", num)
	}
}

func TestFileUint24(t *testing.T) {

	file, err := os.OpenFile("testfile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("Failed to create testfile: %s", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	if err = WriteWriterUint24(file, 12365234); err != nil {
		t.Fatalf("Filed to write first uint24: %T", err)
	}

	if err = WriteWriterUint24(file, 13256123); err != nil {
		t.Fatalf("Filed to write second uint24: %T", err)
	}

	file.Seek(0, 0)

	if num, err := ReadReaderUint24(file); err != nil && err != io.EOF {
		t.Fatalf("Failed to read first uint24: %s", err)
	} else if num != 12365234 {
		t.Fatalf("first num mismatch: 12365234 vs %d", num)
	}

	if num, err := ReadReaderUint24(file); err != nil && err != io.EOF {
		t.Fatalf("Failed to read second uint24: %s", err)
	} else if num != 13256123 {
		t.Fatalf("second num mismatch: 13256123 vs %d", num)
	}
}

func TestFileUint32(t *testing.T) {

	file, err := os.OpenFile("testfile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("Failed to create testfile: %s", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	if err = WriteWriterUint32(file, 4294967295); err != nil {
		t.Fatalf("Filed to write first uint32: %T", err)
	}

	if err = WriteWriterUint32(file, 4294967294); err != nil {
		t.Fatalf("Filed to write second uint32: %T", err)
	}

	file.Seek(0, 0)

	if num, err := ReadReaderUint32(file); err != nil && err != io.EOF {
		t.Fatalf("Failed to read first uint32: %s", err)
	} else if num != 4294967295 {
		t.Fatalf("first num mismatch: 4294967295 vs %d", num)
	}

	if num, err := ReadReaderUint32(file); err != nil && err != io.EOF {
		t.Fatalf("Failed to read second uint32: %s", err)
	} else if num != 4294967294 {
		t.Fatalf("second num mismatch: 4294967294 vs %d", num)
	}
}
