// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import api "github.com/hyperledger/udo/gossip/api"
import mock "github.com/stretchr/testify/mock"

// SecurityAdvisor is an autogenerated mock type for the SecurityAdvisor type
type SecurityAdvisor struct {
	mock.Mock
}

// OrgByPeerIdentity provides a mock function with given fields: _a0
func (_m *SecurityAdvisor) OrgByPeerIdentity(_a0 api.PeerIdentityType) api.OrgIdentityType {
	ret := _m.Called(_a0)

	var r0 api.OrgIdentityType
	if rf, ok := ret.Get(0).(func(api.PeerIdentityType) api.OrgIdentityType); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.OrgIdentityType)
		}
	}

	return r0
}
