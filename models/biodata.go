package models

type Biodata struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Full_Name    string `json:"fullName"`
	Age          string `json:"age"`
	Address      string `json:"address"`
	Phone_Number string `json:"phoneNumber"`
	Email        string `json:"email"`
	Time_Slot    string `json:"timeSlot"`
}
