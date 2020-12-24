package monster

import (
	"testing"
)

func TestSerialization(t *testing.T) {

	monster := Monster{
		Name:  "张三",
		Age:   20,
		Skill: "电工",
	}

	res, err := monster.serialization()
	if !res {
		t.Fatalf("执行失败，期望值=%v，实际值=%v,错误原因：%v", true, res, err)
	} else {
		t.Logf("执行正确，期望值=%v，实际值=%v，", true, res)
	}

}
