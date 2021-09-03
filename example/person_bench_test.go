package example

import (
	"encoding/json"
	"strings"
	"testing"
)

var testData = Person{
	Name: strings.Repeat("Fufu@中 文", 22),
	Age:  18,
}

func BenchmarkPerson_Marshal_Gencode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = testData.Marshal(nil)
	}
}

func BenchmarkPerson_Marshal_Gencode_buf(b *testing.B) {
	buf := make([]byte, testData.Size())
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = testData.Marshal(buf[0:0])
	}
}

func BenchmarkPerson_Marshal_Msgp(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = testData.MarshalMsg(nil)
	}
}

func BenchmarkPerson_Marshal_Msgp_buf(b *testing.B) {
	buf := make([]byte, testData.Msgsize())
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = testData.MarshalMsg(buf[0:0])
	}
}

func BenchmarkPerson_Marshal_JSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(testData)
	}
}

func BenchmarkPerson_Unmarshal_Gencode(b *testing.B) {
	var data Person
	bs, _ := testData.Marshal(nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = data.Unmarshal(bs)
	}
}

func BenchmarkPerson_Unmarshal_Msgp(b *testing.B) {
	var data Person
	bs, _ := testData.MarshalMsg(nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = data.UnmarshalMsg(bs)
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
// BenchmarkPerson_Marshal_Gencode-8               10333237               118.3 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Marshal_Gencode-8               11193163               110.0 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Marshal_Gencode-8               10886071               108.8 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Marshal_Gencode_buf-8           64314195                20.14 ns/op            0 B/op          0 allocs/op
// BenchmarkPerson_Marshal_Gencode_buf-8           60844830                19.67 ns/op            0 B/op          0 allocs/op
// BenchmarkPerson_Marshal_Gencode_buf-8           61803423                19.39 ns/op            0 B/op          0 allocs/op
// BenchmarkPerson_Marshal_Msgp-8                   9115250               125.9 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Marshal_Msgp-8                   9416580               124.6 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Marshal_Msgp-8                   8828001               125.9 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Marshal_Msgp_buf-8              45592878                26.55 ns/op            0 B/op          0 allocs/op
// BenchmarkPerson_Marshal_Msgp_buf-8              44223489                27.01 ns/op            0 B/op          0 allocs/op
// BenchmarkPerson_Marshal_Msgp_buf-8              44727219                26.76 ns/op            0 B/op          0 allocs/op
// BenchmarkPerson_Marshal_JSON-8                   1000000              1036 ns/op             312 B/op          2 allocs/op
// BenchmarkPerson_Marshal_JSON-8                   1000000              1112 ns/op             312 B/op          2 allocs/op
// BenchmarkPerson_Marshal_JSON-8                   1000000              1044 ns/op             312 B/op          2 allocs/op
// BenchmarkPerson_Unmarshal_Gencode-8             11196860               112.2 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_Gencode-8             10466701               111.0 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_Gencode-8             10469376               111.6 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_Msgp-8                 8685506               137.1 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_Msgp-8                 8862432               137.1 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_Msgp-8                 8683934               137.2 ns/op           288 B/op          1 allocs/op
// BenchmarkPerson_Unmarshal_JSON-8                  343899              3591 ns/op             512 B/op          6 allocs/op
// BenchmarkPerson_Unmarshal_JSON-8                  336992              3636 ns/op             512 B/op          6 allocs/op
// BenchmarkPerson_Unmarshal_JSON-8                  341617              3673 ns/op             512 B/op          6 allocs/op
