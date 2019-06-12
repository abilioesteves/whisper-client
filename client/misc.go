package client

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ory/x/randx"
)

// getAccessTokenFromRequest is a helper method to recover an Access Token from a http request
func getAccessTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authURLParam := r.URL.Query().Get("token")
	var t string

	if len(authHeader) == 0 && len(authURLParam) == 0 {
		return "", fmt.Errorf("No Authorization Header or URL Param found")
	}

	if len(authHeader) > 0 {
		data := strings.Split(authHeader, " ")

		if len(data) != 2 {
			return "", fmt.Errorf("Bad Authorization Header")
		}

		t = data[0]

		if len(t) == 0 || t != "Bearer" {
			return "", fmt.Errorf("No Bearer Token found")
		}

		t = data[1]

	} else {
		t = authURLParam
	}

	if len(t) == 0 {
		return "", fmt.Errorf("Bad Authorization Token")
	}

	return t, nil
}

func getStateAndNonce() (state, nonce []rune, err error) {
	state, err = randx.RuneSequence(24, randx.AlphaLower)
	if err == nil {
		nonce, err = randx.RuneSequence(24, randx.AlphaLower)

		if err == nil {
			return state, nonce, err
		}
	}
	return nil, nil, nil
}