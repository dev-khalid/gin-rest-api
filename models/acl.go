// Read write permissions of resources.
// Resources should be enum type. "Event"
// Roles should be enum type. "Admin", "User"
package models

type Resource string

const (
	EventResource Resource = "Event"
)

type Role string

const (
	AdminRole Role = "Admin"
	UserRole  Role = "User"
)

type ACL struct {
	ID       uint     `json:"id" gorm:"primarykey"`
	Resource Resource `json:"resource" gorm:"type:enum('Event')"`
	Role     Role     `json:"role" gorm:"type:enum('Admin', 'User')"`
}
