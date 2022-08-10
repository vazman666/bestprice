package main

import (
	"bestprice/models"
	"bestprice/pkg"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var article string

func home_page(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		fmt.Printf("Post")
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		article = r.FormValue("model")
		fmt.Printf("Article =%v\n", article)
		tmp, err := pkg.Tiss(article)
		pkg.Zapros(article)
		/*for i, j := range tmp {
			models.Rez[i] = append(models.Rez[i], models.Str{"Tiss", j})
			fmt.Printf("Rez=%v\n", models.Rez)
		}
		tmp, err = pkg.Forum(article)
		for i, j := range tmp {
			models.Rez[i] = append(models.Rez[i], models.Str{"Forum", j})
			fmt.Printf("Rez=%v\n", models.Rez)
		}*/
		fmt.Printf("%v\n", tmp)
		//fmt.Printf("%v\n", model)
		http.Redirect(w, r, "/", 301)
	} else {

		tmpl, _ := template.ParseFiles("templates/homepage.html")
		//fmt.Printf("Rez=%v\n", models.Rez)
		tmpl.Execute(w, models.Rez)
	}
	//fmt.Fprint(w, "Go is super")
}
func ajax_intro(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Ajax_Intro")
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Fprint(w, "Go is super")
}
func main() {
	//models.Rez = make(map[string][]models.Str)
	http.HandleFunc("/", home_page)
	http.HandleFunc("/ajax_intro/", ajax_intro)
	fmt.Printf("Слушаем порт 8080\n")
	http.ListenAndServe(":8080", nil)
}
