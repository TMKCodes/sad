package grpcclient

import (
	"github.com/sadnetwork/sad/infrastructure/logger"
	"github.com/sadnetwork/sad/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
