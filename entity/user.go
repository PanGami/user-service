package entity

type User struct {
	ID         int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string         `gorm:"size:255;not null;unique" json:"username"`
	FullName   string         `gorm:"size:255;not null" json:"full_name"`
	Password   string         `gorm:"size:255;not null" json:"password"`
	Activities []UserActivity `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"activities"`
}

type ListUsersResponse struct {
	Users      []*Data `json:"users"`
	TotalCount int32   `json:"total_count"`
}

type Data struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
}
