// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/v2/pkg/cloud/amazon/eksctl (interfaces: EKSCtl)

package eksctl_test

import (
	"reflect"
	"time"

	cluster "github.com/jenkins-x/jx/v2/pkg/cluster"
	pegomock "github.com/petergtz/pegomock"
)

type MockEKSCtl struct {
	fail func(message string, callerSkip ...int)
}

func NewMockEKSCtl(options ...pegomock.Option) *MockEKSCtl {
	mock := &MockEKSCtl{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockEKSCtl) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockEKSCtl) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockEKSCtl) DeleteCluster(_param0 *cluster.Cluster) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEKSCtl().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DeleteCluster", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockEKSCtl) VerifyWasCalledOnce() *VerifierMockEKSCtl {
	return &VerifierMockEKSCtl{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockEKSCtl) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockEKSCtl {
	return &VerifierMockEKSCtl{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockEKSCtl) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockEKSCtl {
	return &VerifierMockEKSCtl{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockEKSCtl) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockEKSCtl {
	return &VerifierMockEKSCtl{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockEKSCtl struct {
	mock                   *MockEKSCtl
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockEKSCtl) DeleteCluster(_param0 *cluster.Cluster) *MockEKSCtl_DeleteCluster_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DeleteCluster", params, verifier.timeout)
	return &MockEKSCtl_DeleteCluster_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEKSCtl_DeleteCluster_OngoingVerification struct {
	mock              *MockEKSCtl
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEKSCtl_DeleteCluster_OngoingVerification) GetCapturedArguments() *cluster.Cluster {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockEKSCtl_DeleteCluster_OngoingVerification) GetAllCapturedArguments() (_param0 []*cluster.Cluster) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*cluster.Cluster, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*cluster.Cluster)
		}
	}
	return
}