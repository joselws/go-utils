package mathutils

import (
	"math"
	"testing"
)

func TestMaxManyArgs(t *testing.T) {
	maximum := Max(3, 5, -2.5, 0)
	if maximum != 5 {
		t.Error("Maximum should be 5, not", maximum)
	}
}

func TestMaxInf(t *testing.T) {
	maximum := Max(3, 5, -2.5, math.Inf(1), 0)
	if !math.IsInf(maximum, 1) {
		t.Error("Maximum should be +Inf, not", maximum)
	}
}

func TestMaxNegativeInf(t *testing.T) {
	maximum := Max(3, 5, -2.5, math.Inf(-1), 0)
	if maximum != 5 {
		t.Error("Maximum should be 5, not", maximum)
	}
}

func TestMaxInfOnlyArg(t *testing.T) {
	maximum := Max(math.Inf(1))
	if !math.IsInf(maximum, 1) {
		t.Error("Maximum should be inf, not", maximum)
	}
}

func TestMaxNaN(t *testing.T) {
	nan := math.NaN()
	maximum := Max(3, 5, -2.5, nan, 0)
	if !math.IsNaN(maximum) {
		t.Errorf("Maximum should be %v, not %v", nan, maximum)
	}
}

func TestMaxNaNOnlyArg(t *testing.T) {
	nan := math.NaN()
	maximum := Max(nan)
	if !math.IsNaN(maximum) {
		t.Errorf("Maximum should be %v, not %v", nan, maximum)
	}
}

func TestMaxWithBothInfAndNaN(t *testing.T) {
	nan := math.NaN()
	maximum := Max(nan, 3.3, math.Inf(1), math.Inf(-1), 0, 9, -5)
	if !math.IsInf(maximum, 1) {
		t.Errorf("Maximum should be +Inf, not %v", maximum)
	}
}

func TestMaxSingleArg(t *testing.T) {
	maximum := Max(3)
	if maximum != 3 {
		t.Error("Maximum should be 3, not", maximum)
	}
}

func TestMaxEqualArgs(t *testing.T) {
	maximum := Max(3, 3, 3)
	if maximum != 3 {
		t.Error("Maximum should be 3, not", maximum)
	}
}

func TestMinManyArgs(t *testing.T) {
	minimum := Min(3, 5, -2.5, 0)
	if minimum != -2.5 {
		t.Error("Minimum should be -2.5, not", minimum)
	}
}

func TestMinInf(t *testing.T) {
	minimum := Min(3, 5, -2.5, math.Inf(-1), 0)
	if minimum != math.Inf(-1) {
		t.Error("Minimum should be -Inf, not", minimum)
	}
}

func TestMinPositiveInf(t *testing.T) {
	minimum := Min(3, 5, -2.5, math.Inf(1), 0)
	if minimum != -2.5 {
		t.Error("Minimum should be -2.5, not", minimum)
	}
}

func TestMinInfOnlyArg(t *testing.T) {
	minimum := Min(math.Inf(-1))
	if minimum != math.Inf(-1) {
		t.Error("Minimum should be -inf, not", minimum)
	}
}

func TestMinNaN(t *testing.T) {
	nan := math.NaN()
	minimum := Min(3, 5, -2.5, nan, 0)
	if !math.IsNaN(minimum) {
		t.Errorf("Minimum should be %v, not %v", nan, minimum)
	}
}

func TestMinNaNOnlyArg(t *testing.T) {
	nan := math.NaN()
	minimum := Min(nan)
	if !math.IsNaN(minimum) {
		t.Errorf("Minimum should be %v, not %v", nan, minimum)
	}
}

func TestMinWithBothInfAndNaN(t *testing.T) {
	nan := math.NaN()
	minimum := Min(nan, 3.3, math.Inf(1), math.Inf(-1), 0, 9, -5)
	if !math.IsInf(minimum, -1) {
		t.Errorf("Minimum should be +Inf, not %v", minimum)
	}
}

func TestMinSingleArg(t *testing.T) {
	minimum := Min(3)
	if minimum != 3 {
		t.Error("Minimum should be 3, not", minimum)
	}
}

func TestMinEqualArgs(t *testing.T) {
	minimum := Min(3, 3, 3)
	if minimum != 3 {
		t.Error("Minimum should be 3, not", minimum)
	}
}
