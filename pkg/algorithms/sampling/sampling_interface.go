package sampling

type AlgoSampling interface {
	Random(k int) ([]interface{}, error)
}
