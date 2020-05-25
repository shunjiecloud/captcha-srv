package services

import (
	"context"

	"github.com/dchest/captcha"
	"github.com/shunjiecloud-proto/captcha/proto"
)

type CaptchaService struct{}

func (*CaptchaService) CaptchaVerfify(ctx context.Context, in *proto.CaptchaVerfifyRequest, out *proto.CaptchaVerfifyResponse) error {
	out.Result = captcha.VerifyString(in.CaptchaId, in.Solution)
	//out.Result = true
	return nil
}
