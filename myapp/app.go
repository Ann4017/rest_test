package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Uesr struct {
	S_f_name  string    `json:"first_name"`
	S_l_name  string    `json:"last_name"`
	S_email   string    `json:"email"`
	Create_at time.Time `json:"create_at"`
}

type Foo_handler struct{}

func Index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func (f *Foo_handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(Uesr)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.Create_at = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func Bar_handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "hello %s", name)
}

func New_http_handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Index_handler)

	mux.HandleFunc("/bar", Bar_handler)

	mux.Handle("/foo", &Foo_handler{})

	return mux
}
