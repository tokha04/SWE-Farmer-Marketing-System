package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/tokha04/swe-farmer-market-system/db/sqlc"
	"github.com/tokha04/swe-farmer-market-system/util"
)

type SignupRequest struct {
	Name        string `json:"name" binding:"required,alphanum"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	PhoneNumber string `json:"phone_number"`
}

type UserResponse struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	IsAdmin     bool   `json:"id_admin"`
}

func userResponse(user db.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		IsAdmin:     user.IsAdmin,
	}
}

func Signup(q *db.Queries) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req SignupRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := util.HashPassword(req.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		arg := db.CreateUserParams{
			Name:           req.Name,
			Email:          req.Email,
			HashedPassword: hashedPassword,
			PhoneNumber:    req.PhoneNumber,
		}

		user, err := q.CreateUser(ctx, arg)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "unique_violation":
					ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
					return
				}
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, userResponse(user))
	}
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken string
	User        UserResponse
}

func Login(q *db.Queries) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req LoginUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := q.GetUserByEmail(ctx, req.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = util.CheckPassword(req.Password, user.HashedPassword)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		accessToken, err := GenerateToken(user.ID, user.IsAdmin)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, LoginUserResponse{
			AccessToken: accessToken,
			User:        userResponse(user),
		})
	}
}
