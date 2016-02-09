package main
import (
	"github.com/and-hom/go-test/conf"
	"github.com/and-hom/go-test/alg"
	"net/http"
	"fmt"
	"github.com/go-martini/martini"
)

func main() {
	fmt.Printf("2^2 = %d\n", alg.Pow2(2))
	conf.SayHello()
	m := martini.Classic()

	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})
}

