A simple URL dispatcher in Go
================================

The Go http package in the standard library has support only for Static URL paths. The purpose of this project is to add support for Dynamic matching of URL's to Go Standard library.


To provide dynamic dispatch support, Import the udis package into your project.


    import("github.com/sriluyarlagadda/udis")



Create a new Router object, this object takes care of dispatching the http request to handlers, based on the http method type, and the url request type.

    router := udis.NewRouter()
    

To create a HTTP Get Request handler which accepts any value of type ***value***
