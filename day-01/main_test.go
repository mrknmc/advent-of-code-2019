package main

import (
	"testing"
)

func TestIterFuel1(t *testing.T) {
	result := IterFuel(654)
	if result != 966. {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", result, 966.)
	}
}

func TestIterFuel2(t *testing.T) {
	result := IterFuel(33583)
	if result != 50346. {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", result, 50346.)
	}
}

func TestIterFuel3(t *testing.T) {
	result := IterFuel(0)
	if result != 0. {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", result, 0.)
	}
}

func TestIterFuel4(t *testing.T) {
	result := IterFuel(6)
	if result != 6. {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", result, 6.)
	}
}

func TestIterFuel5(t *testing.T) {
	result := IterFuel(2)
	if result != 2. {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", result, 2.)
	}
}

func TestCalcFuel1(t *testing.T) {
	result := CalcFuel(12)
	if result != 2. {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", result, 2.)
	}
}

func TestCalcFuel2(t *testing.T) {
	result := CalcFuel(1969)
	if result != 654. {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", result, 654.)
	}
}
