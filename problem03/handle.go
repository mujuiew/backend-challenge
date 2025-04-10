package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var Url = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

type PieFireDire struct {
	Beef map[string]int `json:"beef"`
}

func main() {
	http.HandleFunc("/beef/summary", GetBeef)

	fmt.Println("Start API ที่ http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func GetBeef(w http.ResponseWriter, r *http.Request) {
	var mapCount = make(map[string]int)
	var pieFireDire PieFireDire
	var err error

	body, errgetDataList := getDataList()
	if errgetDataList != nil {
		err = errgetDataList
	}

	wordsList := strings.Fields(strings.ToLower(body))

	for _, word := range wordsList {
		mapCount[word]++
	}

	pieFireDire.Beef = mapCount

	writeResponse(w, &pieFireDire, err)
}

func getDataList() (string, error) {

	resp, err := http.Get(Url)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการเรียก API:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการอ่านข้อมูล:", err)
		return "", err
	}
	return regexp.MustCompile(`[^a-zA-Z\-]+`).ReplaceAllString(string(body), " "), nil
}

func writeResponse(w http.ResponseWriter, obj interface{}, errProcess error) {
	output := ""

	b, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, "error marshal json", http.StatusBadRequest)
	} else {
		output = string(b)
	}

	if errProcess != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, output)
	}

}
