package request

import (
	"io/ioutil"
	"net/http"
)

func Q(url string) (string, error) {
	reqs, err := http.Get(url)
	defer reqs.Body.Close()
	body, err := ioutil.ReadAll(reqs.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
