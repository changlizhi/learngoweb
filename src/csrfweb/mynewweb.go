package csrfweb

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"html/template"
	"net/http"
	"os"
)

func get_request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("get action")
		t, _ := template.ParseFiles("src/csrfweb/templates/mysite.gtpl")
		t.Execute(w, nil)
	}
}

func whoami(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		name := r.Form.Get("name")
		CleanMap := make(map[string]interface{}, 0)
		if name == "xyz1" || name == "xyz2" || name == "xyz3" {
			CleanMap["name"] = "hello " + name
		}
		if len(CleanMap) > 0 {
			fmt.Fprintf(w, "%v", CleanMap["name"])
		}
	}
}
func Use_request() {
	http.HandleFunc("/mysite", get_request)
	http.HandleFunc("/whoami", whoami)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("致命错误！", err)
		os.Exit(0)
	}
}
