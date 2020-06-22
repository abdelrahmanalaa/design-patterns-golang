package main

type User struct {
	name      string
	followers []*User
}

func (u *User) follow(following *User) {
	following.followers = append(following.followers, u)
}

func (u *User) post(text string) {
	println("hello from " + u.name + " feed, post: " + text)
	for _, follower := range u.followers {
		follower.incomingFollowingPosts(text, u)
	}
}

func (u *User) incomingFollowingPosts(text string, followingUser *User) {
	println("hello from " + u.name + " feed, " + followingUser.name + "'s post: " + text)
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
