package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"time"

	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type Signer struct {
	key    *rsa.PrivateKey
	issuer string
	s      jose.Signer
	alg    jose.SignatureAlgorithm
}

func NewSigner() *Signer {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	alg := jose.RS256
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: alg, Key: key}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		panic(err)
	}
	return &Signer{
		key:    key,
		issuer: "kubernetes-serviceaccount-authority",
		s:      sig,
		alg:    alg,
	}
}

func (s *Signer) Sign(c PublicClaims, p PrivateClaims) string {
	var b [18]byte
	rand.Read(b[:])

	cl := jwt.Claims{
		Subject:   c.Subject,
		Audience:  jwt.Audience(c.Audience),
		Issuer:    s.issuer,
		Expiry:    jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		NotBefore: jwt.NewNumericDate(time.Now().Add(-30 * time.Minute)),
		ID:        base64.URLEncoding.EncodeToString(b[:]),
	}
	raw, err := jwt.Signed(s.s).Claims(cl).Claims(p).CompactSerialize()
	if err != nil {
		panic(err)
	}
	return raw
}

func (s *Signer) Verify(data string) (*PublicClaims, *PrivateClaims, error) {
	token, err := jwt.ParseSigned(data)
	if err != nil {
		return nil, nil, err
	}
	cl := jwt.Claims{}
	p := &PrivateClaims{}
	if err := token.Claims(&s.key.PublicKey, &cl, &p); err != nil {
		return nil, nil, err
	}
	if err := cl.Validate(jwt.Expected{Time: time.Now()}); err != nil {
		return nil, nil, err
	}
	return &PublicClaims{Subject: cl.Subject, Audience: cl.Audience}, p, nil
}

func (s *Signer) JWKs() jose.JSONWebKeySet {
	return jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			jose.JSONWebKey{
				Key:       s.key,
				Algorithm: string(s.alg),
			},
		},
	}
}

type PublicClaims struct {
	Subject  string
	Audience []string
}

type PrivateClaims struct {
	Kubernetes KubernetesClaim `json:"k8s,omitempty"`
}

type KubernetesClaim struct {
	Groups []string `json:"groups,omitempty"`
}
