// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a cost anomaly subscription.
func (c *Client) DeleteAnomalySubscription(ctx context.Context, params *DeleteAnomalySubscriptionInput, optFns ...func(*Options)) (*DeleteAnomalySubscriptionOutput, error) {
	if params == nil {
		params = &DeleteAnomalySubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAnomalySubscription", params, optFns, addOperationDeleteAnomalySubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAnomalySubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAnomalySubscriptionInput struct {

	// The unique identifier of the cost anomaly subscription that you want to delete.
	//
	// This member is required.
	SubscriptionArn *string
}

type DeleteAnomalySubscriptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAnomalySubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAnomalySubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAnomalySubscription{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteAnomalySubscriptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAnomalySubscription(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAnomalySubscription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "DeleteAnomalySubscription",
	}
}