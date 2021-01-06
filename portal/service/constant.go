package service

import verpb "videoWeb/verify-service/proto"

var strategy = map[string]verpb.VerifyTargetStrategy{
	"phone": verpb.VerifyTargetStrategy_PHONE,
	"email": verpb.VerifyTargetStrategy_EMAIL,
}
