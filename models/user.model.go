package models

type User struct {
	Model
	Username     string   `gorm:"size:50;unique;not null" json:"username"`
	Email        string   `gorm:"size:100;unique;not null" json:"email"`
	Password     string   `gorm:"size:50" json:"password"`
	RefreshToken string   `gorm:"size:100" json:"refresh_token"`
	ProfileUrl   string   `gorm:"size:100" json:"profile_url"`
	Role         UserRole `gorm:"not null;default:member" json:"role"`
	// Relation
	MessagesSent      []Message          `gorm:"foreignKey:SenderID"`
	MessagesReceived  []Message          `gorm:"foreignKey:RecipientID"`
	Classes           []Class            `gorm:"foreignKey:CreatedByID"`
	UserClasses       []UserClass        `gorm:"foreignKey:UserID"`
	Discussions       []Discussion       `gorm:"foreignKey:UserID"`
	UserQuizResponses []UserQuizResponse `gorm:"foreignKey:UserID"`
}

type UserRole string

const (
	AdminUserRole   UserRole = "admin"
	TeacherUserRole UserRole = "teacher"
	MemberUserRole  UserRole = "member"
)
