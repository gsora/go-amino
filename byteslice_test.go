package amino

import (
	"bytes"
	"testing"
)

func TestReadByteSliceEquality(t *testing.T) {

	var encoded []byte
	var err error
	var cdc = NewCodec()
	type byteWrapper struct {
		Val []byte
	}
	// Write a byteslice
	var testBytes = byteWrapper{[]byte("ThisIsSomeTestArrayEmbeddedInAStruct")}
	encoded, err = cdc.MarshalBinaryLengthPrefixed(testBytes)
	if err != nil {
		t.Error(err.Error())
	}

	// Read the byteslice, should return the same byteslice
	var testBytes2 byteWrapper
	err = cdc.UnmarshalBinaryLengthPrefixed(encoded, &testBytes2)
	if err != nil {
		t.Error(err.Error())
	}

	if !bytes.Equal(testBytes.Val, testBytes2.Val) {
		t.Error("Returned the wrong bytes")
	}

}
