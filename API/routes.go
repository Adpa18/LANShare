package API

var routes = Routes{
	Route{
		"List", "GET", "/api/list", listRootHandler,
	}, Route{
		"List", "GET", "/api/list/{path:.*}", listHandler,
	}, Route{
		"Download", "GET", "/api/download/{path:.*}", downloadHandler,
	}, Route{
		"Upload", "POST", "/api/upload/{path:.*}", uploadHandler,
	}, Route{
		"Zips", "GET", "/api/zips/{zipID:.*}", zipHandler,
	}, Route{
		"Home", "GET", "/", homeHandler,
	},Route{
		"Public", "GET", "/public/{path:.*}", publicHandler,
	},Route{
		"Home", "GET", "/{path:.*}", homeHandler,
	},
}
