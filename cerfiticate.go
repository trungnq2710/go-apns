// Created at 10/21/2021 1:03 AM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_apns

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
)

var (
	errFoundUnknownPrivateKey = errors.New("found unknown private key type in PKCS#8 wrapping")
	errFailedParsePrivateKey  = errors.New("failed to parse private key")
	errNoCertificateFound     = errors.New("no certificate found")
	errPrivateKeyFound        = errors.New("no private key found")
)

func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey:
			return key, nil
		default:
			return nil, errFoundUnknownPrivateKey
		}
	}
	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, nil
	}
	return nil, errFailedParsePrivateKey
}

func CertificateFromPemFile(path string) (*tls.Certificate, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cert tls.Certificate
	for {
		block, rest := pem.Decode(raw)
		if block == nil {
			break
		}
		if block.Type == "CERTIFICATE" {
			cert.Certificate = append(cert.Certificate, block.Bytes)
		} else {
			// TODO correct handle of the keyword
			cert.PrivateKey, err = parsePrivateKey(block.Bytes)
			if err != nil {
				return nil, err
			}
		}
		raw = rest
	}

	if len(cert.Certificate) == 0 {
		return nil, errNoCertificateFound
	} else if cert.PrivateKey == nil {
		return nil, errPrivateKeyFound
	}

	return &cert, nil
}

func CertificateFromP12File(path string) (*tls.Certificate, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	key, cert, err := pkcs12.Decode(raw, "")
	if err != nil {
		return nil, err
	}

	certificate := &tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  key,
		Leaf:        cert,
	}

	return certificate, nil
}
