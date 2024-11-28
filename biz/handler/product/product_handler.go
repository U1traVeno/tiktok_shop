// Code generated by hertz generator.

package product

import (
	"context"

	product "github.com/U1traVeno/tiktok-shop/biz/model/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListProducts .
// @router /products [GET]
func ListProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ListProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(product.ListProductsResp)

	c.JSON(consts.StatusOK, resp)
}

// GetProduct .
// @router /products/{id} [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.GetProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(product.GetProductResp)

	c.JSON(consts.StatusOK, resp)
}

// SearchProducts .
// @router /products/search [GET]
func SearchProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(product.SearchProductsResp)

	c.JSON(consts.StatusOK, resp)
}

// CreateProduct .
// @router /products [POST]
func CreateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(product.CreateProductResp)

	c.JSON(consts.StatusOK, resp)
}

// UpdateProduct .
// @router /products/{id} [PUT]
func UpdateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.UpdateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(product.UpdateProductResp)

	c.JSON(consts.StatusOK, resp)
}

// DeleteProduct .
// @router /products/{id} [DELETE]
func DeleteProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.DeleteProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(product.DeleteProductResp)

	c.JSON(consts.StatusOK, resp)
}