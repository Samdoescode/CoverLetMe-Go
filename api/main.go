package handler
 
import (
  "fmt"
  "net/http"
	"bytes"
	"io"
	"log"
	"os"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
body := []byte(`{"model":"mistral"}`)
	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(body))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Fprintf(w, string(responseData))

}


