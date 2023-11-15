package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Manager struct {
	FullName       string `json:"full_name,omitempty"`
	Position       string `json:"position,omitempty"`
	Age            int32  `json:"age,omitempty"`
	YearsInCompany int32  `json:"years_in_company,omitempty"`
}

func EncodeManager(manager *Manager) (io.Reader, error) {
	myManager := Manager{
		FullName:       manager.FullName,
		Age:            manager.Age,
		Position:       manager.Position,
		YearsInCompany: manager.YearsInCompany,
	}

	jsonBytes, err := json.Marshal(myManager)
	if err != nil {
		return nil, err
	}

	jsonString := string(jsonBytes)

	reader := strings.NewReader(jsonString)

	return reader, nil
}

func main() {
	myManager := Manager{
		FullName:       "John Smith",
		Position:       "Manager",
		Age:            40,
		YearsInCompany: 10,
	}

	reader, err := EncodeManager(&myManager)
	if err != nil {
		fmt.Println(err)
	}

	var result string
	buf := make([]byte, 5)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}

		result += string(buf[:n])
	}

	fmt.Println(result)
}
