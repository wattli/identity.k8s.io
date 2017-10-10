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
	issuer string
	s      jose.Signer
}

func NewSigner() *Signer {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: key}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		panic(err)
	}
	return &Signer{
		issuer: "kubernetes-serviceaccount-authority",
		s:      sig,
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
