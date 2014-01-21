A simple URL dispatcher in Go
================================

The Go http package in the standard library has support only for Static URL paths. The purpose of this project is to add support for Dynamic matching of URL's to Go Standard library.


To provide dynamic dispatch support, Import the udis package into your project.


    import("github.com/sriluyarlagadda/udis")



Create a new Router object, this object takes care of dispatching the http request to handlers, based on the http method type, and the url request type.

    router := udis.NewRouter()
    

To create a HTTP Get Request handler which accepts any value of type ***value***

    router.Get("/:value", nameHandler)
    
    func nameHandler(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintf(w, "hello %s", r.FormValue("name"))
    }
    
Here the router registers the given path, associates the corresponding handler function to that particular path, the ***value*** passed in the url is append to the form values of the http request.

