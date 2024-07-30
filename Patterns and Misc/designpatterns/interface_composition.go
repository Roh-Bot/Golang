package designpatterns

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

type HashReaderInterface interface {
	io.Reader
	hash() string
}

func InterFaceComposition() {
	payload := []byte("Dhebug Bhai")
	err := hashAndBroadcast(NewHashReader(payload))
	if err != nil {
		return
	}
}

type HashReaderStruct struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *HashReaderStruct {
	return &HashReaderStruct{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *HashReaderStruct) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(r HashReaderInterface) error {
	//hash := r.(*HashReaderStruct).hash()
	hash := r.hash()
	fmt.Println(hash)
	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Printf("String: %s", b)
	return nil
}
