package resolver

import (
	"adjust/app"
	. "adjust/hasher"
	. "adjust/logger"
	. "adjust/provider"
	. "adjust/writer"
)

func ResolveAppDependency() (app.Hasher, *HttpProvider, *Logger, *StdOutWriter){
	hasher := NewHash()
	provider := NewHttpProvider()
	logger := NewLogger()
	writer := NewStdOutWriter()

	return hasher, provider, logger, writer
}
