# Some common algorithms implemented in Golang

## Running Tests
To run tests for all packages, cd to this project's directory and run -
```bash
go test ./...
```

To test an individual package (let's say sorting) -
```bash
go test ./sorting
```

## Peak Element
'peak_element' contains methods for finding the any peak element in a 1-d and 2-d array.

## Sorting
'sorting' contains the methods for insertion and merge sorting an array of integers.

## Heaps
'heap' contains the struct Heap and PriorityQueue. Heap is implemented using min-heap.
PriorityQueue repackages Heap struct to work as a max-heap.