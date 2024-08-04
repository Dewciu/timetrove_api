package users

import (
	_ "github.com/dewciu/timetrove_api/docs"
	"github.com/dewciu/timetrove_api/pkg/middleware"
	perm "github.com/dewciu/timetrove_api/pkg/permissions"
	"github.com/gin-gonic/gin"
)

const (
	UsersEndpoint       = "/users"
	AuthEndpoint        = "/auth"
	LoginEndpoint       = "/login"
	PermissionsEndpoint = "/permissions"
)

var UsersPermissions = []perm.PermissionModel{}

func AddUsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group(UsersEndpoint)
	{
		users.GET("/", middleware.AuthMiddleware(), GetAllUsersController)
		users.POST("/", middleware.AuthMiddleware(), CreateUserController)
		users.GET("/:id", middleware.AuthMiddleware(), GetUserByIDController)
		users.DELETE("/:id", middleware.AuthMiddleware(), DeleteUserByIDController)
		users.PUT("/:id", middleware.AuthMiddleware(), UpdateUserController)
		users.GET("/:id"+PermissionsEndpoint, middleware.AuthMiddleware(), GetUserWithPermissionsController)
	}
}

func GetUserPermissions() []perm.PermissionModel {
	return []perm.PermissionModel{
		{
			Endpoint: UsersEndpoint,
			Method:   "GET",
		},
		{
			Endpoint: UsersEndpoint,
			Method:   "POST",
		},
		{
			Endpoint: UsersEndpoint,
			Method:   "PUT",
		},
		{
			Endpoint: UsersEndpoint,
			Method:   "DELETE",
		},
		{
			Endpoint: UsersEndpoint + "/:id",
			Method:   "GET",
		},
		{
			Endpoint: UsersEndpoint + "/:id",
			Method:   "DELETE",
		},
		{
			Endpoint: UsersEndpoint + "/:id",
			Method:   "PUT",
		},
	}
}

func AddAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group(AuthEndpoint)
	{
		auth.POST(LoginEndpoint, LoginController)
	}
}

func GetAuthPermissions() []perm.PermissionModel {
	return []perm.PermissionModel{
		{
			Endpoint: LoginEndpoint,
			Method:   "POST",
		},
	}
}
