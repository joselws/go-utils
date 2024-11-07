package set

import (
	"fmt"
	"slices"
	"sync"
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
	set := NewSet[rune]()
	set.Add('a')
	if set.Length() != 1 {
		t.Error("Error at AddNewItem", set)
	}
}

func TestAddExistingItem(t *testing.T) {
	set := NewSet[rune]()
	set.Add('a')
	set.Add('a')
	if set.Length() != 1 {
		t.Error("Set should have 1 item, got", set)
	}
}

func TestContainsExistingItem(t *testing.T) {
	set := NewSet[rune]()
	set.Add('a')
	if !set.Contains('a') {
		t.Error("Set should contain a", set)
	}
}

func TestNotContainsItem(t *testing.T) {
	set := NewSet[rune]()
	set.Add('a')
	if set.Contains('b') {
		t.Error("Set should not contain b", set)
	}
}

func TestIsSupersetOf(t *testing.T) {
	set1 := NewSet[int]()
	set1.AddFromSlice([]int{1, 2, 3})
	set2 := NewSet[int]()
	set2.AddFromSlice([]int{1, 2})
	if !set1.IsSupersetOf(set2) {
		t.Errorf("%v should be superset of %v", set1, set2)
	}
}

func TestIsSupersetOfEqualSet(t *testing.T) {
	set1 := NewSet[int]()
	set1.AddFromSlice([]int{1, 2, 3})
	set2 := NewSet[int]()
	set2.AddFromSlice([]int{1, 2, 3})
	if !set1.IsSupersetOf(set2) {
		t.Errorf("%v should be superset of %v", set1, set2)
	}
}

func TestIsSupersetOfSubset(t *testing.T) {
	set1 := NewSet[int]()
	set1.AddFromSlice([]int{1, 2})
	set2 := NewSet[int]()
	set2.AddFromSlice([]int{1, 2, 3})
	if set1.IsSupersetOf(set2) {
		t.Errorf("%v should be superset of %v", set1, set2)
	}
}

func TestIsSupersetOfEmpty(t *testing.T) {
	set1 := NewSet[int]()
	set2 := NewSet[int]()
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
	set1 := &Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set2 := &Set[int]{Items: map[int]bool{2: true, 4: true}}
	var intersection *Set[int] = set1.Intersection(set2)
	var intersectionLength int = intersection.Length()
	if intersectionLength != 1 {
		t.Error("Intersection should be 1, not", intersectionLength)
	}
}

func TestUnion(t *testing.T) {
	set1 := &Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set2 := &Set[int]{Items: map[int]bool{2: true, 4: true}}
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
	set1 := &Set[int]{Items: map[int]bool{1: true, 2: true}}
	set2 := &Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	if !set1.IsSubsetOf(set2) {
		t.Errorf("%v should be subset of %v", set1, set2)
	}
}

func TestIsSubsetOfEqualSet(t *testing.T) {
	set1 := &Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set2 := &Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	if !set1.IsSubsetOf(set2) {
		t.Errorf("%v should be subset of %v", set1, set2)
	}
}

func TestIsSubsetOfSuperset(t *testing.T) {
	set1 := &Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	set2 := &Set[int]{Items: map[int]bool{1: true, 2: true}}
	if set1.IsSubsetOf(set2) {
		t.Errorf("%v should be subset of %v", set2, set1)
	}
}

func TestIsSubsetOfEmpty(t *testing.T) {
	set1 := &Set[int]{Items: map[int]bool{}}
	set2 := &Set[int]{Items: map[int]bool{}}
	if !set1.IsSubsetOf(set2) {
		t.Errorf("%v should be subset of %v", set1, set2)
	}
}

func TestString(t *testing.T) {
	set1 := &Set[int]{Items: map[int]bool{1: true, 2: true, 3: true}}
	var setString string = fmt.Sprint(set1)
	if setString != "Set[int]{1, 2, 3}" {
		t.Error("Expected Set[int]{1, 2, 3} String() output, not", setString)
	}
}

func TestEmptyString(t *testing.T) {
	set1 := &Set[int]{Items: map[int]bool{}}
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

func TestIsDisjoint(t *testing.T) {
	set1 := NewSet[string]()
	set1.AddFromSlice([]string{"apple", "pear", "orange"})
	set2 := NewSet[string]()
	set2.AddFromSlice([]string{"coconut", "avocado", "grape"})
	if !set1.IsDisjoint(set2) {
		t.Error("Sets should be disjointed")
	}
	if !set2.IsDisjoint(set1) {
		t.Error("Sets should be disjointed")
	}
}

func TestIsDisjointEmptySet(t *testing.T) {
	set1 := NewSet[string]()
	set2 := NewSet[string]()
	set2.AddFromSlice([]string{"coconut", "avocado", "grape"})
	if !set1.IsDisjoint(set2) {
		t.Error("Sets should be disjointed")
	}
	if !set2.IsDisjoint(set1) {
		t.Error("Sets should be disjointed")
	}
}

func TestCopy(t *testing.T) {
	set1 := NewSet[int]()
	set1.AddFromSlice([]int{1, 2, 3})
	set2 := set1.Copy()
	if set1 == set2 {
		t.Error("Sets should be different")
	}
	if !set1.Equal(set2) {
		t.Error("Sets should have the same elements")
	}
}

func TestEqual(t *testing.T) {
	set1 := NewSet[int]()
	set1.AddFromSlice([]int{1, 2, 3})
	set2 := NewSet[int]()
	set2.AddFromSlice([]int{1, 2, 3})
	set3 := NewSet[int]()
	set3.AddFromSlice([]int{1, 2, 4})

	if !set1.Equal(set2) {
		t.Error("set1 and set2 should be equal")
	}
	if set1.Equal(set3) {
		t.Error("set1 and set3 should be different")
	}
}

func TestDifference(t *testing.T) {
	set1 := NewSet[int]()
	set1.AddFromSlice([]int{1, 2, 3})
	set2 := NewSet[int]()
	set2.AddFromSlice([]int{1, 3})

	result := NewSet[int]()
	result.Add(2)

	if !result.Equal(set1.Difference(set2)) {
		t.Error("Difference of both sets should be a set with only the element 2")
	}
}

func TestDifferenceEmptySets(t *testing.T) {
	set := NewSet[int]()
	set.AddFromSlice([]int{1, 2, 3})
	empty := NewSet[int]()

	if !set.Equal(set.Difference(empty)) {
		t.Error("Empty sets shouldn't take away elements from other sets")
	}
	if !empty.Equal(empty.Difference(set)) {
		t.Error("Empty sets shouldn't take away elements from other sets")
	}
}

func TestSymmetricDifference(t *testing.T) {
	set1 := NewSet[int]()
	set1.AddFromSlice([]int{1, 2, 3})
	set2 := NewSet[int]()
	set2.AddFromSlice([]int{3, 4, 5})

	expected := NewSet[int]()
	expected.AddFromSlice([]int{1, 2, 4, 5})

	result := set1.SymmetricDifference(set2)
	if !result.Equal(expected) {
		t.Error("Symmetric difference should be 1, 2, 4, 5, not", result)
	}
	if !set1.SymmetricDifference(set2).Equal(set2.SymmetricDifference(set1)) {
		t.Error("Symmetric difference should yield the same result both ways")
	}
}

func TestConcurrencyAdd(t *testing.T) {
	t.Run("Check overwrites on Add", func(t *testing.T) {
		numbers := make([]int, 0, 1000)
		for i := 0; i < 100; i++ {
			s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			numbers = slices.Insert(numbers, 0, s...)
		}

		set := NewSet[int]()
		var wg sync.WaitGroup
		wg.Add(1000)
		for _, number := range numbers {
			go func(n int) {
				set.Add(n)
				wg.Done()
			}(number)
		}
		wg.Wait()

		if set.Length() != 10 {
			t.Error("set.Add() is not concurrent safe.")
		}
	})

	t.Run("Check concurrent writes for different numbers on Add", func(t *testing.T) {
		numbers := make([]int, 0, 1000)
		for i := 0; i < 1000; i++ {
			numbers = append(numbers, i)
		}

		set := NewSet[int]()
		var wg sync.WaitGroup
		wg.Add(1000)
		for _, number := range numbers {
			go func(n int) {
				set.Add(n)
				wg.Done()
			}(number)
		}
		wg.Wait()

		if set.Length() != 1000 {
			t.Error("set.Add() is not concurrent safe.")
		}
	})
}

func TestConcurrencyRemove(t *testing.T) {
	t.Run("Check removing repeated elements", func(t *testing.T) {
		numbers := make([]int, 0, 1000)
		for i := 0; i < 100; i++ {
			s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			numbers = slices.Insert(numbers, 0, s...)
		}

		set := NewSet[int]()
		set.AddFromSlice(numbers)
		var wg sync.WaitGroup
		wg.Add(1000)
		for _, number := range numbers {
			go func(n int) {
				set.Remove(n)
				wg.Done()
			}(number)
		}
		wg.Wait()

		if set.Length() != 0 {
			t.Error("set.Remove() is not concurrent safe.")
		}
	})

	t.Run("Check removing different elements", func(t *testing.T) {
		numbers := make([]int, 0, 1000)
		for i := 0; i < 1000; i++ {
			numbers = append(numbers, i)
		}

		set := NewSet[int]()
		set.AddFromSlice(numbers)
		var wg sync.WaitGroup
		wg.Add(1000)
		for _, number := range numbers {
			go func(n int) {
				set.Remove(n)
				wg.Done()
			}(number)
		}
		wg.Wait()

		if set.Length() != 0 {
			t.Error("set.Remove() is not concurrent safe.")
		}
	})
}

// Only applicable with the -race flag to check for race conditions
func TestSetConcurrencyAddRemove(t *testing.T) {
	numbers := make([]int, 0, 1000)
	for i := 0; i < 100; i++ {
		s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		numbers = slices.Insert(numbers, 0, s...)
	}

	set := NewSet[int]()
	var wg sync.WaitGroup
	wg.Add(2000)
	for _, number := range numbers {
		go func(n int) {
			set.Add(n)
			wg.Done()
		}(number)
	}
	for _, number := range numbers {
		go func(n int) {
			set.Remove(n)
			wg.Done()
		}(number)
	}
	wg.Wait()
	if !(set.Length() >= 0) {
		t.Error("Set is not concurrent safe.")
	}
}
