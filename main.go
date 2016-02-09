package main
import (
	"conf"
	"net/http"
	"github.com/go-martini/martini"
)

func main() {
	conf.SayHello()
	m := martini.Classic()

	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})
}
