package runtime

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func CheckModelStatus(modelName string) (bool, error) {
	url := "http://localhost::11434/api/show"
	payload := map[string]string{"name": modelName}
	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}
