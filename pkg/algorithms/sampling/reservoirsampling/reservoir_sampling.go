package reservoirsampling

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type ReservoirSampling struct {
	data []interface{}
}

func Init(data []interface{}) ReservoirSampling {
	rand.Seed(time.Now().UnixNano())
	return ReservoirSampling{
		data: data,
	}
}

func (rs *ReservoirSampling) Random(k int) ([]interface{}, error) {
	if k <= 0 {
		return nil, errors.New("k needed to set more than 0")
	}
	if len(rs.data) < k {
		return nil, fmt.Errorf("length of data is less than %d", k)
	}

	r := make([]interface{}, k)
	for i := 0; i < k; i++ {
		r[i] = rs.data[i]
	}
	for i := k; i < len(rs.data); i++ {
		j := rand.Intn(i + 1)
		if j < k {
			r[j] = rs.data[i]
		}
	}

	return r, nil
}
