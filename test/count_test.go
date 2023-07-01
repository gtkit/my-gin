package test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"

	"github.com/stretchr/testify/assert"

	"godata/db"
)

func TestCountShop(t *testing.T) {
	active := sq.Select("count(*)").From(db.ShopTableName()).Where(sq.Eq{"source": 4})
	ql, args, err := active.ToSql()
	assert.NoError(t, err)

	expectedSql := "SELECT count(*) FROM assistant_new_shops WHERE source = ?"
	assert.Equal(t, ql, expectedSql)

	expectedArgs := []interface{}{4}
	assert.Equal(t, expectedArgs, args)
}

func TestNumShop(t *testing.T) {
	subQ := sq.Select("COUNT(DISTINCT s2.down_amount)").From(db.ShopTableName() + " as s2").Where("s2.down_amount > s1.down_amount").Where("s2.type=s1.type").Where(sq.Eq{"s2.source": 4})
	sql, args, _ := subQ.ToSql()
	args = append(args, 5)

	query := sq.Select("s1.id", "s1.shop_title", "s1.`type`", "s1.down_amount").
		From(db.ShopTableName()+" as s1").
		Where("("+sql+")< ?", args...).
		Where(sq.Eq{"s1.source": 4}).
		OrderBy("s1.type, s1.down_amount DESC")
	sql2, _, err := query.ToSql()
	assert.NoError(t, err)

	expectedSql := "SELECT s1.id, s1.shop_title, s1.`type`, s1.down_amount FROM assistant_new_shops as s1 WHERE (SELECT COUNT(DISTINCT s2.down_amount) FROM assistant_new_shops as s2 WHERE s2.down_amount > s1.down_amount AND s2.type=s1.type AND s2.source = ?)< ? AND s1.source = ? ORDER BY s1.type, s1.down_amount DESC"
	assert.Equal(t, sql2, expectedSql)
	// expectedArgs := []interface{}{4}
	// assert.Equal(t, expectedArgs, args)
}

func ExampleSelectBuilder_Columns() {
	subQ := sq.Select("COUNT(DISTINCT s2.down_amount)").From(db.ShopTableName() + " as s2").Where("s2.down_amount > s1.down_amount").Where("s2.type=s1.type").Where(sq.Eq{"s2.source": 4})
	sql, args, _ := subQ.ToSql()
	args = append(args, 5)

	query := sq.Select("s1.id", "s1.shop_title", "s1.`type`", "s1.down_amount").
		From(db.ShopTableName()+" as s1").
		Where("("+sql+")< ?", args...).
		Where(sq.Eq{"s1.source": 4}).
		OrderBy("s1.type, s1.down_amount DESC")
	sql2, _, _ := query.ToSql()
	fmt.Println(sql2)
	// Output: SELECT s1.id, s1.shop_title, s1.`type`, s1.down_amount FROM assistant_new_shops as s1 WHERE (SELECT COUNT(DISTINCT s2.down_amount) FROM assistant_new_shops as s2 WHERE s2.down_amount > s1.down_amount AND s2.type=s1.type AND s2.source = ?)< ? AND s1.source = ? ORDER BY s1.type, s1.down_amount DESC
}
