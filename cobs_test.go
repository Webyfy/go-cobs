package cobs

import (
	"bytes"
	"crypto/rand"
	"testing"
)

type dataEncodedPair struct {
	data    []byte
	encoded []byte
}

func TestCobsEncoding(t *testing.T) {
	dataEncodedPairs := []dataEncodedPair{
		{
			[]byte{0x00},
			[]byte{0x01, 0x01, 0x00},
		},
		{
			[]byte{0x00, 0x00},
			[]byte{0x01, 0x01, 0x01, 0x00},
		},
		{
			[]byte{0x00, 0x11, 0x00},
			[]byte{01, 0x02, 0x11, 0x01, 0x00},
		},
		{
			[]byte{0x11, 0x22, 0x00, 0x33},
			[]byte{0x03, 0x11, 0x22, 0x02, 0x33, 0x00},
		},
		{
			[]byte{0x11, 0x22, 0x33, 0x44},
			[]byte{0x05, 0x11, 0x22, 0x33, 0x44, 0x00},
		},
		{
			[]byte{0x11, 0x00, 0x00, 0x00},
			[]byte{0x02, 0x11, 0x01, 0x01, 0x01, 0x00},
		},
	}

	for _, item := range dataEncodedPairs {
		encoded := Encode(item.data)
		if !bytes.Equal(encoded, item.encoded) {
			t.Errorf("Encoding test failed for input: %v", item.data)
			t.Errorf("Expected: %v, got: %v", item.encoded, encoded)
		}
	}

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
		t.Errorf("Cobs Decoding test failed")
	}
}
