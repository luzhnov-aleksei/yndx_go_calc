package application

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContainsLetters(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123+456", false},
		{"1a+2b", true},
		{"5*6", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := containsLetters(test.input)
			if got != test.expected {
				t.Errorf("containsLetters(%q) = %v; want %v", test.input, got, test.expected)
			}
		})
	}
}

func TestHandleCalculate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		wantCode   int
		wantBody   string
	}{
		{
			name:       "Valid expression",
			expression: `{"expression": "3+2"}`,
			wantCode:   http.StatusOK,
			wantBody:   `{"result": "5.000000"}`,
		},
		{
			name:       "Invalid expression",
			expression: `{"expression": "3/0"}`,
			wantCode:   http.StatusUnprocessableEntity,
			wantBody:   `{"error": "Invalid expression"}`,
		},
		{
			name:       "Invalid characters in expression",
			expression: `{"expression": "3+2a"}`,
			wantCode:   http.StatusBadRequest,
			wantBody:   `{"error": "Expression contains invalid characters"}`,
		},
		{
			name:       "Invalid JSON",
			expression: `{"expression": "3+2"`,
			wantCode:   http.StatusInternalServerError,
			wantBody:   `{"error": "Internal server error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer([]byte(tt.expression)))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handleCalculate)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantCode)
			}

			if rr.Body.String() != tt.wantBody {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestStartServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handleHello))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("could not send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}
}
