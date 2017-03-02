package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type jobIDResponse struct {
	JobID string
}

var data = `
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">    
 <SOAP-ENV:Header/>    
  <SOAP-ENV:Body>       
    <ns10:JobIDResponse xmlns:ns10="urn:com:ssn:schema:service:v2.1:JobManager.xsd">
      <ns10:JobID>4686</ns10:JobID>       
    </ns10:JobIDResponse>    
  </SOAP-ENV:Body> 
</SOAP-ENV:Envelope>`

func main() {
	fmt.Printf("Endpoint listening for http requests here:\n")
	fmt.Printf("http://localhost:8081/soapserver\n")
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
		log.Println(k, ": ", v)
	}

	//Print the body
	fmt.Printf("\n%s\n", b)

	//Render some xml
	// response := jobIDResponse{"123456"}
	// x, err := xml.MarshalIndent(response, "", "  ")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	var soapresponse = []byte(data)

	w.Header().Set("Content-Type", "text/xml")
	w.Write(soapresponse)
}
