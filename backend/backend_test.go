package backend

import (
	"testing"
)

func TestMakeMatrix(t *testing.T) {
	a := MakeMatrix(10)

	if len(a) != 10 {
		t.Fatal("a should be a 10 by 10 matrix")
	}

	for i := range a {
		for j := range a[i] {
			if len(a[i]) != 10 {
				t.Fatal("a should be a 10 by 10 matrix")
			}
			if a[i][j] != 0 {
				t.Fatal("a should only have zeros")
			}
		}
	}
}

func TestSameMatrix(t *testing.T) {
	a := MakeMatrix(10)
	b := MakeMatrix(10)

	if !SameMatrix(a, b) {
		t.Fatalf("a and b should be the same matrix")
	}

	a[0][1] = 1

	if SameMatrix(a, b) {
		t.Fatalf("a and b should not be the same matrix")
	}
}

func TestCopyMatrix(t *testing.T) {
	a := MakeMatrix(10)
	b := CopyMatrix(a)

	if !SameMatrix(a, b) {
		t.Fatalf("a and b should be the same matrix")
	}
}

func TestIsAlive(t *testing.T) {
	if isAlive(0) {
		t.Fatal("0 is not alive")
	}
	if !isAlive(1) {
		t.Fatal("1 is alive")
	}
}
