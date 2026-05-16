package openapi

import (
	"IDCard2FA/internal/config"
	"IDCard2FA/internal/logger"

	"go.gh.ink/json"
	"go.gh.ink/openapi/sdk/20260512/v3/client"
	"go.uber.org/zap"
)

var C *client.Client

// Init inits Ghink OpenAPI client
func Init(spaceName string) {
	// Create a new client
	c, err := client.NewClient(
		config.Get().GhinkOpenAPI[spaceName].SecretID,
		config.Get().GhinkOpenAPI[spaceName].SecretKey,
		client.WithLogger(&ZapLogger{logger: logger.L}),
		client.WithExponentialBackoff(false),
		client.WithMarshal(json.Marshal),
		client.WithUnmarshal(json.Unmarshal),
	)
	if err != nil {
		logger.L.Fatal("failed to init client of Ghink OpenAPI", zap.Error(err))
	}

	C = c

	logger.L.Debug("Ghink OpenAPI client initialized")
}
