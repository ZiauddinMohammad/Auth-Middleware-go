package data

//user struct
type User struct {
	Fullname   string
	Email      string
	Username   string
	Password   string
	Createdate string
	Role       string
}

// user data
var users = []User{
	{
		Fullname:   "userone",
		Email:      "userone@abc.com",
		Username:   "user1",
		Password:   "pass1",
		Createdate: "1780506070",
		Role:       "admin",
	},
	{
		Fullname:   "usertwo",
		Email:      "usertwo@abc.com",
		Username:   "user2",
		Password:   "pass2",
		Createdate: "1780506070",
		Role:       "user"},
}

//returns a user object based on user email
func GetUser(email string) (User, bool) {
	//loop thru users data
	for _, user := range users {
		if user.Email == email {
			return user, true
		}
	}
	return User{}, false
}

//validate password
func (u *User) ValidatePassword(password string) bool {
	return u.Password == password
}

// Add user
func AddUser(fullname string, email string, username string, password string, role string) bool {
	//check if email or username already exists
	for _, exist_user := range users {
		if exist_user.Email == email || exist_user.Username == username {
			return false
		}
	}
	//Create and append new user to users data list
	user := User{
		Fullname: fullname,
		Email:    email,
		Username: username,
		Password: password,
		Role:     role,
	}
	users = append(users, user)
	return true
}
