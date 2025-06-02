package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func ResponseJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func ReadFileData(path string) ([]string, error) {
	var items []string

	file, err := os.Open(path)
	if err != nil {
		return []string{}, fmt.Errorf("failed to load %s: %v", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}
	return items, nil
}
