package ginwebserver

import "github.com/gin-gonic/gin"

type resetHandler interface {
	Reset()
	IsInterfaceNil() bool
}

// WebServerHandler defines the basic actions that a web server should be able to do
type WebServerHandler interface {
	Start()
	Close() error
	IsInterfaceNil() bool
}

// MiddlewareProcessor defines a processor used internally by the web server when processing requests
type MiddlewareProcessor interface {
	MiddlewareHandlerFunc() gin.HandlerFunc
	IsInterfaceNil() bool
}

// MainApiHandler interface defines methods that can be used from `elrondFacade` context variable
type MainApiHandler interface {
	RestApiInterface() string
	RestAPIServerDebugMode() bool
	PprofEnabled() bool
	IsInterfaceNil() bool
}
