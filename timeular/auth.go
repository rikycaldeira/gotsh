package timeular

import (
	"fmt"

	"github.com/rikycaldeira/gotsh/common"
	"github.com/rikycaldeira/gotsh/http"
)

func Login(apiKey string, apiSecret string) Auth {
	url := TIMEULAR_API + "/developer/sign-in"
	authRequest := AuthRequest{
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
	auth := Auth{}

	http.Post(url, authRequest, &auth, make(map[string]string))
	return auth
}

func Logout(token string) {
	url := TIMEULAR_API + "/developer/logout"

	http.Post(url, common.Empty{}, common.Empty{}, map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
}
