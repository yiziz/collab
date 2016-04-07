package config

import (
	"os"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/yiziz/collab/config/initializers"
	"github.com/yiziz/collab/services/env"
)

// AppEnv determines what env the app is running in
var AppEnv string

// SetAppEnv sets the server's envirnoment var
func SetAppEnv() {
	AppEnv = os.Getenv("GIN_ENV")
	if AppEnv == "" {
		AppEnv = "development"
	}
}

func setAppMode() {
	switch AppEnv {
	case "development":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	}
}

// Run spins up the app
func Run() {
	env.SetEnvVars()
	SetAppEnv()
	setAppMode()

	jobrunner.Start()

	// rm := data.RedempData()
	// data.Test()
	// recommend.Foo()
	// pm := data.PerkTermScores()

	// um := data.CalcUserTermScore(rm, pm)
	// fmt.Println(um)

	// recommend.Bar()

	r := gin.Default()
	initializers.Initialize(r, AppEnv)
	if AppEnv == "development" {
		r.Run(":9000")
	} else {
		r.Run(":80")
	}
}
