package userest

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func get_user(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}
func modify_user(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are modify user %s", uid)
}
func delete_user(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are delete user %s", uid)
}
func add_user(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are add user %s", uid)
}

func Use_rest() {
	mux := routes.New()
	mux.Get("/user/:uid", get_user)
	mux.Post("/user/:uid", modify_user)
	mux.Del("/user/:uid", delete_user)
	mux.Put("/user/", add_user)
	http.Handle("/", mux)
	http.ListenAndServe(":8088", nil)
}
