package minihttp

import (
	"reflect"
	"testing"
)

func TestTrimBytes(t *testing.T) {
	buf := []byte{1, 2, 3, 4, 0, 0, 0, 0}
	want := []byte{1, 2, 3, 4}
	got := trimBytes(buf)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v got %v", got, want)
	}
}
