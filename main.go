package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func playSoundFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	app := "mpg123"
	arg1 := r.URL.Path
	// s := []string{"/Users/matt/go/src/github.com/icecreammatt/sound-server", arg1}
	s := []string{"/root", arg1}
	musicCommand := strings.Join(s, "")
	fmt.Println(musicCommand)
	cmd := exec.Command(app, musicCommand)
	stdout, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
	print(string(stdout))
	fmt.Fprintf(w, "Playing Sound\n") // send data to client side
}

func main() {
	http.HandleFunc("/", playSoundFile)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}