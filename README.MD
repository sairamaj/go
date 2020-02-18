* Package => Collection of Source Files in the same directory
* Modules => Collection of Packages
* Repository => Collection of Modules (typically one module)
    * go.mod contains module definition

[Creating Hello World](https://golang.org/doc/code.html)

    * go mod init example.com/user/hello
    * go install example.com/user/hello
    * go run hello.go  (can be used to run immediately)
    * go env

### Allocating memory
* __new__
    * new(T)
    * allocates with zero initialized memroy
    * returns T* (pointer)
* __composite literals__
    * expression creates a new instance
    * ex:*
        * &File{fd: fd , name: name}
* __make__
    * make()
    * for slices, maps and channels
    * ex: make([]int, 10, 100)

### Array vs Slice
* Array
    * size is part of the index
    * ex: [3]int
* Slice
    * pointer to array
    * []int  (note: no size)
[goslices](https://blog.golang.org/slices)


### Type assertions
* x.(T)
    * asserts that x is not nil and that the value stored in x is of type T
    * x.(T) is called a type assertion