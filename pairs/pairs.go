package pairs

type Pair struct {
	Key   string
	Value int
}

type ByValue []Pair

func (bv ByValue) Len() int {
	return len(bv)
}

func (bv ByValue) Less(i, j int) bool {
	if bv[i].Value != bv[j].Value {
		return bv[i].Value < bv[j].Value
	}
	return bv[i].Key < bv[j].Key
}

func (bv ByValue) Swap(i, j int) {
	bv[i], bv[j] = bv[j], bv[i]
}
