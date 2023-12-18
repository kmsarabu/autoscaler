// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package frauddetectoriface provides an interface to enable mocking the Amazon Fraud Detector service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package frauddetectoriface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/frauddetector"
)

// FraudDetectorAPI provides an interface to enable mocking the
// frauddetector.FraudDetector service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Fraud Detector.
//	func myFunc(svc frauddetectoriface.FraudDetectorAPI) bool {
//	    // Make svc.BatchCreateVariable request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := frauddetector.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockFraudDetectorClient struct {
//	    frauddetectoriface.FraudDetectorAPI
//	}
//	func (m *mockFraudDetectorClient) BatchCreateVariable(input *frauddetector.BatchCreateVariableInput) (*frauddetector.BatchCreateVariableOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockFraudDetectorClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type FraudDetectorAPI interface {
	BatchCreateVariable(*frauddetector.BatchCreateVariableInput) (*frauddetector.BatchCreateVariableOutput, error)
	BatchCreateVariableWithContext(aws.Context, *frauddetector.BatchCreateVariableInput, ...request.Option) (*frauddetector.BatchCreateVariableOutput, error)
	BatchCreateVariableRequest(*frauddetector.BatchCreateVariableInput) (*request.Request, *frauddetector.BatchCreateVariableOutput)

	BatchGetVariable(*frauddetector.BatchGetVariableInput) (*frauddetector.BatchGetVariableOutput, error)
	BatchGetVariableWithContext(aws.Context, *frauddetector.BatchGetVariableInput, ...request.Option) (*frauddetector.BatchGetVariableOutput, error)
	BatchGetVariableRequest(*frauddetector.BatchGetVariableInput) (*request.Request, *frauddetector.BatchGetVariableOutput)

	CancelBatchImportJob(*frauddetector.CancelBatchImportJobInput) (*frauddetector.CancelBatchImportJobOutput, error)
	CancelBatchImportJobWithContext(aws.Context, *frauddetector.CancelBatchImportJobInput, ...request.Option) (*frauddetector.CancelBatchImportJobOutput, error)
	CancelBatchImportJobRequest(*frauddetector.CancelBatchImportJobInput) (*request.Request, *frauddetector.CancelBatchImportJobOutput)

	CancelBatchPredictionJob(*frauddetector.CancelBatchPredictionJobInput) (*frauddetector.CancelBatchPredictionJobOutput, error)
	CancelBatchPredictionJobWithContext(aws.Context, *frauddetector.CancelBatchPredictionJobInput, ...request.Option) (*frauddetector.CancelBatchPredictionJobOutput, error)
	CancelBatchPredictionJobRequest(*frauddetector.CancelBatchPredictionJobInput) (*request.Request, *frauddetector.CancelBatchPredictionJobOutput)

	CreateBatchImportJob(*frauddetector.CreateBatchImportJobInput) (*frauddetector.CreateBatchImportJobOutput, error)
	CreateBatchImportJobWithContext(aws.Context, *frauddetector.CreateBatchImportJobInput, ...request.Option) (*frauddetector.CreateBatchImportJobOutput, error)
	CreateBatchImportJobRequest(*frauddetector.CreateBatchImportJobInput) (*request.Request, *frauddetector.CreateBatchImportJobOutput)

	CreateBatchPredictionJob(*frauddetector.CreateBatchPredictionJobInput) (*frauddetector.CreateBatchPredictionJobOutput, error)
	CreateBatchPredictionJobWithContext(aws.Context, *frauddetector.CreateBatchPredictionJobInput, ...request.Option) (*frauddetector.CreateBatchPredictionJobOutput, error)
	CreateBatchPredictionJobRequest(*frauddetector.CreateBatchPredictionJobInput) (*request.Request, *frauddetector.CreateBatchPredictionJobOutput)

	CreateDetectorVersion(*frauddetector.CreateDetectorVersionInput) (*frauddetector.CreateDetectorVersionOutput, error)
	CreateDetectorVersionWithContext(aws.Context, *frauddetector.CreateDetectorVersionInput, ...request.Option) (*frauddetector.CreateDetectorVersionOutput, error)
	CreateDetectorVersionRequest(*frauddetector.CreateDetectorVersionInput) (*request.Request, *frauddetector.CreateDetectorVersionOutput)

	CreateList(*frauddetector.CreateListInput) (*frauddetector.CreateListOutput, error)
	CreateListWithContext(aws.Context, *frauddetector.CreateListInput, ...request.Option) (*frauddetector.CreateListOutput, error)
	CreateListRequest(*frauddetector.CreateListInput) (*request.Request, *frauddetector.CreateListOutput)

	CreateModel(*frauddetector.CreateModelInput) (*frauddetector.CreateModelOutput, error)
	CreateModelWithContext(aws.Context, *frauddetector.CreateModelInput, ...request.Option) (*frauddetector.CreateModelOutput, error)
	CreateModelRequest(*frauddetector.CreateModelInput) (*request.Request, *frauddetector.CreateModelOutput)

	CreateModelVersion(*frauddetector.CreateModelVersionInput) (*frauddetector.CreateModelVersionOutput, error)
	CreateModelVersionWithContext(aws.Context, *frauddetector.CreateModelVersionInput, ...request.Option) (*frauddetector.CreateModelVersionOutput, error)
	CreateModelVersionRequest(*frauddetector.CreateModelVersionInput) (*request.Request, *frauddetector.CreateModelVersionOutput)

	CreateRule(*frauddetector.CreateRuleInput) (*frauddetector.CreateRuleOutput, error)
	CreateRuleWithContext(aws.Context, *frauddetector.CreateRuleInput, ...request.Option) (*frauddetector.CreateRuleOutput, error)
	CreateRuleRequest(*frauddetector.CreateRuleInput) (*request.Request, *frauddetector.CreateRuleOutput)

	CreateVariable(*frauddetector.CreateVariableInput) (*frauddetector.CreateVariableOutput, error)
	CreateVariableWithContext(aws.Context, *frauddetector.CreateVariableInput, ...request.Option) (*frauddetector.CreateVariableOutput, error)
	CreateVariableRequest(*frauddetector.CreateVariableInput) (*request.Request, *frauddetector.CreateVariableOutput)

	DeleteBatchImportJob(*frauddetector.DeleteBatchImportJobInput) (*frauddetector.DeleteBatchImportJobOutput, error)
	DeleteBatchImportJobWithContext(aws.Context, *frauddetector.DeleteBatchImportJobInput, ...request.Option) (*frauddetector.DeleteBatchImportJobOutput, error)
	DeleteBatchImportJobRequest(*frauddetector.DeleteBatchImportJobInput) (*request.Request, *frauddetector.DeleteBatchImportJobOutput)

	DeleteBatchPredictionJob(*frauddetector.DeleteBatchPredictionJobInput) (*frauddetector.DeleteBatchPredictionJobOutput, error)
	DeleteBatchPredictionJobWithContext(aws.Context, *frauddetector.DeleteBatchPredictionJobInput, ...request.Option) (*frauddetector.DeleteBatchPredictionJobOutput, error)
	DeleteBatchPredictionJobRequest(*frauddetector.DeleteBatchPredictionJobInput) (*request.Request, *frauddetector.DeleteBatchPredictionJobOutput)

	DeleteDetector(*frauddetector.DeleteDetectorInput) (*frauddetector.DeleteDetectorOutput, error)
	DeleteDetectorWithContext(aws.Context, *frauddetector.DeleteDetectorInput, ...request.Option) (*frauddetector.DeleteDetectorOutput, error)
	DeleteDetectorRequest(*frauddetector.DeleteDetectorInput) (*request.Request, *frauddetector.DeleteDetectorOutput)

	DeleteDetectorVersion(*frauddetector.DeleteDetectorVersionInput) (*frauddetector.DeleteDetectorVersionOutput, error)
	DeleteDetectorVersionWithContext(aws.Context, *frauddetector.DeleteDetectorVersionInput, ...request.Option) (*frauddetector.DeleteDetectorVersionOutput, error)
	DeleteDetectorVersionRequest(*frauddetector.DeleteDetectorVersionInput) (*request.Request, *frauddetector.DeleteDetectorVersionOutput)

	DeleteEntityType(*frauddetector.DeleteEntityTypeInput) (*frauddetector.DeleteEntityTypeOutput, error)
	DeleteEntityTypeWithContext(aws.Context, *frauddetector.DeleteEntityTypeInput, ...request.Option) (*frauddetector.DeleteEntityTypeOutput, error)
	DeleteEntityTypeRequest(*frauddetector.DeleteEntityTypeInput) (*request.Request, *frauddetector.DeleteEntityTypeOutput)

	DeleteEvent(*frauddetector.DeleteEventInput) (*frauddetector.DeleteEventOutput, error)
	DeleteEventWithContext(aws.Context, *frauddetector.DeleteEventInput, ...request.Option) (*frauddetector.DeleteEventOutput, error)
	DeleteEventRequest(*frauddetector.DeleteEventInput) (*request.Request, *frauddetector.DeleteEventOutput)

	DeleteEventType(*frauddetector.DeleteEventTypeInput) (*frauddetector.DeleteEventTypeOutput, error)
	DeleteEventTypeWithContext(aws.Context, *frauddetector.DeleteEventTypeInput, ...request.Option) (*frauddetector.DeleteEventTypeOutput, error)
	DeleteEventTypeRequest(*frauddetector.DeleteEventTypeInput) (*request.Request, *frauddetector.DeleteEventTypeOutput)

	DeleteEventsByEventType(*frauddetector.DeleteEventsByEventTypeInput) (*frauddetector.DeleteEventsByEventTypeOutput, error)
	DeleteEventsByEventTypeWithContext(aws.Context, *frauddetector.DeleteEventsByEventTypeInput, ...request.Option) (*frauddetector.DeleteEventsByEventTypeOutput, error)
	DeleteEventsByEventTypeRequest(*frauddetector.DeleteEventsByEventTypeInput) (*request.Request, *frauddetector.DeleteEventsByEventTypeOutput)

	DeleteExternalModel(*frauddetector.DeleteExternalModelInput) (*frauddetector.DeleteExternalModelOutput, error)
	DeleteExternalModelWithContext(aws.Context, *frauddetector.DeleteExternalModelInput, ...request.Option) (*frauddetector.DeleteExternalModelOutput, error)
	DeleteExternalModelRequest(*frauddetector.DeleteExternalModelInput) (*request.Request, *frauddetector.DeleteExternalModelOutput)

	DeleteLabel(*frauddetector.DeleteLabelInput) (*frauddetector.DeleteLabelOutput, error)
	DeleteLabelWithContext(aws.Context, *frauddetector.DeleteLabelInput, ...request.Option) (*frauddetector.DeleteLabelOutput, error)
	DeleteLabelRequest(*frauddetector.DeleteLabelInput) (*request.Request, *frauddetector.DeleteLabelOutput)

	DeleteList(*frauddetector.DeleteListInput) (*frauddetector.DeleteListOutput, error)
	DeleteListWithContext(aws.Context, *frauddetector.DeleteListInput, ...request.Option) (*frauddetector.DeleteListOutput, error)
	DeleteListRequest(*frauddetector.DeleteListInput) (*request.Request, *frauddetector.DeleteListOutput)

	DeleteModel(*frauddetector.DeleteModelInput) (*frauddetector.DeleteModelOutput, error)
	DeleteModelWithContext(aws.Context, *frauddetector.DeleteModelInput, ...request.Option) (*frauddetector.DeleteModelOutput, error)
	DeleteModelRequest(*frauddetector.DeleteModelInput) (*request.Request, *frauddetector.DeleteModelOutput)

	DeleteModelVersion(*frauddetector.DeleteModelVersionInput) (*frauddetector.DeleteModelVersionOutput, error)
	DeleteModelVersionWithContext(aws.Context, *frauddetector.DeleteModelVersionInput, ...request.Option) (*frauddetector.DeleteModelVersionOutput, error)
	DeleteModelVersionRequest(*frauddetector.DeleteModelVersionInput) (*request.Request, *frauddetector.DeleteModelVersionOutput)

	DeleteOutcome(*frauddetector.DeleteOutcomeInput) (*frauddetector.DeleteOutcomeOutput, error)
	DeleteOutcomeWithContext(aws.Context, *frauddetector.DeleteOutcomeInput, ...request.Option) (*frauddetector.DeleteOutcomeOutput, error)
	DeleteOutcomeRequest(*frauddetector.DeleteOutcomeInput) (*request.Request, *frauddetector.DeleteOutcomeOutput)

	DeleteRule(*frauddetector.DeleteRuleInput) (*frauddetector.DeleteRuleOutput, error)
	DeleteRuleWithContext(aws.Context, *frauddetector.DeleteRuleInput, ...request.Option) (*frauddetector.DeleteRuleOutput, error)
	DeleteRuleRequest(*frauddetector.DeleteRuleInput) (*request.Request, *frauddetector.DeleteRuleOutput)

	DeleteVariable(*frauddetector.DeleteVariableInput) (*frauddetector.DeleteVariableOutput, error)
	DeleteVariableWithContext(aws.Context, *frauddetector.DeleteVariableInput, ...request.Option) (*frauddetector.DeleteVariableOutput, error)
	DeleteVariableRequest(*frauddetector.DeleteVariableInput) (*request.Request, *frauddetector.DeleteVariableOutput)

	DescribeDetector(*frauddetector.DescribeDetectorInput) (*frauddetector.DescribeDetectorOutput, error)
	DescribeDetectorWithContext(aws.Context, *frauddetector.DescribeDetectorInput, ...request.Option) (*frauddetector.DescribeDetectorOutput, error)
	DescribeDetectorRequest(*frauddetector.DescribeDetectorInput) (*request.Request, *frauddetector.DescribeDetectorOutput)

	DescribeModelVersions(*frauddetector.DescribeModelVersionsInput) (*frauddetector.DescribeModelVersionsOutput, error)
	DescribeModelVersionsWithContext(aws.Context, *frauddetector.DescribeModelVersionsInput, ...request.Option) (*frauddetector.DescribeModelVersionsOutput, error)
	DescribeModelVersionsRequest(*frauddetector.DescribeModelVersionsInput) (*request.Request, *frauddetector.DescribeModelVersionsOutput)

	DescribeModelVersionsPages(*frauddetector.DescribeModelVersionsInput, func(*frauddetector.DescribeModelVersionsOutput, bool) bool) error
	DescribeModelVersionsPagesWithContext(aws.Context, *frauddetector.DescribeModelVersionsInput, func(*frauddetector.DescribeModelVersionsOutput, bool) bool, ...request.Option) error

	GetBatchImportJobs(*frauddetector.GetBatchImportJobsInput) (*frauddetector.GetBatchImportJobsOutput, error)
	GetBatchImportJobsWithContext(aws.Context, *frauddetector.GetBatchImportJobsInput, ...request.Option) (*frauddetector.GetBatchImportJobsOutput, error)
	GetBatchImportJobsRequest(*frauddetector.GetBatchImportJobsInput) (*request.Request, *frauddetector.GetBatchImportJobsOutput)

	GetBatchImportJobsPages(*frauddetector.GetBatchImportJobsInput, func(*frauddetector.GetBatchImportJobsOutput, bool) bool) error
	GetBatchImportJobsPagesWithContext(aws.Context, *frauddetector.GetBatchImportJobsInput, func(*frauddetector.GetBatchImportJobsOutput, bool) bool, ...request.Option) error

	GetBatchPredictionJobs(*frauddetector.GetBatchPredictionJobsInput) (*frauddetector.GetBatchPredictionJobsOutput, error)
	GetBatchPredictionJobsWithContext(aws.Context, *frauddetector.GetBatchPredictionJobsInput, ...request.Option) (*frauddetector.GetBatchPredictionJobsOutput, error)
	GetBatchPredictionJobsRequest(*frauddetector.GetBatchPredictionJobsInput) (*request.Request, *frauddetector.GetBatchPredictionJobsOutput)

	GetBatchPredictionJobsPages(*frauddetector.GetBatchPredictionJobsInput, func(*frauddetector.GetBatchPredictionJobsOutput, bool) bool) error
	GetBatchPredictionJobsPagesWithContext(aws.Context, *frauddetector.GetBatchPredictionJobsInput, func(*frauddetector.GetBatchPredictionJobsOutput, bool) bool, ...request.Option) error

	GetDeleteEventsByEventTypeStatus(*frauddetector.GetDeleteEventsByEventTypeStatusInput) (*frauddetector.GetDeleteEventsByEventTypeStatusOutput, error)
	GetDeleteEventsByEventTypeStatusWithContext(aws.Context, *frauddetector.GetDeleteEventsByEventTypeStatusInput, ...request.Option) (*frauddetector.GetDeleteEventsByEventTypeStatusOutput, error)
	GetDeleteEventsByEventTypeStatusRequest(*frauddetector.GetDeleteEventsByEventTypeStatusInput) (*request.Request, *frauddetector.GetDeleteEventsByEventTypeStatusOutput)

	GetDetectorVersion(*frauddetector.GetDetectorVersionInput) (*frauddetector.GetDetectorVersionOutput, error)
	GetDetectorVersionWithContext(aws.Context, *frauddetector.GetDetectorVersionInput, ...request.Option) (*frauddetector.GetDetectorVersionOutput, error)
	GetDetectorVersionRequest(*frauddetector.GetDetectorVersionInput) (*request.Request, *frauddetector.GetDetectorVersionOutput)

	GetDetectors(*frauddetector.GetDetectorsInput) (*frauddetector.GetDetectorsOutput, error)
	GetDetectorsWithContext(aws.Context, *frauddetector.GetDetectorsInput, ...request.Option) (*frauddetector.GetDetectorsOutput, error)
	GetDetectorsRequest(*frauddetector.GetDetectorsInput) (*request.Request, *frauddetector.GetDetectorsOutput)

	GetDetectorsPages(*frauddetector.GetDetectorsInput, func(*frauddetector.GetDetectorsOutput, bool) bool) error
	GetDetectorsPagesWithContext(aws.Context, *frauddetector.GetDetectorsInput, func(*frauddetector.GetDetectorsOutput, bool) bool, ...request.Option) error

	GetEntityTypes(*frauddetector.GetEntityTypesInput) (*frauddetector.GetEntityTypesOutput, error)
	GetEntityTypesWithContext(aws.Context, *frauddetector.GetEntityTypesInput, ...request.Option) (*frauddetector.GetEntityTypesOutput, error)
	GetEntityTypesRequest(*frauddetector.GetEntityTypesInput) (*request.Request, *frauddetector.GetEntityTypesOutput)

	GetEntityTypesPages(*frauddetector.GetEntityTypesInput, func(*frauddetector.GetEntityTypesOutput, bool) bool) error
	GetEntityTypesPagesWithContext(aws.Context, *frauddetector.GetEntityTypesInput, func(*frauddetector.GetEntityTypesOutput, bool) bool, ...request.Option) error

	GetEvent(*frauddetector.GetEventInput) (*frauddetector.GetEventOutput, error)
	GetEventWithContext(aws.Context, *frauddetector.GetEventInput, ...request.Option) (*frauddetector.GetEventOutput, error)
	GetEventRequest(*frauddetector.GetEventInput) (*request.Request, *frauddetector.GetEventOutput)

	GetEventPrediction(*frauddetector.GetEventPredictionInput) (*frauddetector.GetEventPredictionOutput, error)
	GetEventPredictionWithContext(aws.Context, *frauddetector.GetEventPredictionInput, ...request.Option) (*frauddetector.GetEventPredictionOutput, error)
	GetEventPredictionRequest(*frauddetector.GetEventPredictionInput) (*request.Request, *frauddetector.GetEventPredictionOutput)

	GetEventPredictionMetadata(*frauddetector.GetEventPredictionMetadataInput) (*frauddetector.GetEventPredictionMetadataOutput, error)
	GetEventPredictionMetadataWithContext(aws.Context, *frauddetector.GetEventPredictionMetadataInput, ...request.Option) (*frauddetector.GetEventPredictionMetadataOutput, error)
	GetEventPredictionMetadataRequest(*frauddetector.GetEventPredictionMetadataInput) (*request.Request, *frauddetector.GetEventPredictionMetadataOutput)

	GetEventTypes(*frauddetector.GetEventTypesInput) (*frauddetector.GetEventTypesOutput, error)
	GetEventTypesWithContext(aws.Context, *frauddetector.GetEventTypesInput, ...request.Option) (*frauddetector.GetEventTypesOutput, error)
	GetEventTypesRequest(*frauddetector.GetEventTypesInput) (*request.Request, *frauddetector.GetEventTypesOutput)

	GetEventTypesPages(*frauddetector.GetEventTypesInput, func(*frauddetector.GetEventTypesOutput, bool) bool) error
	GetEventTypesPagesWithContext(aws.Context, *frauddetector.GetEventTypesInput, func(*frauddetector.GetEventTypesOutput, bool) bool, ...request.Option) error

	GetExternalModels(*frauddetector.GetExternalModelsInput) (*frauddetector.GetExternalModelsOutput, error)
	GetExternalModelsWithContext(aws.Context, *frauddetector.GetExternalModelsInput, ...request.Option) (*frauddetector.GetExternalModelsOutput, error)
	GetExternalModelsRequest(*frauddetector.GetExternalModelsInput) (*request.Request, *frauddetector.GetExternalModelsOutput)

	GetExternalModelsPages(*frauddetector.GetExternalModelsInput, func(*frauddetector.GetExternalModelsOutput, bool) bool) error
	GetExternalModelsPagesWithContext(aws.Context, *frauddetector.GetExternalModelsInput, func(*frauddetector.GetExternalModelsOutput, bool) bool, ...request.Option) error

	GetKMSEncryptionKey(*frauddetector.GetKMSEncryptionKeyInput) (*frauddetector.GetKMSEncryptionKeyOutput, error)
	GetKMSEncryptionKeyWithContext(aws.Context, *frauddetector.GetKMSEncryptionKeyInput, ...request.Option) (*frauddetector.GetKMSEncryptionKeyOutput, error)
	GetKMSEncryptionKeyRequest(*frauddetector.GetKMSEncryptionKeyInput) (*request.Request, *frauddetector.GetKMSEncryptionKeyOutput)

	GetLabels(*frauddetector.GetLabelsInput) (*frauddetector.GetLabelsOutput, error)
	GetLabelsWithContext(aws.Context, *frauddetector.GetLabelsInput, ...request.Option) (*frauddetector.GetLabelsOutput, error)
	GetLabelsRequest(*frauddetector.GetLabelsInput) (*request.Request, *frauddetector.GetLabelsOutput)

	GetLabelsPages(*frauddetector.GetLabelsInput, func(*frauddetector.GetLabelsOutput, bool) bool) error
	GetLabelsPagesWithContext(aws.Context, *frauddetector.GetLabelsInput, func(*frauddetector.GetLabelsOutput, bool) bool, ...request.Option) error

	GetListElements(*frauddetector.GetListElementsInput) (*frauddetector.GetListElementsOutput, error)
	GetListElementsWithContext(aws.Context, *frauddetector.GetListElementsInput, ...request.Option) (*frauddetector.GetListElementsOutput, error)
	GetListElementsRequest(*frauddetector.GetListElementsInput) (*request.Request, *frauddetector.GetListElementsOutput)

	GetListElementsPages(*frauddetector.GetListElementsInput, func(*frauddetector.GetListElementsOutput, bool) bool) error
	GetListElementsPagesWithContext(aws.Context, *frauddetector.GetListElementsInput, func(*frauddetector.GetListElementsOutput, bool) bool, ...request.Option) error

	GetListsMetadata(*frauddetector.GetListsMetadataInput) (*frauddetector.GetListsMetadataOutput, error)
	GetListsMetadataWithContext(aws.Context, *frauddetector.GetListsMetadataInput, ...request.Option) (*frauddetector.GetListsMetadataOutput, error)
	GetListsMetadataRequest(*frauddetector.GetListsMetadataInput) (*request.Request, *frauddetector.GetListsMetadataOutput)

	GetListsMetadataPages(*frauddetector.GetListsMetadataInput, func(*frauddetector.GetListsMetadataOutput, bool) bool) error
	GetListsMetadataPagesWithContext(aws.Context, *frauddetector.GetListsMetadataInput, func(*frauddetector.GetListsMetadataOutput, bool) bool, ...request.Option) error

	GetModelVersion(*frauddetector.GetModelVersionInput) (*frauddetector.GetModelVersionOutput, error)
	GetModelVersionWithContext(aws.Context, *frauddetector.GetModelVersionInput, ...request.Option) (*frauddetector.GetModelVersionOutput, error)
	GetModelVersionRequest(*frauddetector.GetModelVersionInput) (*request.Request, *frauddetector.GetModelVersionOutput)

	GetModels(*frauddetector.GetModelsInput) (*frauddetector.GetModelsOutput, error)
	GetModelsWithContext(aws.Context, *frauddetector.GetModelsInput, ...request.Option) (*frauddetector.GetModelsOutput, error)
	GetModelsRequest(*frauddetector.GetModelsInput) (*request.Request, *frauddetector.GetModelsOutput)

	GetModelsPages(*frauddetector.GetModelsInput, func(*frauddetector.GetModelsOutput, bool) bool) error
	GetModelsPagesWithContext(aws.Context, *frauddetector.GetModelsInput, func(*frauddetector.GetModelsOutput, bool) bool, ...request.Option) error

	GetOutcomes(*frauddetector.GetOutcomesInput) (*frauddetector.GetOutcomesOutput, error)
	GetOutcomesWithContext(aws.Context, *frauddetector.GetOutcomesInput, ...request.Option) (*frauddetector.GetOutcomesOutput, error)
	GetOutcomesRequest(*frauddetector.GetOutcomesInput) (*request.Request, *frauddetector.GetOutcomesOutput)

	GetOutcomesPages(*frauddetector.GetOutcomesInput, func(*frauddetector.GetOutcomesOutput, bool) bool) error
	GetOutcomesPagesWithContext(aws.Context, *frauddetector.GetOutcomesInput, func(*frauddetector.GetOutcomesOutput, bool) bool, ...request.Option) error

	GetRules(*frauddetector.GetRulesInput) (*frauddetector.GetRulesOutput, error)
	GetRulesWithContext(aws.Context, *frauddetector.GetRulesInput, ...request.Option) (*frauddetector.GetRulesOutput, error)
	GetRulesRequest(*frauddetector.GetRulesInput) (*request.Request, *frauddetector.GetRulesOutput)

	GetRulesPages(*frauddetector.GetRulesInput, func(*frauddetector.GetRulesOutput, bool) bool) error
	GetRulesPagesWithContext(aws.Context, *frauddetector.GetRulesInput, func(*frauddetector.GetRulesOutput, bool) bool, ...request.Option) error

	GetVariables(*frauddetector.GetVariablesInput) (*frauddetector.GetVariablesOutput, error)
	GetVariablesWithContext(aws.Context, *frauddetector.GetVariablesInput, ...request.Option) (*frauddetector.GetVariablesOutput, error)
	GetVariablesRequest(*frauddetector.GetVariablesInput) (*request.Request, *frauddetector.GetVariablesOutput)

	GetVariablesPages(*frauddetector.GetVariablesInput, func(*frauddetector.GetVariablesOutput, bool) bool) error
	GetVariablesPagesWithContext(aws.Context, *frauddetector.GetVariablesInput, func(*frauddetector.GetVariablesOutput, bool) bool, ...request.Option) error

	ListEventPredictions(*frauddetector.ListEventPredictionsInput) (*frauddetector.ListEventPredictionsOutput, error)
	ListEventPredictionsWithContext(aws.Context, *frauddetector.ListEventPredictionsInput, ...request.Option) (*frauddetector.ListEventPredictionsOutput, error)
	ListEventPredictionsRequest(*frauddetector.ListEventPredictionsInput) (*request.Request, *frauddetector.ListEventPredictionsOutput)

	ListEventPredictionsPages(*frauddetector.ListEventPredictionsInput, func(*frauddetector.ListEventPredictionsOutput, bool) bool) error
	ListEventPredictionsPagesWithContext(aws.Context, *frauddetector.ListEventPredictionsInput, func(*frauddetector.ListEventPredictionsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*frauddetector.ListTagsForResourceInput) (*frauddetector.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *frauddetector.ListTagsForResourceInput, ...request.Option) (*frauddetector.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*frauddetector.ListTagsForResourceInput) (*request.Request, *frauddetector.ListTagsForResourceOutput)

	ListTagsForResourcePages(*frauddetector.ListTagsForResourceInput, func(*frauddetector.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *frauddetector.ListTagsForResourceInput, func(*frauddetector.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	PutDetector(*frauddetector.PutDetectorInput) (*frauddetector.PutDetectorOutput, error)
	PutDetectorWithContext(aws.Context, *frauddetector.PutDetectorInput, ...request.Option) (*frauddetector.PutDetectorOutput, error)
	PutDetectorRequest(*frauddetector.PutDetectorInput) (*request.Request, *frauddetector.PutDetectorOutput)

	PutEntityType(*frauddetector.PutEntityTypeInput) (*frauddetector.PutEntityTypeOutput, error)
	PutEntityTypeWithContext(aws.Context, *frauddetector.PutEntityTypeInput, ...request.Option) (*frauddetector.PutEntityTypeOutput, error)
	PutEntityTypeRequest(*frauddetector.PutEntityTypeInput) (*request.Request, *frauddetector.PutEntityTypeOutput)

	PutEventType(*frauddetector.PutEventTypeInput) (*frauddetector.PutEventTypeOutput, error)
	PutEventTypeWithContext(aws.Context, *frauddetector.PutEventTypeInput, ...request.Option) (*frauddetector.PutEventTypeOutput, error)
	PutEventTypeRequest(*frauddetector.PutEventTypeInput) (*request.Request, *frauddetector.PutEventTypeOutput)

	PutExternalModel(*frauddetector.PutExternalModelInput) (*frauddetector.PutExternalModelOutput, error)
	PutExternalModelWithContext(aws.Context, *frauddetector.PutExternalModelInput, ...request.Option) (*frauddetector.PutExternalModelOutput, error)
	PutExternalModelRequest(*frauddetector.PutExternalModelInput) (*request.Request, *frauddetector.PutExternalModelOutput)

	PutKMSEncryptionKey(*frauddetector.PutKMSEncryptionKeyInput) (*frauddetector.PutKMSEncryptionKeyOutput, error)
	PutKMSEncryptionKeyWithContext(aws.Context, *frauddetector.PutKMSEncryptionKeyInput, ...request.Option) (*frauddetector.PutKMSEncryptionKeyOutput, error)
	PutKMSEncryptionKeyRequest(*frauddetector.PutKMSEncryptionKeyInput) (*request.Request, *frauddetector.PutKMSEncryptionKeyOutput)

	PutLabel(*frauddetector.PutLabelInput) (*frauddetector.PutLabelOutput, error)
	PutLabelWithContext(aws.Context, *frauddetector.PutLabelInput, ...request.Option) (*frauddetector.PutLabelOutput, error)
	PutLabelRequest(*frauddetector.PutLabelInput) (*request.Request, *frauddetector.PutLabelOutput)

	PutOutcome(*frauddetector.PutOutcomeInput) (*frauddetector.PutOutcomeOutput, error)
	PutOutcomeWithContext(aws.Context, *frauddetector.PutOutcomeInput, ...request.Option) (*frauddetector.PutOutcomeOutput, error)
	PutOutcomeRequest(*frauddetector.PutOutcomeInput) (*request.Request, *frauddetector.PutOutcomeOutput)

	SendEvent(*frauddetector.SendEventInput) (*frauddetector.SendEventOutput, error)
	SendEventWithContext(aws.Context, *frauddetector.SendEventInput, ...request.Option) (*frauddetector.SendEventOutput, error)
	SendEventRequest(*frauddetector.SendEventInput) (*request.Request, *frauddetector.SendEventOutput)

	TagResource(*frauddetector.TagResourceInput) (*frauddetector.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *frauddetector.TagResourceInput, ...request.Option) (*frauddetector.TagResourceOutput, error)
	TagResourceRequest(*frauddetector.TagResourceInput) (*request.Request, *frauddetector.TagResourceOutput)

	UntagResource(*frauddetector.UntagResourceInput) (*frauddetector.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *frauddetector.UntagResourceInput, ...request.Option) (*frauddetector.UntagResourceOutput, error)
	UntagResourceRequest(*frauddetector.UntagResourceInput) (*request.Request, *frauddetector.UntagResourceOutput)

	UpdateDetectorVersion(*frauddetector.UpdateDetectorVersionInput) (*frauddetector.UpdateDetectorVersionOutput, error)
	UpdateDetectorVersionWithContext(aws.Context, *frauddetector.UpdateDetectorVersionInput, ...request.Option) (*frauddetector.UpdateDetectorVersionOutput, error)
	UpdateDetectorVersionRequest(*frauddetector.UpdateDetectorVersionInput) (*request.Request, *frauddetector.UpdateDetectorVersionOutput)

	UpdateDetectorVersionMetadata(*frauddetector.UpdateDetectorVersionMetadataInput) (*frauddetector.UpdateDetectorVersionMetadataOutput, error)
	UpdateDetectorVersionMetadataWithContext(aws.Context, *frauddetector.UpdateDetectorVersionMetadataInput, ...request.Option) (*frauddetector.UpdateDetectorVersionMetadataOutput, error)
	UpdateDetectorVersionMetadataRequest(*frauddetector.UpdateDetectorVersionMetadataInput) (*request.Request, *frauddetector.UpdateDetectorVersionMetadataOutput)

	UpdateDetectorVersionStatus(*frauddetector.UpdateDetectorVersionStatusInput) (*frauddetector.UpdateDetectorVersionStatusOutput, error)
	UpdateDetectorVersionStatusWithContext(aws.Context, *frauddetector.UpdateDetectorVersionStatusInput, ...request.Option) (*frauddetector.UpdateDetectorVersionStatusOutput, error)
	UpdateDetectorVersionStatusRequest(*frauddetector.UpdateDetectorVersionStatusInput) (*request.Request, *frauddetector.UpdateDetectorVersionStatusOutput)

	UpdateEventLabel(*frauddetector.UpdateEventLabelInput) (*frauddetector.UpdateEventLabelOutput, error)
	UpdateEventLabelWithContext(aws.Context, *frauddetector.UpdateEventLabelInput, ...request.Option) (*frauddetector.UpdateEventLabelOutput, error)
	UpdateEventLabelRequest(*frauddetector.UpdateEventLabelInput) (*request.Request, *frauddetector.UpdateEventLabelOutput)

	UpdateList(*frauddetector.UpdateListInput) (*frauddetector.UpdateListOutput, error)
	UpdateListWithContext(aws.Context, *frauddetector.UpdateListInput, ...request.Option) (*frauddetector.UpdateListOutput, error)
	UpdateListRequest(*frauddetector.UpdateListInput) (*request.Request, *frauddetector.UpdateListOutput)

	UpdateModel(*frauddetector.UpdateModelInput) (*frauddetector.UpdateModelOutput, error)
	UpdateModelWithContext(aws.Context, *frauddetector.UpdateModelInput, ...request.Option) (*frauddetector.UpdateModelOutput, error)
	UpdateModelRequest(*frauddetector.UpdateModelInput) (*request.Request, *frauddetector.UpdateModelOutput)

	UpdateModelVersion(*frauddetector.UpdateModelVersionInput) (*frauddetector.UpdateModelVersionOutput, error)
	UpdateModelVersionWithContext(aws.Context, *frauddetector.UpdateModelVersionInput, ...request.Option) (*frauddetector.UpdateModelVersionOutput, error)
	UpdateModelVersionRequest(*frauddetector.UpdateModelVersionInput) (*request.Request, *frauddetector.UpdateModelVersionOutput)

	UpdateModelVersionStatus(*frauddetector.UpdateModelVersionStatusInput) (*frauddetector.UpdateModelVersionStatusOutput, error)
	UpdateModelVersionStatusWithContext(aws.Context, *frauddetector.UpdateModelVersionStatusInput, ...request.Option) (*frauddetector.UpdateModelVersionStatusOutput, error)
	UpdateModelVersionStatusRequest(*frauddetector.UpdateModelVersionStatusInput) (*request.Request, *frauddetector.UpdateModelVersionStatusOutput)

	UpdateRuleMetadata(*frauddetector.UpdateRuleMetadataInput) (*frauddetector.UpdateRuleMetadataOutput, error)
	UpdateRuleMetadataWithContext(aws.Context, *frauddetector.UpdateRuleMetadataInput, ...request.Option) (*frauddetector.UpdateRuleMetadataOutput, error)
	UpdateRuleMetadataRequest(*frauddetector.UpdateRuleMetadataInput) (*request.Request, *frauddetector.UpdateRuleMetadataOutput)

	UpdateRuleVersion(*frauddetector.UpdateRuleVersionInput) (*frauddetector.UpdateRuleVersionOutput, error)
	UpdateRuleVersionWithContext(aws.Context, *frauddetector.UpdateRuleVersionInput, ...request.Option) (*frauddetector.UpdateRuleVersionOutput, error)
	UpdateRuleVersionRequest(*frauddetector.UpdateRuleVersionInput) (*request.Request, *frauddetector.UpdateRuleVersionOutput)

	UpdateVariable(*frauddetector.UpdateVariableInput) (*frauddetector.UpdateVariableOutput, error)
	UpdateVariableWithContext(aws.Context, *frauddetector.UpdateVariableInput, ...request.Option) (*frauddetector.UpdateVariableOutput, error)
	UpdateVariableRequest(*frauddetector.UpdateVariableInput) (*request.Request, *frauddetector.UpdateVariableOutput)
}

var _ FraudDetectorAPI = (*frauddetector.FraudDetector)(nil)
