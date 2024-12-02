package circular

import (
	"container/list"
	"errors"
)

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.

type Buffer struct {
	size int
	ll   *list.List
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		size,
		list.New(),
	}
}

func (b *Buffer) pop() (byte, error) {
	e := b.ll.Back()
	if e == nil {
		return 0, errors.New("trying to read from empty buffer")
	}

	v := b.ll.Remove(e).(byte)

	return v, nil
}

func (b *Buffer) ReadByte() (byte, error) {
	return b.pop()
}

func (b *Buffer) WriteByte(c byte) error {
	if b.ll.Len() == b.size {
		return errors.New("trying to write to full buffer")
	}

	b.ll.PushFront(c)

	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.ll.Len() == b.size {
		b.pop()
	}

	b.WriteByte(c)
}

func (b *Buffer) Reset() {
	b.ll = list.New()
}
