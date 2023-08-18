# Stack

Stack data structure implementation for the Go programming language. Implement LIFO (last in, first out) solutions to your code with this easy to use package.

## Features

- Easy to use
- 100% test coverage
- All data types available

## High level overview

```Go
// Create new stack
numberStack := stack.Stack[int]{}

// Add items to the stack
numberStack.Push(500)

// Easy to understand string representation
fmt.Println(numberStack) // Stack[int][500]

// Get items from the stack
number, err := numberStack.Pop()
```

## Usage

### Import

After installation, add this `import` statement to your Go file:

```Go
import "github.com/joselws/go-utils/stack"
```

### Data types

All data types are available to use for `Stack`:

### Initialization

Create a new stack with the `Stack[T]{}` structure, where `T` is any data type. For example:

```Go
myStack := stack.Stack[int]{}
```

### String representation

`Stack` has an easy to interpret print format when shown using the `fmt` module. For example, given an `int` stack with numbers:

```Go
fmt.Println(numberStack)
// Stack[int][1 2 3 4 5]
```

### Operations

**Length**

The `Length() int` method returns the number of items in the stack as an `int`:

```Go
fmt.Println(numberStack) // Stack[int][1 2 3 4 5]
fmt.Println(numberStack.Length()) // 5
```

**Push**

`Push(item T)` Adds a new item on top of the stack.

```Go
fmt.Println(numberStack) // Stack[int][1 2 3 4 5]
numberStack.Push(6)
fmt.Println(numberStack) // Stack[int][1 2 3 4 5 6]
```

**Pop**

`Pop() (item T, err error)` Returns the item on top of the stack and possibly an error if the stack is empty.

```Go
// Valid operation
fmt.Println(numberStack) // Stack[int][1]
nextNumber, err := numberStack.Pop() // returns (1, nil)

// Error operation
fmt.Println(numberStack) // Stack[int][]
nextNumber, err := numberStack.Pop() // error is not nil, handle it!
```

**ExtractMany**

`ExtractMany(int) (items []T, err error)` Removes many items from the Stack. It takes an integer as an argument, which is the amount of items you want to take off the stack, and it returns a slice of that many items and an error. The error is not `nil` when you request for more items than the stack holds.

```Go
fmt.Println(numberStack) // Stack[int][1 2 3 4 5]
numberSlice, err := numberStack.ExtrackMany(3) // returns [3 4 5], nil

fmt.Println(numberStack) // Stack[int][1 2]
numberSlice, err := numberStack.ExtrackMany(3) // handle not nil error!
```

**InsertFromSlice**

`InsertFromSlice(items []T)` inserts many items on top of the stack, taking as an argument a slice of items.

```Go
fmt.Println(numberStack) // Stack[int][1 2 3 4 5]
numberSlice, err := numberStack.InsertFromSlice([]int{6, 7, 8}) 
fmt.Println(numberStack) // Stack[int][1 2 3 4 5 6 7 8]
```

**InspectNextItem**

`InspectNextItem() (item T, err error)` is similar to `Pop()`, in the sense that you get the next item from the stack and an error if the stack is empty. However, the stack size doesn't change when you use this method.

```Go
// Valid operation
fmt.Println(numberStack) // Stack[int][1]
nextNumber, err := numberStack.InspectNextItem() // returns (1, nil)
fmt.Println(numberStack) // Stack[int][1]

// Error operation
fmt.Println(numberStack) // Stack[int][]
nextNumber, err := numberStack.Pop() // error is not nil, handle it!
```
