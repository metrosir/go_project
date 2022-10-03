package request

import (
	"io/ioutil"
	"net/http"
)

func Q(url string) (string, error) {
	reqs, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer reqs.Body.Close()
	body, err := ioutil.ReadAll(reqs.Body)
	if err != nil {
		return "", err
	}
	res := string(body)
	return res, nil
}
