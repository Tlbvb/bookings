package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"bookings/pkg/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)
var session* scs.SessionManager
//repository pattern-allows to swap inside the project with minimum cost

func main() {
	var App config.AllCache
	App.InProduction=false

	session=scs.New()
	session.Lifetime=24*time.Hour
	session.Cookie.Persist=true
	session.Cookie.SameSite=http.SameSiteLaxMode
	session.Cookie.Secure=App.InProduction

	App.TemplateCache=render.CreateCache()
	App.UseCache=false
	App.Session=session

	repo:=handlers.NewRepo(&App)
	handlers.NewHandlers(repo)
	
	render.GetCache(&App)   //-> good way of doing variable transfer from one pkg to another
	//http.HandleFunc("/",handlers.Repo.Home) -> good way of doing it, but it is better if router is in another file-> we will use 3rd party package-> pat routing
	//http.HandleFunc("/divide",handlers.Divide)
	
	/*fmt.Print("df")
	http.HandleFunc("/", func(w http.ResponseWriter,r* http.Request){
		fmt.Fprintf(w,"Hello world")
	})*/

	//http.ListenAndServe(":8080",nil)

	srv:=&http.Server{
		Addr:"localhost:8080",
		Handler:routes(&App),
	}
	srv.ListenAndServe()
}


//go mod init app_name
//Emmet, formerly known as Zen Coding, is a web development tool and abbreviation engine that makes it easier and faster to write HTML and CSS code.
//go run .  -> run multiple files at the same time
//go run ./cmd/web  works in both mac and windows