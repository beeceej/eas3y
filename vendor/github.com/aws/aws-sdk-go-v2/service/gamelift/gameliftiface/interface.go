// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package gameliftiface provides an interface to enable mocking the Amazon GameLift service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package gameliftiface

import (
	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

// GameLiftAPI provides an interface to enable mocking the
// gamelift.GameLift service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon GameLift.
//    func myFunc(svc gameliftiface.GameLiftAPI) bool {
//        // Make svc.AcceptMatch request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := gamelift.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockGameLiftClient struct {
//        gameliftiface.GameLiftAPI
//    }
//    func (m *mockGameLiftClient) AcceptMatch(input *gamelift.AcceptMatchInput) (*gamelift.AcceptMatchOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockGameLiftClient{}
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
type GameLiftAPI interface {
	AcceptMatchRequest(*gamelift.AcceptMatchInput) gamelift.AcceptMatchRequest

	CreateAliasRequest(*gamelift.CreateAliasInput) gamelift.CreateAliasRequest

	CreateBuildRequest(*gamelift.CreateBuildInput) gamelift.CreateBuildRequest

	CreateFleetRequest(*gamelift.CreateFleetInput) gamelift.CreateFleetRequest

	CreateGameSessionRequest(*gamelift.CreateGameSessionInput) gamelift.CreateGameSessionRequest

	CreateGameSessionQueueRequest(*gamelift.CreateGameSessionQueueInput) gamelift.CreateGameSessionQueueRequest

	CreateMatchmakingConfigurationRequest(*gamelift.CreateMatchmakingConfigurationInput) gamelift.CreateMatchmakingConfigurationRequest

	CreateMatchmakingRuleSetRequest(*gamelift.CreateMatchmakingRuleSetInput) gamelift.CreateMatchmakingRuleSetRequest

	CreatePlayerSessionRequest(*gamelift.CreatePlayerSessionInput) gamelift.CreatePlayerSessionRequest

	CreatePlayerSessionsRequest(*gamelift.CreatePlayerSessionsInput) gamelift.CreatePlayerSessionsRequest

	CreateVpcPeeringAuthorizationRequest(*gamelift.CreateVpcPeeringAuthorizationInput) gamelift.CreateVpcPeeringAuthorizationRequest

	CreateVpcPeeringConnectionRequest(*gamelift.CreateVpcPeeringConnectionInput) gamelift.CreateVpcPeeringConnectionRequest

	DeleteAliasRequest(*gamelift.DeleteAliasInput) gamelift.DeleteAliasRequest

	DeleteBuildRequest(*gamelift.DeleteBuildInput) gamelift.DeleteBuildRequest

	DeleteFleetRequest(*gamelift.DeleteFleetInput) gamelift.DeleteFleetRequest

	DeleteGameSessionQueueRequest(*gamelift.DeleteGameSessionQueueInput) gamelift.DeleteGameSessionQueueRequest

	DeleteMatchmakingConfigurationRequest(*gamelift.DeleteMatchmakingConfigurationInput) gamelift.DeleteMatchmakingConfigurationRequest

	DeleteScalingPolicyRequest(*gamelift.DeleteScalingPolicyInput) gamelift.DeleteScalingPolicyRequest

	DeleteVpcPeeringAuthorizationRequest(*gamelift.DeleteVpcPeeringAuthorizationInput) gamelift.DeleteVpcPeeringAuthorizationRequest

	DeleteVpcPeeringConnectionRequest(*gamelift.DeleteVpcPeeringConnectionInput) gamelift.DeleteVpcPeeringConnectionRequest

	DescribeAliasRequest(*gamelift.DescribeAliasInput) gamelift.DescribeAliasRequest

	DescribeBuildRequest(*gamelift.DescribeBuildInput) gamelift.DescribeBuildRequest

	DescribeEC2InstanceLimitsRequest(*gamelift.DescribeEC2InstanceLimitsInput) gamelift.DescribeEC2InstanceLimitsRequest

	DescribeFleetAttributesRequest(*gamelift.DescribeFleetAttributesInput) gamelift.DescribeFleetAttributesRequest

	DescribeFleetCapacityRequest(*gamelift.DescribeFleetCapacityInput) gamelift.DescribeFleetCapacityRequest

	DescribeFleetEventsRequest(*gamelift.DescribeFleetEventsInput) gamelift.DescribeFleetEventsRequest

	DescribeFleetPortSettingsRequest(*gamelift.DescribeFleetPortSettingsInput) gamelift.DescribeFleetPortSettingsRequest

	DescribeFleetUtilizationRequest(*gamelift.DescribeFleetUtilizationInput) gamelift.DescribeFleetUtilizationRequest

	DescribeGameSessionDetailsRequest(*gamelift.DescribeGameSessionDetailsInput) gamelift.DescribeGameSessionDetailsRequest

	DescribeGameSessionPlacementRequest(*gamelift.DescribeGameSessionPlacementInput) gamelift.DescribeGameSessionPlacementRequest

	DescribeGameSessionQueuesRequest(*gamelift.DescribeGameSessionQueuesInput) gamelift.DescribeGameSessionQueuesRequest

	DescribeGameSessionsRequest(*gamelift.DescribeGameSessionsInput) gamelift.DescribeGameSessionsRequest

	DescribeInstancesRequest(*gamelift.DescribeInstancesInput) gamelift.DescribeInstancesRequest

	DescribeMatchmakingRequest(*gamelift.DescribeMatchmakingInput) gamelift.DescribeMatchmakingRequest

	DescribeMatchmakingConfigurationsRequest(*gamelift.DescribeMatchmakingConfigurationsInput) gamelift.DescribeMatchmakingConfigurationsRequest

	DescribeMatchmakingRuleSetsRequest(*gamelift.DescribeMatchmakingRuleSetsInput) gamelift.DescribeMatchmakingRuleSetsRequest

	DescribePlayerSessionsRequest(*gamelift.DescribePlayerSessionsInput) gamelift.DescribePlayerSessionsRequest

	DescribeRuntimeConfigurationRequest(*gamelift.DescribeRuntimeConfigurationInput) gamelift.DescribeRuntimeConfigurationRequest

	DescribeScalingPoliciesRequest(*gamelift.DescribeScalingPoliciesInput) gamelift.DescribeScalingPoliciesRequest

	DescribeVpcPeeringAuthorizationsRequest(*gamelift.DescribeVpcPeeringAuthorizationsInput) gamelift.DescribeVpcPeeringAuthorizationsRequest

	DescribeVpcPeeringConnectionsRequest(*gamelift.DescribeVpcPeeringConnectionsInput) gamelift.DescribeVpcPeeringConnectionsRequest

	GetGameSessionLogUrlRequest(*gamelift.GetGameSessionLogUrlInput) gamelift.GetGameSessionLogUrlRequest

	GetInstanceAccessRequest(*gamelift.GetInstanceAccessInput) gamelift.GetInstanceAccessRequest

	ListAliasesRequest(*gamelift.ListAliasesInput) gamelift.ListAliasesRequest

	ListBuildsRequest(*gamelift.ListBuildsInput) gamelift.ListBuildsRequest

	ListFleetsRequest(*gamelift.ListFleetsInput) gamelift.ListFleetsRequest

	PutScalingPolicyRequest(*gamelift.PutScalingPolicyInput) gamelift.PutScalingPolicyRequest

	RequestUploadCredentialsRequest(*gamelift.RequestUploadCredentialsInput) gamelift.RequestUploadCredentialsRequest

	ResolveAliasRequest(*gamelift.ResolveAliasInput) gamelift.ResolveAliasRequest

	SearchGameSessionsRequest(*gamelift.SearchGameSessionsInput) gamelift.SearchGameSessionsRequest

	StartFleetActionsRequest(*gamelift.StartFleetActionsInput) gamelift.StartFleetActionsRequest

	StartGameSessionPlacementRequest(*gamelift.StartGameSessionPlacementInput) gamelift.StartGameSessionPlacementRequest

	StartMatchBackfillRequest(*gamelift.StartMatchBackfillInput) gamelift.StartMatchBackfillRequest

	StartMatchmakingRequest(*gamelift.StartMatchmakingInput) gamelift.StartMatchmakingRequest

	StopFleetActionsRequest(*gamelift.StopFleetActionsInput) gamelift.StopFleetActionsRequest

	StopGameSessionPlacementRequest(*gamelift.StopGameSessionPlacementInput) gamelift.StopGameSessionPlacementRequest

	StopMatchmakingRequest(*gamelift.StopMatchmakingInput) gamelift.StopMatchmakingRequest

	UpdateAliasRequest(*gamelift.UpdateAliasInput) gamelift.UpdateAliasRequest

	UpdateBuildRequest(*gamelift.UpdateBuildInput) gamelift.UpdateBuildRequest

	UpdateFleetAttributesRequest(*gamelift.UpdateFleetAttributesInput) gamelift.UpdateFleetAttributesRequest

	UpdateFleetCapacityRequest(*gamelift.UpdateFleetCapacityInput) gamelift.UpdateFleetCapacityRequest

	UpdateFleetPortSettingsRequest(*gamelift.UpdateFleetPortSettingsInput) gamelift.UpdateFleetPortSettingsRequest

	UpdateGameSessionRequest(*gamelift.UpdateGameSessionInput) gamelift.UpdateGameSessionRequest

	UpdateGameSessionQueueRequest(*gamelift.UpdateGameSessionQueueInput) gamelift.UpdateGameSessionQueueRequest

	UpdateMatchmakingConfigurationRequest(*gamelift.UpdateMatchmakingConfigurationInput) gamelift.UpdateMatchmakingConfigurationRequest

	UpdateRuntimeConfigurationRequest(*gamelift.UpdateRuntimeConfigurationInput) gamelift.UpdateRuntimeConfigurationRequest

	ValidateMatchmakingRuleSetRequest(*gamelift.ValidateMatchmakingRuleSetInput) gamelift.ValidateMatchmakingRuleSetRequest
}

var _ GameLiftAPI = (*gamelift.GameLift)(nil)
