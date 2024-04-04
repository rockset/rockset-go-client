package rockset

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// CreateScheduledLambda creates a new scheduled lambda, with optional arguments.
//
// REST API documentation https://docs.rockset.com/documentation/reference/createscheduledlambda
func (rc *RockClient) CreateScheduledLambda(ctx context.Context, workspace, apikey string, cronString string, qlName string,
	options ...option.ScheduledLambdaOption) (openapi.ScheduledLambda, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ScheduledLambdaResponse

	q := rc.ScheduledLambdasApi.CreateScheduledLambda(ctx, workspace)

	opts := option.ScheduledLambdaOptions{}
	for _, o := range options {
		o(&opts)
	}
	opts.CreateRequest.Apikey = apikey;
	opts.CreateRequest.CronString = cronString;
	opts.CreateRequest.QlName = qlName;

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(opts.CreateRequest).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.ScheduledLambda{}, err
	}

	log := zerolog.Ctx(ctx)
	log.Trace().Str("status", httpResp.Status).Msg("scheduled lambda created")

	return resp.GetData(), nil
}

// UpdateScheduledLambda updates an existing scheduled lambda, with optional arguments.
//
// REST API documentation https://docs.rockset.com/documentation/reference/updatescheduledlambda
func (rc *RockClient) UpdateScheduledLambda(ctx context.Context, workspace string, scheduledLambdaRRN string,
options ...option.ScheduledLambdaOption) (openapi.ScheduledLambda, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ScheduledLambdaResponse

	q := rc.ScheduledLambdasApi.UpdateScheduledLambda(ctx, workspace, scheduledLambdaRRN)

	opts := option.ScheduledLambdaOptions{}
	for _, o := range options {
		o(&opts)
	}

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(opts.UpdateRequest).Execute()

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
func (rc *RockClient) DeleteScheduledLambda(ctx context.Context, workspace, scheduledLambdaRRN string) error {
	var err error
	var httpResp *http.Response

	q := rc.ScheduledLambdasApi.DeleteScheduledLambda(ctx, workspace, scheduledLambdaRRN)

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
func (rc *RockClient) GetScheduledLambda(ctx context.Context, workspace, scheduledLambdaRRN string) (openapi.ScheduledLambda, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ScheduledLambdaResponse
	log := zerolog.Ctx(ctx)

	q := rc.ScheduledLambdasApi.GetScheduledLambda(ctx, workspace, scheduledLambdaRRN)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.ScheduledLambda{}, err
	}

	data := resp.GetData()
	log.Debug().Str("RRN", data.GetRrn()).Msg("get scheduled lambda successful")

	return data, nil
}
