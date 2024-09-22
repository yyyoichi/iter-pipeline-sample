package main

import (
	"log"
)

type Router struct {
	s Service
}

func (r Router) HandleWithIter() {
	var count int
	for sum := range r.s.Iter() {
		_ = sum
		count++
	}
	if count != 100 {
		log.Println("Iter", count)
	}
}

func (r Router) HandleWithPipeline() {
	var count int
	for sum := range r.s.Pipeline() {
		_ = sum
		count++
	}
	if count != 100 {
		log.Println("Pipeline", count)
	}
}

func (r Router) HandleWithFunOut() {
	var count int
	for sum := range r.s.FunOut() {
		_ = sum
		count++
	}
	if count != 100 {
		log.Println("FunOut", count)
	}
}
