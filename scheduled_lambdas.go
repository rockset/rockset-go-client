package rockset

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// CreateScheduledLambda creates a new scheduled lambda, with an optional description.
//
// REST API documentation https://docs.rockset.com/documentation/reference/createscheduledlambda
func (rc *RockClient) CreateScheduledLambda(ctx context.Context, workspace, apikey string, cronString string, qlName string,
	options ...option.ScheduledLambdaOption) (openapi.ScheduledLambda, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ScheduledLambdaResponse

	q := rc.ScheduledLambdasApi.CreateScheduledLambda(ctx, workspace)
	req := openapi.NewCreateScheduledLambdaRequest(apikey, cronString, qlName)

	opts := option.ScheduledLambdaOptions{}
	for _, o := range options {
		o(&opts)
	}

	if opts.Tag != nil {
		req.Tag = opts.Tag
	}
	if opts.TotalTimesToExecute != nil {
		req.TotalTimesToExecute = opts.TotalTimesToExecute
	}
	if opts.Version != nil {
		req.Version= opts.Version
	}
	if opts.WebhookAuthHeader != nil {
		req.WebhookAuthHeader = opts.WebhookAuthHeader
	}
	if opts.WebhookPayload != nil {
		req.WebhookPayload = opts.WebhookPayload
	}
	if opts.WebhookUrl != nil {
		req.WebhookUrl = opts.WebhookUrl
	}

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.ScheduledLambda{}, err
	}

	log := zerolog.Ctx(ctx)
	log.Trace().Str("status", httpResp.Status).Msg("scheduled lambda created")

	return resp.GetData(), nil
}

// UpdateScheduledLambda updates an existing scheduled lambda, with an optional description.
//
// REST API documentation https://docs.rockset.com/documentation/reference/updatescheduledlambda
func (rc *RockClient) UpdateScheduledLambda(ctx context.Context, workspace string, scheduledLambdaRrn string,
options ...option.ScheduledLambdaOption) (openapi.ScheduledLambda, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ScheduledLambdaResponse

	q := rc.ScheduledLambdasApi.UpdateScheduledLambda(ctx, workspace, scheduledLambdaRrn)
	req := openapi.NewUpdateScheduledLambdaRequest()

	opts := option.ScheduledLambdaOptions{}
	for _, o := range options {
		o(&opts)
	}

	if opts.TotalTimesToExecute != nil {
		req.TotalTimesToExecute = opts.TotalTimesToExecute
	}
	if opts.WebhookAuthHeader != nil {
		req.WebhookAuthHeader = opts.WebhookAuthHeader
	}
	if opts.WebhookPayload != nil {
		req.WebhookPayload = opts.WebhookPayload
	}
	if opts.WebhookUrl != nil {
		req.WebhookUrl = opts.WebhookUrl
	}
	if opts.ResumePermanentError != nil {
		req.ResumePermanentError = opts.ResumePermanentError
	}

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.ScheduledLambda{}, err
	}

	log := zerolog.Ctx(ctx)
	log.Trace().Str("status", httpResp.Status).Msg("scheduled lambda updated")

	return resp.GetData(), nil
}

// DeleteScheduledLambda marks the scheduled lambda for deletion, which will take place in the background. Use the
// WaitUntilScheduledLambdaGone() call to block until the scheduled lambda has been deleted.
//
// REST API documentation https://docs.rockset.com/documentation/reference/deletescheduledlambda
func (rc *RockClient) DeleteScheduledLambda(ctx context.Context, workspace, scheduledLambdaRrn string) error {
	var err error
	var httpResp *http.Response

	q := rc.ScheduledLambdasApi.DeleteScheduledLambda(ctx, workspace, scheduledLambdaRrn)

	err = rc.Retry(ctx, func() error {
		_, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return err
	}

	log := zerolog.Ctx(ctx)
	log.Trace().Str("status", httpResp.Status).Msg("scheduled lambda deleted")

	return nil
}

// GetScheduledLambda gets details about a scheduled lambda.
//
// REST API documentation https://docs.rockset.com/documentation/reference/getscheduledlambda
func (rc *RockClient) GetScheduledLambda(ctx context.Context, workspace, scheduledLambdaRrn string) (openapi.ScheduledLambda, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ScheduledLambdaResponse
	log := zerolog.Ctx(ctx)

	q := rc.ScheduledLambdasApi.GetScheduledLambda(ctx, workspace, scheduledLambdaRrn)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.ScheduledLambda{}, err
	}

	data := resp.GetData()
	log.Debug().Str("rrn", data.GetRrn()).Msg("get scheduled lambda successful")

	return data, nil
}
