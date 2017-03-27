package useerr

import "net/http"

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandlerNewTwo func(http.ResponseWriter, *http.Request) *appError

func (fn appHandlerNewTwo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		c := appengine.NewContext(r)
		c.Errorf("%v", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

func viewRecordTwo(w http.ResponseWriter, r *http.Request) *appError {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return &appError{err, "Record not found", 404}
	}
	if err := viewTemplate.Execute(w, record); err != nil {
		return &appError{err, "Cant display record", 500}
	}
	return nil
}
