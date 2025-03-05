# AppError: Structured Error Handling for Go

## Overview

The `AppError` package provides a structured and robust approach to error handling in Go applications. It captures stack traces, supports error chaining, and includes metadata such as HTTP status codes and additional details to enhance debugging and error reporting.

## Features

- **Structured Error Representation**: Includes error codes, messages, and HTTP status codes.
- **Automatic Stack Trace Capture**: Provides detailed debugging information.
- **Error Wrapping & Chaining**: Preserves the original error context.
- **JSON Serialization Support**: Enables logging and API-friendly error responses.
- **Go Standard Compatibility**: Implements `Unwrap()` for use with `errors.Is` and `errors.As`.

## Installation

To install the package, use:

```sh
go get github.com/meshashank-b/go-errors
```

## Usage

### Creating a New AppError

Create an error with a custom error code and message:

```go
err := apperror.NewAppError("NOT_FOUND", "Resource not found", 404, nil)
fmt.Println(err.Error())
```

### Retrieving the Stack Trace

Obtain a detailed stack trace for debugging:

```go
fmt.Println(err.StackTrace().String())
```

## Example Output

```
[DB_ERROR] Database error occurred: database connection failed
Stack Trace:
---------------------------------------------------
#1 - Function: main.someFunction
     Location: /path/to/file.go:42
     Pointer: 0x123abc
     Function Pointer: 0x456def
---------------------------------------------------
```

## License

This package is licensed under the MIT License.
