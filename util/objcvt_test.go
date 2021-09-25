package util

import (
	"testing"
)

func TestCopyStructFields(t *testing.T) {

	type A struct {
		AName string `copy:"BName"`
		AAge  int    `copy:"BAge"`
	}

	type B struct {
		BName string
		BAge  int
	}

	a, b := A{"张三", 1}, B{}
	err := CopyStructFields(&a, &b)

	if err != nil {
		t.Fatal("执行失败", err)
	}

	if b.BName != a.AName || b.BAge != a.AAge {
		t.Fatalf("执行失败，期望值%v，目标值 %v", a, b)
	}

}

func TestConvertForComplexStruct(t *testing.T) {

	type A struct {
	}

	type B struct {
		BStr   string
		BInt   int64
		BFloat float32
		BA     A
	}

	type C struct {
	}

	type D struct {
	}

}
