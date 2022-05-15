package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Sum int `json:"sum"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	// Start server
	srv := &http.Server{
		Addr:         "0.0.0.0:8090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      appRouter(),
	}
	srv.ListenAndServe()
}

// appRouter returns the router for the application
func appRouter() http.Handler {
	rt := http.NewServeMux()
	rt.HandleFunc("/", handleBinaryTree)
	return rt
}

// handleBinaryTree handles the request for the binary tree
func handleBinaryTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var p = map[string]interface{}{}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	var arr = NodeArr{}
	arr = p["tree"].(map[string]interface{})["nodes"].([]interface{})
	rootData := p["tree"].(map[string]interface{})["root"].(string)

	rootNum, err := strconv.Atoi(rootData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	sum, err := arr.maxPathSum(arr.Find(rootNum))

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	result := Response{Sum: sum}
	json.NewEncoder(w).Encode(result)
}
