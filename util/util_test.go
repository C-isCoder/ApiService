package util

import "testing"

// 	单元测试
func TestGenShortId(t *testing.T) {
	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		t.Error("GenShortId failed!")
	}

	t.Log("GenShortId test pass")
}

// 性能测试
func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

// 性能压力测试
func BenchmarkGenShortIdTimeConsuming(b *testing.B) {
	b.StopTimer() // 调用该函数停止压力测试的时间计数

	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		b.Error(err)
	}

	b.StartTimer() // 重新开始时间

	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}
