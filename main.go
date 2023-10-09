package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const githubAPIURL = "https://api.github.com/user"

type GithubResponse struct {
	Message string `json:"message"`
	Login   string `json:"login"`
}

func checkTokenValidity(token string) {
	req, err := http.NewRequest("GET", githubAPIURL, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating request:", err)
		return
	}

	authValue := base64.StdEncoding.EncodeToString([]byte("user:" + token))
	req.Header.Set("Authorization", "Basic "+authValue)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading response:", err)
		return
	}

	var response GithubResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing JSON:", err)
		return
	}

	if response.Login != "" {
		fmt.Printf("Token %s: Valid (User: %s)\n", token, response.Login)
	} else {
		fmt.Printf("Token %s: Invalid\n", token)
	}
}

func customUsage() {
	fmt.Println("Usage: ")
	fmt.Println("ghchk checks the validity of GitHub access tokens.")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Println("  Check a single token: ghchk -token YOUR_ACCESS_TOKEN")
	fmt.Println("  Check tokens from a file: ghchk -file tokens.txt")
	fmt.Println("  Check tokens via stdin: cat tokens.txt | ghchk")
}

func main() {
	tokenPtr := flag.String("token", "", "Provide a single GitHub access token.")
	filePtr := flag.String("file", "", "Provide a file path containing GitHub access tokens, one per line.")
	flag.Usage = customUsage
	flag.Parse()

	if *tokenPtr != "" {
		checkTokenValidity(*tokenPtr)
	} else if *filePtr != "" {
		file, err := os.Open(*filePtr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			token := strings.TrimSpace(scanner.Text())
			if token != "" {
				checkTokenValidity(token)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from file:", err)
		}
	} else if fileInfo, _ := os.Stdin.Stat(); fileInfo.Mode()&os.ModeCharDevice == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			token := strings.TrimSpace(scanner.Text())
			if token != "" {
				checkTokenValidity(token)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		}
	} else {
		flag.Usage()
	}
}
