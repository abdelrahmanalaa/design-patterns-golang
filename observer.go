package main

type User struct {
	name      string
	followers []*User
}

func (u *User) follow(following *User) {
	following.followers = append(following.followers, u)
}

func (u *User) post(text string) {
	for _, follower := range u.followers {
		follower.incomingFollowingPosts(text, u)
	}
}

func (u *User) incomingFollowingPosts(text string, followingUser *User) {
	println("hey " + u.name + " there is an incoming post from " + followingUser.name + ": " + text)
}

func createUser(name string) *User {
	return &User{name: name, followers: []*User{}}
}

func main() {
	abdelrahman := createUser("abdelrahman")
	sondos := createUser("sondos")

	abdelrahman.follow(sondos)
	sondos.post("my first post")
}
