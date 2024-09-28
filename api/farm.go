package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/tokha04/swe-farmer-market-system/db/sqlc"
)

type CreateFarmRequest struct {
	Address      string  `json:"address"`
	Size         float64 `json:"size"`
	GovernmentID int32   `json:"government_id"`
}

func CreateFarm(q *db.Queries) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateFarmRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		authPayload := ctx.MustGet("authorization_payload").(*Payload)

		arg := db.CreateFarmParams{
			FarmerID:     pgtype.Int4{Int32: authPayload.UserID, Valid: true},
			Address:      pgtype.Text{String: req.Address, Valid: true},
			Size:         pgtype.Float8{Float64: req.Size, Valid: true},
			GovernmentID: pgtype.Int4{Int32: req.GovernmentID, Valid: true},
		}

		farm, err := q.CreateFarm(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, farm)
	}
}

type GetFarmRequest struct {
	ID int32 `uri:"id"`
}

func GetFarm(q *db.Queries) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetFarmRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		farm, err := q.GetFarm(ctx, pgtype.Int4{Int32: req.ID, Valid: true})
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		authPayload := ctx.MustGet("authorization_payload").(*Payload)
		if farm.FarmerID.Int32 != authPayload.UserID {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "farm does not belong to the authenticated user"})
			return
		}

		ctx.JSON(http.StatusOK, farm)
	}
}

type ListFarmsRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func ListFarms(q *db.Queries) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ListFarmsRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		authPayload := ctx.MustGet("authorization_payload").(*Payload)

		arg := db.ListFarmsParams{
			FarmerID: pgtype.Int4{Int32: authPayload.UserID, Valid: true},
			Limit:    req.Limit,
			Offset:   req.Offset,
		}

		farms, err := q.ListFarms(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, farms)
	}
}

func UpdateFarm(q *db.Queries) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetFarmRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		farm, err := q.GetFarm(ctx, pgtype.Int4{Int32: req.ID, Valid: true})
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		authPayload := ctx.MustGet("authorization_payload").(*Payload)
		if farm.FarmerID.Int32 != authPayload.UserID {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "farm does not belong to the authenticated user"})
			return
		}

		var updReq CreateFarmRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		arg := db.UpdateFarmParams{
			ID:      req.ID,
			Column2: updReq.Address,
			Column3: updReq.Size,
			Column4: updReq.GovernmentID,
		}

		updFarm, err := q.UpdateFarm(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, updFarm)
	}
}

func DeleteFarm(q *db.Queries) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetFarmRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		farm, err := q.GetFarm(ctx, pgtype.Int4{Int32: req.ID, Valid: true})
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		authPayload := ctx.MustGet("authorization_payload").(*Payload)
		if farm.FarmerID.Int32 != authPayload.UserID {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "farm does not belong to the authenticated user"})
			return
		}

		err = q.DeleteFarm(ctx, req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"msg": "successfully deleted"})
	}
}
