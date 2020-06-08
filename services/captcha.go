package services

import (
	"context"

	"github.com/dchest/captcha"
	merr "github.com/micro/go-micro/v2/errors"
	"github.com/shunjiecloud-proto/captcha/proto"
	"github.com/shunjiecloud/errors"
)

type CaptchaService struct{}

func (*CaptchaService) CaptchaId(ctx context.Context, in *proto.CaptchaIdRequest, out *proto.CaptchaIdResponse) error {
	id := captcha.New()
	out.CaptchaId = id
	return nil
}

func (*CaptchaService) CaptchaVerfify(ctx context.Context, in *proto.CaptchaVerfifyRequest, out *proto.CaptchaVerfifyResponse) error {
	isOk := captcha.VerifyString(in.CaptchaId, in.Solution)
	if isOk == false {
		return errors.New(merr.BadRequest("captcha verfify failed", "input.CaptchaId:%v, input.Solution:%v", in.CaptchaId, in.Solution))
	}
	return nil
}
