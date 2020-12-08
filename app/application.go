package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	router *gin.Engine
	lock   = &sync.Mutex{}
)

func init() {
	router, _ = generateEngine()
}

func StartApplication() {
	mapURLs()
	router.Run(":8080")
}

func generateEngine() (*gin.Engine, error) {
	if router == nil {
		lock.Lock()
		defer lock.Unlock()
		router = gin.Default()
	} else {
		return nil, errors.New("engine already exist")
	}
	return router, nil
}
