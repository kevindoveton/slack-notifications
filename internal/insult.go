package internal

import (
	"io/ioutil"
	"net/http"
)

// GetInsult comment
func GetInsult() string {
	resp, err := http.Get("https://evilinsult.com/generate_insult.php?lang=en&type=text")
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
