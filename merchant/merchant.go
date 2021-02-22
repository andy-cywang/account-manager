package merchant

type Merchant struct {
	ID      string            `json:"merchantID" bson:"_id"`
	logo    string            `json:"logo" bson:"logo"`
	Members map[string]Member `json:"members" bson:"members"`
}

type Member struct {
	Email string `json:"email" bson:"email"`
	Name  string `json:"name" bson:"name"`
}

func (m *Merchant) getTeamMembers() []Member {
	return nil
}
