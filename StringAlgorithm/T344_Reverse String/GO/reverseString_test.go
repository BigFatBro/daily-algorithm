package main

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	t.Run("reverse a byte array", func(t *testing.T) {
		s := []byte{'h', 'e', 'l', 'l', 'o'}
		want := []byte{'o', 'l', 'l', 'e', 'h'}
		reverseString(s)
		if !reflect.DeepEqual(s, want) {
			t.Errorf("got %v want %v", s, want)
		}
	})
}
