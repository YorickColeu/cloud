package test

import (
	"sync"
	"testing"

	"github.com/plgd-dev/cloud/cloud2cloud-gateway/refImpl"
	testCfg "github.com/plgd-dev/cloud/test/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
)

func MakeConfig(t *testing.T) refImpl.Config {
	var cfg refImpl.Config
	err := envconfig.Process("", &cfg)
	require.NoError(t, err)
	cfg.Service.Addr = testCfg.C2C_GW_HOST
	cfg.JwksURL = testCfg.JWKS_URL
	cfg.Service.ResourceDirectoryAddr = testCfg.RESOURCE_DIRECTORY_HOST
	cfg.Service.FQDN = "cloud2cloud-gateway-" + t.Name()
	cfg.Listen.File.DisableVerifyClientCertificate = true
	cfg.Service.OAuth.ClientID = testCfg.OAUTH_MANAGER_CLIENT_ID
	cfg.Service.OAuth.Endpoint.TokenURL = testCfg.OAUTH_MANAGER_ENDPOINT_TOKENURL
	return cfg
}

func SetUp(t *testing.T) (TearDown func()) {
	return New(t, MakeConfig(t))
}

func New(t *testing.T, cfg refImpl.Config) func() {

	s, err := refImpl.Init(cfg)
	require.NoError(t, err)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Serve()
	}()

	return func() {
		s.Close()
		wg.Wait()
	}
}
