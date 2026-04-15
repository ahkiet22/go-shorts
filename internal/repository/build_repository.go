package repository

import (
	"fmt"
	"strings"
)

// Builder pattern
type QueryBuilder struct {
	table   string
	wheres  []string
	selects []string
	limit   int
	order   string
}

func NewQueryBuilder(table string) *QueryBuilder {
	return &QueryBuilder{table: table}
}

func (qd *QueryBuilder) Select(cols ...string) *QueryBuilder {
	qd.selects = append(qd.selects, cols...)

	return qd
}

func (qb *QueryBuilder) Where(w string) *QueryBuilder {
	qb.wheres = append(qb.wheres, w)
	return qb
}

func (qb *QueryBuilder) Limit(n int) *QueryBuilder {
	qb.limit = n
	return qb
}

func (qb *QueryBuilder) Order(o string) *QueryBuilder {
	qb.order = o
	return qb
}

func (qb *QueryBuilder) Build() string {

	selectDefault := "*"

	if len(qb.selects) > 0 {
		selectDefault = strings.Join(qb.selects, ",")
	}

	q := "SELECT " + selectDefault + "FROM " + qb.table

	if len(qb.wheres) > 0 {
		q += "WHERE" + strings.Join(qb.wheres, "AND")
	}

	if qb.order != "" {
		q += "ORDER BY" + qb.order
	}

	if qb.limit > 0 {
		q += fmt.Sprintf(" LIMIT %d", qb.limit)
	}

	return q
}
