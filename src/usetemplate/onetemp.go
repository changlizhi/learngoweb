package usetemplate

import (
	"os"
	"text/template"
)

/*
如果要访问当前对象的字段通过{{.FieldName}},
但是需要注意一点：这个字段必须是导出的(字段首字母必须是大写的),否则在渲染的时候就不显示
*/
type Person struct {
	UserName string
	email    string //未导出的字段
}

func TempPerson() {
	t := template.New("filedname example")
	t, _ = t.Parse("hello {{.UserName}}!email:{{.email}}")
	p := Person{UserName: "xyz", email: "ccds@as.com"}
	t.Execute(os.Stdout, p)
}

type Friend struct {
	Fname string
}
type PersonNew struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func QianTaoMoban() {
	f1 := Friend{Fname: "xinmux.ma"}
	f2 := Friend{Fname: "xasd2.wei"}
	t := template.New("filedname exapmple")
	t, _ = t.Parse(`
		hello {{.UserName}}!
		{{range .Emails}}
			an email {{.}}
		{{end}}
		{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
		{{end}}
		`)
	p := PersonNew{
		UserName: "axisa",
		Emails:   []string{"sdfe@bego.co", "ass@aliy.com"},
		Friends:  []*Friend{&f1, &f2},
	}
	t.Execute(os.Stdout, p)
}
