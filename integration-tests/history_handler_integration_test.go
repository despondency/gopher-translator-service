package integration_tests

import (
	"context"
	testClientV1 "gopher-translator-service/integration-tests/client/v1"
	"gopher-translator-service/integration-tests/helper"
	v1 "gopher-translator-service/internal/api/v1"
	"gopher-translator-service/internal/history"
	"testing"
)

func TestGetTranslationHistory(t *testing.T) {
	container, err := helper.SetupService(context.Background())
	if err != nil {
		panic(err)
	}
	testClient := testClientV1.NewTestClient()
	defer container.Terminate(context.Background())
	inputTranslations := []*v1.GopherWordRequest{
		{
			EnglishWord: "c",
		},
		{
			EnglishWord: "b",
		},
		{
			EnglishWord: "a",
		},
	}
	prepare(t, testClient, container.URI, inputTranslations)
	historyActual, err := testClient.GetTranslationHistory(container.URI)
	expectedHistory := v1.TranslationHistory{
		History: []*history.Entry{
			{
				Word:        "a",
				Translation: "ga",
			},
			{
				Word:        "b",
				Translation: "bogo",
			},
			{
				Word:        "c",
				Translation: "cogo",
			},
		},
	}
	if len(historyActual.History) != len(expectedHistory.History) {
		t.Errorf("expected len of %d for history but was %d", len(expectedHistory.History), len(historyActual.History))
	}
	for i, hist := range historyActual.History {
		if *hist != *expectedHistory.History[i] {
			t.Errorf("histories are not equal, expected %s/%s, actual %s/%s", expectedHistory.History[i].Word, expectedHistory.History[i].Translation, hist.Word, hist.Translation)
		}
	}
}

func prepare(t *testing.T, testClient *testClientV1.TestClient, containerURI string, inputTranslations []*v1.GopherWordRequest) {
	for _, inp := range inputTranslations {
		_, err := testClient.Translate(containerURI, inp)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}
