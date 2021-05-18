package service

import (
	"context"
	"gomod/mig_one/user"
)

var UserHandle *User

func init()  {
	UserHandle = new(User)
}

type User struct {

}

func (u *User) 	IsRiskUser(ctx context.Context, in *user.Param) (resp *user.Response, err error){
	resp = &user.Response{
		Val:   true,
	}
	return
}
