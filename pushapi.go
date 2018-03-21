package pushapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseURL = "http://push-api.news-host.pw"

// OverrideBaseURL ...
func OverrideBaseURL(url string) {
	baseURL = url
}

// SendPush ...
/*
	push := &pushapi.Push{
		Subscriber: &pushapi.Subscriber{
			Subtool: "fcm",
		},
		Message:  &pushapi.Message{},
		Postback: []pushapi.Postback{},
	}
	b, _ := json.Marshal(push)

	err := pushapi.SendPush(push)

*/
func SendPush(push *Push) error {
	url := fmt.Sprintf(baseURL + "/pushes/send/")

	j, err := json.Marshal(push)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if 202 != resp.StatusCode {
		return fmt.Errorf("%s", body)
	}

	return err
}
