package main

import "testing"

func BenchmarkRouter(b *testing.B) {
	var router = Router{s: Service{r: NewRepository()}}

	b.Run("Iter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			router.HandleWithIter()
		}
	})

	b.Run("Pipeline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			router.HandleWithPipeline()
		}
	})

	b.Run("FunOut", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			router.HandleWithFunOut()
		}
	})
}
