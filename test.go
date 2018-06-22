package main

import (
	"fmt"
	"strconv"
)


func main() {

	var s1 []string
	s1 = append(s1, "1")
	s1 = append(s1, "2")
	s1 = append(s1, "3")

	var s2 = make([]string, len(s1))
	copy(s2, s1)

	fmt.Println(s1)
	fmt.Println(s1[0])
	fmt.Println(&s1[0])

	fmt.Println(s2)
	fmt.Println(s2[0])
	fmt.Println(&s2[0])

	s1[0] = "haha"
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := []string{"1"}
	fmt.Println("s3 length", len(s3))
	for i :=0 ; i<5 ;i++  {
		fmt.Println("s3 length", len(s3))
		s3 = append(s3, strconv.Itoa(i))
	}

	for i := 0; i<5; i++ {
		s3 = append(s3[:i], s3[i+1:]...)
		fmt.Println(s3, "length=", len(s3))
	}

}
