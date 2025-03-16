package client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// RunClient выполняет HTTP запрос с указанными параметрами.
func RunClient(method, url, data string) {
	if url == "" {
		log.Fatal("Необходимо указать URL")
	}

	var req *http.Request
	var err error
	if method == "POST" {
		req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}
	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf("Ошибка чтения ответа: %v", err)
	}

	fmt.Println("Ответ от сервера:")
	fmt.Println(string(body))
}
