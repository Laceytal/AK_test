package handles

import (
	"aktest/internal/ak"
	"fmt"
	"net/http"
)

func HandlOne(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello AK!")
}
func HandlTwo(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello Gevorg!")
}
func HandlThree(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, ak.SymFunc())
}
