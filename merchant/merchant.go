package merchant

import "github.com/google/uuid"

type Merchant struct {
	ID   uuid.UUID
	logo string
	Team map[string]Member
}

type Member struct {
	Email string
	Name  string
}


