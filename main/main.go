package main

import (
	"github.com/spacetimi/server/bonda/app_src/app_init"
	"github.com/spacetimi/server/bonda/dummy_controller"
	"github.com/spacetimi/server/bonda/testing"
	"github.com/spacetimi/server/timi_shared/code/core/shared_init"
	"github.com/spacetimi/server/timi_shared/code/api_server"
)

func main() {

	shared_init.SharedInit(app_init.GetAppInitializer())

	// TODO: Remove these
	api_server.StartServer(testing.TestingController, dummy_controller.DummyController)
}
