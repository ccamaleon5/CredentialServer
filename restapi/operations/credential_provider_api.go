// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ccamaleon5/CredentialServer/restapi/operations/credential"
	"github.com/ccamaleon5/CredentialServer/restapi/operations/did"
)

// NewCredentialProviderAPI creates a new CredentialProvider instance
func NewCredentialProviderAPI(spec *loads.Document) *CredentialProviderAPI {
	return &CredentialProviderAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		XMLConsumer:         runtime.XMLConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		XMLProducer:         runtime.XMLProducer(),
		CredentialCreateCredentialHandler: credential.CreateCredentialHandlerFunc(func(params credential.CreateCredentialParams) middleware.Responder {
			return middleware.NotImplemented("operation CredentialCreateCredential has not yet been implemented")
		}),
		CredentialGetCredentialByIDHandler: credential.GetCredentialByIDHandlerFunc(func(params credential.GetCredentialByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation CredentialGetCredentialByID has not yet been implemented")
		}),
		CredentialRenewalCredentialHandler: credential.RenewalCredentialHandlerFunc(func(params credential.RenewalCredentialParams) middleware.Responder {
			return middleware.NotImplemented("operation CredentialRenewalCredential has not yet been implemented")
		}),
		CredentialRevokeCredentialHandler: credential.RevokeCredentialHandlerFunc(func(params credential.RevokeCredentialParams) middleware.Responder {
			return middleware.NotImplemented("operation CredentialRevokeCredential has not yet been implemented")
		}),
		DidValidateDidHandler: did.ValidateDidHandlerFunc(func(params did.ValidateDidParams) middleware.Responder {
			return middleware.NotImplemented("operation DidValidateDid has not yet been implemented")
		}),
		CredentialVerifyCredentialHandler: credential.VerifyCredentialHandlerFunc(func(params credential.VerifyCredentialParams) middleware.Responder {
			return middleware.NotImplemented("operation CredentialVerifyCredential has not yet been implemented")
		}),
	}
}

/*CredentialProviderAPI This is a Provider Credential Server that validates, signs, generates, revokes and updates credential to identify persons, institutions and objects.

The Provider Credential Server sign a credential using its own keys, it is configurable.

The Provider Credential needs to manage its own repository of credentials, default is smart contract that is deployed when server init.

The Credentials are verifiable against blockchain default, but you can configure and choose your proof and revocation list service. */
type CredentialProviderAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// XMLConsumer registers a consumer for a "application/xml" mime type
	XMLConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer
	// XMLProducer registers a producer for a "application/xml" mime type
	XMLProducer runtime.Producer

	// CredentialCreateCredentialHandler sets the operation handler for the create credential operation
	CredentialCreateCredentialHandler credential.CreateCredentialHandler
	// CredentialGetCredentialByIDHandler sets the operation handler for the get credential by Id operation
	CredentialGetCredentialByIDHandler credential.GetCredentialByIDHandler
	// CredentialRenewalCredentialHandler sets the operation handler for the renewal credential operation
	CredentialRenewalCredentialHandler credential.RenewalCredentialHandler
	// CredentialRevokeCredentialHandler sets the operation handler for the revoke credential operation
	CredentialRevokeCredentialHandler credential.RevokeCredentialHandler
	// DidValidateDidHandler sets the operation handler for the validate did operation
	DidValidateDidHandler did.ValidateDidHandler
	// CredentialVerifyCredentialHandler sets the operation handler for the verify credential operation
	CredentialVerifyCredentialHandler credential.VerifyCredentialHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *CredentialProviderAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CredentialProviderAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CredentialProviderAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CredentialProviderAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CredentialProviderAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CredentialProviderAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CredentialProviderAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CredentialProviderAPI
func (o *CredentialProviderAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.XMLConsumer == nil {
		unregistered = append(unregistered, "XMLConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.XMLProducer == nil {
		unregistered = append(unregistered, "XMLProducer")
	}

	if o.CredentialCreateCredentialHandler == nil {
		unregistered = append(unregistered, "credential.CreateCredentialHandler")
	}

	if o.CredentialGetCredentialByIDHandler == nil {
		unregistered = append(unregistered, "credential.GetCredentialByIDHandler")
	}

	if o.CredentialRenewalCredentialHandler == nil {
		unregistered = append(unregistered, "credential.RenewalCredentialHandler")
	}

	if o.CredentialRevokeCredentialHandler == nil {
		unregistered = append(unregistered, "credential.RevokeCredentialHandler")
	}

	if o.DidValidateDidHandler == nil {
		unregistered = append(unregistered, "did.ValidateDidHandler")
	}

	if o.CredentialVerifyCredentialHandler == nil {
		unregistered = append(unregistered, "credential.VerifyCredentialHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CredentialProviderAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CredentialProviderAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *CredentialProviderAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *CredentialProviderAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "application/xml":
			result["application/xml"] = o.XMLConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *CredentialProviderAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		case "application/xml":
			result["application/xml"] = o.XMLProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *CredentialProviderAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the credential provider API
func (o *CredentialProviderAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CredentialProviderAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/credential"] = credential.NewCreateCredential(o.context, o.CredentialCreateCredentialHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/credential/{credentialId}"] = credential.NewGetCredentialByID(o.context, o.CredentialGetCredentialByIDHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/credential/{credentialId}"] = credential.NewRenewalCredential(o.context, o.CredentialRenewalCredentialHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/credential/{credentialId}"] = credential.NewRevokeCredential(o.context, o.CredentialRevokeCredentialHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/did/{did}"] = did.NewValidateDid(o.context, o.DidValidateDidHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/credential/verify"] = credential.NewVerifyCredential(o.context, o.CredentialVerifyCredentialHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CredentialProviderAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *CredentialProviderAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *CredentialProviderAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *CredentialProviderAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}