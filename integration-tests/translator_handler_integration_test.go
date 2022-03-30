package integration_tests

import (
	"context"
	testClientV1 "gopher-translator-service/integration-tests/client/v1"
	"gopher-translator-service/integration-tests/helper"
	v1 "gopher-translator-service/internal/api/v1"
	"testing"
)

func Test_TranslateWord(t *testing.T) {
	container, err := helper.SetupService(context.Background())
	if err != nil {
		panic(err)
	}
	testClient := testClientV1.NewTestClient()
	defer container.Terminate(context.Background())
	var tests = []struct {
		input    *v1.GopherWordRequest
		expected *v1.GopherWordResponse
	}{
		{
			input: &v1.GopherWordRequest{
				EnglishWord: "apple",
			},
			expected: &v1.GopherWordResponse{
				GopherWord: "gapple",
			},
		},
	}
	for _, tt := range tests {
		resp, err := testClient.Translate(container.URI, tt.input)
		if err != nil {
			t.Errorf(err.Error())
		}
		if *resp != *tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, resp)
		}
	}
}

func Test_TranslateSentence(t *testing.T) {
	container, err := helper.SetupService(context.Background())
	if err != nil {
		panic(err)
	}
	testClient := testClientV1.NewTestClient()
	var tests = []struct {
		input    *v1.GopherSentenceRequest
		expected *v1.GopherSentenceResponse
	}{
		{
			input: &v1.GopherSentenceRequest{
				EnglishSentence: "Apples grow on trees.",
			},
			expected: &v1.GopherSentenceResponse{
				GopherSentence: "gApples owgrogo gon eestrogo.",
			},
		},
	}
	for _, tt := range tests {
		resp, err := testClient.TranslateSentence(container.URI, tt.input)
		if err != nil {
			t.Errorf(err.Error())
		}
		if *resp != *tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, resp)
		}
	}
	container.Terminate(context.Background())
}
