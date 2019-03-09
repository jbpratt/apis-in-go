package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

// function to execute system command and return output
func getCommandOutput(command string, arguments ...string) string {
	// args... unpacks args array into elements
	cmd := exec.Command(command, arguments...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	return out.String()
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("/usr/bin/go", "version"))
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("/bin/cat", params.ByName("name")))
}

func main() {
	// create new router
	router := httprouter.New()
	// GET takes two args, URL path and handler func
	// mapping to methods is possible with httprouter
	router.GET("/api/v1/go-version", goVersion)
	// path var called name used here
	router.GET("/api/vi/show-file/:name", getFileContent)
	// router is then passed to ListenAndServe
	log.Fatal(http.ListenAndServe(":8000", router))
}
