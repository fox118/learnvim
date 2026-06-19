package day11

type Profile struct {
	Username string
	Email    string
}

func DisplayName(p Profile) string {
	return p.Username
}
