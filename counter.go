package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}
var counter = 1

//UpdateCounter increment counter and display new value
func UpdateCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()

	fmt.Fprintf(w, strconv.Itoa(counter))
	counter++
	mutex.Unlock()
}
