package api

import (
	"net/http"
	"net/url"
	"strings"

	models "github.com/arielmorelli/servus-api/models"
	route "github.com/arielmorelli/servus-api/route"
	"github.com/gin-gonic/gin"
)

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
	routeInfo := models.RegisterSchema{
		ResponseCode: 200,
	}
	if err := c.BindJSON(&routeInfo); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	route.RegisterRoute(routeInfo)
	// for _, method := range routeInfo.Methods {
	// 	route.RegisterRoute(routeInfo.Route, method, routeInfo.ResponseCode, routeInfo.Headers, routeInfo.Parameters, routeInfo.Response)
	// }

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
	deleteInfo := models.DeleteSchema{
		Methods: []string{},
	}
	if err := c.BindJSON(&deleteInfo); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if route.DeleteRoute(deleteInfo) {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusNotFound, nil)
	}

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
