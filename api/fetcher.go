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

	// リクエストデータをマップに設定
	requestData := map[string]string{
		"query": query,
	}

	// リクエストデータをJSON形式に変換
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshalling request data:", err)
		return "", nil
	}

	// HTTP POSTリクエストの作成
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return "", err
	}

	// リクエストヘッダーの設定
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// HTTPクライアントの生成
	client := &http.Client{}

	// HTTPリクエストの送信
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// レスポンスの読み取り
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	// レスポンスの出力
	fmt.Println("Response:", string(respBody))

	return string(respBody), nil

}
