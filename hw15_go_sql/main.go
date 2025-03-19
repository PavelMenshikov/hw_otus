package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/PavelMenshikov/hw_otus/hw15_go_sql/client"
	"github.com/PavelMenshikov/hw_otus/hw15_go_sql/db"
	"github.com/PavelMenshikov/hw_otus/hw15_go_sql/server"
)

func main() {
	mode := flag.String("mode", "all", "Режим работы: server, client или all")
	addr := flag.String("addr", "localhost", "Адрес сервера")
	port := flag.Int("port", 8080, "Порт сервера")
	method := flag.String("method", "GET", "HTTP метод: GET или POST")
	url := flag.String("url", "", "URL для отправки запроса")
	data := flag.String("data", "", "Данные для POST запроса (если применимо)")
	flag.Parse()

	dbConfig := db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "q1w2e3r4t5",
		DBName:   "online_shop",
		SSLMode:  "disable",
	}
	if err := db.InitDB(dbConfig); err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}

	users, err := db.GetAllUsers()
	if err != nil {
		log.Printf("Ошибка получения пользователей: %v", err)
	} else {
		fmt.Println("Пользователи в БД:")
		for _, u := range users {
			fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
		}
	}

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
		time.Sleep(500 * time.Millisecond)
		if *url == "" {
			*url = "http://" + *addr + ":" + strconv.Itoa(*port)
		}
		client.RunClient(*method, *url, *data)
		select {}
	default:
		log.Fatal("Неверный режим. Используйте 'server', 'client' или 'all'.")
	}
}
