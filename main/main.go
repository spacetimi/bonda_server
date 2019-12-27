package main

import (
	"github.com/spacetimi/bonda_server/app_src/app_init"
	"github.com/spacetimi/bonda_server/dummy_controller"
	"github.com/spacetimi/bonda_server/testing"
	"github.com/spacetimi/timi_shared_server/code/core/shared_init"
	"github.com/spacetimi/timi_shared_server/code/api_server"
)

func main() {

	shared_init.SharedInit(app_init.GetAppInitializer())

	// TODO: Remove these
	api_server.StartServer(testing.TestingController, dummy_controller.DummyController)
}
