// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmeetings

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon Chime SDK meeting in the specified media Region, with
// attendees. For more information about specifying media Regions, see Amazon Chime
// SDK Media Regions
// (https://docs.aws.amazon.com/chime/latest/dg/chime-sdk-meetings-regions.html) in
// the Amazon Chime Developer Guide. For more information about the Amazon Chime
// SDK, see Using the Amazon Chime SDK
// (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html) in the Amazon
// Chime Developer Guide.
func (c *Client) CreateMeetingWithAttendees(ctx context.Context, params *CreateMeetingWithAttendeesInput, optFns ...func(*Options)) (*CreateMeetingWithAttendeesOutput, error) {
	if params == nil {
		params = &CreateMeetingWithAttendeesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMeetingWithAttendees", params, optFns, c.addOperationCreateMeetingWithAttendeesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMeetingWithAttendeesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMeetingWithAttendeesInput struct {

	// The attendee information, including attendees' IDs and join tokens.
	//
	// This member is required.
	Attendees []types.CreateAttendeeRequestItem

	// The unique identifier for the client request. Use a different token for
	// different meetings.
	//
	// This member is required.
	ClientRequestToken *string

	// The external meeting ID.
	//
	// This member is required.
	ExternalMeetingId *string

	// The Region in which to create the meeting.
	//
	// This member is required.
	MediaRegion *string

	// Reserved.
	MeetingHostId *string

	// The configuration for resource targets to receive notifications when meeting and
	// attendee events occur.
	NotificationsConfiguration *types.NotificationsConfiguration

	noSmithyDocumentSerde
}

type CreateMeetingWithAttendeesOutput struct {

	// The attendee information, including attendees' IDs and join tokens.
	Attendees []types.Attendee

	// If the action fails for one or more of the attendees in the request, a list of
	// the attendees is returned, along with error codes and error messages.
	Errors []types.CreateAttendeeError

	// The meeting information, including the meeting ID and MediaPlacement.
	Meeting *types.Meeting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMeetingWithAttendeesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMeetingWithAttendees{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMeetingWithAttendees{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateMeetingWithAttendeesMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMeetingWithAttendeesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMeetingWithAttendees(options.Region), middleware.Before); err != nil {
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
	return nil
}

type idempotencyToken_initializeOpCreateMeetingWithAttendees struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMeetingWithAttendees) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMeetingWithAttendees) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMeetingWithAttendeesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMeetingWithAttendeesInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMeetingWithAttendeesMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMeetingWithAttendees{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMeetingWithAttendees(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateMeetingWithAttendees",
	}
}
