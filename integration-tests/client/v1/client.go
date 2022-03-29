package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	v1 "gopher-translator-service/internal/api/v1"
	"io/ioutil"
	"net/http"
)

type TestClient struct {
}

func NewTestClient() *TestClient {
	return &TestClient{}
}

func (c *TestClient) Translate(uri string, req *v1.GopherWordRequest) (*v1.GopherWordResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(b)
	resp, err := http.Post(fmt.Sprintf(uri+"/v1/word"), "application/json", body)
	if err != nil {
		return nil, err
	}
	defer func() {
		resp.Body.Close()
	}()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(respBody) > 0 {
		translatedWord := &v1.GopherWordResponse{}
		err = json.Unmarshal(respBody, translatedWord)
		if err != nil {
			return nil, err
		}
		return translatedWord, nil
	}
	return nil, nil
}

func (c *TestClient) TranslateSentence(uri string, req *v1.GopherSentenceRequest) (*v1.GopherSentenceResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(b)
	resp, err := http.Post(fmt.Sprintf(uri+"/v1/sentence"), "application/json", body)
	if err != nil {
		return nil, err
	}
	defer func() {
		resp.Body.Close()
	}()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	translatedWord := &v1.GopherSentenceResponse{}
	err = json.Unmarshal(respBody, translatedWord)
	if err != nil {
		return nil, err
	}
	return translatedWord, nil
}
