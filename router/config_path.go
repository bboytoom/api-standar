package router

import "net/http"

type PathRoute struct {
	rules map[string]map[string]http.HandlerFunc
}

func ConfigPath() *PathRoute {
	return &PathRoute{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (pr *PathRoute) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := pr.rules[path]

	handler, methodExists := pr.rules[path][method]

	return handler, methodExists, exist
}

func (pr *PathRoute) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, exist := pr.FindHandler(request.URL.Path, request.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
