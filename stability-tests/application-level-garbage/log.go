package main

import (
	"github.com/sadnetwork/sad/infrastructure/logger"
	"github.com/sadnetwork/sad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("APLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
