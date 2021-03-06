// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// Shop is an object representing the database table.
type Shop struct {
	ID          int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	Shopname    string            `boil:"shopname" json:"shopname" toml:"shopname" yaml:"shopname"`
	Name        null.String       `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Desc        null.String       `boil:"desc" json:"desc,omitempty" toml:"desc" yaml:"desc,omitempty"`
	PrefectureN null.String       `boil:"prefecture_n" json:"prefecture_n,omitempty" toml:"prefecture_n" yaml:"prefecture_n,omitempty"`
	CityN       null.String       `boil:"city_n" json:"city_n,omitempty" toml:"city_n" yaml:"city_n,omitempty"`
	ZipCD       null.String       `boil:"zip_cd" json:"zip_cd,omitempty" toml:"zip_cd" yaml:"zip_cd,omitempty"`
	Tel         null.String       `boil:"tel" json:"tel,omitempty" toml:"tel" yaml:"tel,omitempty"`
	Email       null.String       `boil:"email" json:"email,omitempty" toml:"email" yaml:"email,omitempty"`
	HouseNumber null.String       `boil:"house_number" json:"house_number,omitempty" toml:"house_number" yaml:"house_number,omitempty"`
	MapLevel    null.String       `boil:"map_level" json:"map_level,omitempty" toml:"map_level" yaml:"map_level,omitempty"`
	AdID        int               `boil:"ad_id" json:"ad_id" toml:"ad_id" yaml:"ad_id"`
	Point       int               `boil:"point" json:"point" toml:"point" yaml:"point"`
	Lat         types.NullDecimal `boil:"lat" json:"lat,omitempty" toml:"lat" yaml:"lat,omitempty"`
	Lon         types.NullDecimal `boil:"lon" json:"lon,omitempty" toml:"lon" yaml:"lon,omitempty"`
	URL         null.String       `boil:"url" json:"url,omitempty" toml:"url" yaml:"url,omitempty"`
	CreatedAt   null.Time         `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt   null.Time         `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedNote null.String       `boil:"deleted_note" json:"deleted_note,omitempty" toml:"deleted_note" yaml:"deleted_note,omitempty"`
	DeletedFLG  bool              `boil:"deleted_flg" json:"deleted_flg" toml:"deleted_flg" yaml:"deleted_flg"`

	R *shopR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L shopL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ShopColumns = struct {
	ID          string
	Shopname    string
	Name        string
	Desc        string
	PrefectureN string
	CityN       string
	ZipCD       string
	Tel         string
	Email       string
	HouseNumber string
	MapLevel    string
	AdID        string
	Point       string
	Lat         string
	Lon         string
	URL         string
	CreatedAt   string
	UpdatedAt   string
	DeletedNote string
	DeletedFLG  string
}{
	ID:          "id",
	Shopname:    "shopname",
	Name:        "name",
	Desc:        "desc",
	PrefectureN: "prefecture_n",
	CityN:       "city_n",
	ZipCD:       "zip_cd",
	Tel:         "tel",
	Email:       "email",
	HouseNumber: "house_number",
	MapLevel:    "map_level",
	AdID:        "ad_id",
	Point:       "point",
	Lat:         "lat",
	Lon:         "lon",
	URL:         "url",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedNote: "deleted_note",
	DeletedFLG:  "deleted_flg",
}

// Generated where

type whereHelpertypes_NullDecimal struct{ field string }

func (w whereHelpertypes_NullDecimal) EQ(x types.NullDecimal) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_NullDecimal) NEQ(x types.NullDecimal) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_NullDecimal) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_NullDecimal) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpertypes_NullDecimal) LT(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_NullDecimal) LTE(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_NullDecimal) GT(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_NullDecimal) GTE(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ShopWhere = struct {
	ID          whereHelperint
	Shopname    whereHelperstring
	Name        whereHelpernull_String
	Desc        whereHelpernull_String
	PrefectureN whereHelpernull_String
	CityN       whereHelpernull_String
	ZipCD       whereHelpernull_String
	Tel         whereHelpernull_String
	Email       whereHelpernull_String
	HouseNumber whereHelpernull_String
	MapLevel    whereHelpernull_String
	AdID        whereHelperint
	Point       whereHelperint
	Lat         whereHelpertypes_NullDecimal
	Lon         whereHelpertypes_NullDecimal
	URL         whereHelpernull_String
	CreatedAt   whereHelpernull_Time
	UpdatedAt   whereHelpernull_Time
	DeletedNote whereHelpernull_String
	DeletedFLG  whereHelperbool
}{
	ID:          whereHelperint{field: "`shop`.`id`"},
	Shopname:    whereHelperstring{field: "`shop`.`shopname`"},
	Name:        whereHelpernull_String{field: "`shop`.`name`"},
	Desc:        whereHelpernull_String{field: "`shop`.`desc`"},
	PrefectureN: whereHelpernull_String{field: "`shop`.`prefecture_n`"},
	CityN:       whereHelpernull_String{field: "`shop`.`city_n`"},
	ZipCD:       whereHelpernull_String{field: "`shop`.`zip_cd`"},
	Tel:         whereHelpernull_String{field: "`shop`.`tel`"},
	Email:       whereHelpernull_String{field: "`shop`.`email`"},
	HouseNumber: whereHelpernull_String{field: "`shop`.`house_number`"},
	MapLevel:    whereHelpernull_String{field: "`shop`.`map_level`"},
	AdID:        whereHelperint{field: "`shop`.`ad_id`"},
	Point:       whereHelperint{field: "`shop`.`point`"},
	Lat:         whereHelpertypes_NullDecimal{field: "`shop`.`lat`"},
	Lon:         whereHelpertypes_NullDecimal{field: "`shop`.`lon`"},
	URL:         whereHelpernull_String{field: "`shop`.`url`"},
	CreatedAt:   whereHelpernull_Time{field: "`shop`.`created_at`"},
	UpdatedAt:   whereHelpernull_Time{field: "`shop`.`updated_at`"},
	DeletedNote: whereHelpernull_String{field: "`shop`.`deleted_note`"},
	DeletedFLG:  whereHelperbool{field: "`shop`.`deleted_flg`"},
}

// ShopRels is where relationship names are stored.
var ShopRels = struct {
}{}

// shopR is where relationships are stored.
type shopR struct {
}

// NewStruct creates a new relationship struct
func (*shopR) NewStruct() *shopR {
	return &shopR{}
}

// shopL is where Load methods for each relationship are stored.
type shopL struct{}

var (
	shopAllColumns            = []string{"id", "shopname", "name", "desc", "prefecture_n", "city_n", "zip_cd", "tel", "email", "house_number", "map_level", "ad_id", "point", "lat", "lon", "url", "created_at", "updated_at", "deleted_note", "deleted_flg"}
	shopColumnsWithoutDefault = []string{"shopname", "name", "desc", "prefecture_n", "city_n", "zip_cd", "tel", "email", "house_number", "map_level", "lat", "lon", "url", "created_at", "updated_at", "deleted_note"}
	shopColumnsWithDefault    = []string{"id", "ad_id", "point", "deleted_flg"}
	shopPrimaryKeyColumns     = []string{"id"}
)

type (
	// ShopSlice is an alias for a slice of pointers to Shop.
	// This should generally be used opposed to []Shop.
	ShopSlice []*Shop
	// ShopHook is the signature for custom Shop hook methods
	ShopHook func(context.Context, boil.ContextExecutor, *Shop) error

	shopQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	shopType                 = reflect.TypeOf(&Shop{})
	shopMapping              = queries.MakeStructMapping(shopType)
	shopPrimaryKeyMapping, _ = queries.BindMapping(shopType, shopMapping, shopPrimaryKeyColumns)
	shopInsertCacheMut       sync.RWMutex
	shopInsertCache          = make(map[string]insertCache)
	shopUpdateCacheMut       sync.RWMutex
	shopUpdateCache          = make(map[string]updateCache)
	shopUpsertCacheMut       sync.RWMutex
	shopUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var shopBeforeInsertHooks []ShopHook
var shopBeforeUpdateHooks []ShopHook
var shopBeforeDeleteHooks []ShopHook
var shopBeforeUpsertHooks []ShopHook

var shopAfterInsertHooks []ShopHook
var shopAfterSelectHooks []ShopHook
var shopAfterUpdateHooks []ShopHook
var shopAfterDeleteHooks []ShopHook
var shopAfterUpsertHooks []ShopHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Shop) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Shop) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Shop) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Shop) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Shop) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Shop) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Shop) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Shop) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Shop) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddShopHook registers your hook function for all future operations.
func AddShopHook(hookPoint boil.HookPoint, shopHook ShopHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		shopBeforeInsertHooks = append(shopBeforeInsertHooks, shopHook)
	case boil.BeforeUpdateHook:
		shopBeforeUpdateHooks = append(shopBeforeUpdateHooks, shopHook)
	case boil.BeforeDeleteHook:
		shopBeforeDeleteHooks = append(shopBeforeDeleteHooks, shopHook)
	case boil.BeforeUpsertHook:
		shopBeforeUpsertHooks = append(shopBeforeUpsertHooks, shopHook)
	case boil.AfterInsertHook:
		shopAfterInsertHooks = append(shopAfterInsertHooks, shopHook)
	case boil.AfterSelectHook:
		shopAfterSelectHooks = append(shopAfterSelectHooks, shopHook)
	case boil.AfterUpdateHook:
		shopAfterUpdateHooks = append(shopAfterUpdateHooks, shopHook)
	case boil.AfterDeleteHook:
		shopAfterDeleteHooks = append(shopAfterDeleteHooks, shopHook)
	case boil.AfterUpsertHook:
		shopAfterUpsertHooks = append(shopAfterUpsertHooks, shopHook)
	}
}

// One returns a single shop record from the query.
func (q shopQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Shop, error) {
	o := &Shop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for shop")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Shop records from the query.
func (q shopQuery) All(ctx context.Context, exec boil.ContextExecutor) (ShopSlice, error) {
	var o []*Shop

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Shop slice")
	}

	if len(shopAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Shop records in the query.
func (q shopQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count shop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q shopQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if shop exists")
	}

	return count > 0, nil
}

// Shops retrieves all the records using an executor.
func Shops(mods ...qm.QueryMod) shopQuery {
	mods = append(mods, qm.From("`shop`"))
	return shopQuery{NewQuery(mods...)}
}

// FindShop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindShop(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Shop, error) {
	shopObj := &Shop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `shop` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, shopObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from shop")
	}

	return shopObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Shop) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no shop provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shopColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	shopInsertCacheMut.RLock()
	cache, cached := shopInsertCache[key]
	shopInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			shopAllColumns,
			shopColumnsWithDefault,
			shopColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(shopType, shopMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(shopType, shopMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `shop` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `shop` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `shop` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, shopPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into shop")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == shopMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for shop")
	}

CacheNoHooks:
	if !cached {
		shopInsertCacheMut.Lock()
		shopInsertCache[key] = cache
		shopInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Shop.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Shop) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	shopUpdateCacheMut.RLock()
	cache, cached := shopUpdateCache[key]
	shopUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			shopAllColumns,
			shopPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update shop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `shop` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, shopPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(shopType, shopMapping, append(wl, shopPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update shop row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for shop")
	}

	if !cached {
		shopUpdateCacheMut.Lock()
		shopUpdateCache[key] = cache
		shopUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q shopQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for shop")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for shop")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ShopSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shopPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `shop` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, shopPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in shop slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all shop")
	}
	return rowsAff, nil
}

var mySQLShopUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Shop) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no shop provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shopColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLShopUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	shopUpsertCacheMut.RLock()
	cache, cached := shopUpsertCache[key]
	shopUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			shopAllColumns,
			shopColumnsWithDefault,
			shopColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			shopAllColumns,
			shopPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert shop, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "shop", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `shop` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(shopType, shopMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(shopType, shopMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for shop")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == shopMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(shopType, shopMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for shop")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for shop")
	}

CacheNoHooks:
	if !cached {
		shopUpsertCacheMut.Lock()
		shopUpsertCache[key] = cache
		shopUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Shop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Shop) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Shop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), shopPrimaryKeyMapping)
	sql := "DELETE FROM `shop` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from shop")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for shop")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q shopQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no shopQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from shop")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for shop")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ShopSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(shopBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shopPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `shop` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, shopPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from shop slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for shop")
	}

	if len(shopAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Shop) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindShop(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ShopSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ShopSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shopPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `shop`.* FROM `shop` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, shopPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ShopSlice")
	}

	*o = slice

	return nil
}

// ShopExists checks if the Shop row exists.
func ShopExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `shop` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if shop exists")
	}

	return exists, nil
}
