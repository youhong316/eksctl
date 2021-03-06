package cloudformation

import (
	"encoding/json"
)

// AWSS3Bucket_LambdaConfiguration AWS CloudFormation Resource (AWS::S3::Bucket.LambdaConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html
type AWSS3Bucket_LambdaConfiguration struct {

	// Event AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-event
	Event *Value `json:"Event,omitempty"`

	// Filter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-filter
	Filter *AWSS3Bucket_NotificationFilter `json:"Filter,omitempty"`

	// Function AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-function
	Function *Value `json:"Function,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_LambdaConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.LambdaConfiguration"
}

func (r *AWSS3Bucket_LambdaConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
