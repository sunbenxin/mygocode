package main

import "testing"

var a = []string{
	"教练",
	"主任",
	"厂长",
	"行长",
	"主管",
	"所长",
	"师父",
	"木匠",
	"客户",
	"站长",
	"女士",
	"监理",
	"中介",
	"店长",
	"出租",
	"助理",
	"三轮",
	"镇长",
	"空调",
	"师付",
	"教授",
	"同学",
	"园长",
	"股长",
	"政委",
	"物业",
	"参谋",
	"摩托",
}

func BenchmarkRange1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		range1()
	}
}
func BenchmarkRange2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		range2()
	}
}

func range1() {
	for k, v := range a {
		_ = k
		_ = v
	}
}
func range2() {
	for k, _ := range a {
		_ = k
		_ = a[k]
	}
}
