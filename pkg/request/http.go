package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Q(url string) map[string]interface{} {
	ch := make(chan string, 100)
	errCh := make(chan error, 100)
	go func() {
		reqs, err := http.Get(url)
		if err != nil {
			errCh <- err
			return
		}
		defer reqs.Body.Close()
		body, err := ioutil.ReadAll(reqs.Body)
		if err != nil {
			errCh <- err
		}
		ch <- string(body)
	}()
	var res map[string]interface{}
	go func() {
		var temp []string
		for v := range ch {
			temp = append(temp, v)
		}
		res["res"] = temp

		var errTemp []error
		for err := range errCh {
			errTemp = append(errTemp, err)
		}
		res["error"] = errTemp
	}()
	fmt.Println("res", res)
	return res
}
