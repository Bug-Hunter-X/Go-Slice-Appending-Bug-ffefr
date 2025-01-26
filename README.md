# Go Slice Appending Bug
This repository demonstrates a common bug in Go related to appending to slices. When appending to a slice whose capacity is exceeded, Go reallocates the underlying array, potentially invalidating existing pointers and causing unexpected behavior or crashes.  This example shows how to address this issue.

## Bug Description
The `bug.go` file contains a simple program that appends elements to a slice.  However, it does not explicitly set the capacity, resulting in multiple reallocations which can be inefficient and potentially lead to errors in more complex programs where pointers to the original slice are held.

## Solution
The `bugSolution.go` file provides a corrected version. By setting the initial capacity appropriately, we minimize reallocations, improving performance and reliability. 
