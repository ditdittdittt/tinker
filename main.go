package main

import (
	"fmt"
	"io"
	"net/http"
)

var (
	url = "http://localhost:4440/api"
)

func main() {
	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, url+"/44/project/tinker/jobs", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Rundeck-Auth-Token", "CgOYvB5MO5Lul8GYcW61JR02uCE3WyUq")

	resp, _ := client.Do(req)

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
