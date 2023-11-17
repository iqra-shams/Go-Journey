package main
import(
	"github.com/gorilla/mux"
	"net/http"
	"log"
)
func setupServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
  }
  func main() {
	setupServer()
  }
  func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gorilla!"))
  }