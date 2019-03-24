package bench

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// name of the file ends with *_test.go

// name of tests should be TestXXXX
func TestFib(t *testing.T) {
	assert.Equal(t, 0, fib(1), "fib test")
	assert.Equal(t, 1, fib(2), "fib test")
	assert.Equal(t, 1, fib(3), "fib test")
	assert.Equal(t, 2, fib(4), "fib test")
}

func TestFibNicely(t *testing.T) {
	tests := []struct {
		nth int
		out int
	}{
		{nth: 5, out: 3},
		{nth: 6, out: 5},
		{nth: 7, out: 8},
	}

	for _, test := range tests {
		res := fib(test.nth)
		assert.Equal(t, test.out, res, "fib test")
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(100000)
		// change above to i
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello(1000000)
	}
}

var res string
var j int

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, j = tryJson("ashish", i)
		assert.Equal(b, res, "ashish")
		assert.Equal(b, j, i)
	}
}

// $ go test -v
// === RUN   TestFib
// --- PASS: TestFib (0.00s)
// === RUN   TestFibNicely
// --- PASS: TestFibNicely (0.00s)
// PASS
// ok      github.com/ashishnegi/golang-tour/bench 0.058s

// $ go test -bench Fib
// goos: windows
// goarch: amd64
// pkg: github.com/ashishnegi/golang-tour/bench
// BenchmarkFib-8          200000000                7.67 ns/op
// PASS
// ok      github.com/ashishnegi/golang-tour/bench 2.356s

// C:\Users\asnegi\go\src\github.com\ashishnegi\golang-tour\bench>go test -bench Hello -benchmem
// goos: windows
// goarch: amd64
// pkg: github.com/ashishnegi/golang-tour/bench
// BenchmarkHello-4            2000            711997 ns/op               0 B/op          0 allocs/op
// PASS
// ok      github.com/ashishnegi/golang-tour/bench 2.953s

// change code in hello to allocate a string.

// C:\Users\asnegi\go\src\github.com\ashishnegi\golang-tour\bench>go test -bench Hello -benchmem
// goos: windows
// goarch: amd64
// pkg: github.com/ashishnegi/golang-tour/bench
// BenchmarkHello-4            2000            814000 ns/op               8 B/op          1 allocs/op
// PASS
// ok      github.com/ashishnegi/golang-tour/bench 3.112s

// $ go test -bench Hello -benchmem -cpuprofile cpu.out
// $ go tool pprof bench.test.exe cpu.out
// File: bench.test.exe
// Type: cpu
// Time: Mar 22, 2018 at 8:16pm (IST)
// Duration: 1.52s, Total samples = 1.31s (86.41%)
// Entering interactive mode (type "help" for commands, "o" for options)
// (pprof) top10
// Showing nodes accounting for 1.31s, 100% of 1.31s total
//       flat  flat%   sum%        cum   cum%
//      1.30s 99.24% 99.24%      1.30s 99.24%  github.com/ashishnegi/golang-tour/bench.fib
//      0.01s  0.76%   100%      0.01s  0.76%  sync.(*Pool).Get
//          0     0%   100%      0.01s  0.76%  fmt.Fprintln
//          0     0%   100%      0.01s  0.76%  fmt.Println
//          0     0%   100%      0.01s  0.76%  fmt.newPrinter
//          0     0%   100%      1.31s   100%  github.com/ashishnegi/golang-tour/bench.BenchmarkHello
//          0     0%   100%      1.31s   100%  github.com/ashishnegi/golang-tour/bench.hello
//          0     0%   100%      0.01s  0.76%  github.com/ashishnegi/golang-tour/bench.world
//          0     0%   100%      1.31s   100%  testing.(*B).launch
//          0     0%   100%      1.31s   100%  testing.(*B).runN
// (pprof)
// (pprof) list hello
// Total: 1.31s
// ROUTINE ======================== github.com/ashishnegi/golang-tour/bench.hello in D:\GoPath\src\github.com\ashishnegi\golang-tour\bench\code.go
//          0      1.31s (flat, cum)   100% of Total
//          .          .     15:   return a
//          .          .     16:}
//          .          .     17:
//          .          .     18:func hello(nth int) int {
//          .          .     19:   world("hello world..")
//          .      1.30s     20:   n := fib(nth)
//          .       10ms     21:   world(" nth fib")
//          .          .     22:   return n
//          .          .     23:}
//          .          .     24:
//          .          .     25:func world(s string) {
//          .          .     26:   fmt.Println(s)
// (pprof)
// (pprof) web
// (pprof) weblist fib
// (pprof)

// ashish@DESKTOP-133N35M MINGW64 /d/GoPath/src/github.com/ashishnegi/golang-tour/bench (master)
// $ go test -bench Json -benchtime 5s -memprofile memjson.out
// goos: windows
// goarch: amd64
// pkg: github.com/ashishnegi/golang-tour/bench
// BenchmarkJson-8           300000             21432 ns/op
// PASS
// ok      github.com/ashishnegi/golang-tour/bench 6.686s

// ashish@DESKTOP-133N35M MINGW64 /d/GoPath/src/github.com/ashishnegi/golang-tour/bench (master)
// $  go tool pprof --alloc_space bench.test.exe memjson.out
// File: bench.test.exe
// Type: alloc_space
// Time: Mar 22, 2018 at 9:24pm (IST)
// Entering interactive mode (type "help" for commands, "o" for options)

// (pprof) web

// (pprof) top10 -cum
// Showing nodes accounting for 410.58MB, 32.13% of 1277.74MB total
// Dropped 8 nodes (cum <= 6.39MB)
// Showing top 10 nodes out of 37
//       flat  flat%   sum%        cum   cum%
//       13MB  1.02%  1.02%  1277.24MB   100%  github.com/ashishnegi/golang-tour/bench.BenchmarkJson
//          0     0%  1.02%  1277.24MB   100%  testing.(*B).launch
//          0     0%  1.02%  1277.24MB   100%  testing.(*B).runN
//   204.55MB 16.01% 17.03%  1166.23MB 91.27%  github.com/ashishnegi/golang-tour/bench.tryJson
//    86.02MB  6.73% 23.76%   777.64MB 60.86%  encoding/json.Unmarshal
//          0     0% 23.76%   684.61MB 53.58%  encoding/json.(*decodeState).array
//          0     0% 23.76%   684.61MB 53.58%  encoding/json.(*decodeState).object
//          0     0% 23.76%   684.61MB 53.58%  encoding/json.(*decodeState).unmarshal
//          0     0% 23.76%   684.61MB 53.58%  encoding/json.(*decodeState).value
//      107MB  8.37% 32.13%   638.61MB 49.98%  reflect.MakeSlice
// (pprof)
