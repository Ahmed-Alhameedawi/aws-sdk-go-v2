// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the properties associated with an asynchronous batch translation job
// including name, ID, status, source and target languages, input/output S3
// buckets, and so on.
func (c *Client) DescribeTextTranslationJob(ctx context.Context, params *DescribeTextTranslationJobInput, optFns ...func(*Options)) (*DescribeTextTranslationJobOutput, error) {
	if params == nil {
		params = &DescribeTextTranslationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTextTranslationJob", params, optFns, c.addOperationDescribeTextTranslationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTextTranslationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTextTranslationJobInput struct {

	// The identifier that Amazon Translate generated for the job. The
	// StartTextTranslationJob operation returns this identifier in its response.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type DescribeTextTranslationJobOutput struct {

	// An object that contains the properties associated with an asynchronous batch
	// translation job.
	TextTranslationJobProperties *types.TextTranslationJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTextTranslationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTextTranslationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTextTranslationJob{}, middleware.After)
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
	if err = addOpDescribeTextTranslationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTextTranslationJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTextTranslationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "DescribeTextTranslationJob",
	}
}
