# 比較

`pipeline`パターンは有効？

`repository` -> `service` -> `router`
それぞれを`iter`, `loop`, `Pipeline`, `Fun-Out`で計測。

詳細は[blog.yyyoichi.com](https://blog.yyyoichi.com/posts/202409230745/)から。

## 重い処理ベンチマーク(100_000)

`service`で、無駄な処理を10万回繰り返すパターンで計測。

```text
goos: linux
goarch: amd64
pkg: github.com/yyyoichi/iter-pipeline-sample
cpu: 13th Gen Intel(R) Core(TM) i7-1360P
=== RUN   BenchmarkRouter
BenchmarkRouter
=== RUN   BenchmarkRouter/Iter
BenchmarkRouter/Iter
BenchmarkRouter/Iter-16                        2         591060918 ns/op             160 B/op          7 allocs/op
=== RUN   BenchmarkRouter/Loop
BenchmarkRouter/Loop
BenchmarkRouter/Loop-16                        2         595090432 ns/op             896 B/op          1 allocs/op
=== RUN   BenchmarkRouter/Pipeline
BenchmarkRouter/Pipeline
BenchmarkRouter/Pipeline-16                    3         465564291 ns/op             224 B/op          7 allocs/op
=== RUN   BenchmarkRouter/FunOut
BenchmarkRouter/FunOut
BenchmarkRouter/FunOut-16                     20          55785737 ns/op            8379 B/op         78 allocs/op
PASS
ok      github.com/yyyoichi/iter-pipeline-sample        8.420s
```

`Fun-Out`が速い。メモリをうまく使って処理を終了できたよう。`iter`や`loop`と比較して約1/10のスピードで完了。
`Pipeline`が次点。やや速い。

## 軽い処理ベンチマーク

合計値の計算を単に、単価x個数で計算する。

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
