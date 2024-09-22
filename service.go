package main

import (
	"iter"
	"runtime"
	"sync"
)

type Service struct {
	r Repository
}

func (s Service) Iter() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i, item := range s.r.Generate() {
			p := s.sumPrice(item[0], item[1])
			_ = yield(i, p)
		}
	}
}

func (s Service) Pipeline() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		var wg sync.WaitGroup
		for i, item := range s.r.Generate() {
			wg.Add(1)
			p := s.sumPrice(item[0], item[1])
			go func() {
				defer wg.Done()
				_ = yield(i, p)
			}()
		}
		wg.Wait()
	}
}

func (s Service) FunOut() iter.Seq2[int, int] {
	source := s.r.Generate()

	// Fun-Out
	procs := runtime.GOMAXPROCS(0)
	seq2s := make([]iter.Seq2[int, int], procs)
	for p := range runtime.GOMAXPROCS(0) {
		seq2s[p] = func(yield func(int, int) bool) {
			go func() {
				for i, item := range source {
					p := s.sumPrice(item[0], item[1])
					go func() {
						_ = yield(i, p)
					}()
				}
			}()
		}
	}

	// Fun-In
	return func(yield func(int, int) bool) {
		var wg sync.WaitGroup
		wg.Add(procs)

		for _, seq2 := range seq2s {
			go func() {
				defer wg.Done()
				for i, sum := range seq2 {
					_ = yield(i, sum)
				}
			}()
		}
		wg.Wait()
	}
}

func (s *Service) sumPrice(price, num int) int {
	return price * num
}
