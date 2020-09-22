package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
type host struct {
	address string
}
var hostNameToIP map[string]host
/*
 * /api/hostname/a.b.com/ip/a.b.c.d
 */
func addHost(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	hostname := vars["hostname"]
	ip := vars["ip"]

	h, ok := hostNameToIP[hostname]
	msg := ""
	if (ok){
		msg = "add host: already exists. " + hostname + " is " + h.address
	}else{
		msg = "add host: " + hostname + " is " + ip
		hostNameToIP[hostname] = host{ip}
	}
	fmt.Println(msg)
	fmt.Fprintf(w, msg)
}
// /api/4
func oneItem(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "api one here for id= " + id)
}
// /api
func apiPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "api here")
}
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Println("/ called")
	fmt.Fprintf(w, "hello")
}
func handleRequests(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/api", apiPage)
	router.HandleFunc("/api/{id}", oneItem)
	router.HandleFunc("/api/hostname/{hostname}/ip/{ip}", addHost)
	log.Fatal(http.ListenAndServe(":4000", router))
}
func main() {
	hostNameToIP = make(map[string]host)
	fmt.Println("Listen on 4000")
	handleRequests()
}
