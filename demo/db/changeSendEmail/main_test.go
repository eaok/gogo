package main

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestBeforeUpdate(t *testing.T) {
	e := &email.Email{
		From: "wang <kcoewoys@aliyun.com>",
		To: []string{"kcoewoys@qq.com"},
		Subject: "go email test",
		Text: []byte("Text Body is, of course, supported!"),
	}

	auth := smtp.PlainAuth("", "kcoewoys@aliyun.com", "xxx", "smtp.aliyun.com")
	err := e.Send("smtp.aliyun.com:25", auth)
	if err != nil {
		t.Errorf(err.Error())
	}
}
