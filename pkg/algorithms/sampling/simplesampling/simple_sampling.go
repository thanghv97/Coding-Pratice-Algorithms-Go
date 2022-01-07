// package simplesampling
// simple sampling algorithm use to pick random a element with same probability and know size n
package simplesampling

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type SimpleSampling struct {
	data []interface{}
	size int
}

func Init(data []interface{}) SimpleSampling {
	rand.Seed(time.Now().UnixNano())
	return SimpleSampling{
		data: data,
		size: len(data),
	}
}

func (ss *SimpleSampling) Random(k int) ([]interface{}, error) {
	if k <= 0 {
		return nil, errors.New("k needed to set more than 0")
	}
	if ss.size < k {
		return nil, fmt.Errorf("length of data is less than %d", k)
	}

	r := make([]interface{}, k)
	for i := 0; i < k; i++ {
		ii := rand.Intn(ss.size)
		r[i] = ss.data[ii]
	}
	return r, nil
}
