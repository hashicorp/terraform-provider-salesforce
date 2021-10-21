package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nimajalali/go-force/force"
)

const (
	salesforceLoginServer   = "https://hashicorp--terraform2.my.salesforce.com"
	salesforceOAuthEndpoint = "/services/oauth2/token"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	InstanceUrl string `json:"instance_url"`
	Id          string `json:"id"`
	TokenType   string `json:"token_type"`
}

func SignJWT(privateKey []byte, user string, clientId string) (string, error) {
	priv, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.StandardClaims{
		ExpiresAt: time.Now().UTC().Add(3 * time.Minute).Unix(),
		Subject:   user,
		Issuer:    clientId,
		Audience:  salesforceLoginServer,
	})

	return token.SignedString(priv)
}

func Authenticate(signedJwt string) (AuthResponse, error) {
	var oauth AuthResponse

	payload := url.Values{}
	payload.Add("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer")
	payload.Add("assertion", signedJwt)

	// Build Body
	body := strings.NewReader(payload.Encode())

	// Build Request
	req, err := http.NewRequest("POST", salesforceLoginServer+salesforceOAuthEndpoint, body)
	if err != nil {
		return oauth, fmt.Errorf("Error creating authentication request: %v", err)
	}

	// Add Headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return oauth, fmt.Errorf("Error sending authentication request: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return oauth, fmt.Errorf("Error reading authentication response bytes: %v", err)
	}

	// Attempt to parse response as a force.com api error
	apiError := &force.ApiError{}
	if err := json.Unmarshal(respBytes, apiError); err == nil {
		// Check if api error is valid
		if apiError.Validate() {
			return oauth, apiError
		}
	}

	if err := json.Unmarshal(respBytes, &oauth); err != nil {
		return oauth, fmt.Errorf("Unable to unmarshal authentication response: %v", err)
	}
	return oauth, nil
}

type Config struct {
	ClientId   string
	PrivateKey string
	ApiVersion string
	Username   string
}

func Client(config Config) (*force.ForceApi, error) {
	// try to read private key as file
	privateKeyBytes, err := ioutil.ReadFile(config.PrivateKey)
	if os.IsNotExist(err) {
		// assume private key was passed directly
		privateKeyBytes = []byte(config.PrivateKey)
	} else if err != nil {
		return nil, err
	}

	signedJwt, err := SignJWT(privateKeyBytes, config.Username, config.ClientId)
	if err != nil {
		return nil, err
	}

	resp, err := Authenticate(signedJwt)
	if err != nil {
		return nil, err
	}

	return force.CreateWithAccessToken(config.ApiVersion, config.ClientId, resp.AccessToken, resp.InstanceUrl)
}
