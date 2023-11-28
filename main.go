package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1><p>To get in touch, email me at <a href=\"mailto:matthew.t.mckenney@gmail.com\">email</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w,
		`<h1>FAQ Page</h1>
		<ul>
			<li>
				<p>Q: Is there a free version?</p>
				<p>A: Yes! We offer a free trial for 30 days on any paid plan.</p>
			</li>
			<li>
				<p>Q: What are your support hours?</p>
				<p>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</p>
			</li>
			<li>
				<p>Q: How do I contact support?</p>
				<p>A: Email us!</p>
			</li>
		</ul>
		`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
