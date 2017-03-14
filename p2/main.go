package main // hello world, the web server
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	data, _ := ioutil.ReadAll(req.Body)
	fmt.Println(data)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
