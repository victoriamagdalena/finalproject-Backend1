package controllers

import (
	"finalProject/models"

	"github.com/gin-gonic/gin"
)

func GetBiodatas(context *gin.Context) {
	var biodatas []models.Biodata
	models.DB.Find(&biodatas)

	context.JSON(200, gin.H{"success": true, "result": biodatas, "errorMessage": nil})
}

func PostBiodata(context *gin.Context) {
	type postBiodata struct {
		ID           uint   `json:"id" gorm:"primary_key"`
		Full_Name    string `json:"fullName"`
		Age          string `json:"age"`
		Address      string `json:"address"`
		Phone_Number string `json:"phoneNumber"`
		Email        string `json:"email"`
		Time_Slot    string `json:"timeSlot"`
	}

	var input postBiodata
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(500, gin.H{"success": false, "result": nil, "errorMessage": err.Error()})
		return
	}

	// count(input.Time_Slot)
	biodata := models.Biodata{Full_Name: input.Full_Name, Age: input.Age, Address: input.Address, Phone_Number: input.Phone_Number, Email: input.Email, Time_Slot: input.Time_Slot}
	models.DB.Create(&biodata)

	context.JSON(201, gin.H{"success": true, "result": biodata, "errorMessage": nil})

}

// func count(context *gin.Context) {
// 	count1 = db.Model(&Biodata{}).Where("time_slot = ?", "07.15").Count(&count)
// 	count2 = db.Model(&Biodata{}).Where("time_slot = ?", "09.00").Count(&count)
// 	count3 = db.Model(&Biodata{}).Where("time_slot = ?", "13.00").Count(&count)
// 	count4 = db.Model(&Biodata{}).Where("time_slot = ?", "16.00").Count(&count)

// 	if count1 > 5 {
// 		context.JSON(201, gin.H{"success": true, "result": "Maaf, kuota jemaat untuk sesi ibadah jam 07.15 sudah penuh.", "errorMessage": nil})
// 		return
// 	}

// 	if count2 > 50 {
// 		context.JSON(201, gin.H{"success": true, "result": "Maaf, kuota jemaat untuk sesi ibadah jam 09.00 sudah penuh.", "errorMessage": nil})
// 		return
// 	}

// 	if count3 > 50 {
// 		context.JSON(201, gin.H{"success": true, "result": "Maaf, kuota jemaat untuk sesi ibadah jam 13.00 sudah penuh.", "errorMessage": nil})
// 		return
// 	}

// 	if count4 > 50 {
// 		context.JSON(201, gin.H{"success": true, "result": "Maaf, kuota jemaat untuk sesi ibadah jam 16.00 sudah penuh.", "errorMessage": nil})
// 		return
// 	}
// }
