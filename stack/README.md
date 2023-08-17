# Stack

Stack data structure implementation for the Go programming language. Implement LIFO (last in, first out) solutions to your code with this easy to use package.

## Features

- Easy to use
- 100% test coverage
- All built-in data types available

## High level overview

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

**Push**

**Pop**

**ExtractMany**

**InsertFromSlice**

**InspectNextItem**
