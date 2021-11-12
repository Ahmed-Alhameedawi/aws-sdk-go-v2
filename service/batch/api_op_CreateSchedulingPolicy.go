// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Batch scheduling policy.
func (c *Client) CreateSchedulingPolicy(ctx context.Context, params *CreateSchedulingPolicyInput, optFns ...func(*Options)) (*CreateSchedulingPolicyOutput, error) {
	if params == nil {
		params = &CreateSchedulingPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSchedulingPolicy", params, optFns, c.addOperationCreateSchedulingPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSchedulingPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSchedulingPolicyInput struct {

	// The name of the scheduling policy. Up to 128 letters (uppercase and lowercase),
	// numbers, hyphens, and underscores are allowed.
	//
	// This member is required.
	Name *string

	// The fair share policy of the scheduling policy.
	FairsharePolicy *types.FairsharePolicy

	// The tags that you apply to the scheduling policy to help you categorize and
	// organize your resources. Each tag consists of a key and an optional value. For
	// more information, see Tagging Amazon Web Services Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in Amazon Web
	// Services General Reference. These tags can be updated or removed using the
	// TagResource
	// (https://docs.aws.amazon.com/batch/latest/APIReference/API_TagResource.html) and
	// UntagResource
	// (https://docs.aws.amazon.com/batch/latest/APIReference/API_UntagResource.html)
	// API operations.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSchedulingPolicyOutput struct {

	// The Amazon Resource Name (ARN) of the scheduling policy. The format is
	// aws:Partition:batch:Region:Account:scheduling-policy/Name . For example,
	// aws:aws:batch:us-west-2:012345678910:scheduling-policy/MySchedulingPolicy.
	//
	// This member is required.
	Arn *string

	// The name of the scheduling policy.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSchedulingPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSchedulingPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSchedulingPolicy{}, middleware.After)
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
	if err = addOpCreateSchedulingPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSchedulingPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSchedulingPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "CreateSchedulingPolicy",
	}
}
