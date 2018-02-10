package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

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

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	u, _ := uuid.NewV4()
	tasks[u.String()] = getBodyString(r)

	fmt.Fprintln(w, u.String())
}

func taskDelete(w http.ResponseWriter, r *http.Request) {
	if allowMethod(r, w, "DELETE") != nil {
		return
	}

	u := getBodyString(r)
	delete(tasks, u)
}

func taskRetrieve(w http.ResponseWriter, r *http.Request) {
	if allowMethod(r, w, "POST") != nil {
		return
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	var u = getBodyString(r)
	fmt.Printf(u)
	fmt.Fprintln(w, tasks[u])
}

func taskRetrieveAll(w http.ResponseWriter, r *http.Request) {
	if allowMethod(r, w, "POST") != nil {
		return
	}

	jsonString, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, string(jsonString))
}

func main() {
	tasks = make(map[string]string)
	handle()
	http.ListenAndServe(":8000", nil)
}
