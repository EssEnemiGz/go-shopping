package request

import (
	"fmt"
	"net/http"
)

func contains(method string, avaible_methods []string) bool {
	for _, value := range avaible_methods {
		if method == value {
			return true
		}
	}

	return false
}

func Method_verificator(r *http.Request, avaible_methods []string) bool {
	return contains(r.Method, avaible_methods)
}

func Args(r *http.Request) {
	fmt.Println(r)
}
