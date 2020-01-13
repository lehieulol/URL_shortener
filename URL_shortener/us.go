package main

import(
	"net/http"
	"log"
	"io/ioutil"
	"html/template"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Path[len("/goto/"):];
	dest, err := ioutil.ReadFile("URL\\"+source+".txt");
	if(err != nil){
		http.Redirect(w, r, "/edit/"+source, http.StatusFound);
		return
	}else{
		http.Redirect(w, r, string(dest), http.StatusFound);
	}
}
func add(w http.ResponseWriter, r *http.Request){
	source := r.URL.Path[len("/edit/"):];
	t, _ := template.ParseFiles("edit.html");
	t.Execute(w, source)
}
func save(w http.ResponseWriter, r *http.Request){
	source := r.URL.Path[len("/save/"):];
	a := r.FormValue("body");
	ioutil.WriteFile("URL\\"+source+".txt",[]byte(a),0600);
	http.Redirect(w,r,"/goto/"+source,http.StatusFound);
}
func main() {
	http.HandleFunc("/edit/", add);
	http.HandleFunc("/goto/", redirect);
	http.HandleFunc("/save/", save);
    log.Fatal(http.ListenAndServe("", nil));
}