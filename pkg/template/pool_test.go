package template

import (
	"fmt"
	"testing"
)

func Test_poll_add(t *testing.T) {
	Default.Add("test", "测试模板", "test—dir")
	fmt.Println(Default.Allkind())
	fmt.Println(Default.Get("dockerfile"))
}
