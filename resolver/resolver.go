package resolver

import (
	"adjust/app"
	. "adjust/hasher"
	. "adjust/logger"
	. "adjust/provider"
	"log"
)

func ResolveAppDependency() (app.Hasher, *HttpProvider, log.Logger){
	hasher := NewHash()
	provider := NewHttpProvider()
	logger := GetLogger()

	return hasher, provider, logger
}
