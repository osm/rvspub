package buffer

import (
	"errors"
)

var ErrBadRead = errors.New("bad read")

type Buffer struct {
	buf  []byte
	off  int
	size int
}

func New(buf []byte) *Buffer {
	return &Buffer{
		buf:  buf,
		size: len(buf),
	}
}

func (b *Buffer) Offset() int {
	return b.off
}

func (b *Buffer) Size() int {
	return b.size
}

func (b *Buffer) IsEnd() bool {
	return b.off >= b.size
}

func (b *Buffer) ReadBytes(n int) ([]byte, error) {
	if b.off+n > b.size {
		return nil, ErrBadRead
	}

	r := b.buf[b.off : b.off+n]
	b.off += n
	return r, nil
}

func (b *Buffer) SkipBytes(n int) error {
	if b.off+n > b.size {
		return ErrBadRead
	}

	b.off += n
	return nil
}
