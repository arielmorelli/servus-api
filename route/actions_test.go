package route

import (
	"testing"

	models "github.com/arielmorelli/servus-api/models"
)

// RegisterRoute
func TestRegisterRouteSingleMethod(t *testing.T) {
	// given
	Routes = make(models.Route)

	// when
	RegisterRoute(models.RegisterSchema{
		Route:        "/a",
		Methods:      []string{"GET"},
		ResponseCode: 204,
		Response:     "Hello world",
	})

	// then
	if len(Routes) != 1 {
		t.Fatalf("Routes was not set")
	}

	_, exists := Routes["a"]
	if !exists {
		t.Fatalf("Route /a was not created")
	}

	values, exists := Routes["a"]["get"]
	if !exists {
		t.Fatalf("Method get for route /a was not created")
	}
	if len(values) != 1 {
		t.Fatalf("Get list not set correctly")
	}

	value := values[0]
	if len(value.Headers) != 0 {
		t.Fatalf("Header not set up correctly")
	}

	if len(value.Parameters) != 0 {
		t.Fatalf("Parameters not set up correctly")
	}

	if value.ResponseCode != 204 {
		t.Fatalf("ResponseCode not set up correctly")
	}

	if value.Response != "Hello world" {
		t.Fatalf("Response not set up correctly")
	}
}

func TestRegisterRouteMultipleRoutes(t *testing.T) {
	// given
	Routes = make(models.Route)

	// when
	RegisterRoute(models.RegisterSchema{
		Route:        "/a",
		Methods:      []string{"GET"},
		ResponseCode: 204,
		Response:     "Hello world",
	})

	RegisterRoute(models.RegisterSchema{
		Route:        "/b",
		Methods:      []string{"GET"},
		ResponseCode: 204,
		Response:     "Hello world",
	})

	// then
	if len(Routes) != 2 {
		t.Fatalf("Routes was not set")
	}

	_, exists := Routes["a"]
	if !exists {
		t.Fatalf("Route /a was not created")
	}

	_, exists = Routes["b"]
	if !exists {
		t.Fatalf("Route /b was not created")
	}

}

func TestRegisterRouteMultipleMethods(t *testing.T) {
	// given
	Routes = make(models.Route)

	// when
	RegisterRoute(models.RegisterSchema{
		Route:        "/a",
		Methods:      []string{"GET"},
		ResponseCode: 204,
		Response:     "Hello world",
	})
	RegisterRoute(models.RegisterSchema{
		Route:        "/a",
		Methods:      []string{"POST"},
		ResponseCode: 204,
		Response:     "Hello world",
	})

	// then
	if len(Routes) != 1 {
		t.Fatalf("Routes was not set")
	}

	_, exists := Routes["a"]
	if !exists {
		t.Fatalf("Route /a was not created")
	}

	_, exists = Routes["a"]["get"]
	if !exists {
		t.Fatalf("Method get for route /a was not created")
	}

	_, exists = Routes["a"]["post"]
	if !exists {
		t.Fatalf("Method get for route /a was not created")
	}

}

func TestRegisterRouteAppendEmptyHeaderAndParameterInTheEnd(t *testing.T) {
	// given
	Routes = make(models.Route)

	// when
	RegisterRoute(models.RegisterSchema{
		Route:        "/a",
		Methods:      []string{"GET"},
		Headers:      map[string]string{"a": "b"},
		ResponseCode: 204,
		Response:     "Hello world",
	})
	RegisterRoute(models.RegisterSchema{
		Route:        "/a",
		Methods:      []string{"GET"},
		ResponseCode: 204,
		Response:     "Hello world",
	})

	// then
	values, _ := Routes["a"]["get"]
	if len(values[0].Headers) == 0 {
		t.Fatalf("First element has empty header")
	}
	if len(values[1].Headers) != 0 {
		t.Fatalf("Second element doesn't have empty header")
	}
}

func TestRegisterRouteAppendMethodWithHeaderAndParameterInTheBeggining(t *testing.T) {
	// given
	Routes = make(models.Route)

	// when
	RegisterRoute(models.RegisterSchema{
		Route:        "/a",
		Methods:      []string{"GET"},
		ResponseCode: 204,
		Response:     "Hello world",
	})
	RegisterRoute(models.RegisterSchema{
		Route:        "/a",
		Methods:      []string{"GET"},
		Headers:      map[string]string{"a": "b"},
		ResponseCode: 204,
		Response:     "Hello world",
	})

	// then
	values, _ := Routes["a"]["get"]
	if len(values[0].Headers) == 0 {
		t.Fatalf("First element has empty header")
	}
	if len(values[1].Headers) != 0 {
		t.Fatalf("Second element doesn't have empty header")
	}
}

func TestRegisterRouteStripName(t *testing.T) {
	// given
	Routes = make(models.Route)

	// when
	RegisterRoute(models.RegisterSchema{
		Route:        "/a",
		Methods:      []string{"GET"},
		ResponseCode: 204,
		Response:     "Hello world",
	})
	RegisterRoute(models.RegisterSchema{
		Route:        "/b",
		Methods:      []string{"GET"},
		ResponseCode: 204,
		Response:     "Hello world",
	})
	RegisterRoute(models.RegisterSchema{
		Route:        "/c",
		Methods:      []string{"GET"},
		ResponseCode: 204,
		Response:     "Hello world",
	})

	// then
	_, exists := Routes["a"]
	if !exists {
		t.Fatalf("Route /a was not created")
	}
	_, exists = Routes["b"]
	if !exists {
		t.Fatalf("Route /a was not created")
	}
	_, exists = Routes["c"]
	if !exists {
		t.Fatalf("Route /a was not created")
	}
}

// FindRoute
func TestFindRouteEmptyRoutes(t *testing.T) {
	// given
	Routes = make(models.Route)
	var emptyMap map[string]string

	// when
	_, found := FindRoute("/a", "GET", emptyMap, emptyMap)

	// then
	if found {
		t.Fatalf("Route /a should not be found")
	}

}

func TestFindRouteRouteNotFound(t *testing.T) {
	// given
	Routes = make(models.Route)
	Routes["a/b"] = make(models.Method)
	var emptyMap map[string]string

	// when
	_, found := FindRoute("/a", "GET", emptyMap, emptyMap)

	// then
	if found {
		t.Fatalf("Route /a should not be found")
	}
}

func TestFindRouteMethodNotFound(t *testing.T) {
	// given
	Routes = make(models.Route)
	Routes["a/b"] = make(models.Method)
	Routes["a/b"]["get"] = make([]models.MethodValue, 0)
	var emptyMap map[string]string

	// when
	_, found := FindRoute("/a", "POST", emptyMap, emptyMap)

	// then
	if found {
		t.Fatalf("Route /a should not be found")
	}
}

func TestFindRouteWithParametersButWithoutHeaders(t *testing.T) {
	// given
	Routes = make(models.Route)
	Routes["a"] = make(models.Method)
	Routes["a"]["get"] = make([]models.MethodValue, 0)
	Routes["a"]["get"] = append(Routes["a"]["get"], models.MethodValue{
		Headers:      map[string]string{"a": "a"},
		Parameters:   map[string]string{},
		ResponseCode: 200,
		Response:     "test",
	})
	var emptyMap map[string]string

	// when
	_, found := FindRoute("/a", "get", emptyMap, emptyMap)

	// then
	if found {
		t.Fatalf("Route /a should not be found")
	}
}

func TestFindRouteWithHeaderButWithoutParameters(t *testing.T) {
	// given
	Routes = make(models.Route)
	Routes["a"] = make(models.Method)
	Routes["a"]["get"] = make([]models.MethodValue, 0)
	Routes["a"]["get"] = append(Routes["a"]["get"], models.MethodValue{
		Headers:      map[string]string{},
		Parameters:   map[string]string{"a": "a"},
		ResponseCode: 200,
		Response:     "test",
	})
	var emptyMap map[string]string

	// when
	_, found := FindRoute("/a", "get", emptyMap, emptyMap)

	// then
	if found {
		t.Fatalf("Route /a should not be found")
	}
}

func TestFindRouteWithEmptyParamaAndGeaders(t *testing.T) {
	// given
	Routes = make(models.Route)
	Routes["a"] = make(models.Method)
	Routes["a"]["get"] = make([]models.MethodValue, 0)
	Routes["a"]["get"] = append(Routes["a"]["get"], models.MethodValue{
		Headers:      map[string]string{},
		Parameters:   map[string]string{},
		ResponseCode: 200,
		Response:     "test",
	})
	var emptyMap map[string]string

	// when
	_, found := FindRoute("/a", "get", emptyMap, emptyMap)

	// then
	if !found {
		t.Fatalf("Method value should be found")
	}
}

func TestFindRouteMatchParamaAndGeaders(t *testing.T) {
	// given
	Routes = make(models.Route)
	Routes["a"] = make(models.Method)
	Routes["a"]["get"] = make([]models.MethodValue, 0)
	Routes["a"]["get"] = append(Routes["a"]["get"], models.MethodValue{
		Headers:      map[string]string{"a": "a"},
		Parameters:   map[string]string{"b": "b"},
		ResponseCode: 200,
		Response:     "test",
	})

	// when
	_, found := FindRoute("/a", "get", map[string]string{"a": "a"}, map[string]string{"b": "b"})

	// then
	if !found {
		t.Fatalf("Method value should be found")
	}
}

// DeleteRoute
func TestDeleteRouteEmptyRoutes(t *testing.T) {
	// given
	Routes = make(models.Route)

	// when + then
	deleteSchema := models.DeleteSchema{
		Route:   "/a",
		Methods: []string{"GET"},
	}
	if DeleteRoute(deleteSchema) {
		t.Fatalf("Route should not exist")
	}

}

func TestDeleteRouteMethodNotFound(t *testing.T) {
	// given
	Routes = make(models.Route)
	Routes["a"] = make(models.Method)
	Routes["a"]["get"] = make([]models.MethodValue, 0)
	Routes["a"]["get"] = append(Routes["a"]["get"], models.MethodValue{
		Headers:      map[string]string{},
		Parameters:   map[string]string{},
		ResponseCode: 200,
		Response:     "test",
	})

	// when + then
	deleteSchema := models.DeleteSchema{
		Route:   "/a",
		Methods: []string{"POST"},
	}
	if DeleteRoute(deleteSchema) {
		t.Fatalf("Method should not exist")
	}
}

func TestDeleteRoute(t *testing.T) {
	// given
	Routes = make(models.Route)
	Routes["a"] = make(models.Method)
	Routes["a"]["get"] = make([]models.MethodValue, 0)
	Routes["a"]["get"] = append(Routes["a"]["get"], models.MethodValue{
		Headers:      map[string]string{},
		Parameters:   map[string]string{},
		ResponseCode: 200,
		Response:     "test",
	})

	// when + then
	deleteSchema := models.DeleteSchema{
		Route:   "/a",
		Methods: []string{"get"},
	}
	if !DeleteRoute(deleteSchema) {
		t.Fatalf("Method should not exist")
	}
}
