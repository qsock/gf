// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TimeZoneTransitionForDao is the golang structure of table time_zone_transition for DAO operations like Where/Data.
type TimeZoneTransitionForDao struct {
	g.Meta           `orm:"dto:true"`
	TimeZoneId       interface{} //
	TransitionTime   interface{} //
	TransitionTypeId interface{} //
}