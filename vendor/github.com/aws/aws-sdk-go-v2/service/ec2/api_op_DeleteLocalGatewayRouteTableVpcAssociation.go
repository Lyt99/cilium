// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteLocalGatewayRouteTableVpcAssociationInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the association.
	//
	// LocalGatewayRouteTableVpcAssociationId is a required field
	LocalGatewayRouteTableVpcAssociationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLocalGatewayRouteTableVpcAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLocalGatewayRouteTableVpcAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLocalGatewayRouteTableVpcAssociationInput"}

	if s.LocalGatewayRouteTableVpcAssociationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LocalGatewayRouteTableVpcAssociationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteLocalGatewayRouteTableVpcAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the association.
	LocalGatewayRouteTableVpcAssociation *LocalGatewayRouteTableVpcAssociation `locationName:"localGatewayRouteTableVpcAssociation" type:"structure"`
}

// String returns the string representation
func (s DeleteLocalGatewayRouteTableVpcAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteLocalGatewayRouteTableVpcAssociation = "DeleteLocalGatewayRouteTableVpcAssociation"

// DeleteLocalGatewayRouteTableVpcAssociationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified association between a VPC and local gateway route table.
//
//    // Example sending a request using DeleteLocalGatewayRouteTableVpcAssociationRequest.
//    req := client.DeleteLocalGatewayRouteTableVpcAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteLocalGatewayRouteTableVpcAssociation
func (c *Client) DeleteLocalGatewayRouteTableVpcAssociationRequest(input *DeleteLocalGatewayRouteTableVpcAssociationInput) DeleteLocalGatewayRouteTableVpcAssociationRequest {
	op := &aws.Operation{
		Name:       opDeleteLocalGatewayRouteTableVpcAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLocalGatewayRouteTableVpcAssociationInput{}
	}

	req := c.newRequest(op, input, &DeleteLocalGatewayRouteTableVpcAssociationOutput{})
	return DeleteLocalGatewayRouteTableVpcAssociationRequest{Request: req, Input: input, Copy: c.DeleteLocalGatewayRouteTableVpcAssociationRequest}
}

// DeleteLocalGatewayRouteTableVpcAssociationRequest is the request type for the
// DeleteLocalGatewayRouteTableVpcAssociation API operation.
type DeleteLocalGatewayRouteTableVpcAssociationRequest struct {
	*aws.Request
	Input *DeleteLocalGatewayRouteTableVpcAssociationInput
	Copy  func(*DeleteLocalGatewayRouteTableVpcAssociationInput) DeleteLocalGatewayRouteTableVpcAssociationRequest
}

// Send marshals and sends the DeleteLocalGatewayRouteTableVpcAssociation API request.
func (r DeleteLocalGatewayRouteTableVpcAssociationRequest) Send(ctx context.Context) (*DeleteLocalGatewayRouteTableVpcAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLocalGatewayRouteTableVpcAssociationResponse{
		DeleteLocalGatewayRouteTableVpcAssociationOutput: r.Request.Data.(*DeleteLocalGatewayRouteTableVpcAssociationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLocalGatewayRouteTableVpcAssociationResponse is the response type for the
// DeleteLocalGatewayRouteTableVpcAssociation API operation.
type DeleteLocalGatewayRouteTableVpcAssociationResponse struct {
	*DeleteLocalGatewayRouteTableVpcAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLocalGatewayRouteTableVpcAssociation request.
func (r *DeleteLocalGatewayRouteTableVpcAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
