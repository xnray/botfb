// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReviewDao is the data access object for table review.
type ReviewDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ReviewColumns // columns contains all the column names of Table for convenient usage.
}

// ReviewColumns defines and stores column names for table review.
type ReviewColumns struct {
	Id            string //
	CustomerId    string //
	ReviewText    string //
	Created       string //
	TransactionId string //
}

// reviewColumns holds the columns for table review.
var reviewColumns = ReviewColumns{
	Id:            "id",
	CustomerId:    "customer_id",
	ReviewText:    "review_text",
	Created:       "created",
	TransactionId: "transaction_id",
}

// NewReviewDao creates and returns a new DAO object for table data access.
func NewReviewDao() *ReviewDao {
	return &ReviewDao{
		group:   "default",
		table:   "review",
		columns: reviewColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReviewDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReviewDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReviewDao) Columns() ReviewColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReviewDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReviewDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReviewDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
