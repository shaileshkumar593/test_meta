package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func request() error {
	resp, err := http.Get("http://shailesh.com.me")

	if err != nil {
		return &RequestError{
			StatusCode: 404,
			Err:        errors.New("Server Not Found"),
		}
	}

	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)

	if err != nil {
		return &RequestError{
			StatusCode: 404,
			Err:        err,
		}
	}
	return nil
}

func main() {

	err := request()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Request Completed")
}
