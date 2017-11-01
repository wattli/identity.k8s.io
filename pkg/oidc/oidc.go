package oidc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	jose "gopkg.in/square/go-jose.v2"
	"k8s.io/apiserver/pkg/authorization/authorizer"
)

type OIDCProviderMetadata struct {
	// REQUIRED. URL using the https scheme with no query or fragment
	// component that the OP asserts as its Issuer Identifier. If Issuer
	// discovery is supported (see Section 2), this value MUST be identical to
	// the issuer value returned by WebFinger. This also MUST be identical to
	// the iss Claim value in ID Tokens issued from this Issuer.
	Issuer string `json:"issuer"`

	//  REQUIRED. URL of the OP's OAuth 2.0 Authorization Endpoint [OpenID.Core].
	AuthorizationEndpoint string `json:"authorization_endpoint"`

	// URL of the OP's OAuth 2.0 Token Endpoint [OpenID.Core]. This is
	// REQUIRED unless only the Implicit Flow is used.
	TokenEndpoint string `json:"token_endpoint"`

	// REQUIRED. URL of the OP's JSON Web Key Set [JWK] document. This
	// contains the signing key(s) the RP uses to validate signatures from the
	// OP. The JWK Set MAY also contain the Server's encryption key(s), which
	// are used by RPs to encrypt requests to the Server. When both signing
	// and encryption keys are made available, a use (Key Use) parameter value
	// is REQUIRED for all keys in the referenced JWK Set to indicate each
	// key's intended usage. Although some algorithms allow the same key to be
	// used for both signatures and encryption, doing so is NOT RECOMMENDED,
	// as it is less secure. The JWK x5c parameter MAY be used to provide
	// X.509 representations of keys provided. When used, the bare key values
	// MUST still be present and MUST match those in the certificate.
	JwksURI string `json:"jwks_uri"`

	// REQUIRED. JSON array containing a list of the OAuth 2.0
	// response_type values that this OP supports. Dynamic OpenID Providers
	// MUST support the code, id_token, and the token id_token Response Type
	// values.
	ResponseTypesSupported []string `json:"response_types_supported"`

	//response_modes_supported
	//OPTIONAL. JSON array containing a list of the OAuth 2.0 response_mode values that this OP supports, as specified in OAuth 2.0 Multiple Response Type Encoding Practices [OAuth.Responses]. If omitted, the default for Dynamic OpenID Providers is ["query", "fragment"].

	// REQUIRED. JSON array containing a list of the Subject Identifier
	// types that this OP supports. Valid types include pairwise and public.
	SubjectTypesSupported []string `json:"subject_types_supported"`

	// REQUIRED. JSON array containing a list of the JWS signing algorithms
	// (alg values) supported by the OP for the ID Token to encode the Claims
	// in a JWT [JWT]. The algorithm RS256 MUST be included. The value none
	// MAY be supported, but MUST NOT be used unless the Response Type used
	// returns no ID Token from the Authorization Endpoint (such as when using
	// the Authorization Code Flow).
	IDTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`

	// OPTIONAL. JSON array containing a list of Client Authentication
	// methods supported by this Token Endpoint. The options are
	// client_secret_post, client_secret_basic, client_secret_jwt, and
	// private_key_jwt, as described in Section 9 of OpenID Connect Core 1.0
	// [OpenID.Core]. Other authentication methods MAY be defined by
	// extensions. If omitted, the default is client_secret_basic -- the HTTP
	// Basic Authentication Scheme specified in Section 2.3.1 of OAuth 2.0
	// [RFC6749].
	//
	// We support none.
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`

	// OPTIONAL. JSON array containing a list of the Claim Types that the
	// OpenID Provider supports. These Claim Types are described in Section
	// 5.6 of OpenID Connect Core 1.0 [OpenID.Core]. Values defined by this
	// specification are normal, aggregated, and distributed. If omitted, the
	// implementation supports only normal Claims.
	//claim_types_supported

	// RECOMMENDED. JSON array containing a list of the Claim Names of the
	// Claims that the OpenID Provider MAY be able to supply values for. Note
	// that for privacy or other reasons, this might not be an exhaustive
	// list.
	//claims_supported

	//OPTIONAL. Boolean value specifying whether the OP supports use of the claims parameter, with true indicating support. If omitted, the default value is false.
	//claims_parameter_supported

	//OPTIONAL. Boolean value specifying whether the OP supports use of the request parameter, with true indicating support. If omitted, the default value is false.
	//request_parameter_supported

	//OPTIONAL. Boolean value specifying whether the OP supports use of the request_uri parameter, with true indicating support. If omitted, the default value is true.
	//request_uri_parameter_supported

	//OPTIONAL. Boolean value specifying whether the OP requires any request_uri values used to be pre-registered using the request_uris registration parameter. Pre-registration is REQUIRED when the value is true. If omitted, the default value is false.
	//require_request_uri_registration
}

func Provider(issuer string) OIDCProviderMetadata {
	return OIDCProviderMetadata{
		Issuer:                issuer,
		AuthorizationEndpoint: issuer + "/oauth2/v0/auth",
		TokenEndpoint:         issuer + "/oauth2/v0/token",
		JwksURI:               issuer + "/oauth2/v0/certs",
		ResponseTypesSupported: []string{
			"code",
			"token",
			"id_token",
			"code token",
			"code id_token",
			"token id_token",
			"code token id_token",
			"none",
		},
		SubjectTypesSupported: []string{
			"public",
		},
		IDTokenSigningAlgValuesSupported: []string{
			"RS256",
		},
		TokenEndpointAuthMethodsSupported: []string{},
	}
}

type OIDCMeta struct {
	Issuer string
	JWKs   jose.JSONWebKeySet
}

func (m *OIDCMeta) WriteDiscoveryDir(path string) error {
	mode := os.FileMode(0600)
	metapath := filepath.Join(path, ".well-known/openid-configuration")
	certspath := filepath.Join(path, "oauth2/v0/certs")

	metab, err := json.MarshalIndent(Provider(m.Issuer), "", "  ")
	if err != nil {
		return fmt.Errorf("couldn't marshal oidc-configuration: %v", err)
	}
	jwksb, err := json.MarshalIndent(m.JWKs, "", "  ")
	if err != nil {
		return fmt.Errorf("couldn't marshal jwks: %v", err)
	}

	for _, p := range []string{
		path,
		filepath.Dir(metapath),
		filepath.Dir(certspath),
	} {
		if err := os.MkdirAll(p, mode); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(metapath, metab, mode); err != nil {
		return fmt.Errorf("couldn't write oidc-configuration: %v", err)
	}
	if err := ioutil.WriteFile(certspath, jwksb, mode); err != nil {
		return fmt.Errorf("couldn't write jwks: %v", err)
	}
	return nil
}

type muxer interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

func (m *OIDCMeta) InstallHandlers(mux muxer) error {
	metapath := filepath.Join("/oidc", ".well-known/openid-configuration")
	certspath := filepath.Join("/oidc", "oauth2/v0/certs")
	metab, err := json.MarshalIndent(Provider(m.Issuer), "", "  ")
	if err != nil {
		return fmt.Errorf("couldn't marshal oidc-configuration: %v", err)
	}
	jwksb, err := json.MarshalIndent(m.JWKs, "", "  ")
	if err != nil {
		return fmt.Errorf("couldn't marshal jwks: %v", err)
	}
	mux.HandleFunc(metapath, func(w http.ResponseWriter, r *http.Request) {
		w.Write(metab)
	})
	mux.HandleFunc(certspath, func(w http.ResponseWriter, r *http.Request) {
		w.Write(jwksb)
	})
	return nil
}

// TODO: once unequivocal deny merges, deny /token and /auth endpoints
func Authorizer() authorizer.Authorizer {
	return authorizer.AuthorizerFunc(func(a authorizer.Attributes) (bool, string, error) {
		for _, p := range []string{
			"/oidc/oauth2/v0/certs",
			"/oidc/.well-known/openid-configuration",
		} {
			if a.GetPath() == p {
				return true, "", nil
			}
		}
		return false, "", nil
	})
}
