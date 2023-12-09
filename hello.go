// main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

type RequestBody struct {
	ToSort [][]int `json:"to_sort"`
}

type Response struct {
	SortedArrays [][]int       `json:"sorted_arrays"`
	TimeTaken    time.Duration `json:"time_taken"`
}

func processSequential(numbers [][]int) Response {
	startTime := time.Now()
	sorted := make([][]int, len(numbers))
	for i, arr := range numbers {
		sorted[i] = make([]int, len(arr))
		copy(sorted[i], arr)
		sort.Ints(sorted[i])
	}
	elapsed := time.Since(startTime)

	return Response{
		SortedArrays: sorted,
		TimeTaken:    elapsed,
	}
}

func processConcurrent(numbers [][]int) Response {
	startTime := time.Now()

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	sorted := make([][]int, len(numbers))

	for i, arr := range numbers {
		go func(i int, arr []int) {
			defer wg.Done()
			sorted[i] = make([]int, len(arr))
			copy(sorted[i], arr)
			sort.Ints(sorted[i])
		}(i, arr)
	}

	wg.Wait()
	elapsed := time.Since(startTime).Round(time.Nanosecond)

	return Response{
		SortedArrays: sorted,
		TimeTaken:    elapsed,
	}
}

func processSingleHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := processSequential(requestBody.ToSort)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func processConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := processConcurrent(requestBody.ToSort)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/process-single", processSingleHandler)
	http.HandleFunc("/process-concurrent", processConcurrentHandler)

	port := 8000
	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
