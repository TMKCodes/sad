package main

import (
	"github.com/sadnetwork/sad/infrastructure/logger"
	"github.com/sadnetwork/sad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RPIC")
	spawn      = panics.GoroutineWrapperFunc(log)
)
