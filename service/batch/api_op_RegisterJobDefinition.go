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

// Registers an Batch job definition.
func (c *Client) RegisterJobDefinition(ctx context.Context, params *RegisterJobDefinitionInput, optFns ...func(*Options)) (*RegisterJobDefinitionOutput, error) {
	if params == nil {
		params = &RegisterJobDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterJobDefinition", params, optFns, c.addOperationRegisterJobDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterJobDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for RegisterJobDefinition.
type RegisterJobDefinitionInput struct {

	// The name of the job definition to register. Up to 128 letters (uppercase and
	// lowercase), numbers, hyphens, and underscores are allowed.
	//
	// This member is required.
	JobDefinitionName *string

	// The type of job definition. For more information about multi-node parallel jobs,
	// see Creating a multi-node parallel job definition
	// (https://docs.aws.amazon.com/batch/latest/userguide/multi-node-job-def.html) in
	// the Batch User Guide. If the job is run on Fargate resources, then multinode
	// isn't supported.
	//
	// This member is required.
	Type types.JobDefinitionType

	// An object with various properties specific to single-node container-based jobs.
	// If the job definition's type parameter is container, then you must specify
	// either containerProperties or nodeProperties. If the job runs on Fargate
	// resources, then you must not specify nodeProperties; use only
	// containerProperties.
	ContainerProperties *types.ContainerProperties

	// An object with various properties specific to multi-node parallel jobs. If you
	// specify node properties for a job, it becomes a multi-node parallel job. For
	// more information, see Multi-node Parallel Jobs
	// (https://docs.aws.amazon.com/batch/latest/userguide/multi-node-parallel-jobs.html)
	// in the Batch User Guide. If the job definition's type parameter is container,
	// then you must specify either containerProperties or nodeProperties. If the job
	// runs on Fargate resources, then you must not specify nodeProperties; use
	// containerProperties instead.
	NodeProperties *types.NodeProperties

	// Default parameter substitution placeholders to set in the job definition.
	// Parameters are specified as a key-value pair mapping. Parameters in a SubmitJob
	// request override any corresponding parameter defaults from the job definition.
	Parameters map[string]string

	// The platform capabilities required by the job definition. If no value is
	// specified, it defaults to EC2. To run the job on Fargate resources, specify
	// FARGATE.
	PlatformCapabilities []types.PlatformCapability

	// Specifies whether to propagate the tags from the job or job definition to the
	// corresponding Amazon ECS task. If no value is specified, the tags are not
	// propagated. Tags can only be propagated to the tasks during task creation. For
	// tags with the same name, job tags are given priority over job definitions tags.
	// If the total number of combined tags from the job and job definition is over 50,
	// the job is moved to the FAILED state.
	PropagateTags bool

	// The retry strategy to use for failed jobs that are submitted with this job
	// definition. Any retry strategy that's specified during a SubmitJob operation
	// overrides the retry strategy defined here. If a job is terminated due to a
	// timeout, it isn't retried.
	RetryStrategy *types.RetryStrategy

	// The scheduling priority for jobs that are submitted with this job definition.
	// This will only affect jobs in job queues with a fair share policy. Jobs with a
	// higher scheduling priority will be scheduled before jobs with a lower scheduling
	// priority. The minimum supported value is 0 and the maximum supported value is
	// 9999.
	SchedulingPriority int32

	// The tags that you apply to the job definition to help you categorize and
	// organize your resources. Each tag consists of a key and an optional value. For
	// more information, see Tagging Amazon Web Services Resources
	// (https://docs.aws.amazon.com/batch/latest/userguide/using-tags.html) in Batch
	// User Guide.
	Tags map[string]string

	// The timeout configuration for jobs that are submitted with this job definition,
	// after which Batch terminates your jobs if they have not finished. If a job is
	// terminated due to a timeout, it isn't retried. The minimum value for the timeout
	// is 60 seconds. Any timeout configuration that's specified during a SubmitJob
	// operation overrides the timeout configuration defined here. For more
	// information, see Job Timeouts
	// (https://docs.aws.amazon.com/batch/latest/userguide/job_timeouts.html) in the
	// Batch User Guide.
	Timeout *types.JobTimeout

	noSmithyDocumentSerde
}

type RegisterJobDefinitionOutput struct {

	// The Amazon Resource Name (ARN) of the job definition.
	//
	// This member is required.
	JobDefinitionArn *string

	// The name of the job definition.
	//
	// This member is required.
	JobDefinitionName *string

	// The revision of the job definition.
	//
	// This member is required.
	Revision int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterJobDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterJobDefinition{}, middleware.After)
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
	if err = addOpRegisterJobDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterJobDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterJobDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "RegisterJobDefinition",
	}
}
