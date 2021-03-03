package find_dup_num

import "testing"

func benchmarkFindDup(b *testing.B, n int)  {
	b.StopTimer()
	nums := GenerateRandomArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FindDuplication(nums)
	}
}

func BenchmarkFindDup10(b *testing.B) { benchmarkFindDup(b, 10)}
func BenchmarkFindDup100(b *testing.B) { benchmarkFindDup(b, 100)}
func BenchmarkFindDup1000(b *testing.B) {benchmarkFindDup(b, 1000)}
func BenchmarkFindDup10000(b *testing.B) {benchmarkFindDup(b, 10000)}

// 执行 基准测试 基于内存
// go test -bench=. -benchmen
// 上述代码结果如下:
/*
BenchmarkFindDup10-4            312930950                8.60 ns/op            0 B/op          0 allocs/op
BenchmarkFindDup100-4           341443082                3.48 ns/op            0 B/op          0 allocs/op
BenchmarkFindDup1000-4          347663949                3.46 ns/op            0 B/op          0 allocs/op
BenchmarkFindDup10000-4         343853840                3.42 ns/op            0 B/op          0 allocs/op
 */

func benchmarkFindByBoolMap(n int, b *testing.B)  {
	b.StopTimer()
	nums := GenerateRandomArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FindDupNumByBoolMap(nums)
	}
}

func BenchmarkFindDupNumByBoolMap10(b *testing.B) { benchmarkFindByBoolMap(10, b)}
func BenchmarkFindDupNumByBoolMap100(b *testing.B) { benchmarkFindByBoolMap(100, b)}
func BenchmarkFindDupNumByBoolMap1000(b *testing.B) { benchmarkFindByBoolMap(1000, b)}
func BenchmarkFindDupNumByBoolMap10000(b *testing.B) { benchmarkFindByBoolMap(10000, b)}
// 执行基准测试, 结果如下:
/*
BenchmarkFindDupNumByBoolMap10-4         3824025               439 ns/op             176 B/op          1 allocs/op
BenchmarkFindDupNumByBoolMap100-4         728170              1405 ns/op            1568 B/op          2 allocs/op
BenchmarkFindDupNumByBoolMap1000-4        110856             11485 ns/op           24608 B/op          2 allocs/op
BenchmarkFindDupNumByBoolMap10000-4        18192             67775 ns/op          196640 B/op          2 allocs/op
*/

// 可以看出来, 进行了内存分配, 而且当数组长度增加, 每次执行会增加运行时间


// 若使用 空结构体map 试试结果
func benchmarkFindByStructMap(n int, b *testing.B)  {
	b.StopTimer()
	nums := GenerateRandomArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FindDupNumByStructMap(nums)
	}
}

func BenchmarkFindDupNumByStructMap10(b *testing.B) { benchmarkFindByStructMap(10, b)}
func BenchmarkFindDupNumByStructMap100(b *testing.B) { benchmarkFindByStructMap(100, b)}
func BenchmarkFindDupNumByStructMap1000(b *testing.B) { benchmarkFindByStructMap(1000, b)}
func BenchmarkFindDupNumByStructMap10000(b *testing.B) { benchmarkFindByStructMap(10000, b)}

// 执行基准测试, 结果如下:
/*
BenchmarkFindDupNumByStructMap10-4               5111673               269 ns/op             160 B/op          1 allocs/op
BenchmarkFindDupNumByStructMap100-4               635659              1660 ns/op            1440 B/op          2 allocs/op
BenchmarkFindDupNumByStructMap1000-4              126697             10610 ns/op           21792 B/op          2 allocs/op
BenchmarkFindDupNumByStructMap10000-4              19399             58873 ns/op          180256 B/op          2 allocs/op
 */