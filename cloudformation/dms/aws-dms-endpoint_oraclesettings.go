package dms

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
	"github.com/awslabs/goformation/v5/cloudformation/utils"
)

var _ utils.Value[struct{}]

// Endpoint_OracleSettings AWS CloudFormation Resource (AWS::DMS::Endpoint.OracleSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html
type Endpoint_OracleSettings struct {

	// AccessAlternateDirectly AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-accessalternatedirectly
	AccessAlternateDirectly *utils.Value[bool] `json:"AccessAlternateDirectly,omitempty"`

	// AddSupplementalLogging AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-addsupplementallogging
	AddSupplementalLogging *utils.Value[bool] `json:"AddSupplementalLogging,omitempty"`

	// AdditionalArchivedLogDestId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-additionalarchivedlogdestid
	AdditionalArchivedLogDestId *utils.Value[int] `json:"AdditionalArchivedLogDestId,omitempty"`

	// AllowSelectNestedTables AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-allowselectnestedtables
	AllowSelectNestedTables *utils.Value[bool] `json:"AllowSelectNestedTables,omitempty"`

	// ArchivedLogDestId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-archivedlogdestid
	ArchivedLogDestId *utils.Value[int] `json:"ArchivedLogDestId,omitempty"`

	// ArchivedLogsOnly AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-archivedlogsonly
	ArchivedLogsOnly *utils.Value[bool] `json:"ArchivedLogsOnly,omitempty"`

	// AsmPassword AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-asmpassword
	AsmPassword string `json:"AsmPassword,omitempty"`

	// AsmServer AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-asmserver
	AsmServer string `json:"AsmServer,omitempty"`

	// AsmUser AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-asmuser
	AsmUser string `json:"AsmUser,omitempty"`

	// CharLengthSemantics AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-charlengthsemantics
	CharLengthSemantics string `json:"CharLengthSemantics,omitempty"`

	// DirectPathNoLog AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-directpathnolog
	DirectPathNoLog *utils.Value[bool] `json:"DirectPathNoLog,omitempty"`

	// DirectPathParallelLoad AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-directpathparallelload
	DirectPathParallelLoad *utils.Value[bool] `json:"DirectPathParallelLoad,omitempty"`

	// EnableHomogenousTablespace AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-enablehomogenoustablespace
	EnableHomogenousTablespace *utils.Value[bool] `json:"EnableHomogenousTablespace,omitempty"`

	// ExtraArchivedLogDestIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-extraarchivedlogdestids
	ExtraArchivedLogDestIds []int `json:"ExtraArchivedLogDestIds,omitempty"`

	// FailTasksOnLobTruncation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-failtasksonlobtruncation
	FailTasksOnLobTruncation *utils.Value[bool] `json:"FailTasksOnLobTruncation,omitempty"`

	// NumberDatatypeScale AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-numberdatatypescale
	NumberDatatypeScale *utils.Value[int] `json:"NumberDatatypeScale,omitempty"`

	// OraclePathPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-oraclepathprefix
	OraclePathPrefix string `json:"OraclePathPrefix,omitempty"`

	// ParallelAsmReadThreads AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-parallelasmreadthreads
	ParallelAsmReadThreads *utils.Value[int] `json:"ParallelAsmReadThreads,omitempty"`

	// ReadAheadBlocks AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-readaheadblocks
	ReadAheadBlocks *utils.Value[int] `json:"ReadAheadBlocks,omitempty"`

	// ReadTableSpaceName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-readtablespacename
	ReadTableSpaceName *utils.Value[bool] `json:"ReadTableSpaceName,omitempty"`

	// ReplacePathPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-replacepathprefix
	ReplacePathPrefix *utils.Value[bool] `json:"ReplacePathPrefix,omitempty"`

	// RetryInterval AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-retryinterval
	RetryInterval *utils.Value[int] `json:"RetryInterval,omitempty"`

	// SecretsManagerAccessRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanageraccessrolearn
	SecretsManagerAccessRoleArn string `json:"SecretsManagerAccessRoleArn,omitempty"`

	// SecretsManagerOracleAsmAccessRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanageroracleasmaccessrolearn
	SecretsManagerOracleAsmAccessRoleArn string `json:"SecretsManagerOracleAsmAccessRoleArn,omitempty"`

	// SecretsManagerOracleAsmSecretId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanageroracleasmsecretid
	SecretsManagerOracleAsmSecretId string `json:"SecretsManagerOracleAsmSecretId,omitempty"`

	// SecretsManagerSecretId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanagersecretid
	SecretsManagerSecretId string `json:"SecretsManagerSecretId,omitempty"`

	// SecurityDbEncryption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-securitydbencryption
	SecurityDbEncryption string `json:"SecurityDbEncryption,omitempty"`

	// SecurityDbEncryptionName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-securitydbencryptionname
	SecurityDbEncryptionName string `json:"SecurityDbEncryptionName,omitempty"`

	// SpatialDataOptionToGeoJsonFunctionName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-spatialdataoptiontogeojsonfunctionname
	SpatialDataOptionToGeoJsonFunctionName string `json:"SpatialDataOptionToGeoJsonFunctionName,omitempty"`

	// StandbyDelayTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-standbydelaytime
	StandbyDelayTime *utils.Value[int] `json:"StandbyDelayTime,omitempty"`

	// UseAlternateFolderForOnline AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usealternatefolderforonline
	UseAlternateFolderForOnline *utils.Value[bool] `json:"UseAlternateFolderForOnline,omitempty"`

	// UseBFile AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usebfile
	UseBFile *utils.Value[bool] `json:"UseBFile,omitempty"`

	// UseDirectPathFullLoad AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usedirectpathfullload
	UseDirectPathFullLoad *utils.Value[bool] `json:"UseDirectPathFullLoad,omitempty"`

	// UseLogminerReader AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-uselogminerreader
	UseLogminerReader *utils.Value[bool] `json:"UseLogminerReader,omitempty"`

	// UsePathPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usepathprefix
	UsePathPrefix string `json:"UsePathPrefix,omitempty"`

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
func (r *Endpoint_OracleSettings) AWSCloudFormationType() string {
	return "AWS::DMS::Endpoint.OracleSettings"
}
