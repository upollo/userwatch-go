package userwatchgo

import (
	context "context"
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type KeyPerRPCCredentials struct {
	key        string
	requireTls bool
	credentials.PerRPCCredentials
}

func (key KeyPerRPCCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"X-API-KEY": key.key}, nil
}
func (key KeyPerRPCCredentials) RequireTransportSecurity() bool {
	return key.requireTls
}

type ClientOptions struct {
	apiKey     string
	url        string
	requireTls bool
}

func NewClientBuilder(apiKey string) *ClientOptions {
	return &ClientOptions{
		apiKey:     apiKey,
		url:        "api.userwat.ch:443",
		requireTls: true,
	}
}

// optionally override the url for testing
func (opts *ClientOptions) WithUrl(url string) *ClientOptions {
	opts.url = url
	return opts
}

// optionally override the tls requirement for testing.
func (opts *ClientOptions) WithRequireTls(requireTls bool) *ClientOptions {
	opts.requireTls = requireTls
	return opts
}

func (opts *ClientOptions) Build() (ShepherdClient, error) {
	dialOptions := []grpc.DialOption{grpc.WithPerRPCCredentials(&KeyPerRPCCredentials{
		key:        opts.apiKey,
		requireTls: opts.requireTls,
	})}

	if opts.requireTls {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: false,
		})))
	} else {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(opts.url, dialOptions...)
	if err != nil {
		return nil, err
	}

	return NewShepherdClient(conn), nil
}
