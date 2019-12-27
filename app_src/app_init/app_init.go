package app_init

import "github.com/spacetimi/timi_shared_server/code/core/shared_init"

func GetAppInitializer() shared_init.IAppInitializer {
	return &appInitializer
}

type AppInitializer struct {	// Implements IAppInit
}
var appInitializer AppInitializer

/********** Begin IAppInitializer implementation **********/
func (appInitializer *AppInitializer) AppInit() {
	// TODO: Nothing to do yet
}
/********** End IAppInitializer implementation **********/


