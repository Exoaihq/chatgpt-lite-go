package main

import (
	"fmt"
	"net/http"
	"os"
)

var config = struct {
	Runtime string
}{
	Runtime: "edge",
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling request: %s\n", r.URL)

	azureOpenAIKey := false
	openAIKey := false

	if apiKey := os.Getenv("AZURE_OPENAI_API_KEY"); apiKey != "" {
		azureOpenAIKey = true
	}

	if apiKey := os.Getenv("OPENAI_API_KEY"); apiKey != "" {
		openAIKey = true
	}

	if !azureOpenAIKey && !openAIKey {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "OpenAI key is empty")
		return
	}

	fmt.Fprint(w, "Ok")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}