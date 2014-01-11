//package udis implements a simple URL Dispatcher for go
package udis

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

//the default regular expression to use if one is not provided
const DEFAULT_REGEXP = "[a-z]*"

//Router is responsible for receiving http requests and
//routing it to appropriate handlers
type Router struct {
	*http.ServeMux
	routes []Route
}

func (r *Router) String() {
	for _, route := range r.routes {
		fmt.Println(route)
	}
}

//Route represents a single route which is registered with the router
type Route struct {
	regExpressionPattern *regexp.Regexp
	methodType           string
	handlerFunc          http.HandlerFunc
	urlParams            map[string]string
}

func (r *Route) String() {
	fmt.Sprintf("pattern:", r.regExpressionPattern, "methodType:", r.methodType, "handlerFunc:", r.handlerFunc)
}

//
func (r *Route) routeMatch(request *http.Request) bool {

	fmt.Println("request method type:", request.Method)

	if request.Method == r.methodType {
		fmt.Println("matching string:", request.URL.Path, " with ", r.regExpressionPattern.String())
		isMatch := r.regExpressionPattern.MatchString(request.URL.Path)
		fmt.Println("isMatch?:", isMatch)
		return isMatch
	}

	return false
}

func (r *Route) populateForm(request *http.Request) error {
	matches := r.regExpressionPattern.FindStringSubmatch(request.URL.Path)
	fmt.Println("url path:", request.URL.Path)
	fmt.Println("matches", matches)
	err := request.ParseForm()
	if err != nil {
		return err
	}

	j := 1
	for key, _ := range r.urlParams {
		fmt.Println("key:", key, " value:", matches[j])
		request.Form.Add(key, matches[j])
		j++
	}
	return nil
}

//create a new router for url dispatch
//
func NewRouter() *Router {
	routes := make([]Route, 0)
	router := &Router{http.NewServeMux(), routes}
	router.ServeMux.HandleFunc("/", router.routesHandler())
	return router
}

//given a url pattern process extracts relavant data
func processPattern(pattern string) (map[string]string, string) {
	//split all urls into parts
	urlParts := strings.Split(pattern, "/")

	urlParameters := make(map[string]string)

	for _, part := range urlParts {
		if part != "" {
			fmt.Println("parts:", part)
			var param, regExp string

			//if each url part starts with a : then it is a parameter
			if strings.HasPrefix(part, ":") {
				colonIndex := strings.Index(part, ":")

				//if it has a custom regular expression followed by the parameter
				if strings.Contains(part, "{") && strings.Contains(part, "}") {

					openBracket := strings.Index(part, "{")
					closeBracket := strings.Index(part, "}")

					//extract the regular expression and the url parameter
					regExp = part[openBracket+1 : closeBracket]
					param = part[colonIndex+1 : openBracket]

				} else {

					param = part[colonIndex+1:]
					regExp = DEFAULT_REGEXP
				}

				regExp = "(" + regExp + ")"
				fmt.Println("param:", param)
				fmt.Println("regexp:", regExp)
				urlParameters[param] = regExp
				pattern = strings.Replace(pattern, part, regExp, 1)
			} else {
				//ignore that part
			}
		}
	}
	//append ^ at the beginning and $ at the end to make exact match
	pattern = strings.Join([]string{"^", pattern, "$"}, "")

	return urlParameters, pattern
}

// Get function registers a http Handler function for a particular pattern for
// http get Requests
func (router *Router) Get(pattern string, f http.HandlerFunc) {

	fmt.Println("pattern before:", pattern)
	urlParams, pattern := processPattern(pattern)

	fmt.Println("pattern after:", pattern)

	regExpPattern := regexp.MustCompile(pattern)
	route := Route{regExpPattern, "GET", f, urlParams}

	router.appendRoute(route)
}

// Post function registers a http Handler function for a particular pattern for
// http Post Requests
func (router *Router) Post(pattern string, f http.HandlerFunc) {
	urlParams, pattern := processPattern(pattern)
	regExpPattern := regexp.MustCompile(pattern)
	fmt.Println("pattern after:", pattern)
	route := Route{regExpPattern, "POST", f, urlParams}
	router.appendRoute(route)
}

// appendRoute appends all the routes to the router
func (router *Router) appendRoute(route Route) {
	router.routes = append(router.routes, route)
}

// routesHandler accepts Http request and initiates dispatch handling logic
func (router *Router) routesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleRoutes(router, w, r)
	}

}

func handleRoutes(router *Router, writer http.ResponseWriter, request *http.Request) {
	routeMatched := false

	//for every route registered with the router, check if there is any match
	for _, route := range router.routes {
		if route.routeMatch(request) {
			routeMatched = true
			err := route.populateForm(request)

			if err != nil {
				panic(err)
			}
			route.handlerFunc(writer, request)
			break
		}
	}

	//is not part of any route
	if routeMatched != true {

	}

}
