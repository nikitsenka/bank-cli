package command

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.AddCommand(googleCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Get access to REST API",
	Long:  `Work with GitHub issues.`,
}

var googleCmd = &cobra.Command{
	Use:   "google",
	Short: "Get google access token",
	Long:  `Work with GitHub issues.`,
	Run:   googleToken,
}

func googleToken(cmd *cobra.Command, args []string) {
	m := http.NewServeMux()
	s := http.Server{Addr: ":8080", Handler: m}
	m.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		openbrowser(fmt.Sprintf("file:///Users/iniki1/projects/bank-cli/success.html"))
		s.Shutdown(context.Background())
	})
	config := &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	url := config.AuthCodeURL("pseudo-random")
	openbrowser(url)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Printf("Finished")
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
