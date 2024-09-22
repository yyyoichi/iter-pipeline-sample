package main

import (
	"iter"
	"runtime"
	"sync"
)

type Service struct {
	r Repository
}

func (s Service) Iter() iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, item := range s.r.Generate() {
			p := s.sumPrice(item[0], item[1])
			_ = yield(p)
		}
	}
}

func (s Service) Pipeline() iter.Seq[int] {
	ch := make(chan [2]int)
	go func() {
		defer close(ch)
		for _, item := range s.r.Generate() {
			ch <- item
		}
	}()
	return func(yield func(int) bool) {
		for item := range ch {
			p := s.sumPrice(item[0], item[1])
			_ = yield(p)
		}
	}
}

func (s Service) FunOut() iter.Seq[int] {
	ch := make(chan [2]int)
	go func() {
		defer close(ch)
		for _, item := range s.r.Generate() {
			ch <- item
		}
	}()

	// Fun-Out
	procs := runtime.GOMAXPROCS(0)
	var outChs = make([]chan int, 0, procs)
	for range procs {
		outCh := make(chan int)
		go func() {
			defer close(outCh)
			for {
				item, ok := <-ch
				if !ok {
					return
				}
				p := s.sumPrice(item[0], item[1])
				outCh <- p
			}
		}()
		outChs = append(outChs, outCh)
	}

	// Fun-In
	var wg sync.WaitGroup

	inCh := make(chan int)
	for _, outCh := range outChs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range outCh {
				inCh <- p
			}
		}()
	}

	go func() {
		wg.Wait()
		close(inCh)
	}()
	return func(yield func(int) bool) {
		for p := range inCh {
			_ = yield(p)
		}
	}
}

func (s *Service) sumPrice(price, num int) int {
	return price * num
}
