package main

import (
	"github.com/samfil-technohub/samuel-nwoye-website/controllers"
)

// program entry point
func main() {
	//Start the API Server
	controllers.ServeAPI()
}
