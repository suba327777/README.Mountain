package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
)

func fetchGrass(username string) (string, error) {

	apiURL := "https://api.github.com/graphql"

	err := godotenv.Load(".env")

	token := os.Getenv("GITHUB_TOKEN")

	query := fmt.Sprintf(`query {
		user(login: "%s") {
			contributionsCollection {
				totalCommitContributions
			}
		}
	}`, username)

	requestData := map[string]string{
		"query": query,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshalling request data:", err)
		return "", nil
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	fmt.Println("Response:", string(respBody))

	return string(respBody), nil

}
