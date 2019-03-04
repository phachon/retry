package retry

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestRetry_RetryCount(t *testing.T) {

	url := "https://api.github.com/"
	req, _ := http.NewRequest("GET", url, nil)
	var res *http.Response

	err := Retry(func() (err error) {
		res, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		return nil
	}, SetRetryCount(2))
	if err != nil {
		t.Errorf("retry err=%s", err.Error())
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Println(string(body))
}

func TestRetry_RetryIntervalTime(t *testing.T) {

	url := "https://api.github.com/"
	req, _ := http.NewRequest("GET", url, nil)
	var res *http.Response

	err := Retry(func() (err error) {
		res, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		return nil
	}, SetRetryCount(3), SetIntervalTime(3 * time.Second))
	if err != nil {
		t.Errorf("retry err=%s", err.Error())
		return
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Println(string(body))
}

func TestRetry_RetryTimeout(t *testing.T) {

	url := "https://api.github.com/"
	req, _ := http.NewRequest("GET", url, nil)
	var res *http.Response

	err := Retry(func() (err error) {
		res, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		return nil
	}, SetRetryCount(3), SetRetryTimeout(800 * time.Millisecond))
	if err != nil {
		t.Errorf("retry err=%s", err.Error())
		return
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Println(string(body))
}
