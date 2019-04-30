package service

import (
	"testing"
	"github.com/tomoyane/grant-n-z/server/di"
	"github.com/stretchr/testify/assert"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
)

func TestGenerateJwt(t *testing.T) {
	assert.NotEmpty(t, di.ProviderTokenService.GenerateJwt(username, userUuid, false))
}

func TestParseJwtOk(t *testing.T) {
	token := di.ProviderTokenService.GenerateJwt(username, userUuid, false)
	_, result := di.ProviderTokenService.ParseJwt(token)

	assert.Equal(t, result, true)
}

func TestParseJwtInValid01(t *testing.T) {
	_, result := di.ProviderTokenService.ParseJwt("test")

	assert.Equal(t, result, false)
}

func TestParseJwtInValid02(t *testing.T) {
	_, result := di.ProviderTokenService.ParseJwt(token)

	assert.Equal(t, result, false)
}

func TestGetTokenByUserUuid(t *testing.T) {
	userUuidStr := "52F6228E-9169-4563-ADE2-07ED697B67BA"
	token := di.ProviderTokenService.GetTokenByUserUuid(userUuidStr)

	assert.Equal(t, token.TokenType, "Bearer")
	assert.Equal(t, token.Token, "testToken")
	assert.Equal(t, token.RefreshToken, "testRefreshToken")
	assert.Equal(t, token.UserUuid, userUuid)
}

func TestInsertToken(t *testing.T) {
	token := entity.Token {
		Id: 1,
		TokenType: "Bearer",
		Token: "testToken",
		RefreshToken: "testRefreshToken",
		UserUuid: userUuid,
	}

	assert.Equal(t, token.TokenType, "Bearer")
	assert.Equal(t, token.Token, "testToken")
	assert.Equal(t, token.RefreshToken, "testRefreshToken")
	assert.Equal(t, token.UserUuid, userUuid)
}