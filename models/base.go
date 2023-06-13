/*
   @TIME: 2023/6/12 17:20
   @Author: Nightz
   @Site:
   @File: base.go
   @SoftWare: GoLand
*/

package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at" time_format:"sql_datetime" time_utc:"false"`
	UpdatedAt time.Time      `json:"updated_at" time_format:"sql_datetime" time_utc:"false"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

/*

 */
