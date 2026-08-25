package main

import(
	"fmt"
)

type Store[K comparable,V any] struct{
	Data map[K]V
}

func main(){

	userStore:=Store[int,string]{
		Data: map[int]string{
			1:"Fatih",
			2:"Ali",
		},
	}

	userStore.Add(3, "Ahmet")
	value, exists := userStore.Get(1)
	if exists {
		fmt.Println(value)
	}

	scoreStore:=Store[string,int]{
		Data: map[string]int{
			"Math": 90,
			"Science": 80,
		},
	}

}

func (s *Store[K,V]) Add(key K,value V){
	if s.Data==nil{
		s.Data=make(map[K]V)
	}
	s.Data[key]=value
}

func (s Store[K,V]) Get(key K) (V, bool) {
	value, exists := s.Data[key]
	return value, exists
}
