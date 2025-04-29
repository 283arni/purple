package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"struct/bins"
	"time"
)

var key = "$2a$10$YXpconxx2hWndtNuufd1xuOwIMuAz/24MLVtybz5fhZRe5KtSaZX2"

func PostBin(name string) string {
	bin, _ := json.Marshal(map[string]string{
		"name": name,
	})

	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b/", bytes.NewBuffer(bin))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", key)

	if err != nil {
		fmt.Println(err.Error())
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(res.Body)

	defer res.Body.Close()

	return string(body)
}

func UpdateBin(bin bins.Bin) string {

	bin.Metadata.CreatedAt = time.Now()

	binJson, _ := json.Marshal(bin)

	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b/", bytes.NewBuffer(binJson))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", key)

	if err != nil {
		fmt.Println(err.Error())
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(res.Body)

	defer res.Body.Close()
	fmt.Println(string(body))
	return string(body)
}
