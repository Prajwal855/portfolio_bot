package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type telegramUpdateResponse struct {
	OK     bool             `json:"ok"`
	Result []telegramUpdate `json:"result"`
}

type telegramUpdate struct {
	UpdateID int             `json:"update_id"`
	Message  *telegramMessage `json:"message"`
}

type telegramMessage struct {
	Text string       `json:"text"`
	Chat telegramChat `json:"chat"`
}

type telegramChat struct {
	ID int64 `json:"id"`
}

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is required")
	}

	appMode := getEnv("APP_MODE", "polling")
	port := getEnv("PORT", "8080")

	api := fmt.Sprintf("https://api.telegram.org/bot%s", token)

	if appMode == "webhook" {
		startWebhookServer(api, port)
		return
	}

	startPolling(api)
}

func startPolling(api string) {
	log.Println("running in polling mode")
	offset := 0

	for {
		updates, err := getUpdates(api, offset)
		if err != nil {
			log.Printf("getUpdates failed: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}

		for _, u := range updates {
			offset = u.UpdateID + 1
			if u.Message == nil || u.Message.Text == "" {
				continue
			}

			msg := responseFor(u.Message.Text)
			if err := sendMessage(api, u.Message.Chat.ID, msg); err != nil {
				log.Printf("sendMessage failed: %v", err)
			}
		}
	}
}

func startWebhookServer(api, port string) {
	log.Printf("running in webhook mode on :%s", port)
	secret := os.Getenv("WEBHOOK_SECRET")

	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		if secret != "" {
			if r.Header.Get("X-Telegram-Bot-Api-Secret-Token") != secret {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
		}

		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		var update telegramUpdate
		if err := json.Unmarshal(body, &update); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		if update.Message != nil && update.Message.Text != "" {
			msg := responseFor(update.Message.Text)
			if err := sendMessage(api, update.Message.Chat.ID, msg); err != nil {
				log.Printf("sendMessage failed: %v", err)
			}
		}

		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func getUpdates(api string, offset int) ([]telegramUpdate, error) {
	url := api + "/getUpdates?timeout=25&offset=" + strconv.Itoa(offset)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("getUpdates status %d: %s", resp.StatusCode, string(body))
	}

	var out telegramUpdateResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	if !out.OK {
		return nil, errors.New("telegram returned ok=false")
	}

	return out.Result, nil
}

func sendMessage(api string, chatID int64, text string) error {
	payload := map[string]any{
		"chat_id": chatID,
		"text":    text,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(api+"/sendMessage", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("sendMessage status %d: %s", resp.StatusCode, string(respBody))
	}

	return nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
