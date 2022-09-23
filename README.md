# File-Handling-Packages-in-Golang
The built-in library of Go offers excellent support to perform file-related operations, such as creating a file, read/write operations, file rename or move, copy, getting metadata information of the file, and so forth. 

The os package provides an API interface for file handling which is uniform across all operating systems. Although the design follows Unix standards, it provides a platform-independent interface to any operating system functionality. However, the error handling mechanism is unique to Go because, on failure, the return values are type errors rather than error numbers that we typically find in Unix systems. One particular advantage to that is that Go error types are loaded with information, which is particularly helpful in the debugging process. The os package provides functionality such as creation, deletion, opening a file, modification of permission, etc. The package therefore can be thought of as something that provides meta functionality related to the file handling process in Go.
