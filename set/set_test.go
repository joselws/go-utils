package set

import (
	"fmt"
	"slices"
	"testing"
)

func TestSetLength(t *testing.T) {
	set := Set[rune]{Items: map[rune]bool{'a': true, 'b': true}}
	var setLength int = set.Length()
	if setLength != 2 {
		t.Error("Set length should be 2, not", setLength)
	}
}

func TestSetNoLength(t *testing.T) {
	set := Set[rune]{}
	var setLength int = set.Length()
	if setLength != 0 {
		t.Error("Set length should be 0, not", setLength)
	}
}

func TestAddNewItem(t *testing.T) {
	set := Set[rune]{Items: make(map[rune]bool)}
	set.Add('a')
	if set.Length() != 1 {
		t.Error("Error at AddNewItem", set)
	}
}

func TestAddExistingItem(t *testing.T) {
	set := Set[rune]{Items: map[rune]bool{'a': true}}
	set.Add('a')
	if set.Length() != 1 {
		t.Error("Set should have 1 item, got", set)
	}
}

func TestContainsExistingItem(t *testing.T) {
	set := Set[rune]{Items: map[rune]bool{'a': true}}
	if !set.Contains('a') {
		t.Error("Set should contain a", set)
	}
}

func TestNotContainsItem(t *testing.T) {
	set := Set[rune]{Items: map[rune]bool{'a': true}}
	if set.Contains('b') {
		t.Error("Set should not contain b", set)
	}
}

func TestIsSupersetOf(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set2 := Set[int]{Items: map[int]bool{1: true, 2: true}}
	if !set1.IsSupersetOf(set2) {
		t.Errorf("%v should be superset of %v", set1, set2)
	}
}

func TestIsSupersetOfEqualSet(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set2 := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	if !set1.IsSupersetOf(set2) {
		t.Errorf("%v should be superset of %v", set1, set2)
	}
}

func TestIsSupersetOfSubset(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{1: true, 2: true}}
	set2 := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	if set1.IsSupersetOf(set2) {
		t.Errorf("%v should be superset of %v", set1, set2)
	}
}

func TestIsSupersetOfEmpty(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{}}
	set2 := Set[int]{Items: map[int]bool{}}
	if !set1.IsSupersetOf(set2) {
		t.Errorf("%v should be superset of %v", set1, set2)
	}
}

func TestAddItems(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	set1 := NewSet[int]()
	set1.AddFromSlice(numbers)
	var length int = set1.Length()
	if length != 5 {
		t.Error("Length should be 5, not", length)
	}
}

func TestAddItemsEmpty(t *testing.T) {
	numbers := []int{}
	set1 := NewSet[int]()
	set1.AddFromSlice(numbers)
	var length int = set1.Length()
	if length != 0 {
		t.Error("Length should be 0, not", length)
	}
}

func TestIntersection(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set2 := Set[int]{Items: map[int]bool{2: true, 4: true}}
	var intersection Set[int] = set1.Intersection(set2)
	var intersectionLength int = intersection.Length()
	if intersectionLength != 1 {
		t.Error("Intersection should be 1, not", intersectionLength)
	}
}

func TestUnion(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set2 := Set[int]{Items: map[int]bool{2: true, 4: true}}
	union := set1.Union(set2)
	if union.Length() != 4 {
		t.Error("Union length should be 4, not", union.Length())
	}
	if set1.Length() != 3 {
		t.Error("Union should not change the calling set items")
	}
}

func TestRemove(t *testing.T) {
	set := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set.Remove(3)
	if set.Length() != 2 {
		t.Error("Set length should be 2, not", set.Length())
	}
	if set.Contains(3) {
		t.Error("Set should have item '3' removed")
	}
}

func TestRemoveNonExistingItem(t *testing.T) {
	set := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set.Remove(4)
	if set.Length() != 3 {
		t.Error("Set length should be 3, not", set.Length())
	}
	if set.Contains(4) {
		t.Error("Set should not have item '4'")
	}
}

func TestRemoveFromEmptySet(t *testing.T) {
	set := NewSet[int]()
	set.Remove(3)
	if set.Length() != 0 {
		t.Error("Set length should be 0, not", set.Length())
	}
}

func TestIsSubsetOf(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{1: true, 2: true}}
	set2 := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	if !set1.IsSubsetOf(set2) {
		t.Errorf("%v should be subset of %v", set1, set2)
	}
}

func TestIsSubsetOfEqualSet(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set2 := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	if !set1.IsSubsetOf(set2) {
		t.Errorf("%v should be subset of %v", set1, set2)
	}
}

func TestIsSubsetOfSuperset(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set2 := Set[int]{Items: map[int]bool{1: true, 2: true}}
	if set1.IsSubsetOf(set2) {
		t.Errorf("%v should be subset of %v", set2, set1)
	}
}

func TestIsSubsetOfEmpty(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{}}
	set2 := Set[int]{Items: map[int]bool{}}
	if !set1.IsSubsetOf(set2) {
		t.Errorf("%v should be subset of %v", set1, set2)
	}
}

func TestString(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	var setString string = fmt.Sprint(set1)
	if setString != "Set[int]{1, 2, 3}" {
		t.Error("Expected Set[int]{1, 2, 3} String() output, not", setString)
	}
}

func TestEmptyString(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{}}
	var setString string = fmt.Sprint(set1)
	if setString != "Set[int]{}" {
		t.Error("Expected Set[int]{} String() output, not", setString)
	}
}

func TestToSlice(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{3: true, 1: true, 2: true}}
	var setSlice []int = set1.ToSlice()
	if len(setSlice) != 3 {
		t.Error("Slice length should be 3, not", len(setSlice))
	}
	if !slices.IsSorted(setSlice) {
		t.Error("Slice should be sorted in ascending order, not", setSlice)
	}
}

func TestToSliceEmpty(t *testing.T) {
	set1 := Set[int]{Items: map[int]bool{}}
	var setSlice []int = set1.ToSlice()
	if len(setSlice) != 0 {
		t.Error("Slice length should be 0, not", len(setSlice))
	}
}
