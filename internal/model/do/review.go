// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Review is the golang structure of table review for DAO operations like Where/Data.
type Review struct {
	g.Meta        `orm:"table:review, do:true"`
	Id            interface{} //
	CustomerId    interface{} //
	ReviewText    interface{} //
	Created       interface{} //
	TransactionId interface{} //
}
