package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var url = "http://localhost:8080"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "Connection failed ::: ", err, "\n")
		os.Exit(1) // request failure, dropped connection 1
	}

	err = handleResp(resp)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	fmt.Println("Exiting ::: Status Code 0")
}

func handleResp(resp *http.Response) error {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to parse response body ::: ", err, "\n")
		os.Exit(2) // failed to
	}

	switch resp.StatusCode {
	case 200:
		fmt.Fprintf(os.Stdout, "Success ::: %s\n", string(body))
	case 429:
		retryAfter := resp.Header.Get("Retry-After")
		// parse as either int, "a while", or a date format

		fmt.Println("RETRY AFTER ::: ", retryAfter)
		var retryAfterDuration time.Duration
		retryAfterDuration, err = time.ParseDuration(fmt.Sprintf("%ss", retryAfter))
		if err != nil {
			retryAfterDuration = 1
		}
		fmt.Printf("Server is busy, Retrying in %d durations\n", retryAfterDuration)
		time.Sleep(retryAfterDuration)
		resp, err = http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "Connection failed ::: ", err)
			return errors.New("connection failed")

		}
		return handleResp(resp)

	case 500:
		fmt.Fprint(os.Stdout, "Internal Server Error")
		return errors.New("internal server error")

	}

	return nil
}
