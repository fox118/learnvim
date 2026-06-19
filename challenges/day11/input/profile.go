package day11

type Profile struct {
	User  string
	Email string
}

func DisplayName(p Profile) string {
	return p.User
}
