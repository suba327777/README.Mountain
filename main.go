package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getHello() string {
	return "hello world"
}

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

func mountainHandler(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	totalCommit, err := fetchGrass(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch GitHub data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"totalCommit": totalCommit})

}

func main() {
	fmt.Println(getHello())

	r := gin.Default()

	r.GET("/mountain", mountainHandler)

	fmt.Println("server is runnning")
	r.Run(":8080")
}
