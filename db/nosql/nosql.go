package nosql

import "account-manager/merchant"

type Nosql interface {

	// GetMerchants gets all merchant accounts
	GetMerchants() ([]merchant.Merchant, error)

	// AddMember adds a team member to merchant account
	AddMember(merchantID string, member merchant.Member) error

	// UpdateMember updates a team member from merchant account
	UpdateMember(merchantID string, member merchant.Member) error

	// DeleteMember deletes a team member from merchant account
	DeleteMember(merchantID string, memberEmail string) error

	// GetMembers gets all team members from merchant account
	GetMembers(merchantID string) ([]merchant.Member, error)
}