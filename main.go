//  http://giantmachines.tumblr.com/post/49002286919/dealing-with-soap-xml-requests-in-golang

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/soapserver", SoapServer)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// SoapServer accepts and prints out the reponse body
func SoapServer(w http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	//io.WriteString(w, "hello, world!\n")
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("*******************\n")

	//Print the headers
	for k, v := range req.Header {
		log.Println(k, "=", v)
	}

	//Print the body
	fmt.Printf("\n%s\n", b)
}
