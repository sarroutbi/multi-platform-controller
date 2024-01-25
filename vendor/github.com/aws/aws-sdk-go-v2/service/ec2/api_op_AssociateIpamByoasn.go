// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates your Autonomous System Number (ASN) with a BYOIP CIDR that you own
// in the same Amazon Web Services Region. For more information, see Tutorial:
// Bring your ASN to IPAM (https://docs.aws.amazon.com/vpc/latest/ipam/tutorials-byoasn.html)
// in the Amazon VPC IPAM guide. After the association succeeds, the ASN is
// eligible for advertisement. You can view the association with DescribeByoipCidrs (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeByoipCidrs.html)
// . You can advertise the CIDR with AdvertiseByoipCidr (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AdvertiseByoipCidr.html)
// .
func (c *Client) AssociateIpamByoasn(ctx context.Context, params *AssociateIpamByoasnInput, optFns ...func(*Options)) (*AssociateIpamByoasnOutput, error) {
	if params == nil {
		params = &AssociateIpamByoasnInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateIpamByoasn", params, optFns, c.addOperationAssociateIpamByoasnMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateIpamByoasnOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateIpamByoasnInput struct {

	// A public 2-byte or 4-byte ASN.
	//
	// This member is required.
	Asn *string

	// The BYOIP CIDR you want to associate with an ASN.
	//
	// This member is required.
	Cidr *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type AssociateIpamByoasnOutput struct {

	// The ASN and BYOIP CIDR association.
	AsnAssociation *types.AsnAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateIpamByoasnMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssociateIpamByoasn{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateIpamByoasn{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateIpamByoasn"); err != nil {
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
	if err = addOpAssociateIpamByoasnValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateIpamByoasn(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateIpamByoasn(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateIpamByoasn",
	}
}
