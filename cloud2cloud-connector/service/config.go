package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-ocf/kit/net/grpc"
	"github.com/go-ocf/kit/security/oauth/manager"
)

type PullStaticDeviceEvents struct {
	CacheSize       int           `envconfig:"CACHE_SIZE" default:"2048"`
	Timeout         time.Duration `envconfig:"TIMEOUT" default:"5s"`
	MaxParallelGets int64         `envconfig:"MAX_PARALLEL_GETS" default:"128"`
}

//Config represent application configuration
type Config struct {
	grpc.Config
	AuthServerAddr         string                 `envconfig:"AUTH_SERVER_ADDRESS" default:"127.0.0.1:9100"`
	ResourceAggregateAddr  string                 `envconfig:"RESOURCE_AGGREGATE_ADDRESS"  default:"127.0.0.1:9100"`
	ResourceDirectoryAddr  string                 `envconfig:"RESOURCE_DIRECTORY_ADDRESS"  default:"127.0.0.1:9100"`
	OAuthCallback          string                 `envconfig:"OAUTH_CALLBACK"`
	EventsURL              string                 `envconfig:"EVENTS_URL"`
	PullDevicesDisabled    bool                   `envconfig:"PULL_DEVICES_DISABLED" default:"false"`
	PullDevicesInterval    time.Duration          `envconfig:"PULL_DEVICES_INTERVAL" default:"5s"`
	PullStaticDeviceEvents PullStaticDeviceEvents `envconfig:"PULL_STATIC_DEVICE_EVENTS"`
	OAuth                  manager.Config         `envconfig:"OAUTH"`
}

//String return string representation of Config
func (c Config) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("config: \n%v\n", string(b))
}
