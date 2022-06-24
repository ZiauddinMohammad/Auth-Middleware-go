package data

type user struct {
	fullname   string
	email      string
	username   string
	password   string
	createdate string
	role       string
}

var users = []user{
	{
		fullname:   "userone",
		email:      "userone@abc.com",
		username:   "user1",
		password:   "pass1",
		createdate: "1780506070",
		role:       "admin",
	},
	{
		fullname:   "usertwo",
		email:      "usertwo@abc.com",
		username:   "user2",
		password:   "pass2",
		createdate: "1780506070",
		role:       "user"},
}

//returns a user object based on user email
func GetUser(email string) (user, bool) {
	for _, user := range users {
		if user.email == email {
			return user, true
		}
	}
	return user{}, false
}

//validate password
func (u *user) ValidatePassword(password string) bool {
	return u.password == password
}

// Add user
func AddUser(fullname string, email string, username string, password string, role string) bool {
	user := user{
		fullname: fullname,
		email:    email,
		username: username,
		password: password,
		role:     role,
	}
	for _, exist_user := range users {
		if exist_user.email == user.email || exist_user.username == user.username {
			return false
		}
	}
	users = append(users, user)
	return true
}
