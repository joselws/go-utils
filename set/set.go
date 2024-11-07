/* Basic implementation of a Set data structure */
package set

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
	"strings"
	"sync"

	"github.com/joselws/go-utils/mytypes"
)

/*
A Set can only be initialized with the following built-in types.
*/
type SetTypes interface {
	mytypes.Number | string
}

// The Set is a data structure that works with a map under the hood
type Set[T SetTypes] struct {
	Items map[T]bool
	mutex sync.Mutex
}

// Use this function whenever you want a new set
func NewSet[T SetTypes]() *Set[T] {
	return &Set[T]{Items: make(map[T]bool)}
}

// Returns the length of the set
func (thisSet *Set[T]) Length() int {
	return len(thisSet.Items)
}

// Add an new item to the set. Nothing happens if it already exists
func (thisSet *Set[T]) Add(item T) {
	thisSet.mutex.Lock()
	defer thisSet.mutex.Unlock()
	thisSet.Items[item] = true
}

// Takes in a slice of items and add them to the set
func (thisSet *Set[T]) AddFromSlice(items []T) {
	for _, item := range items {
		thisSet.Add(item)
	}
}

// Remove item from set. Does nothing if it doesn't exist
func (thisSet *Set[T]) Remove(item T) {
	delete(thisSet.Items, item)
}

/*
Checks whether an item is in the set.
Returns true if the item is in the set, and false otherwise
*/
func (thisSet *Set[T]) Contains(item T) bool {
	return thisSet.Items[item]
}

// Returns a slice of the items of the set sorted in ascending order
func (thisSet *Set[T]) ToSlice() []T {
	var setSlice []T
	for item := range thisSet.Items {
		setSlice = append(setSlice, item)
	}
	slices.Sort(setSlice)
	return setSlice
}

/*
Takes in another set of the same type and checks whether your set
(the one you're calling the method with) is a superset of the set
you passed as an argument

Returns true if it is a superset, and false otherwise
*/
func (thisSet *Set[T]) IsSupersetOf(otherSet *Set[T]) bool {
	for key := range otherSet.Items {
		if !thisSet.Contains(key) {
			return false
		}
	}
	return true
}

/*
Takes in another set of the same type and checks whether your set
(the one you're calling the method with) is a subset of the set
you passed as an argument

Returns true if it is a subset, and false otherwise
*/
func (thisSet *Set[T]) IsSubsetOf(otherSet *Set[T]) bool {
	for key := range thisSet.Items {
		if !otherSet.Contains(key) {
			return false
		}
	}
	return true
}

/*
Calculates the intersection of your set and another set you pass as arguments,
and returns these common values in a new set object
*/
func (thisSet *Set[T]) Intersection(otherSet *Set[T]) *Set[T] {
	intersection := NewSet[T]()
	for key := range otherSet.Items {
		if thisSet.Contains(key) {
			intersection.Add(key)
		}
	}
	return intersection
}

/*
Calculates the union of your set and another set you pass as arguments,
and returns these common values in a new set object
*/
func (thisSet *Set[T]) Union(otherSet *Set[T]) *Set[T] {
	union := NewSet[T]()
	for key := range otherSet.Items {
		union.Add(key)
	}
	for key := range thisSet.Items {
		union.Add(key)
	}
	return union
}

/*
Returns true if both sets are disjointed, otherwise returns false.
Two sets are disjointed if they have no elements in common.
*/
func (thisSet *Set[T]) IsDisjoint(otherSet *Set[T]) bool {
	for element := range thisSet.Items {
		if otherSet.Contains(element) {
			return false
		}
	}
	return true
}

// Returns a copy of the set
func (thisSet *Set[T]) Copy() *Set[T] {
	copy := NewSet[T]()
	maps.Copy(copy.Items, thisSet.Items)
	return copy
}

// Performs thisSet - otherSet
func (thisSet *Set[T]) Difference(otherSet *Set[T]) *Set[T] {
	result := thisSet.Copy()
	for element := range otherSet.Items {
		result.Remove(element)
	}
	return result
}

// Returns a set with the symmetric difference between both sets
func (thisSet *Set[T]) SymmetricDifference(otherSet *Set[T]) *Set[T] {
	intersection := thisSet.Intersection(otherSet)
	set1 := thisSet.Difference(intersection)
	set2 := otherSet.Difference(intersection)
	return set1.Union(set2)
}

// Returns true if both sets contain the same elements.
func (thisSet *Set[T]) Equal(otherSet *Set[T]) bool {
	return maps.Equal(thisSet.Items, otherSet.Items)
}

// Outputs set in this format when printed: Set[T]{item1, item2}
func (thisSet *Set[T]) String() string {
	var itemsStringSlice []string
	setSlice := thisSet.ToSlice()
	for _, item := range setSlice {
		itemsStringSlice = append(itemsStringSlice, fmt.Sprint(item))
	}

	setType := reflect.TypeOf(thisSet.Items).Key()
	var itemsString string = strings.Join(itemsStringSlice, ", ")
	return fmt.Sprintf("Set[%v]{%v}", setType, itemsString)
}
