// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package s3iface provides an interface to enable mocking the Amazon Simple Storage Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package s3iface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3API provides an interface to enable mocking the
// s3.S3 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Simple Storage Service.
//    func myFunc(svc s3iface.S3API) bool {
//        // Make svc.AbortMultipartUpload request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := s3.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockS3Client struct {
//        s3iface.S3API
//    }
//    func (m *mockS3Client) AbortMultipartUpload(input *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockS3Client{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type S3API interface {
	AbortMultipartUpload(*s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error)
	AbortMultipartUploadWithContext(aws.Context, *s3.AbortMultipartUploadInput, ...aws.Option) (*s3.AbortMultipartUploadOutput, error)
	AbortMultipartUploadRequest(*s3.AbortMultipartUploadInput) (*aws.Request, *s3.AbortMultipartUploadOutput)

	CompleteMultipartUpload(*s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error)
	CompleteMultipartUploadWithContext(aws.Context, *s3.CompleteMultipartUploadInput, ...aws.Option) (*s3.CompleteMultipartUploadOutput, error)
	CompleteMultipartUploadRequest(*s3.CompleteMultipartUploadInput) (*aws.Request, *s3.CompleteMultipartUploadOutput)

	CopyObject(*s3.CopyObjectInput) (*s3.CopyObjectOutput, error)
	CopyObjectWithContext(aws.Context, *s3.CopyObjectInput, ...aws.Option) (*s3.CopyObjectOutput, error)
	CopyObjectRequest(*s3.CopyObjectInput) (*aws.Request, *s3.CopyObjectOutput)

	CreateBucket(*s3.CreateBucketInput) (*s3.CreateBucketOutput, error)
	CreateBucketWithContext(aws.Context, *s3.CreateBucketInput, ...aws.Option) (*s3.CreateBucketOutput, error)
	CreateBucketRequest(*s3.CreateBucketInput) (*aws.Request, *s3.CreateBucketOutput)

	CreateMultipartUpload(*s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error)
	CreateMultipartUploadWithContext(aws.Context, *s3.CreateMultipartUploadInput, ...aws.Option) (*s3.CreateMultipartUploadOutput, error)
	CreateMultipartUploadRequest(*s3.CreateMultipartUploadInput) (*aws.Request, *s3.CreateMultipartUploadOutput)

	DeleteBucket(*s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error)
	DeleteBucketWithContext(aws.Context, *s3.DeleteBucketInput, ...aws.Option) (*s3.DeleteBucketOutput, error)
	DeleteBucketRequest(*s3.DeleteBucketInput) (*aws.Request, *s3.DeleteBucketOutput)

	DeleteBucketAnalyticsConfiguration(*s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error)
	DeleteBucketAnalyticsConfigurationWithContext(aws.Context, *s3.DeleteBucketAnalyticsConfigurationInput, ...aws.Option) (*s3.DeleteBucketAnalyticsConfigurationOutput, error)
	DeleteBucketAnalyticsConfigurationRequest(*s3.DeleteBucketAnalyticsConfigurationInput) (*aws.Request, *s3.DeleteBucketAnalyticsConfigurationOutput)

	DeleteBucketCors(*s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error)
	DeleteBucketCorsWithContext(aws.Context, *s3.DeleteBucketCorsInput, ...aws.Option) (*s3.DeleteBucketCorsOutput, error)
	DeleteBucketCorsRequest(*s3.DeleteBucketCorsInput) (*aws.Request, *s3.DeleteBucketCorsOutput)

	DeleteBucketInventoryConfiguration(*s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error)
	DeleteBucketInventoryConfigurationWithContext(aws.Context, *s3.DeleteBucketInventoryConfigurationInput, ...aws.Option) (*s3.DeleteBucketInventoryConfigurationOutput, error)
	DeleteBucketInventoryConfigurationRequest(*s3.DeleteBucketInventoryConfigurationInput) (*aws.Request, *s3.DeleteBucketInventoryConfigurationOutput)

	DeleteBucketLifecycle(*s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error)
	DeleteBucketLifecycleWithContext(aws.Context, *s3.DeleteBucketLifecycleInput, ...aws.Option) (*s3.DeleteBucketLifecycleOutput, error)
	DeleteBucketLifecycleRequest(*s3.DeleteBucketLifecycleInput) (*aws.Request, *s3.DeleteBucketLifecycleOutput)

	DeleteBucketMetricsConfiguration(*s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error)
	DeleteBucketMetricsConfigurationWithContext(aws.Context, *s3.DeleteBucketMetricsConfigurationInput, ...aws.Option) (*s3.DeleteBucketMetricsConfigurationOutput, error)
	DeleteBucketMetricsConfigurationRequest(*s3.DeleteBucketMetricsConfigurationInput) (*aws.Request, *s3.DeleteBucketMetricsConfigurationOutput)

	DeleteBucketPolicy(*s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error)
	DeleteBucketPolicyWithContext(aws.Context, *s3.DeleteBucketPolicyInput, ...aws.Option) (*s3.DeleteBucketPolicyOutput, error)
	DeleteBucketPolicyRequest(*s3.DeleteBucketPolicyInput) (*aws.Request, *s3.DeleteBucketPolicyOutput)

	DeleteBucketReplication(*s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error)
	DeleteBucketReplicationWithContext(aws.Context, *s3.DeleteBucketReplicationInput, ...aws.Option) (*s3.DeleteBucketReplicationOutput, error)
	DeleteBucketReplicationRequest(*s3.DeleteBucketReplicationInput) (*aws.Request, *s3.DeleteBucketReplicationOutput)

	DeleteBucketTagging(*s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error)
	DeleteBucketTaggingWithContext(aws.Context, *s3.DeleteBucketTaggingInput, ...aws.Option) (*s3.DeleteBucketTaggingOutput, error)
	DeleteBucketTaggingRequest(*s3.DeleteBucketTaggingInput) (*aws.Request, *s3.DeleteBucketTaggingOutput)

	DeleteBucketWebsite(*s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error)
	DeleteBucketWebsiteWithContext(aws.Context, *s3.DeleteBucketWebsiteInput, ...aws.Option) (*s3.DeleteBucketWebsiteOutput, error)
	DeleteBucketWebsiteRequest(*s3.DeleteBucketWebsiteInput) (*aws.Request, *s3.DeleteBucketWebsiteOutput)

	DeleteObject(*s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
	DeleteObjectWithContext(aws.Context, *s3.DeleteObjectInput, ...aws.Option) (*s3.DeleteObjectOutput, error)
	DeleteObjectRequest(*s3.DeleteObjectInput) (*aws.Request, *s3.DeleteObjectOutput)

	DeleteObjectTagging(*s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error)
	DeleteObjectTaggingWithContext(aws.Context, *s3.DeleteObjectTaggingInput, ...aws.Option) (*s3.DeleteObjectTaggingOutput, error)
	DeleteObjectTaggingRequest(*s3.DeleteObjectTaggingInput) (*aws.Request, *s3.DeleteObjectTaggingOutput)

	DeleteObjects(*s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error)
	DeleteObjectsWithContext(aws.Context, *s3.DeleteObjectsInput, ...aws.Option) (*s3.DeleteObjectsOutput, error)
	DeleteObjectsRequest(*s3.DeleteObjectsInput) (*aws.Request, *s3.DeleteObjectsOutput)

	GetBucketAccelerateConfiguration(*s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error)
	GetBucketAccelerateConfigurationWithContext(aws.Context, *s3.GetBucketAccelerateConfigurationInput, ...aws.Option) (*s3.GetBucketAccelerateConfigurationOutput, error)
	GetBucketAccelerateConfigurationRequest(*s3.GetBucketAccelerateConfigurationInput) (*aws.Request, *s3.GetBucketAccelerateConfigurationOutput)

	GetBucketAcl(*s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error)
	GetBucketAclWithContext(aws.Context, *s3.GetBucketAclInput, ...aws.Option) (*s3.GetBucketAclOutput, error)
	GetBucketAclRequest(*s3.GetBucketAclInput) (*aws.Request, *s3.GetBucketAclOutput)

	GetBucketAnalyticsConfiguration(*s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error)
	GetBucketAnalyticsConfigurationWithContext(aws.Context, *s3.GetBucketAnalyticsConfigurationInput, ...aws.Option) (*s3.GetBucketAnalyticsConfigurationOutput, error)
	GetBucketAnalyticsConfigurationRequest(*s3.GetBucketAnalyticsConfigurationInput) (*aws.Request, *s3.GetBucketAnalyticsConfigurationOutput)

	GetBucketCors(*s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error)
	GetBucketCorsWithContext(aws.Context, *s3.GetBucketCorsInput, ...aws.Option) (*s3.GetBucketCorsOutput, error)
	GetBucketCorsRequest(*s3.GetBucketCorsInput) (*aws.Request, *s3.GetBucketCorsOutput)

	GetBucketInventoryConfiguration(*s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error)
	GetBucketInventoryConfigurationWithContext(aws.Context, *s3.GetBucketInventoryConfigurationInput, ...aws.Option) (*s3.GetBucketInventoryConfigurationOutput, error)
	GetBucketInventoryConfigurationRequest(*s3.GetBucketInventoryConfigurationInput) (*aws.Request, *s3.GetBucketInventoryConfigurationOutput)

	GetBucketLifecycle(*s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error)
	GetBucketLifecycleWithContext(aws.Context, *s3.GetBucketLifecycleInput, ...aws.Option) (*s3.GetBucketLifecycleOutput, error)
	GetBucketLifecycleRequest(*s3.GetBucketLifecycleInput) (*aws.Request, *s3.GetBucketLifecycleOutput)

	GetBucketLifecycleConfiguration(*s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationWithContext(aws.Context, *s3.GetBucketLifecycleConfigurationInput, ...aws.Option) (*s3.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationRequest(*s3.GetBucketLifecycleConfigurationInput) (*aws.Request, *s3.GetBucketLifecycleConfigurationOutput)

	GetBucketLocation(*s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error)
	GetBucketLocationWithContext(aws.Context, *s3.GetBucketLocationInput, ...aws.Option) (*s3.GetBucketLocationOutput, error)
	GetBucketLocationRequest(*s3.GetBucketLocationInput) (*aws.Request, *s3.GetBucketLocationOutput)

	GetBucketLogging(*s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error)
	GetBucketLoggingWithContext(aws.Context, *s3.GetBucketLoggingInput, ...aws.Option) (*s3.GetBucketLoggingOutput, error)
	GetBucketLoggingRequest(*s3.GetBucketLoggingInput) (*aws.Request, *s3.GetBucketLoggingOutput)

	GetBucketMetricsConfiguration(*s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error)
	GetBucketMetricsConfigurationWithContext(aws.Context, *s3.GetBucketMetricsConfigurationInput, ...aws.Option) (*s3.GetBucketMetricsConfigurationOutput, error)
	GetBucketMetricsConfigurationRequest(*s3.GetBucketMetricsConfigurationInput) (*aws.Request, *s3.GetBucketMetricsConfigurationOutput)

	GetBucketNotification(*s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error)
	GetBucketNotificationWithContext(aws.Context, *s3.GetBucketNotificationConfigurationRequest, ...aws.Option) (*s3.NotificationConfigurationDeprecated, error)
	GetBucketNotificationRequest(*s3.GetBucketNotificationConfigurationRequest) (*aws.Request, *s3.NotificationConfigurationDeprecated)

	GetBucketNotificationConfiguration(*s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error)
	GetBucketNotificationConfigurationWithContext(aws.Context, *s3.GetBucketNotificationConfigurationRequest, ...aws.Option) (*s3.NotificationConfiguration, error)
	GetBucketNotificationConfigurationRequest(*s3.GetBucketNotificationConfigurationRequest) (*aws.Request, *s3.NotificationConfiguration)

	GetBucketPolicy(*s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error)
	GetBucketPolicyWithContext(aws.Context, *s3.GetBucketPolicyInput, ...aws.Option) (*s3.GetBucketPolicyOutput, error)
	GetBucketPolicyRequest(*s3.GetBucketPolicyInput) (*aws.Request, *s3.GetBucketPolicyOutput)

	GetBucketReplication(*s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error)
	GetBucketReplicationWithContext(aws.Context, *s3.GetBucketReplicationInput, ...aws.Option) (*s3.GetBucketReplicationOutput, error)
	GetBucketReplicationRequest(*s3.GetBucketReplicationInput) (*aws.Request, *s3.GetBucketReplicationOutput)

	GetBucketRequestPayment(*s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error)
	GetBucketRequestPaymentWithContext(aws.Context, *s3.GetBucketRequestPaymentInput, ...aws.Option) (*s3.GetBucketRequestPaymentOutput, error)
	GetBucketRequestPaymentRequest(*s3.GetBucketRequestPaymentInput) (*aws.Request, *s3.GetBucketRequestPaymentOutput)

	GetBucketTagging(*s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error)
	GetBucketTaggingWithContext(aws.Context, *s3.GetBucketTaggingInput, ...aws.Option) (*s3.GetBucketTaggingOutput, error)
	GetBucketTaggingRequest(*s3.GetBucketTaggingInput) (*aws.Request, *s3.GetBucketTaggingOutput)

	GetBucketVersioning(*s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error)
	GetBucketVersioningWithContext(aws.Context, *s3.GetBucketVersioningInput, ...aws.Option) (*s3.GetBucketVersioningOutput, error)
	GetBucketVersioningRequest(*s3.GetBucketVersioningInput) (*aws.Request, *s3.GetBucketVersioningOutput)

	GetBucketWebsite(*s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error)
	GetBucketWebsiteWithContext(aws.Context, *s3.GetBucketWebsiteInput, ...aws.Option) (*s3.GetBucketWebsiteOutput, error)
	GetBucketWebsiteRequest(*s3.GetBucketWebsiteInput) (*aws.Request, *s3.GetBucketWebsiteOutput)

	GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error)
	GetObjectWithContext(aws.Context, *s3.GetObjectInput, ...aws.Option) (*s3.GetObjectOutput, error)
	GetObjectRequest(*s3.GetObjectInput) (*aws.Request, *s3.GetObjectOutput)

	GetObjectAcl(*s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error)
	GetObjectAclWithContext(aws.Context, *s3.GetObjectAclInput, ...aws.Option) (*s3.GetObjectAclOutput, error)
	GetObjectAclRequest(*s3.GetObjectAclInput) (*aws.Request, *s3.GetObjectAclOutput)

	GetObjectTagging(*s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error)
	GetObjectTaggingWithContext(aws.Context, *s3.GetObjectTaggingInput, ...aws.Option) (*s3.GetObjectTaggingOutput, error)
	GetObjectTaggingRequest(*s3.GetObjectTaggingInput) (*aws.Request, *s3.GetObjectTaggingOutput)

	GetObjectTorrent(*s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error)
	GetObjectTorrentWithContext(aws.Context, *s3.GetObjectTorrentInput, ...aws.Option) (*s3.GetObjectTorrentOutput, error)
	GetObjectTorrentRequest(*s3.GetObjectTorrentInput) (*aws.Request, *s3.GetObjectTorrentOutput)

	HeadBucket(*s3.HeadBucketInput) (*s3.HeadBucketOutput, error)
	HeadBucketWithContext(aws.Context, *s3.HeadBucketInput, ...aws.Option) (*s3.HeadBucketOutput, error)
	HeadBucketRequest(*s3.HeadBucketInput) (*aws.Request, *s3.HeadBucketOutput)

	HeadObject(*s3.HeadObjectInput) (*s3.HeadObjectOutput, error)
	HeadObjectWithContext(aws.Context, *s3.HeadObjectInput, ...aws.Option) (*s3.HeadObjectOutput, error)
	HeadObjectRequest(*s3.HeadObjectInput) (*aws.Request, *s3.HeadObjectOutput)

	ListBucketAnalyticsConfigurations(*s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error)
	ListBucketAnalyticsConfigurationsWithContext(aws.Context, *s3.ListBucketAnalyticsConfigurationsInput, ...aws.Option) (*s3.ListBucketAnalyticsConfigurationsOutput, error)
	ListBucketAnalyticsConfigurationsRequest(*s3.ListBucketAnalyticsConfigurationsInput) (*aws.Request, *s3.ListBucketAnalyticsConfigurationsOutput)

	ListBucketInventoryConfigurations(*s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error)
	ListBucketInventoryConfigurationsWithContext(aws.Context, *s3.ListBucketInventoryConfigurationsInput, ...aws.Option) (*s3.ListBucketInventoryConfigurationsOutput, error)
	ListBucketInventoryConfigurationsRequest(*s3.ListBucketInventoryConfigurationsInput) (*aws.Request, *s3.ListBucketInventoryConfigurationsOutput)

	ListBucketMetricsConfigurations(*s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error)
	ListBucketMetricsConfigurationsWithContext(aws.Context, *s3.ListBucketMetricsConfigurationsInput, ...aws.Option) (*s3.ListBucketMetricsConfigurationsOutput, error)
	ListBucketMetricsConfigurationsRequest(*s3.ListBucketMetricsConfigurationsInput) (*aws.Request, *s3.ListBucketMetricsConfigurationsOutput)

	ListBuckets(*s3.ListBucketsInput) (*s3.ListBucketsOutput, error)
	ListBucketsWithContext(aws.Context, *s3.ListBucketsInput, ...aws.Option) (*s3.ListBucketsOutput, error)
	ListBucketsRequest(*s3.ListBucketsInput) (*aws.Request, *s3.ListBucketsOutput)

	ListMultipartUploads(*s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error)
	ListMultipartUploadsWithContext(aws.Context, *s3.ListMultipartUploadsInput, ...aws.Option) (*s3.ListMultipartUploadsOutput, error)
	ListMultipartUploadsRequest(*s3.ListMultipartUploadsInput) (*aws.Request, *s3.ListMultipartUploadsOutput)

	ListMultipartUploadsPages(*s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool) error
	ListMultipartUploadsPagesWithContext(aws.Context, *s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool, ...aws.Option) error

	ListObjectVersions(*s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error)
	ListObjectVersionsWithContext(aws.Context, *s3.ListObjectVersionsInput, ...aws.Option) (*s3.ListObjectVersionsOutput, error)
	ListObjectVersionsRequest(*s3.ListObjectVersionsInput) (*aws.Request, *s3.ListObjectVersionsOutput)

	ListObjectVersionsPages(*s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool) error
	ListObjectVersionsPagesWithContext(aws.Context, *s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool, ...aws.Option) error

	ListObjects(*s3.ListObjectsInput) (*s3.ListObjectsOutput, error)
	ListObjectsWithContext(aws.Context, *s3.ListObjectsInput, ...aws.Option) (*s3.ListObjectsOutput, error)
	ListObjectsRequest(*s3.ListObjectsInput) (*aws.Request, *s3.ListObjectsOutput)

	ListObjectsPages(*s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool) error
	ListObjectsPagesWithContext(aws.Context, *s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool, ...aws.Option) error

	ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
	ListObjectsV2WithContext(aws.Context, *s3.ListObjectsV2Input, ...aws.Option) (*s3.ListObjectsV2Output, error)
	ListObjectsV2Request(*s3.ListObjectsV2Input) (*aws.Request, *s3.ListObjectsV2Output)

	ListObjectsV2Pages(*s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool) error
	ListObjectsV2PagesWithContext(aws.Context, *s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool, ...aws.Option) error

	ListParts(*s3.ListPartsInput) (*s3.ListPartsOutput, error)
	ListPartsWithContext(aws.Context, *s3.ListPartsInput, ...aws.Option) (*s3.ListPartsOutput, error)
	ListPartsRequest(*s3.ListPartsInput) (*aws.Request, *s3.ListPartsOutput)

	ListPartsPages(*s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool) error
	ListPartsPagesWithContext(aws.Context, *s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool, ...aws.Option) error

	PutBucketAccelerateConfiguration(*s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error)
	PutBucketAccelerateConfigurationWithContext(aws.Context, *s3.PutBucketAccelerateConfigurationInput, ...aws.Option) (*s3.PutBucketAccelerateConfigurationOutput, error)
	PutBucketAccelerateConfigurationRequest(*s3.PutBucketAccelerateConfigurationInput) (*aws.Request, *s3.PutBucketAccelerateConfigurationOutput)

	PutBucketAcl(*s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error)
	PutBucketAclWithContext(aws.Context, *s3.PutBucketAclInput, ...aws.Option) (*s3.PutBucketAclOutput, error)
	PutBucketAclRequest(*s3.PutBucketAclInput) (*aws.Request, *s3.PutBucketAclOutput)

	PutBucketAnalyticsConfiguration(*s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error)
	PutBucketAnalyticsConfigurationWithContext(aws.Context, *s3.PutBucketAnalyticsConfigurationInput, ...aws.Option) (*s3.PutBucketAnalyticsConfigurationOutput, error)
	PutBucketAnalyticsConfigurationRequest(*s3.PutBucketAnalyticsConfigurationInput) (*aws.Request, *s3.PutBucketAnalyticsConfigurationOutput)

	PutBucketCors(*s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error)
	PutBucketCorsWithContext(aws.Context, *s3.PutBucketCorsInput, ...aws.Option) (*s3.PutBucketCorsOutput, error)
	PutBucketCorsRequest(*s3.PutBucketCorsInput) (*aws.Request, *s3.PutBucketCorsOutput)

	PutBucketInventoryConfiguration(*s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error)
	PutBucketInventoryConfigurationWithContext(aws.Context, *s3.PutBucketInventoryConfigurationInput, ...aws.Option) (*s3.PutBucketInventoryConfigurationOutput, error)
	PutBucketInventoryConfigurationRequest(*s3.PutBucketInventoryConfigurationInput) (*aws.Request, *s3.PutBucketInventoryConfigurationOutput)

	PutBucketLifecycle(*s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error)
	PutBucketLifecycleWithContext(aws.Context, *s3.PutBucketLifecycleInput, ...aws.Option) (*s3.PutBucketLifecycleOutput, error)
	PutBucketLifecycleRequest(*s3.PutBucketLifecycleInput) (*aws.Request, *s3.PutBucketLifecycleOutput)

	PutBucketLifecycleConfiguration(*s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationWithContext(aws.Context, *s3.PutBucketLifecycleConfigurationInput, ...aws.Option) (*s3.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationRequest(*s3.PutBucketLifecycleConfigurationInput) (*aws.Request, *s3.PutBucketLifecycleConfigurationOutput)

	PutBucketLogging(*s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error)
	PutBucketLoggingWithContext(aws.Context, *s3.PutBucketLoggingInput, ...aws.Option) (*s3.PutBucketLoggingOutput, error)
	PutBucketLoggingRequest(*s3.PutBucketLoggingInput) (*aws.Request, *s3.PutBucketLoggingOutput)

	PutBucketMetricsConfiguration(*s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error)
	PutBucketMetricsConfigurationWithContext(aws.Context, *s3.PutBucketMetricsConfigurationInput, ...aws.Option) (*s3.PutBucketMetricsConfigurationOutput, error)
	PutBucketMetricsConfigurationRequest(*s3.PutBucketMetricsConfigurationInput) (*aws.Request, *s3.PutBucketMetricsConfigurationOutput)

	PutBucketNotification(*s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error)
	PutBucketNotificationWithContext(aws.Context, *s3.PutBucketNotificationInput, ...aws.Option) (*s3.PutBucketNotificationOutput, error)
	PutBucketNotificationRequest(*s3.PutBucketNotificationInput) (*aws.Request, *s3.PutBucketNotificationOutput)

	PutBucketNotificationConfiguration(*s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error)
	PutBucketNotificationConfigurationWithContext(aws.Context, *s3.PutBucketNotificationConfigurationInput, ...aws.Option) (*s3.PutBucketNotificationConfigurationOutput, error)
	PutBucketNotificationConfigurationRequest(*s3.PutBucketNotificationConfigurationInput) (*aws.Request, *s3.PutBucketNotificationConfigurationOutput)

	PutBucketPolicy(*s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error)
	PutBucketPolicyWithContext(aws.Context, *s3.PutBucketPolicyInput, ...aws.Option) (*s3.PutBucketPolicyOutput, error)
	PutBucketPolicyRequest(*s3.PutBucketPolicyInput) (*aws.Request, *s3.PutBucketPolicyOutput)

	PutBucketReplication(*s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error)
	PutBucketReplicationWithContext(aws.Context, *s3.PutBucketReplicationInput, ...aws.Option) (*s3.PutBucketReplicationOutput, error)
	PutBucketReplicationRequest(*s3.PutBucketReplicationInput) (*aws.Request, *s3.PutBucketReplicationOutput)

	PutBucketRequestPayment(*s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error)
	PutBucketRequestPaymentWithContext(aws.Context, *s3.PutBucketRequestPaymentInput, ...aws.Option) (*s3.PutBucketRequestPaymentOutput, error)
	PutBucketRequestPaymentRequest(*s3.PutBucketRequestPaymentInput) (*aws.Request, *s3.PutBucketRequestPaymentOutput)

	PutBucketTagging(*s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error)
	PutBucketTaggingWithContext(aws.Context, *s3.PutBucketTaggingInput, ...aws.Option) (*s3.PutBucketTaggingOutput, error)
	PutBucketTaggingRequest(*s3.PutBucketTaggingInput) (*aws.Request, *s3.PutBucketTaggingOutput)

	PutBucketVersioning(*s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error)
	PutBucketVersioningWithContext(aws.Context, *s3.PutBucketVersioningInput, ...aws.Option) (*s3.PutBucketVersioningOutput, error)
	PutBucketVersioningRequest(*s3.PutBucketVersioningInput) (*aws.Request, *s3.PutBucketVersioningOutput)

	PutBucketWebsite(*s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error)
	PutBucketWebsiteWithContext(aws.Context, *s3.PutBucketWebsiteInput, ...aws.Option) (*s3.PutBucketWebsiteOutput, error)
	PutBucketWebsiteRequest(*s3.PutBucketWebsiteInput) (*aws.Request, *s3.PutBucketWebsiteOutput)

	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
	PutObjectWithContext(aws.Context, *s3.PutObjectInput, ...aws.Option) (*s3.PutObjectOutput, error)
	PutObjectRequest(*s3.PutObjectInput) (*aws.Request, *s3.PutObjectOutput)

	PutObjectAcl(*s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error)
	PutObjectAclWithContext(aws.Context, *s3.PutObjectAclInput, ...aws.Option) (*s3.PutObjectAclOutput, error)
	PutObjectAclRequest(*s3.PutObjectAclInput) (*aws.Request, *s3.PutObjectAclOutput)

	PutObjectTagging(*s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error)
	PutObjectTaggingWithContext(aws.Context, *s3.PutObjectTaggingInput, ...aws.Option) (*s3.PutObjectTaggingOutput, error)
	PutObjectTaggingRequest(*s3.PutObjectTaggingInput) (*aws.Request, *s3.PutObjectTaggingOutput)

	RestoreObject(*s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error)
	RestoreObjectWithContext(aws.Context, *s3.RestoreObjectInput, ...aws.Option) (*s3.RestoreObjectOutput, error)
	RestoreObjectRequest(*s3.RestoreObjectInput) (*aws.Request, *s3.RestoreObjectOutput)

	UploadPart(*s3.UploadPartInput) (*s3.UploadPartOutput, error)
	UploadPartWithContext(aws.Context, *s3.UploadPartInput, ...aws.Option) (*s3.UploadPartOutput, error)
	UploadPartRequest(*s3.UploadPartInput) (*aws.Request, *s3.UploadPartOutput)

	UploadPartCopy(*s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error)
	UploadPartCopyWithContext(aws.Context, *s3.UploadPartCopyInput, ...aws.Option) (*s3.UploadPartCopyOutput, error)
	UploadPartCopyRequest(*s3.UploadPartCopyInput) (*aws.Request, *s3.UploadPartCopyOutput)

	WaitUntilBucketExists(*s3.HeadBucketInput) error
	WaitUntilBucketExistsWithContext(aws.Context, *s3.HeadBucketInput, ...aws.WaiterOption) error

	WaitUntilBucketNotExists(*s3.HeadBucketInput) error
	WaitUntilBucketNotExistsWithContext(aws.Context, *s3.HeadBucketInput, ...aws.WaiterOption) error

	WaitUntilObjectExists(*s3.HeadObjectInput) error
	WaitUntilObjectExistsWithContext(aws.Context, *s3.HeadObjectInput, ...aws.WaiterOption) error

	WaitUntilObjectNotExists(*s3.HeadObjectInput) error
	WaitUntilObjectNotExistsWithContext(aws.Context, *s3.HeadObjectInput, ...aws.WaiterOption) error
}

var _ S3API = (*s3.S3)(nil)