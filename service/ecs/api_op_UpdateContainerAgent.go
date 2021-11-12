// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the Amazon ECS container agent on a specified container instance.
// Updating the Amazon ECS container agent doesn't interrupt running tasks or
// services on the container instance. The process for updating the agent differs
// depending on whether your container instance was launched with the Amazon
// ECS-optimized AMI or another operating system. The UpdateContainerAgent API
// isn't supported for container instances using the Amazon ECS-optimized Amazon
// Linux 2 (arm64) AMI. To update the container agent, you can update the ecs-init
// package. This updates the agent. For more information, see Updating the Amazon
// ECS container agent
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/agent-update-ecs-ami.html)
// in the Amazon Elastic Container Service Developer Guide. The
// UpdateContainerAgent API requires an Amazon ECS-optimized AMI or Amazon Linux
// AMI with the ecs-init service installed and running. For help updating the
// Amazon ECS container agent on other operating systems, see Manually updating the
// Amazon ECS container agent
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html#manually_update_agent)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) UpdateContainerAgent(ctx context.Context, params *UpdateContainerAgentInput, optFns ...func(*Options)) (*UpdateContainerAgentOutput, error) {
	if params == nil {
		params = &UpdateContainerAgentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateContainerAgent", params, optFns, c.addOperationUpdateContainerAgentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateContainerAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContainerAgentInput struct {

	// The container instance ID or full ARN entries for the container instance where
	// you would like to update the Amazon ECS container agent.
	//
	// This member is required.
	ContainerInstance *string

	// The short name or full Amazon Resource Name (ARN) of the cluster that your
	// container instance is running on. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string

	noSmithyDocumentSerde
}

type UpdateContainerAgentOutput struct {

	// The container instance that the container agent was updated for.
	ContainerInstance *types.ContainerInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateContainerAgentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateContainerAgent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateContainerAgent{}, middleware.After)
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
	if err = addOpUpdateContainerAgentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContainerAgent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateContainerAgent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "UpdateContainerAgent",
	}
}
