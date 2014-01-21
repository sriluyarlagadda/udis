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

The router supports regular expressions for the url matching and it can match multiple values at a time.

    router.Get("/:name{[a-z]*}/:newValue{[1-4]*}", complexHandler)
    
    func complexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s %s times", r.FormValue("name"), r.FormValue("newValue"))
    }
    
Similarly, to intercept a HTTP Post request:

    router.Post("/", postIndexHandler)
    	
    func postIndexHandler(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintf(w, "hello world from http postpost")
    }

We attach the router to the server through:
    http.ListenAndServe(":8080", router)


