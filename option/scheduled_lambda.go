package option

import (
	"github.com/rockset/rockset-go-client/openapi"
)

// ScheduledLambdaOptions is used to hold optional scheduled lambda settings
type ScheduledLambdaOptions struct {
	CreateRequest openapi.CreateScheduledLambdaRequest
	UpdateRequest openapi.UpdateScheduledLambdaRequest
}

type ScheduledLambdaOption func(p *ScheduledLambdaOptions)

func WithScheduledLambdaCreateRequest(r openapi.CreateScheduledLambdaRequest) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.CreateRequest = r
	}
}

func WithScheduledLambdaUpdateRequest(r openapi.UpdateScheduledLambdaRequest) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.UpdateRequest = r
	}
}

func WithScheduledLambdaTag(tag string) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.CreateRequest.Tag = &tag
	}
}

func WithScheduledLambdaTotalTimesToExecute(total int64) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.CreateRequest.TotalTimesToExecute = &total
		o.UpdateRequest.TotalTimesToExecute = &total
	}
}

func WithScheduledLambdaVersion(version string) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.CreateRequest.Version = &version
	}
}

func WithScheduledLambdaWebhookAuthHeader(header string) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.CreateRequest.WebhookAuthHeader = &header
		o.UpdateRequest.WebhookAuthHeader = &header
	}
}

func WithScheduledLambdaWebhookPayload(payload string) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.CreateRequest.WebhookPayload = &payload
		o.UpdateRequest.WebhookPayload = &payload
	}
}

func WithScheduledLambdaWebhookURL(url string) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.CreateRequest.WebhookUrl = &url
		o.UpdateRequest.WebhookUrl = &url
	}
}

func WithScheduledLambdaResumePermanentError(r bool) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.UpdateRequest.ResumePermanentError = &r
	}
}

func WithScheduledLambdaApikey(a string) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.UpdateRequest.Apikey = &a
	}
}