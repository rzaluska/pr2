package compression

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
	"github.com/zaluska/pr2/savegame/compression/lzss"
)

// Reader reads compressed file and provides uncompressed stream
type Reader struct {
	r                io.Reader
	DecompressedSize int
}

func NewReader(r io.Reader) (*Reader, error) {

	// Read size of decompressed data saved as first 4 bytes in compressed file
	sizeBytes := make([]byte, 4)
	_, err := r.Read(sizeBytes)
	if err != nil {
		return nil, errors.Wrap(err, "Can't read save file header")
	}
	nr := new(Reader)
	nr.DecompressedSize = int(binary.LittleEndian.Uint32(sizeBytes))

	nr.r = lzss.NewReader(r)
	return nr, nil
}

func (r *Reader) Read(b []byte) (int, error) {
	return r.r.Read(b)
}
