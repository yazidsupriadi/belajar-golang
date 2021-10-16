package handler

//import package
import (
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

//buat fungsi route

func HomeFunc(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := map[string]interface{}{
		"title":   "belajar golang",
		"content": "sedang belajar golang!",
	}
	template, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func HelloFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Golang!"))
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Test!"))
}
func ProductFunc(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)
	data := map[string]interface{}{
		"content": idNumb,
	}
	template, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, data)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
}
