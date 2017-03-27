package useerr

import "net/http"

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServHttp(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}

}

func init() {
	http.Handle("/view", appHandler(viewRecordNew))
}

func viewRecordNew(w http.ResponseWriter, r *http.Request) err {
	c := appengine.NewContext(r)
	key := datastore.newKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return err
	}
	return viewTemplate.Execute(w, record)
}
