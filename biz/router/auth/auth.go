// Code generated by hertz generator. DO NOT EDIT.

package auth

import (
	auth "github.com/U1traVeno/tiktok-shop/biz/handler/auth"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_auth := root.Group("/auth", _authMw()...)
		_auth.GET("/deliver", append(_delivertokenbyrpcMw(), auth.DeliverTokenByRPC)...)
		_auth.POST("/verify", append(_verifytokenbyrpcMw(), auth.VerifyTokenByRPC)...)
	}
}
