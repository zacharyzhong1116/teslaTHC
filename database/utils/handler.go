package utils

import (
	"log"
	"net/http"
	"teslaTHC/database/persistence"
)

//PostHandler parse the request, get the data and save it into corrsponding table

func PostHandler(w http.ResponseWriter, r *http.Request) {
	p := persistence.GetPers()
	log.Printf("here get the pers and save the data")
}
