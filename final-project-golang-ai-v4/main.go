// PENGERJAAN STANDAR DARI GRADER atau VERSI CLI 1.0
// package main

// import (
// 	"bytes"
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strings"
// )

// type AIModelConnector struct {
// 	Client *http.Client
// }

// type Inputs struct {
// 	Table map[string][]string `json:"table"`
// 	Query string              `json:"query"`
// }

// type Response struct {
// 	Answer      string   `json:"answer"`
// 	Coordinates [][]int  `json:"coordinates"`
// 	Cells       []string `json:"cells"`
// 	Aggregator  string   `json:"aggregator"`
// }

// func CsvToSlice(data string) (map[string][]string, error) {
// 	// TODO: replace this
// 	read := csv.NewReader(strings.NewReader(data))
// 	rows, err := read.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(rows) < 1 {
// 		return nil, fmt.Errorf("CSV data is empty")
// 	}
// 	head := rows[0]
// 	result := make(map[string][]string)
// 	for _, head := range head {
// 		result[head] = []string{}
// 	}
// 	for _, row := range rows[1:] {
// 		for i, value := range row {
// 			result[head[i]] = append(result[head[i]], value)
// 		}
// 	}
// 	return result, nil
// }

// func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
// 	// TODO: replace this
// 	url := "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq"
// 	jsonPayload, err := json.Marshal(payload)
// 	if err != nil {
// 		return Response{}, err
// 	}
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
// 	if err != nil {
// 		return Response{}, err
// 	}
// 	req.Header.Set("Authorization", "Bearer "+token)
// 	req.Header.Set("Content-Type", "application/json")

// 	respon, err := c.Client.Do(req)
// 	if err != nil {
// 		return Response{}, err
// 	}
// 	defer respon.Body.Close()

// 	var response Response
// 	err = json.NewDecoder(respon.Body).Decode(&response)
// 	if err != nil {
// 		return Response{}, err
// 	}
// 	return response, nil
// }

// func main() {
// 	// TODO: answer here
// 	// Kita membuat AIModelConnector baru untuk user
// 	connector := AIModelConnector{
// 		Client: &http.Client{},
// 	}
// 	// Membaca file CSV data-series.csv
// 	csvData, err := ioutil.ReadFile("data-series.csv")
// 	if err != nil {
// 		log.Fatalf("Gagal membaca file CSV: %v", err)
// 	}
// 	// Mengubah data file CSV menjadi slice
// 	table, err := CsvToSlice(string(csvData))
// 	if err != nil {
// 		log.Fatalf("Gagal mengubah file CSV menjadi slice: %v", err)
// 	}
// 	query := "What is the total energy consumption for Refrigerator on 2024-01-01?"
// 	// Buat Payload
// 	payload := Inputs{
// 		Table: table,
// 		Query: query,
// 	}
// 	// Token AI dari Hugging Face Model dari akun saya
// 	TOKEN := "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"
// 	// Mengirim payload ke AI Model, dan mendap
// 	response, err := connector.ConnectAIModel(payload, TOKEN)
// 	if err != nil {
// 		log.Fatalf("Gagal mengirim payload ke model AI: %v", err)
// 	}
// 	// Menampilkan hasil response dari AI Model
// 	fmt.Printf("Answer: %s\n", response.Answer)
// 	fmt.Printf("Coordinates: %v\n", response.Coordinates)
// 	fmt.Printf("Cells: %v\n", response.Cells)
// 	fmt.Printf("Aggregator: %s\n", response.Aggregator)
// }

// VERSI CLI 2.0
// Bisa menambahkan beberapa fitur tambahan seperti:
// 1. Bisa quit seperti ketik perintah "exit"
// 2. Menambahkan fitur untuk menampilkan informasi CSV

// package main

// import (
// 	"bufio"
// 	"bytes"
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/joho/godotenv"
// )

// type AIModelConnector struct {
// 	Client *http.Client
// }

// type Inputs struct {
// 	Table map[string][]string `json:"table"`
// 	Query string              `json:"query"`
// }

// type Response struct {
// 	Answer      string   `json:"answer"`
// 	Coordinates [][]int  `json:"coordinates"`
// 	Cells       []string `json:"cells"`
// 	Aggregator  string   `json:"aggregator"`
// }

// func CsvToSlice(data string) (map[string][]string, error) {
// 	read := csv.NewReader(strings.NewReader(data))
// 	rows, err := read.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(rows) < 1 {
// 		return nil, fmt.Errorf("CSV data is empty")
// 	}

// 	head := rows[0]
// 	result := make(map[string][]string)
// 	for _, head := range head {
// 		result[head] = []string{}
// 	}

// 	for _, row := range rows[1:] {
// 		for i, value := range row {
// 			result[head[i]] = append(result[head[i]], value)
// 		}
// 	}
// 	return result, nil
// }

// func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
// 	url := "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq"
// 	jsonPayload, err := json.Marshal(payload)
// 	if err != nil {
// 		return Response{}, err
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
// 	if err != nil {
// 		return Response{}, err
// 	}
// 	req.Header.Set("Authorization", "Bearer "+token)
// 	req.Header.Set("Content-Type", "application/json")

// 	respon, err := c.Client.Do(req)
// 	if err != nil {
// 		return Response{}, err
// 	}
// 	defer respon.Body.Close()

// 	var response Response
// 	err = json.NewDecoder(respon.Body).Decode(&response)
// 	if err != nil {
// 		return Response{}, err
// 	}
// 	return response, nil
// }

// func ProcessUserInput(query string, table map[string][]string, connector *AIModelConnector, token string) (Response, error) {
// 	payload := Inputs{
// 		Table: table,
// 		Query: query,
// 	}
// 	response, err := connector.ConnectAIModel(payload, token)
// 	if err != nil {
// 		return Response{}, err
// 	}
// 	return response, nil
// }

// func DisplayCsvInfo(data string) {
// 	read := csv.NewReader(strings.NewReader(data))
// 	rows, err := read.ReadAll()
// 	if err != nil {
// 		fmt.Println("Error reading CSV data:", err)
// 		return
// 	}

// 	if len(rows) < 1 {
// 		fmt.Println("CSV data is empty")
// 		return
// 	}

// 	fmt.Println("CSV Columns:", strings.Join(rows[0], ", "))
// 	fmt.Println("Example rows:")
// 	for i, row := range rows[1:] {
// 		if i >= 3 {
// 			break
// 		}
// 		fmt.Println(strings.Join(row, ", "))
// 	}
// }

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println("Error loading .env file:", err)
// 		return
// 	}
// 	token := "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"

// 	csvData, err := os.ReadFile("data-series.csv")
// 	if err != nil {
// 		fmt.Println("Error reading CSV file:", err)
// 		return
// 	}

// 	table, err := CsvToSlice(string(csvData))
// 	if err != nil {
// 		fmt.Println("Error parsing CSV:", err)
// 		return
// 	}

// 	fmt.Println("----------------------------------------------------------------")
// 	fmt.Println("Welcome to Smart Home Energy Management System CLI by BE-9190405")
// 	fmt.Println("----------------------------------------------------------------")
// 	fmt.Println("You can ask questions about the data in the CSV file. For more information, please refer to read the README file.")
// 	fmt.Println("Here are the available columns and some example rows from your CSV data:")
// 	DisplayCsvInfo(string(csvData))

// 	tapasConnector := &AIModelConnector{
// 		Client: &http.Client{},
// 	}

// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {
// 		fmt.Print("\nEnter your question (or type 'exit' to quit): ")
// 		scanner.Scan()
// 		input := scanner.Text()

// 		if strings.ToLower(input) == "exit" {
// 			fmt.Println("You are exiting the CLI system...")
// 			break
// 		}

// 		response, err := ProcessUserInput(input, table, tapasConnector, token)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			continue
// 		}

// 		fmt.Println("Answer:", response.Answer)
// 		fmt.Println("Coordinates:", response.Coordinates)
// 		fmt.Println("Cells:", response.Cells)
// 		fmt.Println("Aggregator:", response.Aggregator)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal("Error reading the standard input:", err)
// 	}
// }

// VERSI CLI 3.0
//Bisa menambahkan beberapa fitur tambahan seperti:
// 1. Bisa quit seperti ketik perintah "exit"
// 2. Menambahkan fitur untuk menampilkan informasi CSV
// 3. Menambahkan fitur untuk verifikasi manusia sebelum mengakses sistem

package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type AIModelConnector struct {
	Client *http.Client
}

type Inputs struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

type Response struct {
	Answer      string   `json:"answer"`
	Coordinates [][]int  `json:"coordinates"`
	Cells       []string `json:"cells"`
	Aggregator  string   `json:"aggregator"`
}

func CsvToSlice(data string) (map[string][]string, error) {
	read := csv.NewReader(strings.NewReader(data))
	rows, err := read.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(rows) < 1 {
		return nil, fmt.Errorf("CSV data is empty")
	}

	head := rows[0]
	result := make(map[string][]string)
	for _, head := range head {
		result[head] = []string{}
	}

	for _, row := range rows[1:] {
		for i, value := range row {
			result[head[i]] = append(result[head[i]], value)
		}
	}
	return result, nil
}

func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
	url := "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq"
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return Response{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return Response{}, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	respon, err := c.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer respon.Body.Close()

	var response Response
	err = json.NewDecoder(respon.Body).Decode(&response)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}

func ProcessUserInput(query string, table map[string][]string, connector *AIModelConnector, token string) (Response, error) {
	payload := Inputs{
		Table: table,
		Query: query,
	}
	response, err := connector.ConnectAIModel(payload, token)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}

func DisplayCsvInfo(data string) {
	read := csv.NewReader(strings.NewReader(data))
	rows, err := read.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV data:", err)
		return
	}

	if len(rows) < 1 {
		fmt.Println("CSV data is empty")
		return
	}

	fmt.Println("CSV Columns:", strings.Join(rows[0], ", "))
	fmt.Println("Example rows:")
	for i, row := range rows[1:] {
		if i >= 3 {
			break
		}
		fmt.Println(strings.Join(row, ", "))
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	token := "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"

	csvData, err := os.ReadFile("data-series.csv")
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	table, err := CsvToSlice(string(csvData))
	if err != nil {
		fmt.Println("Error parsing CSV:", err)
		return
	}

	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Welcome to Smart Home Energy Management System CLI by BE-9190405")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("You can ask questions about the data in the CSV file. For more information, please refer to read the README file.")
	fmt.Println("Here are the available columns and some example rows from your CSV data:")
	DisplayCsvInfo(string(csvData))

	// Simple Human Verification
	rand.Seed(time.Now().UnixNano())
	questions := []struct {
		question string
		answer   int
	}{
		{"What is 3 + 5?", 3 + 5},
		{"What is 7 - 2?", 7 - 2},
		{"What is 6 * 3?", 6 * 3},
		{"What is 10 / 2?", 10 / 2},
		{"What is 9 + 4?", 9 + 4},
		{"What is 8 - 3?", 8 - 3},
		{"What is 5 * 2?", 5 * 2},
		{"What is 12 / 4?", 12 / 4},
		{"What is 15 - 5?", 15 - 5},
		{"What is 4 + 6?", 4 + 6},
	}

	// Select a random question
	question := questions[rand.Intn(len(questions))]
	fmt.Printf("To verify you are human, please answer the following question: %s ", question.question)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer, _ := strconv.Atoi(scanner.Text())

	if answer != question.answer {
		fmt.Println("Verification failed. Exiting the CLI system...")
		return
	}

	fmt.Println("Verification successful. You can now access the system.")

	tapasConnector := &AIModelConnector{
		Client: &http.Client{},
	}

	for {
		fmt.Print("\nEnter your question (or type 'exit' to quit): ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			fmt.Println("You are exiting the CLI system...")
			break
		}

		response, err := ProcessUserInput(input, table, tapasConnector, token)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Answer:", response.Answer)
		fmt.Println("Coordinates:", response.Coordinates)
		fmt.Println("Cells:", response.Cells)
		fmt.Println("Aggregator:", response.Aggregator)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading standard input:", err)
	}
}

/* NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE

Dengan struktur dan data dalam csv ini, berikut adalah beberapa contoh pertanyaan yang dapat Kita ajukan pada program CLI:

1) Total Konsumsi Energi untuk Peralatan pada Tanggal Tertentu:

-"What is the total energy consumption for Refrigerator on 2022-01-01?"
-"What is the total energy consumption for Refrigerator on 2022-01-02?"

2) Konsumsi Energi berdasarkan Waktu:

-"What was the energy consumption of the Refrigerator at 00:00 on 2022-01-01?"
-"What was the energy consumption of the Refrigerator at 01:00 on 2022-01-01?"

3) Status Peralatan pada Waktu Tertentu:

-"What was the status of the Refrigerator at 03:00 on 2022-01-01?"
-"What was the status of the Refrigerator at 04:00 on 2022-01-01?"

Pertanyaan-pertanyaan ini harus sesuai untuk menanyakan data yang diberikan menggunakan program CLI. Kita dapat menjalankan program dengan menggunakan go run main.go dan memasukkan pertanyaan-pertanyaan ini untuk mendapatkan jawaban yang sesuai. ​​​​​​​*/
