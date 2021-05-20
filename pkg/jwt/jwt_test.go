package jwt

import (
	"fmt"
	config2 "opencodes/pkg/config"
	"testing"
)

func TestCreateToken(t *testing.T) {
	config2.Setup("./../../config/settings.dev.yml")
	fmt.Println(config2.JwtConfig.Secret, "----")
	tokens, _ := CreateToken(61)

	t.Log(tokens)
}

func TestParseToken(t *testing.T) {
	config2.Setup("./../config/settings.dev.yml")
	token, _ := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNTk1MzI0MzQ1LCJpYXQiOjE1OTUzMjA3NDV9.nBPBid_vckRKhxLgsTP67V-kY5J_giGrTpLrcLVDQGs")
	t.Log(token.Id)

}
