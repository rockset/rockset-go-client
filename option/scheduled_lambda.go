package option

// ScheduledLambdaOptions is used to hold optional scheduled lambda settings
type ScheduledLambdaOptions struct {
	Tag          *string
	TotalTimesToExecute *int64
	Version *string
	WebhookAuthHeader *string
	WebhookPayload *string
	WebhookUrl *string
	ResumePermanentError *bool
}

type ScheduledLambdaOption func(p *ScheduledLambdaOptions)

func WithScheduledLambdaTag(tag string) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.Tag = &tag
	}
}

func WithScheduledLambdaTotalTimesToExecute(total int64) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.TotalTimesToExecute = &total
	}
}

func WithScheduledLambdaVersion(version string) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.Version = &version
	}
}

func WithScheduledLambdaWebhookAuthHeader(header string) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.WebhookAuthHeader = &header
	}
}

func WithScheduledLambdaWebhookPayload(payload string) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.WebhookPayload = &payload
	}
}

func WithScheduledLambdaWebhookUrl(url string) ScheduledLambdaOption {
	return func(o *ScheduledLambdaOptions) {
		o.WebhookUrl = &url
	}
}