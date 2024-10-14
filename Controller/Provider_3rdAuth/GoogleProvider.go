package provider3rdauth

//Provide Google Authenciation class
import (
	"os"
	config "weather/Config"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func GoogleProvide() {
	var google_client_id string = os.Getenv("Client_ID_Google")
	var google_client_secret string = os.Getenv("Client_secret_Google")
	var google_client_callback string = os.Getenv("Client_callback_Google")

	config.InitSecureCookie()
	goth.UseProviders(
		google.New(google_client_id, google_client_secret, google_client_callback),
	)
}
