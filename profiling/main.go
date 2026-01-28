package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	_ "net/http/pprof"

	"github.com/tanricko-ajaib/workshop/profiling/sorting"
)

func HandleSort(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	qNum := r.Form.Get("num")
	num, err := strconv.Atoi(qNum)
	if err != nil || num <= 0 {
		http.Error(w, "Invalid 'num' parameter", http.StatusBadRequest)
		return
	}

	arr := make([]int, num)
	for i := range arr {
		arr[i] = num - i // Fill with descending numbers for worst-case
	}

	sorting.SortInts(arr)
	w.WriteHeader(http.StatusOK)

	for _, v := range arr {
		w.Write([]byte(strconv.Itoa(v) + "\n"))
	}
}

func main() {
	h := http.DefaultServeMux
	h.HandleFunc("/sort", HandleSort)

	log.Println("Starting http server on :28126")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		err := http.ListenAndServe(":28126", h)
		if err != nil {
			panic(err)
		}
	}()

	<-sig
	log.Println("Shutting down http server...")
}
