package resource

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type AuthResource struct{
	mux *chi.Mux
}

func InitAuthResource(m *chi.Mux){
	rsc := &AuthResource{mux:m}
	rsc.mux.Get("/auth/microsoft", rsc.microsoft)
}

//https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=5a9fbc76-88d3-4e73-bb07-b455160a286e&scope=openid,email,profile&response_type=code
func (rsc *AuthResource) microsoft(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	//body, _ := ioutil.ReadAll(r.Body)
	w.Write([]byte("welcome chi router"))
	log.Printf("request %s",r)
}

