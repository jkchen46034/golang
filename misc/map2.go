// a store with hashmap and slice
package main

import "fmt"

type Store struct {
	m        map[string]string
	snapshot []map[string]string
}

func (s *Store) Set(k string, v string) {
	s.m[k] = v
}

func (s *Store) Get(k string) string {
	return s.m[k]
}

func (s *Store) Snapshot() int {
	snapshot := Copy(s.m)
	s.snapshot = append(s.snapshot, snapshot)
	return len(s.snapshot) - 1
}

func (s *Store) Get_From_Snapshot(id int, k string) (v string) {
	return s.snapshot[id][k]
}

func Copy(a map[string]string) (b map[string]string) {
	b = make(map[string]string)
	for k, v := range a {
		b[k] = v
	}
	return b
}

func newStore() *Store {
	s := new(Store)
	s.m = make(map[string]string)
	s.snapshot = make([]map[string]string, 0)
	return s
}
func main() {
	s := newStore()
	s.Set("k1", "v1")
	s.Set("k2", "v2")
	fmt.Println(s.Get("k1"), s.Get("k2"))
	id := s.Snapshot()
	fmt.Println("snapshot 0, k1: ", s.Get_From_Snapshot(id, "k1"))
	s.Set("k1", "v3")
	id2 := s.Snapshot()
	fmt.Println("snapshot 1, k1: ", s.Get_From_Snapshot(id2, "k1"))
	fmt.Println("snapshot 0, k1: ", s.Get_From_Snapshot(id, "k1"))
}

/*
jk:misc$time go run map2.go
v1 v2
snapshot 0, k1:  v1
snapshot 1, k1:  v3
snapshot 0, k1:  v1

real	0m0.228s
user	0m0.189s
sys	0m0.112s
*/
