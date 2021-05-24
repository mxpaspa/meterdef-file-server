/*
Serve is a very simple static file server in go
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func returnCode200(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200"))
}

func getMeterdefs(w http.ResponseWriter, r *http.Request){

	// var out []byte

	// entries, err := ioutil.ReadDir(path)
	// if err != nil {
	// 	log.Panicf("failed reading directory: %s", err)
	// }

	// for _,fs := range entries {
	// 	fileName := fmt.Sprintf("%s/%s",path,fs.Name())
	// 	dat, err := ioutil.ReadFile(fileName)
	// 	if err != nil {
	// 		log.Panicf("failed reading file: %s", err)
	// 	}

	// 	out = append(out, dat...)
	// }
	path := "./meterdefs"
	name := r.URL.Path[len("/get/"):]
	fileName := fmt.Sprintf("%s/%s.yaml",path,name)
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading file: %s", err)
	}

	fmt.Println(string(dat))
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(dat))
}

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", "./meterdefs", "the directory of static file to host")

	flag.Parse()

	http.HandleFunc("/get/", getMeterdefs)
	http.HandleFunc("/healthz",returnCode200)
	http.HandleFunc("/readyz",returnCode200)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}



