package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"os"
	//"encoding/json"
)

func listUsers(response http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(response, "Welcome, %s!", req.URL.Path)
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	data, _ := ioutil.ReadFile(currentDir+"/users.json")
	//js, _ := json.Marshal(data)
	fmt.Printf("%s", currentDir)
	fmt.Fprintf(response, "%s" , data)
}

func main() {
	http.HandleFunc("/listUsers", listUsers)
	http.ListenAndServe(":8000", nil)
}