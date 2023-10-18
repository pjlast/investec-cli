package cli

import (
	goinvestec "github.com/pjlast/go-investec"
	"github.com/spf13/viper"
)

func NewClient() *goinvestec.Client {
	if viper.GetBool("useSandbox") {
		return goinvestec.NewSandboxClient()
	}

	return goinvestec.NewAuthorizedClient(viper.GetString("clientID"), viper.GetString("secret"), viper.GetString("apiKey"))
}
