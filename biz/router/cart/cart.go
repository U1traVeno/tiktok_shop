// Code generated by hertz generator. DO NOT EDIT.

package cart

import (
	cart "github.com/U1traVeno/tiktok-shop/biz/handler/cart"
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
	root.POST("/cart", append(_createcartMw(), cart.CreateCart)...)
	_cart := root.Group("/cart", _cartMw()...)
	_cart.GET("/get", append(_getcartMw(), cart.GetCart)...)
	_cart.POST("/item", append(_additemMw(), cart.AddItem)...)
	{
		_empty := _cart.Group("/empty", _emptyMw()...)
		_empty.DELETE("/{user_id}", append(_emptycartMw(), cart.EmptyCart)...)
	}
}