package testpackage

import (
	"testing"
	"usetest"
)

func Test_Division_1(t *testing.T) {
	if i, e := usetest.Division(6, 2); i != 3 || e != nil {
		t.Error("除法函数测试失败！")
	} else {
		t.Log("测试通过！")
	}
}

func Test_Division_2(t *testing.T) {
	t.Error("就是不通过！")

}
