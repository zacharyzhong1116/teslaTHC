package database

import (
	"net/http"

	"teslaTHC/database/utils"
)

// suppose get those security infro vault or somewhere else
const (
	HOST     = "myhost"
	PORT     = 5432
	USER     = "myuser"
	PASSWORD = "mypassword"
	ADDRESS  = "myaddress"
)

func main() {
	exitCahnnel := make(chan int)

	per.Init()
	r := utils.NewRouter()
	http.ListenAndServe(ADDRESS, r)
	<-exitCahnnel
}

