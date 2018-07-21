// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudsearchdomainiface provides an interface to enable mocking the Amazon CloudSearch Domain service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudsearchdomainiface

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
)

// CloudSearchDomainAPI provides an interface to enable mocking the
// cloudsearchdomain.CloudSearchDomain service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudSearch Domain.
//    func myFunc(svc cloudsearchdomainiface.CloudSearchDomainAPI) bool {
//        // Make svc.Search request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cloudsearchdomain.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudSearchDomainClient struct {
//        cloudsearchdomainiface.CloudSearchDomainAPI
//    }
//    func (m *mockCloudSearchDomainClient) Search(input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudSearchDomainClient{}
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
type CloudSearchDomainAPI interface {
	SearchRequest(*cloudsearchdomain.SearchInput) cloudsearchdomain.SearchRequest

	SuggestRequest(*cloudsearchdomain.SuggestInput) cloudsearchdomain.SuggestRequest

	UploadDocumentsRequest(*cloudsearchdomain.UploadDocumentsInput) cloudsearchdomain.UploadDocumentsRequest
}

var _ CloudSearchDomainAPI = (*cloudsearchdomain.CloudSearchDomain)(nil)
