package errors

import (
	"errors"
	"testing"
)

var errEx = errors.New("ex")

func TestError_Is(t *testing.T) {
	t.Parallel()

	e1 := WithReason("e1")
	e2 := WithReason("e2")
	e3 := WithReason("e3")
	err := e1.WithErr(e2.WithErr(e3.WithErr(errEx)))

	if !errors.Is(err, e1) {
		t.Errorf("err{%v} is not e1{%v}", err, e1)
	}

	if !errors.Is(err, e2) {
		t.Errorf("err{%v} is not e2{%v}", err, e2)
	}

	if !errors.Is(err, e3) {
		t.Errorf("err{%v} is not e3{%v}", err, e3)
	}

	if !errors.Is(err, errEx) {
		t.Errorf("err{%v} is not ex{%v}", err, errEx)
	}
}

func TestError_WithErr(t *testing.T) {
	t.Parallel()

	e1 := WithReason("e1")
	e2 := e1.WithErr(errEx)
	err := e1.WithErr(e2)

	if !errors.Is(err, e1) {
		t.Errorf("err{%v} is not e1{%v}", err, e1)
	}

	if !errors.Is(err, e2) {
		t.Errorf("err{%v} is not e2{%v}", err, e2)
	}

	if err.Error() != "e1 because: e1 because: ex" {
		t.Errorf("err{%v} is not {%v}", err, "e1 because: e1 because: ex")
	}

	e3 := e1.WithX(1)
	err = e1.WithErr(e3)

	if !errors.Is(err, e1) {
		t.Errorf("err{%v} is not e1{%v}", err, e1)
	}

	if !errors.Is(err, e2) {
		t.Errorf("err{%v} is not e2{%v}", err, e2)
	}

	if err.Error() != "e1 because: e1 with data: int{1}" {
		t.Errorf("err{%v} is not {%v}", err, "e1 because: e1 with data: int{1}")
	}
}

func TestIfErr(t *testing.T) {
	t.Parallel()

	if err := IfErr(nil); err != nil {
		t.Errorf("err{%v} is not nil", err)
	}

	e1 := WithReason("e1")
	err := IfErr(e1)

	if err == nil {
		t.Errorf("err{%v} is nil", err)
	}

	if !errors.Is(err, e1) {
		t.Errorf("err{%v} is not e1{%v}", err, e1)
	}
}

func TestError_IfErr(t *testing.T) {
	t.Parallel()

	e1 := WithReason("e1")

	if err := e1.IfErr(nil); err != nil {
		t.Errorf("err{%v} is not nil", err)
	}

	e2 := WithReason("e2")
	err := e1.IfErr(e2)

	if err == nil {
		t.Errorf("err{%v} is nil", err)
	}

	if !errors.Is(err, e1) {
		t.Errorf("err{%v} is not e1{%v}", err, e1)
	}

	if !errors.Is(err, e2) {
		t.Errorf("err{%v} is not e2{%v}", err, e2)
	}
}
