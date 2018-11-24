package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type mywriter struct{}

func (w mywriter) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}

func main() {
	doit(func(resp *http.Response) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	})
	doit(func(resp *http.Response) {
		io.Copy(os.Stdout, resp.Body)
	})
	doit(func(resp *http.Response) {
		io.Copy(mywriter{}, resp.Body)
	})

}

func doit(h func(*http.Response)) {
	resp, err := http.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
	h(resp)
}
