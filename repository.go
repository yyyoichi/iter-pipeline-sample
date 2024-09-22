package main

import (
	"iter"
	"slices"
)

type Repository struct {
	items [][2]int // [][2]int{price, num}
}

func NewRepository() Repository {
	var r Repository
	r.items = make([][2]int, 100)
	for i := range r.items {
		r.items[i] = [2]int{100 - i, 0 + 1}
	}
	return r
}

func (r Repository) Generate() iter.Seq2[int, [2]int] {
	return slices.All(r.items)
}
