package cli

import (
	"encoding/json"
	"fmt"
	"net/http"

	goinvestec "github.com/pjlast/go-investec"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func NewClient() *goinvestec.Client {
	tk := viper.GetString("token")
	var authToken oauth2.Token
	if err := json.Unmarshal([]byte(tk), &authToken); err != nil {
		fmt.Println("Could not read auth token from config")
	}
	auther := &goinvestec.Authenticator{
		BaseURL:    viper.GetString("apiURL"),
		ClientID:   viper.GetString("clientID"),
		Secret:     viper.GetString("secret"),
		APIKey:     viper.GetString("apiKey"),
		HTTPClient: &http.Client{},
		OAuthToken: &authToken,
		PostRefreshHook: func(t *oauth2.Token) {
			b, err := json.Marshal(t)
			if err != nil {
				fmt.Println("Could not cache auth token", err)
			}
			viper.Set("token", string(b))
			if err := viper.WriteConfig(); err != nil {
				fmt.Println("Error updating config:", err)
			}
		},
	}

	return goinvestec.NewClient(
		viper.GetString("apiURL"),
		auther,
	)
}
