package handlers

import (
	"bookings/pkg/config"
	"bookings/pkg/models"
	"bookings/pkg/render"
	"errors"
	"fmt"
	"net/http"
)

var Repo *Repository
type Repository struct{
	app *config.AllCache
}

func  NewRepo(a* config.AllCache) *Repository{
	return &Repository{
		app: a,
	}
}
func NewHandlers(r *Repository){
	Repo=r
}

func (m* Repository) Home(w http.ResponseWriter, r* http.Request){
	//fmt.Fprintf(w,"Hello world")
	remoteIp:=r.RemoteAddr
	//m.app.Session.Put(r.Context(), "remote_ip", remoteIp)

	stringMap:=map[string]string{
		"test":"hellofdndfcdcvfdj",	
		"remote_ip":remoteIp,
	}

	render.RenderTemplate(w,"home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}


func Divide(w http.ResponseWriter, r* http.Request){
	res,err:=DivideInt(10,0)
	if err==nil{
		fmt.Fprintf(w,"%d",res)
	}else{
		fmt.Fprintf(w,fmt.Sprintf("Error: %s",err))
	}
}

func DivideInt(x, y int) (int, error){
	if y==0{
		return 0, errors.New("Cannot divide by 0")
	}else{
		return x/y, nil
	}
}