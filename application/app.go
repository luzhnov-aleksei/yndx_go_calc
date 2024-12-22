package application

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/luzhnov-aleksei/yndx_go_calc/pkg/calculation"
)

type ServerConfig struct {
	Port string
}

func LoadConfigFromEnv() *ServerConfig {
	cfg := new(ServerConfig)
	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	return cfg
}

type WebApp struct {
	config *ServerConfig
}

func NewApp() *WebApp {
	return &WebApp{
		config: LoadConfigFromEnv(),
	}
}

func (app *WebApp) Start() error {
	for {
		log.Println("Enter an arithmetic expression:")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to read expression from console:", err)
			return err
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			log.Println("Application terminated successfully")
			return nil
		}
		result, err := calculation.Calc(input)
		if err != nil {
			log.Printf("Calculation failed for '%s' with error: %v\n", input, err)
		} else {
			log.Printf("Result for '%s' = %f\n", input, result)
		}
	}
}

func containsLetters(s string) bool {
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return true
		}
	}
	return false
}

type RequestPayload struct {
	Expr string `json:"expression"`
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		payload := new(RequestPayload)
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			log.Println("Error parsing request body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "Internal server error"}`)
			return
		}

		if containsLetters(payload.Expr) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "Expression contains invalid characters"}`)
			return
		}

		result, err := calculation.Calc(payload.Expr)
		if err != nil {
			log.Printf("Calculation failed: %v\n", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, `{"error": "Invalid expression"}`)
		} else {
			log.Printf("Calculation successful: result = %f\n", result)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"result": "%f"}`, result)
		}
	default:
		log.Println("Invalid request method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method Not Allowed"}`))
	}
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!\n")
}

func (app WebApp) StartServer() {
	log.Println("Server starting on port", app.config.Port)
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/api/v1/calculate", handleCalculate)

	log.Fatal(http.ListenAndServe(":"+app.config.Port, nil))
}
