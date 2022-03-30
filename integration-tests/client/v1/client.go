package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	v1 "gopher-translator-service/internal/api/v1"
	"io/ioutil"
	"net/http"
)

type TestClient struct {
	client *http.Client
}

func NewTestClient() *TestClient {
	return &TestClient{
		client: &http.Client{},
	}
}

func (c *TestClient) Translate(uri string, req *v1.GopherWordRequest) (*v1.GopherWordResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(b)
	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf(uri+"/v1/word"), body)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpReq.Close = true
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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
	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf(uri+"/v1/sentence"), body)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpReq.Close = true
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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

func (c *TestClient) GetTranslationHistory(uri string) (*v1.TranslationHistory, error) {
	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf(uri+"/v1/history"), nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	httpReq.Close = true
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	translationHistory := &v1.TranslationHistory{}
	err = json.Unmarshal(respBody, translationHistory)
	if err != nil {
		return nil, err
	}
	return translationHistory, nil
}
