// Created at 10/21/2021 1:13 AM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_apns

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
)

// TODO push batch

type Context interface {
	context.Context
}

type Client struct {
	httpClient *http.Client
	cert       *tls.Certificate
	endpoint   string
}

func NewClient(certificate *tls.Certificate) *Client {
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*certificate},
	}
	transport := &http2.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := &http.Client{
		Transport: transport,
	}

	return &Client{
		httpClient: client,
		cert:       certificate,
		endpoint:   DefaultDevelopmentEndpoint,
	}
}

func (c *Client) Development() *Client {
	c.endpoint = DefaultDevelopmentEndpoint
	return c
}

func (c *Client) Production() *Client {
	c.endpoint = DefaultProductionEndpoint
	return c
}

func (c *Client) generateUrl(token string) string {
	return fmt.Sprintf("%s%s%s", c.endpoint, DefaultPath, token)
}

func (c *Client) Push(notify *Notification) (*Response, error) {
	return c.PushWithCtx(nil, notify)
}

func (c *Client) PushWithCtx(ctx Context, n *Notification) (*Response, error) {
	payload, err := n.payload.marshalJSON()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.generateUrl(n.deviceToken), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if n.topic != "" {
		topic := n.topic
		if n.pushType == PushTypeVOIP {
			topic = fmt.Sprintf("%s.%s", n.topic, PushTypeVOIP)
		}
		req.Header.Set("apns-topic", topic)
	}
	if n.apnsID != "" {
		req.Header.Set("apns-id", n.apnsID)
	}
	if n.collapseID != "" {
		req.Header.Set("apns-collapse-id", n.collapseID)
	}
	if n.priority > 0 {
		req.Header.Set("apns-priority", fmt.Sprintf("%v", n.priority))
	}
	if !n.expiration.IsZero() {
		req.Header.Set("apns-expiration", fmt.Sprintf("%v", n.expiration.Unix()))
	}
	// default PushTypeAlert
	if n.pushType != "" {
		req.Header.Set("apns-push-type", string(n.pushType))
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return parseResponse(resp)
}
