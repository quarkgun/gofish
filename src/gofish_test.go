package main

import (
	"testing"
	"fmt"
)


func removeElem(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}


func TestSliceDelete(t *testing.T) {
	slice := []string{"1", "2", "3"}
	slice = removeElem(slice, 1)
	fmt.Println(slice, "length=", len(slice))
}

