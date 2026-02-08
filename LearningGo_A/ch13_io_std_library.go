package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const TEMP float32 = 0.7
const TOKENS int = 20
const grokMODEL string = "grok-3"

func grokWords(apiKey string, input io.Reader) (string, error) {
	if apiKey == "" {
		fmt.Println("Error: XAI_API_KEY environment variable is not set")
		return "", errors.New("XAI_API_KEY environment variable is not set")
	}

	url := "https://api.x.ai/v1/chat/completions" // "https://api.x.ai/v1/responses"
	var data []string
	buf := make([]byte, 512)
	for {
		n, err := input.Read(buf)
		for _, b := range buf[:n] {
			data = append(data, string(b))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
	}

	requestBody := map[string]interface{}{
		"model": grokMODEL,
		"messages": []map[string]string{
			{"role": "system", "content": "You are Grok, obey me"},
			{"role": "user", "content": strings.Join(data, " ")},
		},
		//"temperature": TEMP,   // Optional: Adjust for creativity (0-2)
		//"max_tokens":  TOKENS, // Optional: Limit response length
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("Error marshaling request: %v\n", err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: API returned status %d - body: %s\n", resp.StatusCode, string(body))
		return "", fmt.Errorf("API returned %d: %s", resp.StatusCode, string(body))
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error unmarshaling response: %v\n", err)
		return "", err
	}

	// Extract and print the assistant's message
	choices := response["choices"].([]interface{})
	firstChoice := choices[0].(map[string]interface{})
	message := firstChoice["message"].(map[string]interface{})
	content := message["content"].(string)
	fmt.Println("Grok's response:", content)
	return content, nil
}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}

	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func writeGZipFile(filename string, content string) error {
	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()
	gZipWriter := gzip.NewWriter(out)
	defer gZipWriter.Close()
	written, err := gZipWriter.Write([]byte(content))
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes to %s.gz\n", written, filename)
	return nil
}

func buildGZipReader(filename string) (*gzip.Reader, func(), error) {
	f, err := os.Open(filename + ".gz")
	if err != nil {
		return nil, nil, err
	}
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, nil, err
	}
	return gz, func() {
		gz.Close()
		f.Close()
	}, nil
}

func main() {
	apiKey := os.Getenv("XAI_API_KEY")
	fmt.Println("XAI_API_KEY:", apiKey)
	gZipFile := "aGZipFile"
	gZipContent := "Scrabble is a game of words"
	err := writeGZipFile(gZipFile, gZipContent)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	gZipReader, closer, err := buildGZipReader(gZipFile)
	defer closer()

	var stringToCount string
	grokResponse, err := grokWords(apiKey, gZipReader)
	if err != nil {
		fmt.Printf("Error getting grok words: %v\n", err)
		s, err := io.ReadAll(gZipReader)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		stringToCount = string(s)
	} else {
		stringToCount = grokResponse
	}
	// fmt.Println(grokResponse)
	letterCount, err := countLetters(strings.NewReader(stringToCount))
	if err != nil {
		fmt.Printf("Error counting letters: %v\n", err)
		return
	}
	for letter, count := range letterCount {
		fmt.Println("Letter:", letter, "Count:", count)
	}
}
