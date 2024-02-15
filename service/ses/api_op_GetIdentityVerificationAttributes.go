// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Given a list of identities (email addresses and/or domains), returns the
// verification status and (for domain identities) the verification token for each
// identity. The verification status of an email address is "Pending" until the
// email address owner clicks the link within the verification email that Amazon
// SES sent to that address. If the email address owner clicks the link within 24
// hours, the verification status of the email address changes to "Success". If the
// link is not clicked within 24 hours, the verification status changes to
// "Failed." In that case, to verify the email address, you must restart the
// verification process from the beginning. For domain identities, the domain's
// verification status is "Pending" as Amazon SES searches for the required TXT
// record in the DNS settings of the domain. When Amazon SES detects the record,
// the domain's verification status changes to "Success". If Amazon SES is unable
// to detect the record within 72 hours, the domain's verification status changes
// to "Failed." In that case, to verify the domain, you must restart the
// verification process from the beginning. This operation is throttled at one
// request per second and can only get verification attributes for up to 100
// identities at a time.
func (c *Client) GetIdentityVerificationAttributes(ctx context.Context, params *GetIdentityVerificationAttributesInput, optFns ...func(*Options)) (*GetIdentityVerificationAttributesOutput, error) {
	if params == nil {
		params = &GetIdentityVerificationAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdentityVerificationAttributes", params, optFns, c.addOperationGetIdentityVerificationAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdentityVerificationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the Amazon SES verification status of a list of
// identities. For domain identities, this request also returns the verification
// token. For information about verifying identities with Amazon SES, see the
// Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html)
// .
type GetIdentityVerificationAttributesInput struct {

	// A list of identities.
	//
	// This member is required.
	Identities []string

	noSmithyDocumentSerde
}

// The Amazon SES verification status of a list of identities. For domain
// identities, this response also contains the verification token.
type GetIdentityVerificationAttributesOutput struct {

	// A map of Identities to IdentityVerificationAttributes objects.
	//
	// This member is required.
	VerificationAttributes map[string]types.IdentityVerificationAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIdentityVerificationAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetIdentityVerificationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetIdentityVerificationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetIdentityVerificationAttributes"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetIdentityVerificationAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityVerificationAttributes(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// GetIdentityVerificationAttributesAPIClient is a client that implements the
// GetIdentityVerificationAttributes operation.
type GetIdentityVerificationAttributesAPIClient interface {
	GetIdentityVerificationAttributes(context.Context, *GetIdentityVerificationAttributesInput, ...func(*Options)) (*GetIdentityVerificationAttributesOutput, error)
}

var _ GetIdentityVerificationAttributesAPIClient = (*Client)(nil)

// IdentityExistsWaiterOptions are waiter options for IdentityExistsWaiter
type IdentityExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// IdentityExistsWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, IdentityExistsWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetIdentityVerificationAttributesInput, *GetIdentityVerificationAttributesOutput, error) (bool, error)
}

// IdentityExistsWaiter defines the waiters for IdentityExists
type IdentityExistsWaiter struct {
	client GetIdentityVerificationAttributesAPIClient

	options IdentityExistsWaiterOptions
}

// NewIdentityExistsWaiter constructs a IdentityExistsWaiter.
func NewIdentityExistsWaiter(client GetIdentityVerificationAttributesAPIClient, optFns ...func(*IdentityExistsWaiterOptions)) *IdentityExistsWaiter {
	options := IdentityExistsWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = identityExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &IdentityExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for IdentityExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *IdentityExistsWaiter) Wait(ctx context.Context, params *GetIdentityVerificationAttributesInput, maxWaitDur time.Duration, optFns ...func(*IdentityExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for IdentityExists waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *IdentityExistsWaiter) WaitForOutput(ctx context.Context, params *GetIdentityVerificationAttributesInput, maxWaitDur time.Duration, optFns ...func(*IdentityExistsWaiterOptions)) (*GetIdentityVerificationAttributesOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetIdentityVerificationAttributes(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for IdentityExists waiter")
}

func identityExistsStateRetryable(ctx context.Context, input *GetIdentityVerificationAttributesInput, output *GetIdentityVerificationAttributesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("VerificationAttributes.*.VerificationStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Success"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected string value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetIdentityVerificationAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetIdentityVerificationAttributes",
	}
}
