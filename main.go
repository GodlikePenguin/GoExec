package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		currentTime := time.Now().UnixNano() / int64(time.Millisecond)
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Fprintln(writer, "Error getting body: "+err.Error())
			return
		}
		err = ioutil.WriteFile(fmt.Sprintf("/tmp/%d.go", currentTime), body, 0644)
		if err != nil {
			fmt.Fprintln(writer, "Error saving temp file: "+err.Error())
			return
		}
		cmdOut, err := exec.Command("go", "run", fmt.Sprintf("/tmp/%d.go", currentTime)).CombinedOutput()
		if err != nil {
			fmt.Fprintln(writer, "Error running go file: "+err.Error())
			return
		}
		fmt.Fprintln(writer, string(cmdOut))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}