package http

import (
	"boobot/dal/repo"
)

type Opts struct {
	UserRepo    repo.User
	SessionRepo repo.Session
}
