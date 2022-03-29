package integration_tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gopher-translator-service/integration-tests/helper"
	v1 "gopher-translator-service/internal/api/v1"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestTranslateWord(t *testing.T) {
	container, err := helper.SetupService(context.Background())
	if err != nil {
		panic(err)
	}
	translateWord := &v1.GopherWordRequest{
		EnglishWord: "apple",
	}
	b, err := json.Marshal(translateWord)
	if err != nil {
		t.Errorf(err.Error())
	}
	body := bytes.NewBuffer(b)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(fmt.Sprintf(container.URI+"/word"), "application/json", body)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer func() {
		resp.Body.Close()
	}()
	//Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}
	translatedWord := &v1.GopherWordResponse{}
	err = json.Unmarshal(respBody, translatedWord)
	if err != nil {
		t.Errorf(err.Error())
	}
	if translatedWord.GopherWord != "gapple" {
		t.Errorf("expected %s, got %s", "gapple", translatedWord.GopherWord)
	}
	container.Terminate(context.Background())
}

func TestTranslateSentence(t *testing.T) {
	container, err := helper.SetupService(context.Background())
	if err != nil {
		panic(err)
	}
	translateWord := &v1.GopherSentenceRequest{
		EnglishSentence: "Apples grow on trees.",
	}
	b, err := json.Marshal(translateWord)
	if err != nil {
		t.Errorf(err.Error())
	}
	body := bytes.NewBuffer(b)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(fmt.Sprintf(container.URI+"/sentence"), "application/json", body)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer func() {
		resp.Body.Close()
	}()
	//Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}
	translatedWord := &v1.GopherWordResponse{}
	err = json.Unmarshal(respBody, translatedWord)
	if err != nil {
		t.Errorf(err.Error())
	}
	if translatedWord.GopherWord != "gApples owgrogo gon eestrogo." {
		t.Errorf("expected %s, got %s", "gapple", translatedWord.GopherWord)
	}
	container.Terminate(context.Background())
}
