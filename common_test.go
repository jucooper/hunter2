package hunter2_test

import (
	"testing"

	"github.com/justinwcooper/hunter2"
)

func TestIsCommon(t *testing.T) {
	password := "hunter2"

	res := hunter2.New().IsCommon(password)

	if want := true; want != res {
		t.Errorf("IsCommon returned the wrong result: got %v, want %v", res, want)
	}
}

func TestIsNotCommon(t *testing.T) {
	password := "3smVBU348%P8qp"

	res := hunter2.New().IsNotCommon(password)

	if want := true; want != res {
		t.Errorf("IsNotCommon returned the wrong result: got %v, want %v", res, want)
	}
}
