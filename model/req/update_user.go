package req

type ReqUpdate struct {
	FullName string `json:"fullName,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Address  string `json:"address,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}