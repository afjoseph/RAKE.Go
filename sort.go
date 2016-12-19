package rake

import (
	"sort"
)

func reverseSortByValue(myMap map[string]float64) PairList {
	pl := make(PairList, len(myMap))
	i := 0
	for k, v := range myMap {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

//Pair is a simple struct for a key-value pair of string and float64
type Pair struct {
	Key   string
	Value float64
}

//PairList is just a slice of Pairs
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
