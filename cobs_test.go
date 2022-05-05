package cobs

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestCobsEncoding(t *testing.T) {

}

func genRandomBytes(size int) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

func TestCobsDecoding(t *testing.T) {
	data, err := genRandomBytes(32)
	if err != nil {
		t.Error(err)
	}

	encodedBytes := Encode(data)
	decodedBytes := Decode(encodedBytes)
	if !bytes.Equal(data, decodedBytes) {
		t.Errorf("Cobs Decoding test with trailing 0 failed")
	}
}
