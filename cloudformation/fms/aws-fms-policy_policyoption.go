package fms

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
	"github.com/awslabs/goformation/v5/cloudformation/utils"
)

var _ utils.Value[struct{}]

// Policy_PolicyOption AWS CloudFormation Resource (AWS::FMS::Policy.PolicyOption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policyoption.html
type Policy_PolicyOption struct {

	// NetworkFirewallPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policyoption.html#cfn-fms-policy-policyoption-networkfirewallpolicy
	NetworkFirewallPolicy *Policy_NetworkFirewallPolicy `json:"NetworkFirewallPolicy,omitempty"`

	// ThirdPartyFirewallPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policyoption.html#cfn-fms-policy-policyoption-thirdpartyfirewallpolicy
	ThirdPartyFirewallPolicy *Policy_ThirdPartyFirewallPolicy `json:"ThirdPartyFirewallPolicy,omitempty"`

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
func (r *Policy_PolicyOption) AWSCloudFormationType() string {
	return "AWS::FMS::Policy.PolicyOption"
}