package readbyte

import (
	"bytes"
	"testing"
)

func TestByteReadLeLengthOfContent(t *testing.T) {
	var (
		in = 3
		expectlength = 3
		expectContent = []byte("abc")
	)
	length, content := ByteRead(in)
	if length == expectlength && bytes.Equal(content, expectContent) {
		t.Log("success")
	} else {
		t.Logf("failed, expectlength = %d, but real is %d, expectContent is %s, but real is %s",
			expectlength, length, expectContent, content)
	}
}

func TestByteReadGtLengthOfContent(t *testing.T) {
	var (
		in = 10
		expectlength = 7
		expectContent = []byte{97, 98, 99, 49, 50, 51, 52, 0, 0, 0}
	)
	length, content := ByteRead(in)
	if length == expectlength && bytes.Equal(content, expectContent) {
		t.Log("success")
	} else {
		t.Logf("failed, expectlength = %d, but real is %d, expectContent is %v, but real is %v",
			expectlength, length, expectContent, content)
	}
}

func TestByteReadAtLeastLeLengthOfContent(t *testing.T) {
	var (
		in = 3
		expectlength = 3
		expectContent = []byte{97, 98, 99}
	)
	length, content := ByteReadAtLeast(in)
	if length == expectlength && bytes.Equal(content, expectContent) {
		t.Log("success")
	} else {
		t.Logf("failed, expectlength = %d, but real is %d, expectContent is %v, but real is %v",
			expectlength, length, expectContent, content)
	}
}

func TestByteReadAtLeastGtLengthOfContent(t *testing.T) {
	var (
		in = 10
		expectlength = 7
		expectContent =  []byte{97, 98, 99, 49, 50, 51, 52, 0, 0, 0}
	)
	length, content := ByteReadAtLeast(in)
	if length == expectlength && bytes.Equal(content, expectContent) {
		t.Log("success")
	} else {
		t.Logf("failed, expectlength = %d, but real is %d, expectContent is %v, but real is %v",
			expectlength, length, expectContent, content)
	}
}