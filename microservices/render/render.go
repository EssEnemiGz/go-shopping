package render

import (
	"fmt"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Render microservice")
}
