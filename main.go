// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	r := mux.NewRouter()
	
	rooms := []Room{}
	rooms = append(rooms, CreateRoom(1, "all"))

	r.HandleFunc("/", serveHome)
	r.HandleFunc("/ws/{roomId}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.RawQuery)
		vars := mux.Vars(r)
		enterChatroom(w, r, vars["roomId"])
	}).Methods("GET")
	
	r.HandleFunc("/room", CreateRoomHandler).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
