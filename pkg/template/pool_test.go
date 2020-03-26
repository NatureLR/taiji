package template_test

import (
	"fmt"
	"testing"

	tpl "github.com/NatureLingRan/go-project/pkg/template"
)

func Test_poll_add(t *testing.T) {
	tpl.DefaultPool.Add("test", "测试模板", "test—dir")
	for kind, t := range tpl.GetDefaul() {
		fmt.Println(kind, t)
	}
}
