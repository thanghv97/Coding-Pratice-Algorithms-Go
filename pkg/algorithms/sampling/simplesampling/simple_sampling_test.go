package simplesampling_test

import (
	"fmt"
	"testing"

	"github.com/thanghv97/Coding-Pratice-Algorithms-Go/pkg/algorithms/sampling/simplesampling"
)

func input() []interface{} {
	return []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func Test(t *testing.T) {
	var data []interface{} = input()

	fmt.Printf("Init Sampling with data: %v\n", data)
	sampling := simplesampling.Init(data)
	for i := 0; i < 10; i++ {
		val, _ := sampling.Random(1)
		fmt.Printf("Get Random: %v\n", val)
	}
}

func Test2(t *testing.T) {
	var data []interface{} = input()

	fmt.Printf("Init Sampling with data: %v\n", data)
	sampling := simplesampling.Init(data)
	for i := 0; i < 10; i++ {
		val, _ := sampling.Random(2)
		fmt.Printf("Get Random: %v\n", val)
	}
}
