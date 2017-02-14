package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

const PORT string = ":8080"
const CERTIFICATE string = "server.crt"
const KEY string = "server.key"

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

func getUserAgent(r *http.Request) string {
	const header string = "User-Agent"
	return r.Header.Get(header)
}

func activityLogHander(w http.ResponseWriter, r *http.Request){
	fmt.Println("Service request from :")
	fmt.Println("\t" + getIpAddress(r))
	fmt.Println("\t" + getUserAgent(r))
}

func main() {

	fileServer := http.FileServer(http.Dir("."))

	http.Handle("/", fileServer)
	//log activity
	http.HandleFunc("/service", activityLogHander)

	//time to serve
	fmt.Println("serving on port " + PORT)

	err := http.ListenAndServeTLS(PORT, CERTIFICATE, KEY, nil)

	if err != nil {
		log.Fatal("ListenAndServer : ", err)
	}
}
