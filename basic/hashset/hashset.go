package main

import "fmt"

type IntSet struct {
	set map[int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(i int) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found	//False if it existed already
}

func (set *IntSet) Existed(i int) bool {
	_, found := set.set[i]
	return found	//true if it existed already
}

func (set *IntSet) Remove(i int) {
	delete(set.set, i)
}

func (set *IntSet) Array() map[int]bool{
	return set.set
}

func main() {
	set := NewIntSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)
	set.Add(3)
	set.Add(3)
	set.Add(3)
	set.Add(4)
	set.Add(5)
	set.Add(5)
	set.Add(5)
	set.Add(1)
	fmt.Println(set.Existed(2))
	set.Remove(2)
	fmt.Println(set.Existed(2))
	for k, v := range set.Array() {
		fmt.Printf("%v=%v;", k, v)
	}


}


//type HashSet struct {
//	set map[string]bool
//}
//
//func NewHashSet() *HashSet {
//	return &HashSet{make(map[string]bool)}
//}
//
//func (set *HashSet) Add(i string) bool {
//	_, found := set.set[i]
//	set.set[i] = true
//	return !found //False if it existed already
//}
//
//func (set *HashSet) Get(i string) bool {
//	_, found := set.set[i]
//	return found //true if it existed already
//}
//
//func (set *HashSet) Get(i string) bool {
//	for k, v := range m {
//		fmt.Printf("%s=%d;", k, v)
//	}
//
//
//	_, found := set.set[i]
//	return found //true if it existed already
//}
//
//func (set *HashSet) Remove(i string) {
//	delete(set.set, i)
//}

//func main() {
//	hashSet := NewHashSet()
//	hashSet.Add("1")
//	hashSet.Add("2")
//	hashSet.Add("3")
//	hashSet.Add("4")
//	hashSet.Add("5")
//
//
//}
