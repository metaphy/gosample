/*
curl -v --http1.0 http://localhost:18888/greeting
curl --http1.0 -H "X-Test: Hello" http://localhost:18888
curl -v --http1.0 -A "Mozilla/65.0" http://localhost:18888
curl -d "{'hello':''world'}" -H "Content-Type: application/json" http://localhost:18888
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler (w http.ResponseWriter, r *http.Request) {
    dump, err := httputil.DumpRequest(r, true)
    if err != nil  {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return 
    }

    fmt.Println(string(dump))
    fmt.Fprintf(w, "<html><body>Hello</body></html>" )
}


func main() {
    var httpServer http.Server
    http.HandleFunc("/", handler)
    log.Println("start http listening: 18888")
    httpServer.Addr = ":18888"
    log.Println(httpServer.ListenAndServe())
}
