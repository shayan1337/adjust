package resolver

import (
	"adjust/app"
	. "adjust/hasher"
	. "adjust/logger"
	. "adjust/provider"
)

func ResolveAppDependency() (app.Hasher, *HttpProvider, *Logger){
	hasher := NewHash()
	provider := NewHttpProvider()
	logger := NewLogger()

	return hasher, provider, logger
}
