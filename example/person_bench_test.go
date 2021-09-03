package example

import (
	"encoding/json"
	"strings"
	"testing"
)

var testData = Person{
	Name: strings.Repeat("Fufu@中 文", 100),
	Age:  18,
}

func BenchmarkPerson_Marshal_Gencode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = testData.Marshal(nil)
	}
}

func BenchmarkPerson_Marshal_Gencode_buf(b *testing.B) {
	buf := make([]byte, len(testData.Name))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = testData.Marshal(buf)
	}
}

func BenchmarkPerson_Marshal_JSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(testData)
	}
}

func BenchmarkPerson_Unmarshal_Gencode(b *testing.B) {
	bs, _ := testData.Marshal(nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data Person
		_, _ = data.Unmarshal(bs)
	}
}

func BenchmarkPerson_Unmarshal_Gencode_buf(b *testing.B) {
	bs, _ := testData.Marshal(nil)
	var data Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = data.Unmarshal(bs)
	}
}

func BenchmarkPerson_Unmarshal_JSON(b *testing.B) {
	bs, _ := json.Marshal(testData)
	var data Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(bs, &data)
	}
}

// cpu: Intel(R) Xeon(R) CPU E3-1230 V2 @ 3.30GHz
// BenchmarkPerson_Marshal_Gencode
// BenchmarkPerson_Marshal_Gencode-8                2961789               382.5 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Marshal_Gencode-8                3249613               370.3 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Marshal_Gencode-8                3286573               367.9 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Marshal_Gencode_buf
// BenchmarkPerson_Marshal_Gencode_buf-8            3246898               372.4 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Marshal_Gencode_buf-8            3252370               367.5 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Marshal_Gencode_buf-8            3310006               379.5 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Marshal_JSON
// BenchmarkPerson_Marshal_JSON-8                    320176              3790 ns/op            1304 B/op          2 allocs/op
// BenchmarkPerson_Marshal_JSON-8                    349522              3592 ns/op            1304 B/op          2 allocs/op
// BenchmarkPerson_Marshal_JSON-8                    333120              3530 ns/op            1304 B/op          2 allocs/op
// BenchmarkPerson_Unmarshal_Gencode
// BenchmarkPerson_Unmarshal_Gencode-8              3193447               359.4 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_Gencode-8              3525896               362.9 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_Gencode-8              3545707               371.0 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_Gencode_buf
// BenchmarkPerson_Unmarshal_Gencode_buf-8          3495943               356.4 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_Gencode_buf-8          3312300               363.2 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_Gencode_buf-8          3451328               355.2 ns/op          1280 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_JSON
// BenchmarkPerson_Unmarshal_JSON-8                  100767             12764 ns/op            1504 B/op          6 allocs/op
// BenchmarkPerson_Unmarshal_JSON-8                   92001             13316 ns/op            1504 B/op          6 allocs/op
// BenchmarkPerson_Unmarshal_JSON-8                   92036             12417 ns/op            1504 B/op          6 allocs/op