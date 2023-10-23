package render

import (
	"bookings/pkg/config"
	"bookings/pkg/models"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate1(w http.ResponseWriter, name string){   //Rendering-> Parsing+ Executing
	parsedTemplate,_:=template.ParseFiles("./templates/"+name, "./templates/base.layout.tmpl" )
	err:=parsedTemplate.Execute(w,nil)  //nil-> data passing to the template
	if err!=nil{
		fmt.Println("error")
	}
}

//building a simple template cache
var cache=make(map[string]*template.Template)
//cache:=map[string]*template.Template{}
func RenderTemplateSimple(w http.ResponseWriter, name string){
	_,inMap:=cache[name]
	if !inMap{
		parsedTemplate,_:=template.ParseFiles("./templates/"+name, "./templates/base.layout.tmpl")
		cache[name]=parsedTemplate
	}else{
		fmt.Println("exists in the cache")
	}
	var tmpl *template.Template
	tmpl=cache[name]
	err:=tmpl.Execute(w,nil)
	if err!=nil{
		fmt.Println(err)
	}
}

var app config.AllCache

func GetCache(a* config.AllCache){
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(&a)
	fmt.Println(a.TemplateCache)
	app=*a
}

//building more complex template cache
//creating the cache immediately for all the files that are inside

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplate(w http.ResponseWriter, name string,k *models.TemplateData){
	//create the cache
	//cache:=CreateCache()  //loading the cache every time-> need some absolute cache which can be accessed everywhere and loaded only once-> new package
	//fmt.Println(cache) -> now no need to create a cache because it is already 
	//find the template u r looking for
	if !app.UseCache{
		app.TemplateCache=CreateCache()
	}
	tmpl,_:=app.TemplateCache[name]
	err:=tmpl.Execute(w,k)   //also can do it with the buffer like buf:=new(bytes.Buffer tmpl.Execute(buf,nil)
	if err!=nil{				//buf.Write(w)
		fmt.Println(err)	
	}

	//*log.Logger can write to anything in the project
	//execute
}

func CreateCache() map[string]*template.Template{
	cache:=map[string]*template.Template{}

	pages,err:=filepath.Glob("./templates/*.page.tmpl")
	if err!=nil{
		fmt.Println(err)
	}
	for _,page :=range pages{
		name:=filepath.Base(page)
		ts, err:=template.New(name).ParseFiles(page)
		fmt.Println(fmt.Sprintf("ts is %v",ts))
		if err!=nil{
			fmt.Println(err)
		}
		matches, err:=filepath.Glob("./templates/*.layout.tmpl")
		if len(matches)>0{
			ts,err=ts.ParseGlob("./templates/*.layout.tmpl")

		}
		cache[name]=ts
	}
	return cache
}