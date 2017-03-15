package usetemplate

import (
	"fmt"
	"text/template"
)

func Mustone() {
	tmOk := template.New("first")
	template.Must(tmOk.Parse(" some static text /* and a comment */"))
	fmt.Println("this first one parse ok.")
	fmt.Println("the next one ough to fail")
	tmErr := template.New("check parse error with Must")
	template.Must(tmErr.Parse(" some static text {{.Name}"))

}
