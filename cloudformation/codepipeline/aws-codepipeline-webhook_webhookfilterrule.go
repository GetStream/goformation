// Code generated by "go generate". Please don't change this file directly.

package codepipeline

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
	"github.com/awslabs/goformation/v6/cloudformation/utils"
)

var _ utils.Value[struct{}]

// Webhook_WebhookFilterRule AWS CloudFormation Resource (AWS::CodePipeline::Webhook.WebhookFilterRule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-webhook-webhookfilterrule.html
type Webhook_WebhookFilterRule struct {

	// JsonPath AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-webhook-webhookfilterrule.html#cfn-codepipeline-webhook-webhookfilterrule-jsonpath
	JsonPath string `json:"JsonPath"`

	// MatchEquals AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-webhook-webhookfilterrule.html#cfn-codepipeline-webhook-webhookfilterrule-matchequals
	MatchEquals *string `json:"MatchEquals,omitempty"`

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
func (r *Webhook_WebhookFilterRule) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Webhook.WebhookFilterRule"
}
