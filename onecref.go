package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

// Структура запроса на запись данных в 1с
type RequestData struct {
	FioUID  string `json:"FioUID" validate:"uuid"`  //UID Сотрудника
	EtapUID string `json:"EtapUID" validate:"uuid"` //UID Этапа производства
	RCUID   string `json:"RCUID" validate:"uuid"`   //UID рабочего центра
	Metrazh string `json:"Metrazh"`                 //Фактически произведенный метраж (преобразование float32 в string)
}

// Структура ответа от 1с
type ResponseData struct {
	RequestData
	Processed bool `json:"Processed"` //Поле результата записи (true - запись удачная)
}

func (s RequestData) validateReq() error {
	if err := validateData(s); err != nil {
		return err
	}

	return nil
}

func get1c(payload any) (map[string]any, error) {

	url := "http://vega64u.mpkabel.ru/erp-copy2/hs/hmi/V1/HMIGet"

	method := "GET"
	// str := payload //fmt.Sprintf("%v", payload)

	pay := strings.NewReader(payload.(string))

	client := &http.Client{}

	req, err := http.NewRequest(method, url, pay)
	if err != nil {
		// fmt.Println(err)
		slog.Error("Ошибка генерации запроса", "Ошибка:", err)
		return nil, err
	}

	// auth := "Basic"
	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Authorization", "Basic "+Config.Token)

	res, err := client.Do(req)
	if err != nil {
		slog.Error("Ошибка обработки запроса", "Ошибка:", err)
		// fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		slog.Info("Тело запроса к 1с обработанное", "Тело: ", string(body))
		// fmt.Println(string(body))
		// return nil, &Error{Code: res.StatusCode, Message: "Ошибка запроса к 1с " + string(body)}
	}

	var result map[string]any
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		slog.Error("Ошибка декодирования", "Ошибка: ", err)
		// fmt.Println("Ошибка декодирования:", err)
	}

	fio, _ := result["ФИО"].(string)
	if strings.Contains(fio, "<Объект не найден>") {
		// Создаем свою ошибку с кодом и текстом
		// apiErr := &Error{
		// 	Code:    404,
		// 	Message: fmt.Sprintf("UID %s отсутствует в базе 1С", result["UID"]),
		// }
		// return nil, apiErr
	}

	return result, nil
}

func write1c(reqData RequestData) (result string, err error) {
	url := Config.URL + Config.WriteAPI
	// url := "http://vega64u.mpkabel.ru/erp-copy2/hs/hmi/V1/HMIPost"
	if err := sendRequest(url, reqData); err != nil {
		slog.Error("Ошибка запроса к 1с", "Ошибка:", err)
		return "Что то пошло не так", err
	}
	return "Надо доделать", nil
}

func sendRequest(url string, reqData RequestData) error {
	if err := reqData.validateReq(); err != nil {
		slog.Error("Ошибка валидации!")
		return err
	}
	// 2. Маршалим в JSON
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return fmt.Errorf("ошибка маршалинга запроса: %v", err)
	}
	// fmt.Println(string(jsonData))
	// Создаём новый POST-запрос
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+Config.Token)

	// Выполняем запрос с помощью клиента (можно использовать http.DefaultClient)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

	// // 3. Отправляем POST-запрос
	// resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	return fmt.Errorf("ошибка отправки запроса: %v", err)
	// }
	// defer resp.Body.Close()

	// 4. Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("получен статус %d, ожидался 200", resp.StatusCode)
	}

	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))

	// 5. Декодируем ответ в структуру
	var respData ResponseData
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	// 6. Выводим результат
	slog.Info("Ответ от 1С", "Тело ответа:", respData)
	// fmt.Printf("Ответ от 1С: %+v\n", respData)
	return nil
}
