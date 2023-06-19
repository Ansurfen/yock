package gateway

import "testing"

const (
	jwtDefaultRule = `{"type": "jwt", "name": "default", "token": {"exp":10,"key":"yockd_key","sub":"yock"}}`
	myJwtRule      = `{"type": "jwt", "name": "myjwt"}`
	pwdDefaultRule = `{"type": "pwd", "name": "default", "token": "123456"}`
	myPwd2Rule     = `{"type": "pwd", "name": "mypwd", "token": "123456"}`
)

func TestRules(t *testing.T) {
	gate := New()
	gate.SetRule("root", jwtDefaultRule, myJwtRule, pwdDefaultRule, myPwd2Rule)
	gate.UnsetRule("root", myPwd2Rule)
}
