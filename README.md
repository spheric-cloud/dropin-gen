# dropin-gen

`dropin-gen` is a command-line tool for Go that generates a "drop-in" replacement for a given package. It creates a new package that mirrors the public API of an existing package, with all functions and methods delegating to the original. This is useful for creating compatibility layers, mocks, or fakes for testing purposes.

## Installation

You can install `dropin-gen` using `go install`:

```sh
go install github.com/spheric-cloud/dropin-gen@latest
```

## Usage

You can invoke `dropin-gen` with the following flags:

- `-i`, `--import-path`: (Required) The import path of the package to generate a drop-in for.
- `-o`, `--output-package`: (Required) The path to the output package where the generated code will be written.
- `-f`, `--output-filename`: (Optional) The name of the generated Go file. Defaults to `zz_generated.dropin.go`.
- `--src-dir`: (Optional) The source directory to run the tool in. Defaults to the current directory.
- `--go-header-file`: (Optional) Go header file to include in the generated code.

### Example

Suppose you want to create a drop-in replacement for the `fmt` package and place it in a local package called `myfmt`. You would run the following command:

```sh
dropin-gen -i fmt -o ./myfmt
```

This will create a new file `myfmt/zz_generated.dropin.go` containing the generated drop-in code for the `fmt` package. The new package will have the name `myfmt`.
