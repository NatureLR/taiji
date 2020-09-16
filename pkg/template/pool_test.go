package template

import (
	"fmt"
	"testing"
)

func Test_poll_add(t *testing.T) {
	DefaultPool.Add("test", "测试模板", "test—dir")
	for kind, t := range GetDefaul() {
		fmt.Println(kind, t)
	}
}
