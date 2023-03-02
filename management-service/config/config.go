package config

import (
	"fmt"
	"time"

	"github.com/gojek/mlp/api/pkg/instrumentation/newrelic"
	"github.com/gojek/mlp/api/pkg/instrumentation/sentry"

	common_config "github.com/caraml-dev/xp/common/config"
)

type Config struct {
	OpenAPISpecsPath string `default:"."`
	Port             int    `default:"3000"`

	AllowedOrigins      []string `default:"*"`
	AuthorizationConfig *AuthorizationConfig
	DbConfig            *DatabaseConfig
	MLPConfig           *MLPConfig
	MessageQueueConfig  *MessageQueueConfig
	SegmenterConfig     map[string]interface{}
	ValidationConfig    ValidationConfig
	DeploymentConfig    DeploymentConfig
	NewRelicConfig      newrelic.Config
	SentryConfig        sentry.Config
	XpUIConfig          *XpUIConfig
}

// AuthorizationConfig captures the config for MLP authz
type AuthorizationConfig struct {
	Enabled bool
	URL     string
}

// DatabaseConfig captures the XP database config
type DatabaseConfig struct {
	Host           string `default:"localhost"`
	Port           int    `default:"5432"`
	User           string `default:"xp"`
	Password       string `default:"xp"`
	Database       string `default:"xp"`
	MigrationsPath string `default:"file://database/db-migrations"`

	ConnMaxIdleTime time.Duration `default:"0s"`
	ConnMaxLifetime time.Duration `default:"0s"`
	MaxIdleConns    int           `default:"0"`
	MaxOpenConns    int           `default:"0"`
}

// MLPConfig captures the configuration used to connect to the MLP API server
type MLPConfig struct {
	URL string
}

// MessageQueueKind describes the message queue for transmitting event updates to and fro Treatment Service
type MessageQueueKind = string

const (
	// NoopMQ is a No-Op Message Queue
	NoopMQ MessageQueueKind = ""
	// PubSubMQ is a PubSub Message Queue
	PubSubMQ MessageQueueKind = "pubsub"
)

type MessageQueueConfig struct {
	// The type of Message Queue for event updates
	Kind MessageQueueKind `default:""`

	// PubSubConfig captures the config related to initializing a PubSub Message Queue
	PubSubConfig *PubSubConfig
}

// PubSubConfig captures the config for the Google PubSub client, to publish messages
// about changes in the experimentation data
type PubSubConfig struct {
	Project   string `default:"dev"`
	TopicName string `default:"xp-update"`
}

// ValidationConfig captures the config related to the validation of schemas
type ValidationConfig struct {
	ValidationUrlTimeoutSeconds int `default:"5"`
}

// DeploymentConfig captures the config related to the deployment of Management Service
type DeploymentConfig struct {
	EnvironmentType string `default:"local"`
}

// XpUIConfig captures config related to serving XP UI files
type XpUIConfig struct {
	// Optional. If configured, xp management service API will serve static files
	// of the xp-ui React app.
	AppDirectory string
	// Optional. Defines the relative path under which the app will be accessible.
	// This should match `homepage` value from the `package.json` file of the CRA app
	Homepage string `default:"/xp"`
}

// ListenAddress returns the Management API app's port
func (c *Config) ListenAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}

func Load(filepaths ...string) (*Config, error) {
	var cfg Config
	err := common_config.ParseConfig(&cfg, filepaths)
	if err != nil {
		return nil, fmt.Errorf("failed to update viper config: %s", err)
	}

	return &cfg, nil
}
