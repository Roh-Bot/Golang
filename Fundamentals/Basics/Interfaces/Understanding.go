package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func InterFaceComposition() {
	err := hashAndBroadcast(myuf())
	if err != nil {
		return
	}
}

type composed struct {
	io.Reader
}

type customType struct {
	composed
}

func myuf() customType {
	return customType{}
}

func hashAndBroadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	h := sha1.Sum(b)
	fmt.Println(hex.EncodeToString(h[:]))
	return nil
}
