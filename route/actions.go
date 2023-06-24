package route

import (
	"fmt"
	"log"
	"strings"

	models "github.com/arielmorelli/servus-api/models"
)

// Routes is a global variable to store all routes
var Routes models.Route = make(models.Route)

// RegisterRoute register a new route in the global context
func RegisterRoute(routeInfo models.RegisterSchema) {

	route := AsRouteName(routeInfo.Route)

	methodValue := models.MethodValue{
		Headers:      routeInfo.Headers,
		Parameters:   routeInfo.Parameters,
		ResponseCode: routeInfo.ResponseCode,
		Response:     routeInfo.Response,
	}

	for _, method := range routeInfo.Methods {
		method = strings.ToLower(method)

		_, exists := Routes[route]
		if !exists {
			Routes[route] = models.Method{}
		}

		_, exists = Routes[route][method]
		if !exists {
			Routes[route][method] = []models.MethodValue{}
		}

		if len(routeInfo.Headers) == 0 && len(routeInfo.Parameters) == 0 {
			Routes[route][method] = append(Routes[route][method], methodValue)
		} else {
			Routes[route][method] = append([]models.MethodValue{methodValue}, Routes[route][method]...)
		}

		log.Println("Added route", fmt.Sprintf("/%s,", route), "for method", strings.ToUpper(method))
	}
}

// FindRoute find route in method in registered Routes
func FindRoute(route, method string, headers, params map[string]string) (models.MethodValue, bool) {
	route = AsRouteName(route)
	method = strings.ToLower(method)

	_, exists := Routes[route]
	if !exists {
		return models.MethodValue{}, false
	}

	_, exists = Routes[route][method]
	if !exists {
		return models.MethodValue{}, false
	}

	for _, methodValue := range Routes[route][method] {
		if isMapSubset(headers, methodValue.Headers) && isMapSubset(params, methodValue.Parameters) {
			return methodValue, true
		}
	}

	return models.MethodValue{}, false
}

// DeleteRoute delete a given route and a method, if provided
func DeleteRoute(deleteInfo models.DeleteSchema) bool {
	route := AsRouteName(deleteInfo.Route)

	mapOfMethods, exists := Routes[route]
	if !exists {
		return false
	}

	if len(deleteInfo.Methods) == 0 {
		// Delete all methods
		delete(Routes, route)
		log.Println("Deleted", fmt.Sprintf("/%s,", route))
		return true
	}

	for _, method := range deleteInfo.Methods {
		method = strings.ToLower(method)

		_, exists = mapOfMethods[method]
		if !exists {
			return false
		}

		delete(mapOfMethods, method)
		log.Println("Deleted method", strings.ToUpper(method), "for route", fmt.Sprintf("/%s,", route), "for method")
	}
	return true
}
