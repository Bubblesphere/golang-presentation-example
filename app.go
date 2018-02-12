package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"os"

	"github.com/nu7hatch/gouuid"
)

var tasks map[string]string

func handle() {
	http.HandleFunc("/task/create", taskCreate)
	http.HandleFunc("/task/retrieve", taskRetrieve)
	http.HandleFunc("/task/list", taskRetrieveAll)
	http.HandleFunc("/task/delete", taskDelete)
}

func getBodyString(r *http.Request) string {

	var bodyBytes []byte
	var err error
	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Printf(err.Error())
		}
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return string(bodyBytes)
}

func allowCORS(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func allowMethod(r *http.Request, w http.ResponseWriter, m string) (err error) {
	err = nil
	if r.Method != m {
		err = errors.New("Method not allowed")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return
}

func taskCreate(w http.ResponseWriter, r *http.Request) {
	if allowMethod(r, w, "POST") != nil {
		return
	}
	allowCORS(w)

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	u, _ := uuid.NewV4()
	tasks[u.String()] = getBodyString(r)

	log.Printf("Create %s - %s\n", u.String(), tasks[u.String()])
	fmt.Fprintln(w, u.String())
}

func taskDelete(w http.ResponseWriter, r *http.Request) {
	if allowMethod(r, w, "POST") != nil {
		return
	}
	allowCORS(w)

	u := getBodyString(r)
	delete(tasks, u)

	log.Printf("Delete %s\n", u)
}

func taskRetrieve(w http.ResponseWriter, r *http.Request) {
	if allowMethod(r, w, "POST") != nil {
		return
	}
	allowCORS(w)

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	var u = getBodyString(r)
	log.Printf("Retrieve %s\n", u)
	fmt.Fprintln(w, tasks[u])
}

func taskRetrieveAll(w http.ResponseWriter, r *http.Request) {
	if allowMethod(r, w, "POST") != nil {
		return
	}
	allowCORS(w)

	jsonString, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Retrieve all\n")
	fmt.Fprintln(w, string(jsonString))
}

func main() {
	 //create your file with desired read/write permissions
	 f, err := os.OpenFile("log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	 if err != nil {
			 log.Fatal(err)
	 }   

	 //defer to close when you're done with it, not because you think it's idiomatic!
	 defer f.Close()

	 //set output of logs to f
	 log.SetOutput(f)

	tasks = make(map[string]string)
	handle()
	http.ListenAndServe(":8000", nil)
}
