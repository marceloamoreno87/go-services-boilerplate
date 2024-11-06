package whatsmeow

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

type WhatsMeowDB struct {
	DSN string
}

func NewWhatsMeowDB() *WhatsMeowDB {

	dsnPostgres := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_NAME"),
		os.Getenv("POSTGRES_SSL_MODE"),
		os.Getenv("POSTGRES_SCHEMA"),
	)

	return &WhatsMeowDB{
		DSN: dsnPostgres,
	}
}

func (wm *WhatsMeowDB) DbSqlConnect() (postgresDb *sqlstore.Container) {

	db, err := sql.Open("postgres", wm.DSN)
	if err != nil {
		log.Error().Msg("Error trying to connect to database: " + err.Error())
		return nil
	}
	defer db.Close()

	postgresDb, err = sqlstore.New("postgres", wm.DSN, nil)
	if err != nil {
		log.Error().Msg("Err to create container with DB" + err.Error())
		return nil
	}

	return
}

type WhatsMeow struct {
	Url string
}

type WhatsMeowMessage struct {
	PhoneNumber      string `json:"phone_number"`
	Message          string `json:"message"`
	NotificationType string `json:"notification_type"`
}

func NewWhatsMeow() *WhatsMeow {
	return &WhatsMeow{
		Url: os.Getenv("WHATSMEOW_URL"),
	}
}

func (wm *WhatsMeow) SendMessage(message WhatsMeowMessage) error {
	// Converte a mensagem para JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("Erro ao converter a mensagem para JSON: %v\n", err)
		return err
	}

	// Cria uma requisição HTTP POST
	req, err := http.NewRequest("POST", wm.Url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Erro ao criar a requisição HTTP: %v\n", err)
		return err
	}

	// Adiciona o cabeçalho de autorização com o JWT
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Secret", os.Getenv("GO_ZAP_SECRET"))

	// Envia a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro ao enviar a requisição HTTP: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	// Verifica o status da resposta
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Erro na resposta HTTP: %v\n", resp.Status)
		return fmt.Errorf("erro na resposta HTTP: %v", resp.Status)
	}

	return nil
}
