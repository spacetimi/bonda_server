package main

import (
	"github.com/spacetimi/bonda_server/app_src/app_init"
	"github.com/spacetimi/bonda_server/dummy_controller"
	"github.com/spacetimi/bonda_server/testing"
	"github.com/spacetimi/timi_shared_server/code/admin_server"
	"github.com/spacetimi/timi_shared_server/code/api_server"
	"github.com/spacetimi/timi_shared_server/code/core/shared_init"
	"sync"
)

func main() {

	shared_init.SharedInit(app_init.GetAppInitializer())

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
	    defer wg.Done()
		api_server.StartServer(testing.TestingController, dummy_controller.DummyController)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		admin_server.StartServer()
	}()

	wg.Wait()
}
