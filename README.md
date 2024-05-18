# viagoli




# Notes



### special note on *internal* directory
any packages which live under this directory can only be imported by code
inside the parent of the internal directory.
looking at other way any packages inside internal *can not be imported by code outside of our project*



### Go file server nice features
- it sanitizes all request paths by running them through path.Clean() function before searching for a path.
remove (. and .. elements from URL path) - stops directory traversal attacks
- Range requests are fully supported.This is great if your application is serving large files
and you want to support resumable downloads.


## http handler function

a handler is an object which satisfies the *http.Handler* interface

    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }

In it's simplest form a handler could look somethign like this
    
    type home struct {}
    func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("This is my home page))
    }

registering then using servemux

    mux := http.NewServeMux()
    mux.Handle("/", &home())


the http.HandlerFunc() adapter tranforms normal function which accept a reponsewriter and request
into a handler like so 

    mux := http.NewServeMux()
    mux.Handle("/, http.HandlerFunc(home))

## a typical request chain handling

When our server receives a new HTTP request, it calls
the servemux’s ServeHTTP() method. This looks up the relevant handler based on the
request URL path, and in turn calls that handler’s ServeHTTP() method. You can think of a Go
web application as a chain of ServeHTTP() methods being called one after another

### requests are handled concurrently

There is one more thing that’s really important to point out: all incoming HTTP requests are
served in their own goroutine. For busy servers, this means it’s very likely that the code in or
called by your handlers will be running concurrently. While this helps make Go blazingly fast,
the downside is that you need to be aware of (and protect against) race conditions when
accessing shared resources from your handlers.


### Logging

logging to standard output and error and redirecting all logs to corresponding files
during runtime.

    go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log


## Dependency Injection
Most web applications need multiple dependencies 
that their handlers need to access, such as database connection pool, 
centralized error handlers and template caches.

For applications where all your handlers are in the same package, like ours, a neat way to
inject dependencies is to put them into a custom application struct, and then define your
handler functions as methods against application.

    type application struct {
        errorLog *log.Logger
        infoLog  *log.Logger
    }


# Dynamic HTML templates
Dynamic content escaping
The html/template package automatically escapes any data that is yielded between {{ }}
tags. This behavior is hugely helpful in avoiding cross-site scripting (XSS) attacks, and is the
reason that you should use the html/template package instead of the more generic
text/template package that Go also provides.

Template actions and functions
In this section we’re going to look at the template actions and functions that Go provides.
We’ve already talked about some of the actions — {{define}}, {{template}} and {{block}}
— but there are three more which you can use to control the display of dynamic data —
{{if}}, {{with}} and {{range}}.
Action Description
{{if .Foo}} C1 {{else}} C2 {{end}} If .Foo is not empty then render the content C1,
otherwise render the content C2.
{{with .Foo}} C1 {{else}} C2 {{end}} If .Foo is not empty, then set dot to the value of
.Foo and render the content C1, otherwise render the
content C2.
{{range .Foo}} C1 {{else}} C2 {{end}} If the length of .Foo is greater than zero then loop
over each element, setting dot to the value of each
element and rendering the content C1. If the length of
.Foo is zero then render the content C2. The underlying
type of .Foo must be an array, slice, map, or channel.