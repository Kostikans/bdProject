// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/kostikans/bdProject/restapi/operations"
)

//go:generate swagger generate server --target ../../bdProject --name Forum --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.ForumAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ForumAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.BinConsumer = runtime.ByteStreamConsumer()
	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.ClearHandler == nil {
		api.ClearHandler = operations.ClearHandlerFunc(func(params operations.ClearParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.Clear has not yet been implemented")
		})
	}
	if api.ForumCreateHandler == nil {
		api.ForumCreateHandler = operations.ForumCreateHandlerFunc(func(params operations.ForumCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ForumCreate has not yet been implemented")
		})
	}
	if api.ForumGetOneHandler == nil {
		api.ForumGetOneHandler = operations.ForumGetOneHandlerFunc(func(params operations.ForumGetOneParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ForumGetOne has not yet been implemented")
		})
	}
	if api.ForumGetThreadsHandler == nil {
		api.ForumGetThreadsHandler = operations.ForumGetThreadsHandlerFunc(func(params operations.ForumGetThreadsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ForumGetThreads has not yet been implemented")
		})
	}
	if api.ForumGetUsersHandler == nil {
		api.ForumGetUsersHandler = operations.ForumGetUsersHandlerFunc(func(params operations.ForumGetUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ForumGetUsers has not yet been implemented")
		})
	}
	if api.PostGetOneHandler == nil {
		api.PostGetOneHandler = operations.PostGetOneHandlerFunc(func(params operations.PostGetOneParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostGetOne has not yet been implemented")
		})
	}
	if api.PostUpdateHandler == nil {
		api.PostUpdateHandler = operations.PostUpdateHandlerFunc(func(params operations.PostUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUpdate has not yet been implemented")
		})
	}
	if api.PostsCreateHandler == nil {
		api.PostsCreateHandler = operations.PostsCreateHandlerFunc(func(params operations.PostsCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostsCreate has not yet been implemented")
		})
	}
	if api.StatusHandler == nil {
		api.StatusHandler = operations.StatusHandlerFunc(func(params operations.StatusParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.Status has not yet been implemented")
		})
	}
	if api.ThreadCreateHandler == nil {
		api.ThreadCreateHandler = operations.ThreadCreateHandlerFunc(func(params operations.ThreadCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ThreadCreate has not yet been implemented")
		})
	}
	if api.ThreadGetOneHandler == nil {
		api.ThreadGetOneHandler = operations.ThreadGetOneHandlerFunc(func(params operations.ThreadGetOneParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ThreadGetOne has not yet been implemented")
		})
	}
	if api.ThreadGetPostsHandler == nil {
		api.ThreadGetPostsHandler = operations.ThreadGetPostsHandlerFunc(func(params operations.ThreadGetPostsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ThreadGetPosts has not yet been implemented")
		})
	}
	if api.ThreadUpdateHandler == nil {
		api.ThreadUpdateHandler = operations.ThreadUpdateHandlerFunc(func(params operations.ThreadUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ThreadUpdate has not yet been implemented")
		})
	}
	if api.ThreadVoteHandler == nil {
		api.ThreadVoteHandler = operations.ThreadVoteHandlerFunc(func(params operations.ThreadVoteParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ThreadVote has not yet been implemented")
		})
	}
	if api.UserCreateHandler == nil {
		api.UserCreateHandler = operations.UserCreateHandlerFunc(func(params operations.UserCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UserCreate has not yet been implemented")
		})
	}
	if api.UserGetOneHandler == nil {
		api.UserGetOneHandler = operations.UserGetOneHandlerFunc(func(params operations.UserGetOneParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UserGetOne has not yet been implemented")
		})
	}
	if api.UserUpdateHandler == nil {
		api.UserUpdateHandler = operations.UserUpdateHandlerFunc(func(params operations.UserUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UserUpdate has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
