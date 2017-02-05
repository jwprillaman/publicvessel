package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strings"
)

var PORT string = ":8080"

func getIpAddress(r *http.Request) string {
	for _, header := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addresses := strings.Split(r.Header.Get(header), ",")
		for i := len(addresses) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addresses[i])

			realIp := net.ParseIP(ip)
			if !realIp.IsGlobalUnicast() {
				continue
			}
			return ip
		}
	}
	return ""
}

func test(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hitting")
	t, _ := template.ParseFiles("assets/index/index.html")
	t.Execute(w, "")

}

func main() {

	fmt.Println("serving on port " + PORT)
	http.HandleFunc("/", test)
	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		log.Fatal("ListenAndServer : ", err)
	}
}
