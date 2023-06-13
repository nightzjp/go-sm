/*
   @TIME: 2023/6/12 14:14
   @Author: Nightz
   @Site:
   @File: base.go
   @SoftWare: GoLand
*/

package constant

import "gorm.io/gorm"

type ByteSize float64

var DB *gorm.DB

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

const JwtSecret = "JwtSecret"

/*

 */
