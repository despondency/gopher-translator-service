// Code generated by MockGen. DO NOT EDIT.
// Source: translator_manager.go

// Package translatormock is a generated GoMock package.
package translatormock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// GopherTranslator is a mock of Manager interface.
type GopherTranslator struct {
	ctrl     *gomock.Controller
	recorder *GopherTranslatorMockRecorder
}

// GopherTranslatorMockRecorder is the mock recorder for GopherTranslator.
type GopherTranslatorMockRecorder struct {
	mock *GopherTranslator
}

// NewGopherTranslator creates a new mock instance.
func NewGopherTranslator(ctrl *gomock.Controller) *GopherTranslator {
	mock := &GopherTranslator{ctrl: ctrl}
	mock.recorder = &GopherTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *GopherTranslator) EXPECT() *GopherTranslatorMockRecorder {
	return m.recorder
}

// Translate mocks base method.
func (m *GopherTranslator) Translate(ctx context.Context, word string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", ctx, word)
	ret0, _ := ret[0].(string)
	return ret0
}

// Translate indicates an expected call of Translate.
func (mr *GopherTranslatorMockRecorder) Translate(ctx, word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*GopherTranslator)(nil).Translate), ctx, word)
}

// TranslateSentence mocks base method.
func (m *GopherTranslator) TranslateSentence(ctx context.Context, sentence string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateSentence", ctx, sentence)
	ret0, _ := ret[0].(string)
	return ret0
}

// TranslateSentence indicates an expected call of TranslateSentence.
func (mr *GopherTranslatorMockRecorder) TranslateSentence(ctx, sentence interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateSentence", reflect.TypeOf((*GopherTranslator)(nil).TranslateSentence), ctx, sentence)
}
