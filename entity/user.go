package entity

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type ListUsersResponse struct {
	Users      []*Data `json:"users"`
	TotalCount int32   `json:"total_count"`
}

type Data struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
}
