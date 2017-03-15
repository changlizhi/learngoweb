package usetemplate

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

//每一个模板函数都有一个唯一值的名字，
//然后与一个Go函数关联，通过如下的方式来关联
//type FuncMap map[string]interface{}

type FriendFunc struct {
	Fname string
}
type PersonFunc struct {
	UserName string
	Emails   []string
	Friends  []*FriendFunc
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	//replace thi @ by " at "
	return (substrs[0] + " at " + substrs[1])
}

func UseEmalDeal() {
	f1 := FriendFunc{Fname: "minxu.ma"}
	f2 := FriendFunc{Fname: "xinshis"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`
	hello {{.UserName}}!
    {{range .Emails}}
	    an emails {{.|emailDeal}}
	{{end}}
	{{with .Friends}}
	{{range .}}
	    my friend name is {{.Fname}}
	{{end}}
	{{end}}
	`)
	p := PersonFunc{
		UserName: "sdfe",
		Emails:   []string{"fewf@fdfe.dd", "fwef@dd.cs"},
		Friends:  []*FriendFunc{&f1, &f2},
	}
	t.Execute(os.Stdout, p)
}
