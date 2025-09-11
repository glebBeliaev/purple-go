package repository

import (
	"encoding/json"
	"os"
)

type Verification struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func LoadVerification(path string) ([]Verification, error) {
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return []Verification{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var verifications []Verification
	err = json.NewDecoder(file).Decode(&verifications)
	if err != nil {
		return []Verification{}, nil
	}
	return verifications, nil
}

func SaveVerifications(path string, verifications []Verification) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(verifications)
}
