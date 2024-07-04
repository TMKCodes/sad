package standalone

import (
	"github.com/sadnetwork/sad/infrastructure/logger"
	"github.com/sadnetwork/sad/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
