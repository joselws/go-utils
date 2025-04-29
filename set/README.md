# Set

Set data structure implementation for the Go programming language. Implement Set Theory solutions to your code with this easy to use package.

## Features

- Easy to use
- 100% test coverage
- Almost all built-in data types available

## High level overview

- Add/remove items from the Set
- Checks existence of elements
- Operations with other sets like union and intersection
- Easy to interpret string representation when printed with the `fmt` library

## Usage

### Import

After installation, add this `import` statement to your Go file:

```Go
import "github.com/joselws/go-utils/set"
```

### Data types

All `comparable` data types are available for `Set`.

### Initialization

Create new sets with the `NewSet[T]()` constructor, where `T` is any of the available data types. It returns a `*Set[T]` type. For example:

```Go
mySet := set.NewSet[int]()
```

### String representation

A Set has an easy to interpret print format when shown using the `fmt` module. For example, given a `string` set with names:

```Go
fmt.Println(nameSet)
// Set[string]{"Jose", "Luis"}
```

### Operations

**Add**

Add individual items to the set with the `Add(item T)` method. If the item already exists in the set, it does nothing. This function returns nothing. For example:

```Go
fmt.Println(nameSet) // Set[string]{"Jose"}
nameSet.Add("Luis") // Luis is successfully added into the set
nameSet.Add("Jose") // Nothing happens because "Jose" already exists!
```

**AddFromSlice**

Add multiple items to the set with the `AddFromSlice(items []T)` method. This function returns nothing. The same rules applies as with the `Add(item T)` method: if some of the items already exist within the set, nothing happens.

```Go
nameSet := set.NewSet[string]()
nameSlice := []string{"Jose", "Luis"}
nameSet.AddFromSlice(nameSlice) // Jose and Luis are added into the set
```

**Remove**

Remove an individual item from the set with `Remove(item T)`. If the item doesn't exist, nothing happens. This function doesn't return anything.

```Go
fmt.Println(nameSet) // Set[string]{"Jose"}
nameSet.Remove("Jose") 
fmt.Println(nameSet) // Set[string]{}
```

**Length**

`Length() int` returns the length of the set as an `int`:

```Go
fmt.Println(numberSet) // Set[int]{1, 2, 3, 4, 5}
numberSet.Length() // 5
```

**Contains**

`Contains(item T)` checks if an item exists withing the slice. Returns `true` if `item` exists, else returns `false`:

```Go
fmt.Println(nameSet) // Set[string]{"Jose"}
nameSet.Contains("Jose") // true
nameSet.Contains("Luis") // false
```

**ToSlice**

`ToSlice() []T` returns the items within the set in a slice of the same type `T` as the set:

```Go
fmt.Println(nameSet) // Set[string]{"Jose", "Luis"}
nameSlice := nameSet.ToSlice()
fmt.Println(nameSlice) // ["Jose" "Luis"]
```

**IsSupersetOf**

The function `IsSupersetOf(other *Set[T]) bool` function takes in another `Set` of the same type and checks whether the `Set` is a super set of this other `Set`. Returns `true` if it's a superset, and `false` otherwise:

```Go
fmt.Println(numberSet1) // Set[int]{1, 2, 3, 4, 5}
fmt.Println(numberSet2) // Set[int]{2, 4}
numberSet1.IsSupersetOf(numberSet2) // true
```

**IsSubsetOf**

The function `IsSubsetOf(other *Set[T]) bool` function takes in another `Set` of the same type and checks whether the `Set` is a subset of this other `Set`. Returns `true` if it's a subset, and `false` otherwise:

```Go
fmt.Println(numberSet1) // Set[int]{2, 4}
fmt.Println(numberSet2) // Set[int]{1, 2, 3, 4, 5}
numberSet1.IsSubsetOf(numberSet2) // true
```

**Intersection**

`Intersection(other *Set[T]) *Set[T]` takes in another set and computes the intersection of both sets, returning a new pointer to a `Set` with the intersected items:

```Go
fmt.Println(numberSet1) // Set[int]{2, 4, 6, 8}
fmt.Println(numberSet2) // Set[int]{1, 2, 3, 4, 5}
intersection := numberSet1.Intersection(numberSet2)
fmt.Println(intersection) // Set[int]{2, 4}
```

**Union**

`Union(other *Set[T]) *Set[T]` takes in another set and computes the Union of both sets, returning a new pointer to a `Set` with the resulting items:

```Go
fmt.Println(numberSet1) // Set[int]{2, 4, 6, 8}
fmt.Println(numberSet2) // Set[int]{1, 2, 3, 4, 5}
union := numberSet1.Union(numberSet2)
fmt.Println(union) // Set[int]{1, 2, 3, 4, 5, 6, 8}
```

**IsDisjoint**

`IsDisjoint(other *Set[T]) bool` takes in another set and returns `true` only if both sets have no elements in common:

```Go
fmt.Println(numberSet1) // Set[int]{1, 3, 5}
fmt.Println(numberSet2) // Set[int]{2, 4, 6}
fmt.Println(numberSet1.IsDisjoint(numberSet2)) // true
```

**Copy**

`Copy() *Set[T]` returns an copy of the set:

```Go
fmt.Println(numberSet1) // Set[int]{1, 3, 5}
numberSet2 := numberSet1.Copy()
fmt.Println(numberSet2) // Set[int]{1, 3, 5}
```

**Difference**

`Difference(other *Set[T]) *Set[T]` takes in another set and returns the difference of the original set substracted by the other set:

```Go
fmt.Println(numberSet1) // Set[int]{1, 2, 3}
fmt.Println(numberSet2) // Set[int]{3}
difference := numberSet1.Difference(numberSet2)
fmt.Println(difference) // Set[int]{1, 2}
```

**Equal**

`Equal(other *Set[T]) bool` returns true if both sets contain the exact same elements. It's different from the `==` operator because it doesn't check whether both sets refer to the same location in memory:

```Go
fmt.Println(numberSet1) // Set[int]{1, 2, 3}
fmt.Println(numberSet2) // Set[int]{1, 2, 3}
fmt.Println(numberSet1.Equal(numberSet2)) // true
fmt.Println(numberSet1 == numberSet2) // false
```

**SymmetricDifference**

`SymmetricDifference(other *Set[T]) *Set[T]` performs the symmetric difference operation between both sets:

```Go
fmt.Println(numberSet1) // Set[int]{1, 2}
fmt.Println(numberSet2) // Set[int]{2, 3}
fmt.Println(numberSet1.SymmetricDifference(numberSet2)) // Set[int]{1, 3}
```
