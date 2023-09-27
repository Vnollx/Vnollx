package main

import (
	"Vnollx/common"
	"Vnollx/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//db := common.InitDB()
	////defer db.Close()
	common.InitDB()
	r := gin.Default()
	router.Routes(r)
	r.Run()
}
