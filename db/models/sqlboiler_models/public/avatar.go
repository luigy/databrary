// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package public

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
)

// Avatar is an object representing the database table.
type Avatar struct {
	Party int `db:"party" json:"avatar_party"`
	Asset int `db:"asset" json:"avatar_asset"`

	R *avatarR `db:"-" json:"-"`
	L avatarL  `db:"-" json:"-"`
}

// avatarR is where relationships are stored.
type avatarR struct {
	Asset *Asset
	Party *Party
}

// avatarL is where Load methods for each relationship are stored.
type avatarL struct{}

var (
	avatarColumns               = []string{"party", "asset"}
	avatarColumnsWithoutDefault = []string{"party", "asset"}
	avatarColumnsWithDefault    = []string{}
	avatarColumnsWithCustom     = []string{}

	avatarPrimaryKeyColumns = []string{"party"}
)

type (
	// AvatarSlice is an alias for a slice of pointers to Avatar.
	// This should generally be used opposed to []Avatar.
	AvatarSlice []*Avatar
	// AvatarHook is the signature for custom Avatar hook methods
	AvatarHook func(boil.Executor, *Avatar) error

	avatarQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	avatarType    = reflect.TypeOf(&Avatar{})
	avatarMapping = queries.MakeStructMapping(avatarType)

	avatarPrimaryKeyMapping, _ = queries.BindMapping(avatarType, avatarMapping, avatarPrimaryKeyColumns)

	avatarInsertCacheMut sync.RWMutex
	avatarInsertCache    = make(map[string]insertCache)
	avatarUpdateCacheMut sync.RWMutex
	avatarUpdateCache    = make(map[string]updateCache)
	avatarUpsertCacheMut sync.RWMutex
	avatarUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var avatarBeforeInsertHooks []AvatarHook
var avatarBeforeUpdateHooks []AvatarHook
var avatarBeforeDeleteHooks []AvatarHook
var avatarBeforeUpsertHooks []AvatarHook

var avatarAfterInsertHooks []AvatarHook
var avatarAfterSelectHooks []AvatarHook
var avatarAfterUpdateHooks []AvatarHook
var avatarAfterDeleteHooks []AvatarHook
var avatarAfterUpsertHooks []AvatarHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Avatar) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range avatarBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Avatar) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range avatarBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Avatar) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range avatarBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Avatar) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range avatarBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Avatar) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range avatarAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Avatar) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range avatarAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Avatar) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range avatarAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Avatar) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range avatarAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Avatar) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range avatarAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAvatarHook registers your hook function for all future operations.
func AddAvatarHook(hookPoint boil.HookPoint, avatarHook AvatarHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		avatarBeforeInsertHooks = append(avatarBeforeInsertHooks, avatarHook)
	case boil.BeforeUpdateHook:
		avatarBeforeUpdateHooks = append(avatarBeforeUpdateHooks, avatarHook)
	case boil.BeforeDeleteHook:
		avatarBeforeDeleteHooks = append(avatarBeforeDeleteHooks, avatarHook)
	case boil.BeforeUpsertHook:
		avatarBeforeUpsertHooks = append(avatarBeforeUpsertHooks, avatarHook)
	case boil.AfterInsertHook:
		avatarAfterInsertHooks = append(avatarAfterInsertHooks, avatarHook)
	case boil.AfterSelectHook:
		avatarAfterSelectHooks = append(avatarAfterSelectHooks, avatarHook)
	case boil.AfterUpdateHook:
		avatarAfterUpdateHooks = append(avatarAfterUpdateHooks, avatarHook)
	case boil.AfterDeleteHook:
		avatarAfterDeleteHooks = append(avatarAfterDeleteHooks, avatarHook)
	case boil.AfterUpsertHook:
		avatarAfterUpsertHooks = append(avatarAfterUpsertHooks, avatarHook)
	}
}

// OneP returns a single avatar record from the query, and panics on error.
func (q avatarQuery) OneP() *Avatar {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single avatar record from the query.
func (q avatarQuery) One() (*Avatar, error) {
	o := &Avatar{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for avatar")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Avatar records from the query, and panics on error.
func (q avatarQuery) AllP() AvatarSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Avatar records from the query.
func (q avatarQuery) All() (AvatarSlice, error) {
	var o AvatarSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Avatar slice")
	}

	if len(avatarAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Avatar records in the query, and panics on error.
func (q avatarQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Avatar records in the query.
func (q avatarQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count avatar rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q avatarQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q avatarQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if avatar exists")
	}

	return count > 0, nil
}

// AssetG pointed to by the foreign key.
func (o *Avatar) AssetG(mods ...qm.QueryMod) assetQuery {
	return o.AssetByFk(boil.GetDB(), mods...)
}

// Asset pointed to by the foreign key.
func (o *Avatar) AssetByFk(exec boil.Executor, mods ...qm.QueryMod) assetQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.Asset),
	}

	queryMods = append(queryMods, mods...)

	query := Assets(exec, queryMods...)
	queries.SetFrom(query.Query, "\"asset\"")

	return query
}

// PartyG pointed to by the foreign key.
func (o *Avatar) PartyG(mods ...qm.QueryMod) partyQuery {
	return o.PartyByFk(boil.GetDB(), mods...)
}

// Party pointed to by the foreign key.
func (o *Avatar) PartyByFk(exec boil.Executor, mods ...qm.QueryMod) partyQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.Party),
	}

	queryMods = append(queryMods, mods...)

	query := Parties(exec, queryMods...)
	queries.SetFrom(query.Query, "\"party\"")

	return query
}

// LoadAsset allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (avatarL) LoadAsset(e boil.Executor, singular bool, maybeAvatar interface{}) error {
	var slice []*Avatar
	var object *Avatar

	count := 1
	if singular {
		object = maybeAvatar.(*Avatar)
	} else {
		slice = *maybeAvatar.(*AvatarSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &avatarR{}
		}
		args[0] = object.Asset
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &avatarR{}
			}
			args[i] = obj.Asset
		}
	}

	query := fmt.Sprintf(
		"select * from \"asset\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Asset")
	}
	defer results.Close()

	var resultSlice []*Asset
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Asset")
	}

	if len(avatarAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Asset = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Asset == foreign.ID {
				local.R.Asset = foreign
				break
			}
		}
	}

	return nil
}

// LoadParty allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (avatarL) LoadParty(e boil.Executor, singular bool, maybeAvatar interface{}) error {
	var slice []*Avatar
	var object *Avatar

	count := 1
	if singular {
		object = maybeAvatar.(*Avatar)
	} else {
		slice = *maybeAvatar.(*AvatarSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &avatarR{}
		}
		args[0] = object.Party
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &avatarR{}
			}
			args[i] = obj.Party
		}
	}

	query := fmt.Sprintf(
		"select * from \"party\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Party")
	}
	defer results.Close()

	var resultSlice []*Party
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Party")
	}

	if len(avatarAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Party = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Party == foreign.ID {
				local.R.Party = foreign
				break
			}
		}
	}

	return nil
}

// SetAssetG of the avatar to the related item.
// Sets o.R.Asset to related.
// Adds o to related.R.Avatars.
// Uses the global database handle.
func (o *Avatar) SetAssetG(insert bool, related *Asset) error {
	return o.SetAsset(boil.GetDB(), insert, related)
}

// SetAssetP of the avatar to the related item.
// Sets o.R.Asset to related.
// Adds o to related.R.Avatars.
// Panics on error.
func (o *Avatar) SetAssetP(exec boil.Executor, insert bool, related *Asset) {
	if err := o.SetAsset(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetAssetGP of the avatar to the related item.
// Sets o.R.Asset to related.
// Adds o to related.R.Avatars.
// Uses the global database handle and panics on error.
func (o *Avatar) SetAssetGP(insert bool, related *Asset) {
	if err := o.SetAsset(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetAsset of the avatar to the related item.
// Sets o.R.Asset to related.
// Adds o to related.R.Avatars.
func (o *Avatar) SetAsset(exec boil.Executor, insert bool, related *Asset) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"avatar\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"asset"}),
		strmangle.WhereClause("\"", "\"", 2, avatarPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Party}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Asset = related.ID

	if o.R == nil {
		o.R = &avatarR{
			Asset: related,
		}
	} else {
		o.R.Asset = related
	}

	if related.R == nil {
		related.R = &assetR{
			Avatars: AvatarSlice{o},
		}
	} else {
		related.R.Avatars = append(related.R.Avatars, o)
	}

	return nil
}

// SetPartyG of the avatar to the related item.
// Sets o.R.Party to related.
// Adds o to related.R.Avatar.
// Uses the global database handle.
func (o *Avatar) SetPartyG(insert bool, related *Party) error {
	return o.SetParty(boil.GetDB(), insert, related)
}

// SetPartyP of the avatar to the related item.
// Sets o.R.Party to related.
// Adds o to related.R.Avatar.
// Panics on error.
func (o *Avatar) SetPartyP(exec boil.Executor, insert bool, related *Party) {
	if err := o.SetParty(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetPartyGP of the avatar to the related item.
// Sets o.R.Party to related.
// Adds o to related.R.Avatar.
// Uses the global database handle and panics on error.
func (o *Avatar) SetPartyGP(insert bool, related *Party) {
	if err := o.SetParty(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetParty of the avatar to the related item.
// Sets o.R.Party to related.
// Adds o to related.R.Avatar.
func (o *Avatar) SetParty(exec boil.Executor, insert bool, related *Party) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"avatar\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"party"}),
		strmangle.WhereClause("\"", "\"", 2, avatarPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Party}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Party = related.ID

	if o.R == nil {
		o.R = &avatarR{
			Party: related,
		}
	} else {
		o.R.Party = related
	}

	if related.R == nil {
		related.R = &partyR{
			Avatar: o,
		}
	} else {
		related.R.Avatar = o
	}

	return nil
}

// AvatarsG retrieves all records.
func AvatarsG(mods ...qm.QueryMod) avatarQuery {
	return Avatars(boil.GetDB(), mods...)
}

// Avatars retrieves all the records using an executor.
func Avatars(exec boil.Executor, mods ...qm.QueryMod) avatarQuery {
	mods = append(mods, qm.From("\"avatar\""))
	return avatarQuery{NewQuery(exec, mods...)}
}

// FindAvatarG retrieves a single record by ID.
func FindAvatarG(party int, selectCols ...string) (*Avatar, error) {
	return FindAvatar(boil.GetDB(), party, selectCols...)
}

// FindAvatarGP retrieves a single record by ID, and panics on error.
func FindAvatarGP(party int, selectCols ...string) *Avatar {
	retobj, err := FindAvatar(boil.GetDB(), party, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAvatar retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAvatar(exec boil.Executor, party int, selectCols ...string) (*Avatar, error) {
	avatarObj := &Avatar{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"avatar\" where \"party\"=$1", sel,
	)

	q := queries.Raw(exec, query, party)

	err := q.Bind(avatarObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from avatar")
	}

	return avatarObj, nil
}

// FindAvatarP retrieves a single record by ID with an executor, and panics on error.
func FindAvatarP(exec boil.Executor, party int, selectCols ...string) *Avatar {
	retobj, err := FindAvatar(exec, party, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Avatar) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Avatar) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Avatar) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Avatar) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no avatar provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(avatarColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	avatarInsertCacheMut.RLock()
	cache, cached := avatarInsertCache[key]
	avatarInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			avatarColumns,
			avatarColumnsWithDefault,
			avatarColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(avatarType, avatarMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(avatarType, avatarMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"avatar\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"avatar\" DEFAULT VALUES"
		}

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into avatar")
	}

	if !cached {
		avatarInsertCacheMut.Lock()
		avatarInsertCache[key] = cache
		avatarInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Avatar record. See Update for
// whitelist behavior description.
func (o *Avatar) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Avatar record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Avatar) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Avatar, and panics on error.
// See Update for whitelist behavior description.
func (o *Avatar) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Avatar.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Avatar) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	avatarUpdateCacheMut.RLock()
	cache, cached := avatarUpdateCache[key]
	avatarUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(avatarColumns, avatarPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update avatar, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"avatar\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, avatarPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(avatarType, avatarMapping, append(wl, avatarPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update avatar row")
	}

	if !cached {
		avatarUpdateCacheMut.Lock()
		avatarUpdateCache[key] = cache
		avatarUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q avatarQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q avatarQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for avatar")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AvatarSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AvatarSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AvatarSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AvatarSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), avatarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"UPDATE \"avatar\" SET %s WHERE (\"party\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(avatarPrimaryKeyColumns), len(colNames)+1, len(avatarPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in avatar slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Avatar) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Avatar) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Avatar) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Avatar) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no avatar provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(avatarColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	avatarUpsertCacheMut.RLock()
	cache, cached := avatarUpsertCache[key]
	avatarUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			avatarColumns,
			avatarColumnsWithDefault,
			avatarColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			avatarColumns,
			avatarPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert avatar, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(avatarPrimaryKeyColumns))
			copy(conflict, avatarPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"avatar\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(avatarType, avatarMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(avatarType, avatarMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert avatar")
	}

	if !cached {
		avatarUpsertCacheMut.Lock()
		avatarUpsertCache[key] = cache
		avatarUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Avatar record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Avatar) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Avatar record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Avatar) DeleteG() error {
	if o == nil {
		return errors.New("models: no Avatar provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Avatar record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Avatar) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Avatar record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Avatar) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Avatar provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), avatarPrimaryKeyMapping)
	query := "DELETE FROM \"avatar\" WHERE \"party\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from avatar")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q avatarQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q avatarQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no avatarQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from avatar")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AvatarSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AvatarSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Avatar slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AvatarSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AvatarSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Avatar slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(avatarBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), avatarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"DELETE FROM \"avatar\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, avatarPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(avatarPrimaryKeyColumns), 1, len(avatarPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from avatar slice")
	}

	if len(avatarAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Avatar) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Avatar) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Avatar) ReloadG() error {
	if o == nil {
		return errors.New("models: no Avatar provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Avatar) Reload(exec boil.Executor) error {
	ret, err := FindAvatar(exec, o.Party)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AvatarSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AvatarSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AvatarSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AvatarSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AvatarSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	avatars := AvatarSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), avatarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"SELECT \"avatar\".* FROM \"avatar\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, avatarPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(avatarPrimaryKeyColumns), 1, len(avatarPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, query, args...)

	err := q.Bind(&avatars)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AvatarSlice")
	}

	*o = avatars

	return nil
}

// AvatarExists checks if the Avatar row exists.
func AvatarExists(exec boil.Executor, party int) (bool, error) {
	var exists bool

	query := "select exists(select 1 from \"avatar\" where \"party\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, party)
	}

	row := exec.QueryRow(query, party)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if avatar exists")
	}

	return exists, nil
}

// AvatarExistsG checks if the Avatar row exists.
func AvatarExistsG(party int) (bool, error) {
	return AvatarExists(boil.GetDB(), party)
}

// AvatarExistsGP checks if the Avatar row exists. Panics on error.
func AvatarExistsGP(party int) bool {
	e, err := AvatarExists(boil.GetDB(), party)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AvatarExistsP checks if the Avatar row exists. Panics on error.
func AvatarExistsP(exec boil.Executor, party int) bool {
	e, err := AvatarExists(exec, party)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
