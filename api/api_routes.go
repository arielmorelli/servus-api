package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/arielmorelli/servus-api/route"
	"github.com/gin-gonic/gin"
)

// RegisterSchema struct to represent a new route
type RegisterSchema struct {
	Route           string            `json:"route"`
	Methods         []string          `json:"methods"`
	ResponseCode    int               `json:"response_code,default=200"`
	Headers         map[string]string `json:"headers"`
	Parameters      map[string]string `json:"parameters"`
	Response        any               `json:"response"`
	ResponseHeaders map[string]string `json:"response_headers"`
}

// DeleteSchema struct to represent a new route
type DeleteSchema struct {
	Route   string   `json:"route"`
	Methods []string `json:"methods"`
}

func headersOrParametersToMap[K http.Header | url.Values](header K) map[string]string {
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
	routeInfo := RegisterSchema{
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

// ShowAPIInfo returns all routes as a json
func ShowAPIInfo(c *gin.Context) {
	c.JSON(http.StatusOK, route.Routes)
}

// FindAPIRoute returns the registered values, if found, 404 otherwise
func FindAPIRoute(targetRoute string, c *gin.Context) {
	methodValue, found := route.FindRoute(targetRoute, c.Request.Method, headersOrParametersToMap(c.Request.Header), headersOrParametersToMap(c.Request.URL.Query()))

	if !found {
		c.JSON(http.StatusNotFound, nil)
	}

	c.JSON(methodValue.ResponseCode, methodValue.Response)
}

// DeleteAPIRoute the registered values, if found, 404 otherwise
func DeleteAPIRoute(c *gin.Context) {
	deleteInfo := RegisterSchema{
		Methods: []string{},
	}
	if err := c.BindJSON(&deleteInfo); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if len(deleteInfo.Methods) > 0 {
		for _, method := range deleteInfo.Methods {
			route.DeleteRoute(deleteInfo.Route, method)
		}
	} else {
		route.DeleteRoute(deleteInfo.Route, "")
	}

	c.JSON(http.StatusOK, nil)
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
		ShowAPIInfo(c)
		return
	}
	if route == "_remove" && c.Request.Method == "PUT" {
		DeleteAPIRoute(c)
		return
	}

	FindAPIRoute(route, c)
}
