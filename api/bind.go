package api

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/mux"
)

func bindQuery(r *http.Request, obj interface{}) error {
	return binding.Query.Bind(r, obj)
}

func bindUri(r *http.Request, obj interface{}) error {
	params := mux.Vars(r)
	m := make(map[string][]string)
	for k, v := range params {
		m[k] = []string{v}
	}
	return binding.Uri.BindUri(m, obj)
}

func bindJSON(r *http.Request, obj interface{}) error {
	return binding.JSON.Bind(r, obj)
}
