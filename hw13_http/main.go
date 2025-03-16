package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/PavelMenshikov/hw_otus/hw13_http/client"
	"github.com/PavelMenshikov/hw_otus/hw13_http/server"
)

func main() {
	mode := flag.String("mode", "all", "Режим работы: server, client или all")
	addr := flag.String("addr", "localhost", "Адрес сервера")
	port := flag.Int("port", 8080, "Порт сервера")
	method := flag.String("method", "GET", "HTTP метод: GET или POST")
	url := flag.String("url", "", "URL для отправки запроса")
	data := flag.String("data", "", "Данные для POST запроса (если применимо)")
	flag.Parse()

	switch *mode {
	case "server":
		server.RunServer(*addr, *port)
	case "client":
		if *url == "" {
			log.Fatal("Необходимо указать URL для клиента")
		}
		client.RunClient(*method, *url, *data)
	case "all":
		go server.RunServer(*addr, *port)
		// Ждём немного, чтобы сервер успел запуститься
		time.Sleep(500 * time.Millisecond)
		if *url == "" {
			*url = "http://" + *addr + ":" + strconv.Itoa(*port)
		}
		client.RunClient(*method, *url, *data)
	default:
		log.Fatal("Неверный режим. Используйте 'server', 'client' или 'all'.")
	}
}
