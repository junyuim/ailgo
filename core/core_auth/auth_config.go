package core_auth

type AuthConfig struct {
	Enabled bool `json:"enabled" yaml:"enabled"`

	IncludeUrls []string `json:"includeUrls" yaml:"include-urls"`

	ExcludeUrls []string `json:"excludeUrls" yaml:"exclude-urls"`

	PublicKey string `json:"publicKey" yaml:"public-key"`

	PrivateKey string `json:"privateKey" yaml:"private-key"`

	TokenIssuer string `json:"tokenIssuer" yaml:"token-issuer"`

	TokenValidity int64 `json:"tokenValidity" yaml:"token-validity"`

	//
	//SsoEnabled bool `yaml:"sso-enabled"`
	//
	//SsoTokenValidity int64 `yaml:"sso-token-validity"`
}
