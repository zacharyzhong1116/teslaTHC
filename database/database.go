package dbopers

type persistence interface {
	Insert(data int) error
	Update(data interface{}) error
	Delete(id string) error
}

//

const (
	host     = "myhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydb"
)

type datebase struct {
}
