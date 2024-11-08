package packagejson

import (
	"encoding/json"
	"fmt"
)

// Define a struct to represent the Name data
type Name struct {
	First  string `json:"first"`
	Middle string `json:"middle"`
	Last   string `json:"last"`
}

// SendJson returns JSON data as a string and an error if any
func SendJson() (string, error) {
	var names []Name

	// Populate the names slice with sample data
	names = append(names, Name{First: "SaiRam", Middle: "SaiRam", Last: "SaiRam"})
	names = append(names, Name{First: "Sai", Middle: "Sai", Last: "Sai"})
	names = append(names, Name{First: "Ram", Middle: "Ram", Last: "Ram"})

	// Convert the names slice to JSON
	jsonData, err := json.Marshal(names)
	if err != nil {
		return "", fmt.Errorf("error converting to JSON: %v", err)
	}

	var name []Name

	err = json.Unmarshal(jsonData, &name)
	if err != nil {
		return "", fmt.Errorf("error converting to JSON: %v", err)
	}
	return string(jsonData), nil
}
