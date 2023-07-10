package server

import (
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/http"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/property"
	"k8s.io/klog/v2"
	"strings"
)

const (
	ApiServerNameField = "server"
)

const (
	BasicAuth = "Basic"
	Ntlm      = "Ntlm"
	OAuth2    = "OAuth2"
)

type ApiServer struct {
	Url                  string `json:"url"`
	AuthenticationMethod string `json:"authenticationMethod"`
	VerifySSL            bool   `json:"insecure"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	Domain               string `json:"domain"`
	ProxyHost            string `json:"proxyHost"`
	ProxyPort            string `json:"proxyPort"`
	ProxyUsername        string `json:"proxyUsername"`
	ProxyPassword        string `json:"proxyPassword"`
	ProxyDomain          string `json:"proxyDomain"`
	AccessTokenUrl       string `json:"accessTokenUrl"`
	ClientId             string `json:"clientId"`
	ClientSecret         string `json:"clientSecret"`
	Scope                string `json:"scope"`
}

func DeserializeApiServer(properties []task.PropertyDefinition) (*ApiServer, error) {
	var apiServer ApiServer
	if err := property.ExtractType(ApiServerNameField, properties, &apiServer); err != nil {
		klog.Errorf("Cannot deserialize server properties %v", err)
		return nil, fmt.Errorf("cannot deserialize server properties: %w", err)
	}
	return &apiServer, nil
}

func (server *ApiServer) GetHttpClient() (*http.HttpClient, error) {
	httpClientConfig := &http.HttpClientConfig{
		Host:          server.Url,
		Insecure:      server.VerifySSL,
		ProxyHost:     server.ProxyHost,
		ProxyPort:     server.ProxyPort,
		ProxyUsername: server.ProxyUsername,
		ProxyPassword: server.ProxyPassword,
		ProxyDomain:   server.ProxyDomain,
	}

	switch server.AuthenticationMethod {
	case BasicAuth:
		httpClientConfig.BasicAuthentication = &http.BasicAuthentication{
			Username: server.Username,
			Password: server.Password,
		}
	case Ntlm:
		httpClientConfig.NtlmAuthentication = &http.NtlmAuthentication{
			Username: server.Username,
			Password: server.Password,
			Domain:   server.Domain,
		}
	case OAuth2:
		httpClientConfig.OAuth2Authentication = &http.OAuth2Authentication{
			ClientID:     server.ClientId,
			ClientSecret: server.ClientSecret,
			Scopes:       strings.Split(server.Scope, " "),
			TokenURL:     server.AccessTokenUrl,
		}
	}

	builder := http.NewHttpClientBuilder().WithHttpClientConfig(httpClientConfig)
	return builder.Build()
}
