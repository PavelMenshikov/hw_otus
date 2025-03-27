package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/PavelMenshikov/hw_otus/hw15_go_sql/db"
	"github.com/PavelMenshikov/hw_otus/hw15_go_sql/server"
)

func main() {
	addr := flag.String("addr", "localhost", "Адрес сервера")
	port := flag.Int("port", 8080, "Порт сервера")
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

	log.Printf("Сервер запущен на %s:%d", *addr, *port)
	server.RunServer(fmt.Sprintf("%s:%d", *addr, *port))
}
