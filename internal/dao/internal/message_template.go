// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MessageTemplateDao is the data access object for table message_template.
type MessageTemplateDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns MessageTemplateColumns // columns contains all the column names of Table for convenient usage.
}

// MessageTemplateColumns defines and stores column names for table message_template.
type MessageTemplateColumns struct {
	TemplateId      string //
	TemplateName    string //
	TemplateContent string //
	Variables       string //
}

// messageTemplateColumns holds the columns for table message_template.
var messageTemplateColumns = MessageTemplateColumns{
	TemplateId:      "template_id",
	TemplateName:    "template_name",
	TemplateContent: "template_content",
	Variables:       "variables",
}

// NewMessageTemplateDao creates and returns a new DAO object for table data access.
func NewMessageTemplateDao() *MessageTemplateDao {
	return &MessageTemplateDao{
		group:   "default",
		table:   "message_template",
		columns: messageTemplateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MessageTemplateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MessageTemplateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MessageTemplateDao) Columns() MessageTemplateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MessageTemplateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MessageTemplateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MessageTemplateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
