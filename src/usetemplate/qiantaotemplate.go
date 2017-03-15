package usetemplate

import (
	"fmt"
	"html/template"
	"os"
)

//{{define "子模板名称"}}内容{{end}}
//{{template "子模板名称"}}

func Use_QianTao_Template() {
	s1, _ := template.ParseFiles("src/usetemplate/templates/header.tmpl", "src/usetemplate/templates/content.tmpl", "src/usetemplate/templates/footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}
