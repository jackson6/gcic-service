package user

import (
	"github.com/captaincodeman/couponcode"
	"github.com/jinzhu/gorm"
	"math/rand"
	"strconv"
)

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	code := couponcode.Generate()
	memberId := strconv.Itoa(rangeIn(10000000, 99999999))
	err := scope.SetColumn("ReferralCode", code)
	if err != nil {
		return err
	}
	err = scope.SetColumn("MemberId", memberId)
	if err != nil {
		return err
	}
	return nil
}
