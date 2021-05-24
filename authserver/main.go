package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	gooidc "github.com/coreos/go-oidc"
	oauth2 "golang.org/x/oauth2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func getToken(Headers http.Header) (token string, err error) {
	token = ""
	err = nil
	for name, headers := range Headers {
		for _, h := range headers {
			if name == "Authorization" {
				token = h
				goto End
			}
		}
	}
End:
	if strings.Contains(token, "Bearer ") {
		token = strings.Replace(token, "Bearer", "", 1)
	}

	if token == "" {
		return "", errors.New("error, token not provided")
	}
	return
}

func Authorize(ctx context.Context, Headers http.Header) error {

	token, err := getToken(Headers)
	fmt.Println("Token is:", token)
	if err != nil {
		return err
	}

	var issClaim struct {
		Issuer string `json:"iss"`
	}

	issTok, err := jwt.ParseSigned(token)
	if err != nil {
		return err
	}

	_ = issTok.UnsafeClaimsWithoutVerification(&issClaim)
	issuer := issClaim.Issuer

	fmt.Println("issuer:", issuer)

	authctx := context.Background()

	client := &http.Client{}

	authctx = context.WithValue(authctx, oauth2.HTTPClient, client)

	provider, err := gooidc.NewProvider(authctx, os.Getenv("OIDCAUTHISSUER"))
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return err
	}

	verifier := provider.Verifier(&gooidc.Config{ClientID: "account"})
	idTok, err := verifier.Verify(authctx, token)
	if err != nil {
		log.Printf("verification failed: %s\n", err.Error())
		return err
	}
	var claims struct {
		Email string `json:"email"`
	}
	if err := idTok.Claims(&claims); err != nil {
		return err
	}

	return nil
}

func alert(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.Write([]byte("\nMethod Not Allowed\n"))
		return
	}
	ctx := req.Context()
	err := Authorize(ctx, req.Header)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("\nAuth Failed\n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("\n"))
	} else {
		w.Write([]byte("\nVerrification Successfull\n"))
	}
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	args := os.Args

	http.HandleFunc("/v1/alerts", alert)
	if len(args) > 1 {
		port := ":" + args[1]
		http.ListenAndServe(port, nil)
	} else {
		http.ListenAndServe(":8093", nil)
	}
}
