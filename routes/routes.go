package routes

import (
	"clms/controllers"
	"clms/logger"
	"clms/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{

		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		//用户操作
		v1.POST("user/register", controllers.UserRegister)
		v1.POST("user/login", controllers.UserLogin)

		//商品操作
		v1.GET("products", controllers.ListProduct)
		v1.GET("product/:id", controllers.ListProductById)
		v1.POST("products", controllers.SearchProduct)
		v1.GET("imgs/:id", controllers.ListProductImg) //商品图片
		v1.GET("categories", controllers.ListCategory) //商品分类
		v1.GET("carousels", controllers.ListCarousels) //轮播图
		r.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
		v1.Use(middlewares.JWTAuthMiddleware())
		{

			// 用户操作
			v1.PUT("user", controllers.UserUpdate)             //改名,发邮箱
			v1.POST("user/valid-email", controllers.UserValid) //接收邮箱，并进行修改
			v1.POST("avatar", controllers.UpLoadAvatar)        //上传头像

			// 商品操作
			v1.POST("product", controllers.AddProduct)
			v1.PUT("product/:id", controllers.UpdateProduct)
			v1.DELETE("product/:id", controllers.DeleteProduct)

			// 收藏夹
			v1.GET("favorites", controllers.ShowFavorite)
			v1.POST("favorites", controllers.AddFavorite)
			v1.DELETE("favorites/:id", controllers.DeleteFavorite)

			// 订单操作
			v1.POST("orders", controllers.AddOrders)
			v1.GET("orders", controllers.GetOrders)
			v1.GET("orders/:id", controllers.GetOrderById)
			v1.DELETE("orders/:id", controllers.DeleteOrder)

			//购物车
			v1.POST("carts", controllers.AddCarts)
			v1.GET("carts/:id", controllers.GetCarts)    // 购物车id
			v1.PUT("carts/:id", controllers.UpdateCarts) // 购物车id
			v1.DELETE("carts/:id", controllers.DeleteCarts)

			//收获地址操作
			v1.POST("addresses")
			v1.GET("addresses/:id")
			v1.GET("addresses")
			v1.PUT("addresses/:id")
			v1.DELETE("addresses/:id")

			// 支付功能
			v1.POST("paydown/:id", controllers.OrderPay)

			// 显示金额
			v1.POST("money")
		}
		return r
	}
}
