# 軽い処理

`repository` -> `service` -> `router`
それぞれを`iter`, `loop`, `Pipeline`, `Fun-Out`で計測。

## ベンチマーク

大体同じ傾向です。

```text

goos: linux
goarch: amd64
pkg: github.com/yyyoichi/iter-pipeline-sample
cpu: 13th Gen Intel(R) Core(TM) i7-1360P
=== RUN   BenchmarkRouter
BenchmarkRouter
=== RUN   BenchmarkRouter/Iter
BenchmarkRouter/Iter
BenchmarkRouter/Iter-16                  1562252               860.9 ns/op           152 B/op          7 allocs/op
=== RUN   BenchmarkRouter/Loop
BenchmarkRouter/Loop
BenchmarkRouter/Loop-16                  4037108               295.3 ns/op           896 B/op          1 allocs/op
=== RUN   BenchmarkRouter/Pipeline
BenchmarkRouter/Pipeline
BenchmarkRouter/Pipeline-16                54279             21092 ns/op             224 B/op          7 allocs/op
=== RUN   BenchmarkRouter/FunOut
BenchmarkRouter/FunOut
BenchmarkRouter/FunOut-16                  12694             94362 ns/op            3325 B/op         59 allocs/op
PASS
ok      github.com/yyyoichi/iter-pipeline-sample        8.147s

```

以上。
軽い処理であれば、普通にループしたほうが良いというのはこれからも変わらないか。
