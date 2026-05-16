package main

import (
	"IDCard2FA/internal/config"
	"IDCard2FA/internal/logger"
	"IDCard2FA/internal/openapi"
	"fmt"

	realname "go.gh.ink/openapi/sdk/20260512/v3/private/real_name"
	"go.uber.org/zap"
)

func main() {
	// Load public config
	config.Init()

	// Init logger
	logger.Init()

	// Read space name
	spaceName := ""
	fmt.Println("Please input space name:")
	if _, err := fmt.Scanln(&spaceName); err != nil {
		logger.L.Error("failed to scan input", zap.Error(err))
		return
	}

	// Init Open API
	openapi.Init(spaceName)

	for {
		// Read name
		fmt.Println("Please input name:")
		name := ""
		if _, err := fmt.Scanln(&name); err != nil {
			logger.L.Error("failed to scan input", zap.Error(err))
			continue
		}

		// Check name content
		if name == "" {
			logger.L.Error("name is empty, please retry")
			continue
		}

		// Read ID
		fmt.Println("Please input id:")
		id := ""
		if _, err := fmt.Scanln(&id); err != nil {
			logger.L.Error("failed to scan input", zap.Error(err))
			continue
		}

		// Check ID content
		if id == "" {
			logger.L.Error("id is empty, please retry")
			continue
		}

		// Request SDK
		ok, err := realname.VerifyCNID(openapi.C, id, name)
		if err != nil {
			logger.L.Error("request api failed, please retry", zap.Error(err))
			continue
		}

		fmt.Println("Verification result:", ok)
	}
}
