package controller

import (
	"context"
	"net/http"
	"server/database"
	"server/models"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func PostStatusKendaraan(ctx *gin.Context) {
	db := database.GetDB()
	VehicleStatus := models.MStatusKendaraan{}

	newId := uuid.New()

	ctx.ShouldBindJSON(&VehicleStatus)

	VehicleStatus.IdStatusKendaraan = newId

	if err := db.Create(&VehicleStatus).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":     VehicleStatus.StatusKendaraan,
		"keterangan": VehicleStatus.Keterangan,
		"message":    "Status Kendaraan Berhasil Ditambahkan",
	})
}

func UpdateStatusKendaraan(ctx *gin.Context) {
	db := database.GetDB()
	var VehicleStatus = models.MStatusKendaraan{}
	var UpdateVehicleStatus = models.MStatusKendaraan{}
	id := ctx.Param("id")

	ctx.ShouldBindJSON(&UpdateVehicleStatus)

	if err := crdbgorm.ExecuteTx(context.Background(), db, nil, func(tx *gorm.DB) error {
		if err := tx.Where("id_status_kendaraan = ?", id).First(&VehicleStatus).Error; err != nil {
			return err
		}

		if err := tx.Where("id_status_kendaraan = ?", VehicleStatus.IdStatusKendaraan).Updates(&UpdateVehicleStatus).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data status kendaraan berhasil diperbarui",
	})
}

func GetAllStatusKendaraan(ctx *gin.Context) {
	db := database.GetDB()
	var ListStatusKendaraan []models.MStatusKendaraan

	if err := db.Find(&ListStatusKendaraan).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, ListStatusKendaraan)
}