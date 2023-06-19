package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/arielmorelli/servus-api/route"
)

// Info struct to represent a new route
type Info struct {
	// json tag to de-serialize json body
	Route        string            `json:"route"`
	Methods      []string          `json:"methods"`
	ResponseCode int               `json:"response_code,default=200"`
	Headers      map[string]string `json:"headers"`
	Parameters   map[string]string `json:"parameters"`
	Response     any               `json:"response"`
}

func headersToMap(header http.Header) map[string]string {
	asMap := make(map[string]string)

	for name, values := range header {
		name = strings.ToLower(name)
		asMap[name] = strings.Join(values, ", ")
	}

	return asMap
}

func paramsToMap(header url.Values) map[string]string {
	asMap := make(map[string]string)

	for name, values := range header {
		name = strings.ToLower(name)
		asMap[name] = strings.Join(values, ", ")
	}

	return asMap
}

// RegisterAPIRoute parse the payload and stores a new route
// This method uses gin.Context to get the payload and process it
func RegisterAPIRoute(c *gin.Context) {
	routeInfo := Info{
		ResponseCode: 200,
	}
	if err := c.BindJSON(&routeInfo); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for _, method := range routeInfo.Methods {
		route.RegisterRoute(routeInfo.Route, method, routeInfo.ResponseCode, routeInfo.Headers, routeInfo.Parameters, routeInfo.Response)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// ShowInfo returns all routes as a json
func ShowInfo(c *gin.Context) {
	c.JSON(http.StatusOK, route.Routes)
}

// FindAndReturnRoute returns the registered values, if found, 404 otherwise
func FindAndReturnRoute(targetRoute string, c *gin.Context) {
	methodValue, found := route.FindRoute(targetRoute, c.Request.Method, headersToMap(c.Request.Header), paramsToMap(c.Request.URL.Query()))

	if !found {
		c.JSON(http.StatusNotFound, nil)
	}

	c.JSON(methodValue.ResponseCode, methodValue.Response)
}

// DynamicRouting is a custom routing on top of gin's routing system.
// This is necessary duo a limitation on Gin and dynamic routes and wildcard the app has only one dynamic route that handles all requests
// https://github.com/gin-gonic/gin/issues/1301
func DynamicRouting(c *gin.Context) {
	route := route.AsRouteName(c.Param("route"))
	if route == "_register" && c.Request.Method == "POST" {
		RegisterAPIRoute(c)
		return
	}
	if route == "_info" && c.Request.Method == "GET" {
		ShowInfo(c)
		return
	}

	FindAndReturnRoute(route, c)
}
