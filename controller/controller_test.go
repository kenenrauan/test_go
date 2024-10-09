package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/model"
	"project/views"
	"strings"
	"testing"
)

// Mock для базы данных
func init() {
	model.Connect() // подключение к тестовой БД или использование mock
}

func TestCreateEmployee(t *testing.T) {
	payload := views.PostRequest{
		Fullname: "Тест Тестов",
		Phone:    "123456789",
		City:     "Алматы",
	}

	// Подготавливаем данные запроса
	jsonPayload, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(create())
	handler.ServeHTTP(rr, req)

	// Проверяем код ответа
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Неверный код ответа: ожидался %v, получен %v", http.StatusCreated, status)
	}

	// Десериализуем ответ в структуру
	var response views.PostRequest
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Не удалось декодировать ответ: %v", err)
	}

	// Проверяем поля, кроме id
	if response.Fullname != payload.Fullname {
		t.Errorf("Неожиданное имя: ожидалось %v, получено %v", payload.Fullname, response.Fullname)
	}
	if response.Phone != payload.Phone {
		t.Errorf("Неожиданный номер телефона: ожидался %v, получен %v", payload.Phone, response.Phone)
	}
	if response.City != payload.City {
		t.Errorf("Неожиданный город: ожидался %v, получен %v", payload.City, response.City)
	}
}

// Тест чтения всех сотрудников
func TestReadAllEmployees(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(readdelete())

	// Запускаем тест
	handler.ServeHTTP(rr, req)

	// Проверяем код ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный код ответа: ожидался %v, получен %v", http.StatusOK, status)
	}
}

// Тест удаления сотрудника
func TestDeleteEmployee(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/?name=Тест Тестов", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(readdelete())

	// Запускаем тест
	handler.ServeHTTP(rr, req)

	// Проверяем код ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный код ответа: ожидался %v, получен %v", http.StatusOK, status)
	}

	// Проверяем содержимое ответа
	expected := `{"Status":"Item deleted"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("Неожиданный ответ: ожидался %v, получен %v", expected, rr.Body.String())
	}
}
