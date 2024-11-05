# Internal Packages in Go

Go's package system makes it easy to structure programs into components with clean boundaries, but there are only two forms of access: local (unexported) and global (exported). Sometimes one wishes to have components that are not exported, for instance to avoid acquiring clients of interfaces to code that is part of a public repository but not intended for use outside the program to which it belongs.

The Go language does not have the power to enforce this distinction, but as of Go 1.4 the `go` command introduces a mechanism to define "internal" packages that may not be imported by packages outside the source subtree in which they reside.

## Creating Internal Packages

To create such a package, place it in a directory named `internal` or in a subdirectory of a directory named `internal`. When the `go` command sees an import of a package with `internal` in its path, it verifies that the package doing the import is within the tree rooted at the parent of the `internal` directory. For example, a package `.../a/b/c/internal/d/e/f` can be imported only by code in the directory tree rooted at `.../a/b/c`. It cannot be imported by code in `.../a/b/g` or in any other repository.

## Enforcement

For Go 1.4, the internal package mechanism is enforced for the main Go repository; from 1.5 and onward it will be enforced for any repository.

## Additional Notes on Current Project

Currently, this project does not have new internal packages beyond the standard structure, but users are encouraged to utilize the internal package mechanism for separating non-exported code modules for better encapsulation and management of code dependencies.

## More Information

Full details of the mechanism are in the [design document](https://golang.org/s/go14internal).
