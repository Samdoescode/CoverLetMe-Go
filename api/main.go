package main 
 
import (
  "fmt"
  "net/http"
	"bytes"
	"io"
	"log"
	"os"
	"context"
"github.com/tmc/langchaingo/llms/openai"
)


func main() {
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}
	prompt := "What would be a good company name for a company that makes colorful socks?"
	completion, err := llm.Call(context.Background(), prompt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(completion)
}
 
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
	fmt.Println(string(responseData))
}


