package incrementor

import (
	"math"
	"testing"
)

func TestNewIncrementor(t *testing.T) {
	incr1 := &Incrementor{Number: 10}
	incr2 := &Incrementor{}

	incr1Gen := NewIncrementor(10)
	incr2Gen := NewIncrementor(0)

	if incr1.Number != incr1Gen.Number {
		t.Errorf("Number was incorrect, got: %d, want: %d.", incr1.Number, incr1Gen.Number)
	}

	if incr2.Number != incr2Gen.Number {
		t.Errorf("Number was incorrect, got: %d, want: %d.", incr2.Number, incr2Gen.Number)
	}
}

func TestGetNumber(t *testing.T) {
	cases := []int32{1, 2, 3, 4, 10, -10, 200, -10000}

	for _, _case := range cases {
		incrementor := NewIncrementor(_case)
		number := incrementor.GetNumber()
		if _case != number{
			t.Errorf("Number was incorrect, got: %d, want: %d.", number, _case)
		}
	}
}

func TestIncrementNumber(t *testing.T) {
	cases := []struct{
		number int32
		increment int32
	}{
		{1, 2},
		{3, 4},
		{-5, -4},
		{math.MaxInt32, 0},
	}
	
	for _, _case := range cases {
		incrementor := NewIncrementor(_case.number)
		incrementor.IncrementNumber()
		if incrementor.GetNumber() != _case.increment {
			t.Errorf("Number was incorrect, got: %d, want: %d.", incrementor.GetNumber(), _case.increment)
		}
	}
}

func TestSetMaximumValue(t *testing.T) {
	cases := []struct{
		number int32
		maxValue int32
	}{
		{1, 1},
		{3, 3},
		{-5, -5},
		{math.MaxInt32, math.MaxInt32},
	}

	for _, _case := range cases {
		incrementor := NewIncrementor(_case.number)
		incrementor.SetMaximumValue(_case.number)
		if incrementor.GetNumber() != _case.maxValue {
			t.Errorf("Number was incorrect, got: %d, want: %d.", incrementor.GetNumber(), _case.maxValue)
		}
	}
	incrementor := NewIncrementor(0)
	incrementor.SetMaximumValue()
	if incrementor.GetNumber() != math.MaxInt32 {
		t.Errorf("Number was incorrect, got: %d, want: %d.", incrementor.GetNumber(), math.MaxInt32)
	}
}