package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
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

func GenerateUnicorn() Unicorn {
	fmt.Println("Producing unicorn...!")

	unicornCapabilities := []string{}
	capabilitiesMap := make(map[string]bool)
	name := fmt.Sprintf(
		"%s-%s",
		Adjectives[rand.Intn(len(Adjectives))],
		PetNames[rand.Intn(len(PetNames))],
	)

	for len(unicornCapabilities) < MAX_CAPABILITY {
		capability := Capabilities[rand.Intn(len(Capabilities))]
		if !capabilitiesMap[capability] {
			capabilitiesMap[capability] = true
			unicornCapabilities = append(unicornCapabilities, capability)
		}
	}
	return Unicorn{Name: name, Capabilities: unicornCapabilities}
}

func LoadData() error {
	nameData, err := ReadFileData("docs/petnames.txt")
	if err != nil {
		return fmt.Errorf("please try later, unicorn factory unavailable")
	}
	PetNames = nameData

	adjData, err := ReadFileData("docs/adj.txt")
	if err != nil {
		return fmt.Errorf("please try later, unicorn factory unavailable")
	}
	Adjectives = adjData

	return nil
}
