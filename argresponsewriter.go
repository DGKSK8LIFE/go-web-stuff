package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ResponseRecorder struct {
	Responses []string
}

func (r *ResponseRecorder) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r.Responses = append(r.Responses, string(body))

	return body, nil
}

func main() {
	recorder := &ResponseRecorder{}
	for _, i := range os.Args[1:] {
		recorder.Get(i)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "http response: %v", recorder.Responses)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

	fmt.Printf("downloaded %d responses\n", len(recorder.Responses))
}
