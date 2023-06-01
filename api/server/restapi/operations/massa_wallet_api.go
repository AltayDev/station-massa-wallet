// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewMassaWalletAPI creates a new MassaWallet instance
func NewMassaWalletAPI(spec *loads.Document) *MassaWalletAPI {
	return &MassaWalletAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		BinProducer: runtime.ByteStreamProducer(),
		CSSProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("css producer has not yet been implemented")
		}),
		HTMLProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("html producer has not yet been implemented")
		}),
		JsProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("js producer has not yet been implemented")
		}),
		JSONProducer: runtime.JSONProducer(),
		TextWebpProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("textWebp producer has not yet been implemented")
		}),

		AccountListHandler: AccountListHandlerFunc(func(params AccountListParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountList has not yet been implemented")
		}),
		BackupAccountHandler: BackupAccountHandlerFunc(func(params BackupAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation BackupAccount has not yet been implemented")
		}),
		CreateAccountHandler: CreateAccountHandlerFunc(func(params CreateAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateAccount has not yet been implemented")
		}),
		DeleteAccountHandler: DeleteAccountHandlerFunc(func(params DeleteAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteAccount has not yet been implemented")
		}),
		ExportAccountFileHandler: ExportAccountFileHandlerFunc(func(params ExportAccountFileParams) middleware.Responder {
			return middleware.NotImplemented("operation ExportAccountFile has not yet been implemented")
		}),
		GetAccountHandler: GetAccountHandlerFunc(func(params GetAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation GetAccount has not yet been implemented")
		}),
		ImportAccountHandler: ImportAccountHandlerFunc(func(params ImportAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation ImportAccount has not yet been implemented")
		}),
		SignHandler: SignHandlerFunc(func(params SignParams) middleware.Responder {
			return middleware.NotImplemented("operation Sign has not yet been implemented")
		}),
		TradeRollsHandler: TradeRollsHandlerFunc(func(params TradeRollsParams) middleware.Responder {
			return middleware.NotImplemented("operation TradeRolls has not yet been implemented")
		}),
		TransferCoinHandler: TransferCoinHandlerFunc(func(params TransferCoinParams) middleware.Responder {
			return middleware.NotImplemented("operation TransferCoin has not yet been implemented")
		}),
		DefaultPageHandler: DefaultPageHandlerFunc(func(params DefaultPageParams) middleware.Responder {
			return middleware.NotImplemented("operation DefaultPage has not yet been implemented")
		}),
		WebHandler: WebHandlerFunc(func(params WebParams) middleware.Responder {
			return middleware.NotImplemented("operation Web has not yet been implemented")
		}),
		WebAppHandler: WebAppHandlerFunc(func(params WebAppParams) middleware.Responder {
			return middleware.NotImplemented("operation WebApp has not yet been implemented")
		}),
	}
}

/*MassaWalletAPI Thyra plugin - Massa Wallet */
type MassaWalletAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// BinProducer registers a producer for the following mime types:
	//   - application/octet-stream
	//   - image/png
	BinProducer runtime.Producer
	// CSSProducer registers a producer for the following mime types:
	//   - text/css
	CSSProducer runtime.Producer
	// HTMLProducer registers a producer for the following mime types:
	//   - text/html
	HTMLProducer runtime.Producer
	// JsProducer registers a producer for the following mime types:
	//   - text/javascript
	JsProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer
	// TextWebpProducer registers a producer for the following mime types:
	//   - text/webp
	TextWebpProducer runtime.Producer

	// AccountListHandler sets the operation handler for the account list operation
	AccountListHandler AccountListHandler
	// BackupAccountHandler sets the operation handler for the backup account operation
	BackupAccountHandler BackupAccountHandler
	// CreateAccountHandler sets the operation handler for the create account operation
	CreateAccountHandler CreateAccountHandler
	// DeleteAccountHandler sets the operation handler for the delete account operation
	DeleteAccountHandler DeleteAccountHandler
	// ExportAccountFileHandler sets the operation handler for the export account file operation
	ExportAccountFileHandler ExportAccountFileHandler
	// GetAccountHandler sets the operation handler for the get account operation
	GetAccountHandler GetAccountHandler
	// ImportAccountHandler sets the operation handler for the import account operation
	ImportAccountHandler ImportAccountHandler
	// SignHandler sets the operation handler for the sign operation
	SignHandler SignHandler
	// TradeRollsHandler sets the operation handler for the trade rolls operation
	TradeRollsHandler TradeRollsHandler
	// TransferCoinHandler sets the operation handler for the transfer coin operation
	TransferCoinHandler TransferCoinHandler
	// DefaultPageHandler sets the operation handler for the default page operation
	DefaultPageHandler DefaultPageHandler
	// WebHandler sets the operation handler for the web operation
	WebHandler WebHandler
	// WebAppHandler sets the operation handler for the web app operation
	WebAppHandler WebAppHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *MassaWalletAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *MassaWalletAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *MassaWalletAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MassaWalletAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MassaWalletAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MassaWalletAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MassaWalletAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MassaWalletAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MassaWalletAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MassaWalletAPI
func (o *MassaWalletAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}
	if o.CSSProducer == nil {
		unregistered = append(unregistered, "CSSProducer")
	}
	if o.HTMLProducer == nil {
		unregistered = append(unregistered, "HTMLProducer")
	}
	if o.JsProducer == nil {
		unregistered = append(unregistered, "JsProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.TextWebpProducer == nil {
		unregistered = append(unregistered, "TextWebpProducer")
	}

	if o.AccountListHandler == nil {
		unregistered = append(unregistered, "AccountListHandler")
	}
	if o.BackupAccountHandler == nil {
		unregistered = append(unregistered, "BackupAccountHandler")
	}
	if o.CreateAccountHandler == nil {
		unregistered = append(unregistered, "CreateAccountHandler")
	}
	if o.DeleteAccountHandler == nil {
		unregistered = append(unregistered, "DeleteAccountHandler")
	}
	if o.ExportAccountFileHandler == nil {
		unregistered = append(unregistered, "ExportAccountFileHandler")
	}
	if o.GetAccountHandler == nil {
		unregistered = append(unregistered, "GetAccountHandler")
	}
	if o.ImportAccountHandler == nil {
		unregistered = append(unregistered, "ImportAccountHandler")
	}
	if o.SignHandler == nil {
		unregistered = append(unregistered, "SignHandler")
	}
	if o.TradeRollsHandler == nil {
		unregistered = append(unregistered, "TradeRollsHandler")
	}
	if o.TransferCoinHandler == nil {
		unregistered = append(unregistered, "TransferCoinHandler")
	}
	if o.DefaultPageHandler == nil {
		unregistered = append(unregistered, "DefaultPageHandler")
	}
	if o.WebHandler == nil {
		unregistered = append(unregistered, "WebHandler")
	}
	if o.WebAppHandler == nil {
		unregistered = append(unregistered, "WebAppHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MassaWalletAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MassaWalletAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *MassaWalletAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *MassaWalletAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *MassaWalletAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer
		case "image/png":
			result["image/png"] = o.BinProducer
		case "text/css":
			result["text/css"] = o.CSSProducer
		case "text/html":
			result["text/html"] = o.HTMLProducer
		case "text/javascript":
			result["text/javascript"] = o.JsProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		case "text/webp":
			result["text/webp"] = o.TextWebpProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MassaWalletAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the massa wallet API
func (o *MassaWalletAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MassaWalletAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/accounts"] = NewAccountList(o.context, o.AccountListHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/accounts/{nickname}/backup"] = NewBackupAccount(o.context, o.BackupAccountHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/accounts/{nickname}"] = NewCreateAccount(o.context, o.CreateAccountHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/api/accounts/{nickname}"] = NewDeleteAccount(o.context, o.DeleteAccountHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/accounts/{nickname}/exportFile"] = NewExportAccountFile(o.context, o.ExportAccountFileHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/accounts/{nickname}"] = NewGetAccount(o.context, o.GetAccountHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/api/accounts"] = NewImportAccount(o.context, o.ImportAccountHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/accounts/{nickname}/sign"] = NewSign(o.context, o.SignHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/accounts/{nickname}/rolls"] = NewTradeRolls(o.context, o.TradeRollsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/accounts/{nickname}/transfer"] = NewTransferCoin(o.context, o.TransferCoinHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"][""] = NewDefaultPage(o.context, o.DefaultPageHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/web/{resource}"] = NewWeb(o.context, o.WebHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/web-app/{resource}"] = NewWebApp(o.context, o.WebAppHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MassaWalletAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MassaWalletAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MassaWalletAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MassaWalletAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *MassaWalletAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}