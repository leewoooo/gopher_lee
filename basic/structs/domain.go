package structs

// User for expose test
type User struct {
	Name string
	pID  int
}

// Tenant for embeded test
type Tenant struct {
	User
	TenantName string
}
