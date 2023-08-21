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

These are the available data types for `Set`:

- All `int` and `uint` types
    - `int`, `int8`, `int16`, `int32`, and `int64`
    - `uint`, `uint8`, `uint16`, `uint32`, and `uint64`
- `float32` and `float64`
- `string`
- `rune`
- `byte`

For the time being, you cannot use `Set` with `maps`, `slices`, `structs`, `Set`, or pointer values.

### Initialization

Create new sets with the `NewSet[T]()` constructor, where `T` is any of the available data types. For example:

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

`ToSlice() []T` returns the items within the set in a slice of the same type `T` as the set in ascending order:

```Go
fmt.Println(nameSet) // Set[string]{"Jose", "Luis"}
nameSlice := nameSet.ToSlice()
fmt.Println(nameSlice) // ["Jose" "Luis"]
```

**IsSupersetOf**

The function `IsSupersetOf(other Set[T]) bool` function takes in another `Set` of the same type and checks whether the `Set` is a super set of this other `Set`. Returns `true` if it's a superset, and `false` otherwise:

```Go
fmt.Println(numberSet1) // Set[int]{1, 2, 3, 4, 5}
fmt.Println(numberSet2) // Set[int]{2, 4}
numberSet1.IsSupersetOf(numberSet2) // true
```

**IsSubsetOf**

The function `IsSubsetOf(other Set[T]) bool` function takes in another `Set` of the same type and checks whether the `Set` is a subset of this other `Set`. Returns `true` if it's a subset, and `false` otherwise:

```Go
fmt.Println(numberSet1) // Set[int]{2, 4}
fmt.Println(numberSet2) // Set[int]{1, 2, 3, 4, 5}
numberSet1.IsSubsetOf(numberSet2) // true
```

**Intersection**

`Intersection(other Set[T]) Set[T]` takes in another set and computes the intersection of both sets, returning a new `Set` with the intersected items:

```Go
fmt.Println(numberSet1) // Set[int]{2, 4, 6, 8}
fmt.Println(numberSet2) // Set[int]{1, 2, 3, 4, 5}
intersection := numberSet1.Intersection(numberSet2)
fmt.Println(intersection) // Set[int]{2, 4}
```

**Union**

`Union(other Set[T]) Set[T]` takes in another set and computes the Union of both sets, returning a new `Set` with the resulting items:

```Go
fmt.Println(numberSet1) // Set[int]{2, 4, 6, 8}
fmt.Println(numberSet2) // Set[int]{1, 2, 3, 4, 5}
union := numberSet1.Union(numberSet2)
fmt.Println(union) // Set[int]{1, 2, 3, 4, 5, 6, 8}
```
