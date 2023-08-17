// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MessageTemplate is the golang structure of table message_template for DAO operations like Where/Data.
type MessageTemplate struct {
	g.Meta          `orm:"table:message_template, do:true"`
	TemplateId      interface{} //
	TemplateName    interface{} //
	TemplateContent interface{} //
	Variables       interface{} //
}
