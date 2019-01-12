package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		currentTime := time.Now().UnixNano() / int64(time.Millisecond)
		filePath := fmt.Sprintf("/tmp/%d.go", currentTime)
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Fprintln(writer, "Error getting body: "+err.Error())
			return
		}
		err = ioutil.WriteFile(filePath, body, 0644)
		if err != nil {
			fmt.Fprintln(writer, "Error saving temp file: "+err.Error())
			return
		}
		cmdOut, err := exec.Command("go", "run", filePath).CombinedOutput()
		if err != nil {
			fmt.Fprintln(writer, "Error running go file: "+err.Error())
			return
		}
		fmt.Fprintln(writer, string(cmdOut))
		go func() {
			err := os.Remove(filePath)
			if err != nil {
				log.Printf("Error removing file %s: %s", filePath, err.Error())
			}
		}()
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}