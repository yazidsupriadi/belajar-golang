package handler

//import package
import (
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
	"webgolang/entity"
)

//buat fungsi route

func HomeFunc(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//	data := map[string]interface{}{
	//		"title":   "belajar golang",
	//		"content": "sedang belajar golang!",
	//	}
	template, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//	data := entity.Product{ID: 1, Name: "Mobilio", Price: 220000, Stock: 3}
	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000, Stock: 3},
		{ID: 2, Name: "Toyota", Price: 230000, Stock: 2},
		{ID: 3, Name: "Xenia", Price: 240000, Stock: 6},
	}

	err = template.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "error is happening keep calm!", http.StatusInternalServerError)
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

func GetPost(w http.ResponseWriter, r *http.Request) {

	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("ini adalah GET"))
	case "POST":
		w.Write([]byte("ini adalah POST"))
	default:
		http.Error(w, "Error is happening", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		template, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = template.Execute(w, nil)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		return

	}
	http.Error(w, "Error is happening", http.StatusBadRequest)
}

func Proses(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		name := r.Form.Get("name")
		message := r.Form.Get("message")

		w.Write([]byte(name))
		w.Write([]byte(message))
		return
	}
	http.Error(w, "Error is happening", http.StatusBadRequest)

}
