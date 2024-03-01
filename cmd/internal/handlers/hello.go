package handlers

import (
	"fmt"
	"net/http"

	"github.com/crewjam/saml/samlsp"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	s := samlsp.SessionFromContext(r.Context())
	if s == nil {
		return
	}
	sa, ok := s.(samlsp.SessionWithAttributes)
	if !ok {
		return
	}
	CreateToken(w, sa.GetAttributes().Get("http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn"))

	fmt.Fprintf(w, "Token contents, %+v!\n", sa.GetAttributes())
	fmt.Fprintf(w, "UPN, %+v!", sa.GetAttributes().Get("http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn"))

}
