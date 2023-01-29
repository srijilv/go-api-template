package openapi

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func (s Server) PostAccountsRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostAccountsRegister api  works")
	var reqBody RegiserRequestBody
	render.Bind(r, &reqBody)
	fmt.Printf("reqBody: %+v\n", reqBody)

}

func (rd RegiserRequestBody) Bind(r *http.Request) (err error) {
	return
}
