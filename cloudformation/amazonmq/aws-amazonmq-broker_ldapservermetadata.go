// Code generated by "go generate". Please don't change this file directly.

package amazonmq

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
	"github.com/awslabs/goformation/v6/cloudformation/utils"
)

var _ utils.Value[struct{}]

// Broker_LdapServerMetadata AWS CloudFormation Resource (AWS::AmazonMQ::Broker.LdapServerMetadata)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html
type Broker_LdapServerMetadata struct {

	// Hosts AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-hosts
	Hosts []string `json:"Hosts"`

	// RoleBase AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-rolebase
	RoleBase string `json:"RoleBase"`

	// RoleName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-rolename
	RoleName *string `json:"RoleName,omitempty"`

	// RoleSearchMatching AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-rolesearchmatching
	RoleSearchMatching string `json:"RoleSearchMatching"`

	// RoleSearchSubtree AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-rolesearchsubtree
	RoleSearchSubtree *utils.Value[bool] `json:"RoleSearchSubtree,omitempty"`

	// ServiceAccountPassword AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-serviceaccountpassword
	ServiceAccountPassword string `json:"ServiceAccountPassword"`

	// ServiceAccountUsername AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-serviceaccountusername
	ServiceAccountUsername string `json:"ServiceAccountUsername"`

	// UserBase AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-userbase
	UserBase string `json:"UserBase"`

	// UserRoleName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-userrolename
	UserRoleName *string `json:"UserRoleName,omitempty"`

	// UserSearchMatching AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-usersearchmatching
	UserSearchMatching string `json:"UserSearchMatching"`

	// UserSearchSubtree AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-usersearchsubtree
	UserSearchSubtree *utils.Value[bool] `json:"UserSearchSubtree,omitempty"`

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
func (r *Broker_LdapServerMetadata) AWSCloudFormationType() string {
	return "AWS::AmazonMQ::Broker.LdapServerMetadata"
}
