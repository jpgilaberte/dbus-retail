package routes

var RoutesAgent = map[string]map[string]Handler{
	"GET": {
		//"/api/v1/list": handleList,
	},
	"POST": {
		//"/api/v1/valkiria": handleShooter,
	},
}

var RoutesOrchestrator = map[string]map[string]Handler{
	"POST": {
		//"/api/v1/login": handleShooter,
		//"/api/v1/valkiria": valkiriaKill,
	},
	"GET": {
		//"/api/v1/install": valkiriaInstall,
		//"/api/v1/uninstall": valkiriaUninstall,
		//"/api/v1/start": valkiriaStart,
		//"/api/v1/stop": valkiriaStop,
		//"/api/v1/nodes": nodes,
		//"/api/v1/list": valkiriaList,
	},
}
