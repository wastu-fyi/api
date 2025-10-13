package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Student struct {
	orm.Model
	Nim               string     `gorm:"column:nim;size:16" json:"nim"`
	Name              string     `gorm:"column:name" json:"name"`
	Gender            uint8      `gorm:"column:gender" json:"gender"`
	Nik               *string    `gorm:"column:nik;size:16" json:"nik,omitempty"`
	Pob               *string    `gorm:"column:pob" json:"pob,omitempty"`
	Dob               *time.Time `gorm:"column:dob;type:date" json:"dob,omitempty"`
	Religion          *string    `gorm:"column:religion" json:"religion,omitempty"`
	Phone             *string    `gorm:"column:phone" json:"phone,omitempty"`
	Email             *string    `gorm:"column:email" json:"email,omitempty"`
	FatherName        *string    `gorm:"column:father_name" json:"father_name,omitempty"`
	FatherDob         *time.Time `gorm:"column:father_dob;type:date" json:"father_dob,omitempty"`
	FatherJob         *string    `gorm:"column:father_job" json:"father_job,omitempty"`
	FatherEducation   *string    `gorm:"column:father_education" json:"father_education,omitempty"`
	MotherName        *string    `gorm:"column:mother_name" json:"mother_name,omitempty"`
	MotherDob         *time.Time `gorm:"column:mother_dob;type:date" json:"mother_dob,omitempty"`
	MotherJob         *string    `gorm:"column:mother_job" json:"mother_job,omitempty"`
	MotherEducation   *string    `gorm:"column:mother_education" json:"mother_education,omitempty"`
	GuardianName      *string    `gorm:"column:guardian_name" json:"guardian_name,omitempty"`
	GuardianDob       *time.Time `gorm:"column:guardian_dob;type:date" json:"guardian_dob,omitempty"`
	GuardianJob       *string    `gorm:"column:guardian_job" json:"guardian_job,omitempty"`
	GuardianEducation *string    `gorm:"column:guardian_education" json:"guardian_education,omitempty"`
	ResidenceType     *string    `gorm:"column:residence_type" json:"residence_type,omitempty"`
	Transportation    *string    `gorm:"column:transportation" json:"transportation,omitempty"`
	Study             int        `gorm:"column:study" json:"study"`
	Status            *string    `gorm:"column:status" json:"status,omitempty"`
	AdmissionYear     int        `gorm:"column:admission_year;type:year" json:"admission_year"`
	ClassName         *string    `gorm:"column:class_name" json:"class_name,omitempty"`
	NameTitle         string     `gorm:"column:name_title" json:"name_title"`
	NameKey           string     `gorm:"column:name_key" json:"name_key"`
	NimSuffix         uint8      `gorm:"column:nim_suffix" json:"nim_suffix"`
	Mobile            *string    `gorm:"column:mobile" json:"mobile,omitempty"`
	Address           *string    `gorm:"column:address" json:"address,omitempty"`
	StreetName        *string    `gorm:"column:street_name" json:"street_name,omitempty"`
	Country           *string    `gorm:"column:country" json:"country,omitempty"`
	Province          *string    `gorm:"column:province" json:"province,omitempty"`
	City              *string    `gorm:"column:city" json:"city,omitempty"`
	Region            *string    `gorm:"column:region" json:"region,omitempty"`
	Village           *string    `gorm:"column:village" json:"village,omitempty"`
	FatherIncome      *string    `gorm:"column:father_income" json:"father_income,omitempty"`
	MotherIncome      *string    `gorm:"column:mother_income" json:"mother_income,omitempty"`
	GuardianIncome    *string    `gorm:"column:guardian_income" json:"guardian_income,omitempty"`

	User *User `gorm:"foreignKey:StudentID;references:ID" json:"user,omitempty"`
}

func (Student) TableName() string {
	return "students"
}
