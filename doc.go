// Package rockset provides a go client to interact with the Rockset database.
//
// The client uses the Rockset REST API https://docs.rockset.com/rest-api/,
// and is an OpenAPI generated code by https://openapi-generator.tech/.
//
// It provides convenience functions for all API calls to make the generated client easier to use,
// by wrapping the API calls in methods that require passing a context.Context and having
// all required arguments in the method signature, so that as many errors as possible are caught at compile time
// instead of at runtime. It uses functional options for any optional arguments.
// The convenience methods return the payload data from the models package, to reduce the indirection.
//
// All methods also automatically retries any retryable error returned by the Rockset API,
// using exponential back-off. The retryable errors are defined in rockset.RetryableErrors.
//
// If a zerolog logger is set in the context, the methods will log to it. E.g.
//
//	console := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
//	log := zerolog.New(console).Level(zerolog.TraceLevel).With().Timestamp().Logger()
//	ctx := log.WithContext(context.Background())
//
//	rc, err := rockset.NewClient()
//	if err != nil {
//	    log.Fatalf("failed to create RockClient: %v," err)
//	}
//
//	wsName := "commons"
//	ws, err := rc.GetWorkspace(ctx, wsName)
//	if err != nil {
//	    log.Fatalf("failed to get workspace %s: %v", wsName, err)
//	}
//
// Example log output
//
//	2021-05-28T13:11:46-07:00 TRC api call curation d=467.371958ms
//	2021-05-28T13:11:46-07:00 DBG total duration d=467.538875ms
//	2021-05-28T13:11:46-07:00 DBG get workspace successful name=commons
package rockset
