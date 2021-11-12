// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the alarm recommendations for a AWS Resilience Hub application.
func (c *Client) ListAlarmRecommendations(ctx context.Context, params *ListAlarmRecommendationsInput, optFns ...func(*Options)) (*ListAlarmRecommendationsOutput, error) {
	if params == nil {
		params = &ListAlarmRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAlarmRecommendations", params, optFns, c.addOperationListAlarmRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAlarmRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAlarmRecommendationsInput struct {

	// The Amazon Resource Name (ARN) of the assessment. The format for this ARN is:
	// arn:partition:dcps:region:account:app-assessment/app-id. For more information
	// about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	AssessmentArn *string

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// Null, or the token from a previous call to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAlarmRecommendationsOutput struct {

	// The alarm recommendations for an AWS Resilience Hub application, returned as an
	// object. This object includes application component names, descriptions,
	// information about whether a recommendation has already been implemented or not,
	// prerequisites, and more.
	//
	// This member is required.
	AlarmRecommendations []types.AlarmRecommendation

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAlarmRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAlarmRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAlarmRecommendations{}, middleware.After)
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
	if err = addOpListAlarmRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAlarmRecommendations(options.Region), middleware.Before); err != nil {
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

// ListAlarmRecommendationsAPIClient is a client that implements the
// ListAlarmRecommendations operation.
type ListAlarmRecommendationsAPIClient interface {
	ListAlarmRecommendations(context.Context, *ListAlarmRecommendationsInput, ...func(*Options)) (*ListAlarmRecommendationsOutput, error)
}

var _ ListAlarmRecommendationsAPIClient = (*Client)(nil)

// ListAlarmRecommendationsPaginatorOptions is the paginator options for
// ListAlarmRecommendations
type ListAlarmRecommendationsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAlarmRecommendationsPaginator is a paginator for ListAlarmRecommendations
type ListAlarmRecommendationsPaginator struct {
	options   ListAlarmRecommendationsPaginatorOptions
	client    ListAlarmRecommendationsAPIClient
	params    *ListAlarmRecommendationsInput
	nextToken *string
	firstPage bool
}

// NewListAlarmRecommendationsPaginator returns a new
// ListAlarmRecommendationsPaginator
func NewListAlarmRecommendationsPaginator(client ListAlarmRecommendationsAPIClient, params *ListAlarmRecommendationsInput, optFns ...func(*ListAlarmRecommendationsPaginatorOptions)) *ListAlarmRecommendationsPaginator {
	if params == nil {
		params = &ListAlarmRecommendationsInput{}
	}

	options := ListAlarmRecommendationsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAlarmRecommendationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAlarmRecommendationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAlarmRecommendations page.
func (p *ListAlarmRecommendationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAlarmRecommendationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListAlarmRecommendations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAlarmRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "ListAlarmRecommendations",
	}
}
