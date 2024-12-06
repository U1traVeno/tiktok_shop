// Code generated by hertz generator.

package product

import (
	"context"
	"github.com/U1traVeno/tiktok-shop/biz/model/product"
	productservice "github.com/U1traVeno/tiktok-shop/biz/service"
	"github.com/U1traVeno/tiktok-shop/pkg/utils"
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
		resp := utils.BuildBaseResp(consts.StatusBadRequest, err, "400 Bad Request")
		c.JSON(consts.StatusBadRequest, product.ListProductsResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	resp, err := productservice.NewProductService(ctx, c).ListProducts(&req)
	if err != nil {
		resp := utils.BuildBaseResp(consts.StatusInternalServerError, err, "find products failed")
		c.JSON(consts.StatusInternalServerError, product.ListProductsResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, resp)

}

// GetProduct .
// @router /products/{id} [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.GetProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(consts.StatusBadRequest, err, "400 Bad Request")
		c.JSON(consts.StatusBadRequest, product.GetProductResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	resp, err := productservice.NewProductService(ctx, c).GetProduct(&req)
	if err != nil {
		resp := utils.BuildBaseResp(consts.StatusInternalServerError, err, "find product failed")
		c.JSON(consts.StatusInternalServerError, product.GetProductResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// SearchProducts .
// @router /products/search [GET]
func SearchProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(consts.StatusInternalServerError, err, "400 Bad Request")
		c.JSON(consts.StatusBadRequest, product.SearchProductsResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	resp, err := productservice.NewProductService(ctx, c).SearchProducts(&req)
	if err != nil {
		resp := utils.BuildBaseResp(consts.StatusInternalServerError, err, "find product failed")
		c.JSON(consts.StatusInternalServerError, product.SearchProductsResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// CreateProduct .
// @router /products [POST]
func CreateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(consts.StatusInternalServerError, err, "400 Bad Request")
		c.JSON(consts.StatusBadRequest, product.CreateProductResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	resp, err := productservice.NewProductService(ctx, c).CreateProduct(&req)
	if err != nil {
		resp := utils.BuildBaseResp(consts.StatusInternalServerError, err, "create product failed")
		c.JSON(consts.StatusInternalServerError, product.CreateProductResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// UpdateProduct .
// @router /products/{id} [PUT]
func UpdateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.UpdateProductReq
	err = c.BindAndValidate(&req)
	//用户权限
	if err != nil {
		resp := utils.BuildBaseResp(consts.StatusInternalServerError, err, "400 Bad Request")
		c.JSON(consts.StatusBadRequest, product.UpdateProductResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	resp, err := productservice.NewProductService(ctx, c).UpdateProduct(&req)
	if err != nil {
		resp := utils.BuildBaseResp(consts.StatusInternalServerError, err, "update product failed")
		c.JSON(consts.StatusInternalServerError, product.UpdateProductResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// DeleteProduct .
// @router /products/{id} [DELETE]
func DeleteProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.DeleteProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(consts.StatusInternalServerError, err, "400 Bad Request")
		c.JSON(consts.StatusBadRequest, product.DeleteProductResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	resp, err := productservice.NewProductService(ctx, c).DeleteProduct(&req)
	if err != nil {
		resp := utils.BuildBaseResp(consts.StatusInternalServerError, err, "delete prooduct failed")
		c.JSON(consts.StatusInternalServerError, product.DeleteProductResp{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}
