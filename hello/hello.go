package hello

import (
    "html/template"
    "net/http"
    "fmt"
    "appengine"
    "appengine/datastore"
)

func init() {
    http.HandleFunc("/", index)
    http.HandleFunc("/new", new_person)
    http.HandleFunc("/add_person", add_person)
}

type EmployeeList struct{
	Employees 	[]Employee
}

type Employee struct{
	First string
	Last string
}

func index(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	q := datastore.NewQuery("employee")

	var employees []Employee
	_ , err := q.GetAll(c, &employees)
	fmt.Println(employees)

	//x["employees"] = employees
	if err != nil{
		fmt.Println(err)
	}
	displayTemplate("people_list", "people_list.html", w, employees)
}

func new_person(w http.ResponseWriter, r *http.Request){
	person := &Employee{First:"",Last:""}

	displayTemplate("new_person", "new_person.html", w, person)
}

func add_person(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w, r, "/new", http.StatusBadRequest)
	}
	c := appengine.NewContext(r)

	e1 := Employee{
		First: 	r.FormValue("first"),
		Last: 	r.FormValue("last"),
	}

	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "employee", nil), &e1)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println("Stored the Employee named ", e1.First)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func displayTemplate(name string, filename string, w http.ResponseWriter, x interface{}){
	t, err := template.New(name).ParseFiles("tmpl/"+filename)
	
	err = t.ExecuteTemplate(w, filename, x)
	if err != nil {
		fmt.Println(err)
	}
}