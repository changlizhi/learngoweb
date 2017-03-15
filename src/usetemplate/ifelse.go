package usetemplate

import (
	"os"
	"text/template"
)

//注意：if里面无法使用条件判断，例如.Mail=="astaxie@gmail.com"，这样的判断是不正确的，if里面只能是bool值
func Useifels() {
	tmEmpty := template.New("template test")
	tmEmpty = template.Must(tmEmpty.Parse("空 pipeline if demo:{{if ``}} 不会输出.{{end}}\n"))
	tmEmpty.Execute(os.Stdout, nil)
	tmWithValue := template.New("template test")
	tmWithValue = template.Must(tmWithValue.Parse("不为空的 pipeline if demo:{{if `anything`}}有内容，会输出{{end}}\n"))
	tmWithValue.Execute(os.Stdout, nil)
	tmIfElse := template.New("template test")
	tmIfElse = template.Must(tmIfElse.Parse("if-else demo :{{if `anything`}} if部分{{else}}else部分.{{end}}\n"))
	tmIfElse.Execute(os.Stdout, nil)
}
