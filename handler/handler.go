package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// data
	// data := map[string]interface{}{
	// 	"title": "I'm Learning Go web",
	// 	"content": "I'm Learing Go Web with Agung Setiawan",
	// }

	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 22000000, Stock: 4}
	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 22000000, Stock: 11},
		{ID: 2, Name: "Xpander", Price: 27000000, Stock: 2},
		{ID: 3, Name: "Pajero Sport", Price: 30000000, Stock: 8},
	}
	
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, saya sedang belajar web golang"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario is Game nittendo"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprintf(w, "Product page : %d", idNumb)
	data := map[string]interface{}{
		"content" : idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"),path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}
