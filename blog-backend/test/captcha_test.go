package test

import (
	"blog-backend/app/service"
	"context"
	"fmt"
	"testing"
)

func TestCaptcha(t *testing.T) {
	id, b64Str, err := service.LoginService.Captcha(context.TODO())
	if err == nil {
		fmt.Println(id)
		fmt.Println(b64Str)
	}
}
