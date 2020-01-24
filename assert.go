package assert

import (
	"errors"
	"testing"
)

func Equal(t *testing.T, want interface{}, got interface{}, msg string) {
	t.Helper()
	if wantErr, ok := want.(error); ok {
		gotErr, ok := got.(error)
		if !ok {
			t.Errorf("%s: want error %v, got non-error %v", msg, wantErr, got)
			return
		}
		if !errors.Is(gotErr, wantErr) {
			t.Errorf("%s: want error %v, got error %v", msg, wantErr, gotErr)
			return
		}
		return
	}
	if want != got {
		t.Errorf("%s: want %v, got %v", msg, want, got)
	}
}
