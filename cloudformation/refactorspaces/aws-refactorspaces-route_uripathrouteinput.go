// Code generated by "go generate". Please don't change this file directly.

package refactorspaces

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
	"github.com/awslabs/goformation/v6/cloudformation/utils"
)

var _ utils.Value[struct{}]

// Route_UriPathRouteInput AWS CloudFormation Resource (AWS::RefactorSpaces::Route.UriPathRouteInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-uripathrouteinput.html
type Route_UriPathRouteInput struct {

	// ActivationState AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-uripathrouteinput.html#cfn-refactorspaces-route-uripathrouteinput-activationstate
	ActivationState string `json:"ActivationState"`

	// IncludeChildPaths AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-uripathrouteinput.html#cfn-refactorspaces-route-uripathrouteinput-includechildpaths
	IncludeChildPaths *utils.Value[bool] `json:"IncludeChildPaths,omitempty"`

	// Methods AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-uripathrouteinput.html#cfn-refactorspaces-route-uripathrouteinput-methods
	Methods *[]string `json:"Methods,omitempty"`

	// SourcePath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-uripathrouteinput.html#cfn-refactorspaces-route-uripathrouteinput-sourcepath
	SourcePath *string `json:"SourcePath,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Route_UriPathRouteInput) AWSCloudFormationType() string {
	return "AWS::RefactorSpaces::Route.UriPathRouteInput"
}
