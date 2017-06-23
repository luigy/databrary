// This file is generated by SQLBoiler (https://github.com/databrary/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package public

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/databrary/sqlboiler/boil"
	"github.com/databrary/sqlboiler/queries"
	"github.com/databrary/sqlboiler/queries/qm"
	"github.com/databrary/sqlboiler/strmangle"
	"github.com/pkg/errors"
	"gopkg.in/nullbio/null.v6"
	"reflect"
	"strings"
	"sync"
	"time"
)

// VolumeCitation is an object representing the database table.
type VolumeCitation struct {
	Volume int         `db:"volume" json:"volumeCitation_volume"`
	Head   string      `db:"head" json:"volumeCitation_head"`
	URL    null.String `db:"url" json:"volumeCitation_url,omitempty"`
	Year   null.Int16  `db:"year" json:"volumeCitation_year,omitempty"`

	R *volumeCitationR `db:"-" json:"-"`
	L volumeCitationL  `db:"-" json:"-"`
}

// volumeCitationR is where relationships are stored.
type volumeCitationR struct {
	Volume *Volume
}

// volumeCitationL is where Load methods for each relationship are stored.
type volumeCitationL struct{}

var (
	volumeCitationColumns               = []string{"volume", "head", "url", "year"}
	volumeCitationColumnsWithoutDefault = []string{"volume", "head", "url", "year"}
	volumeCitationColumnsWithDefault    = []string{}
	volumeCitationColumnsWithCustom     = []string{}

	volumeCitationPrimaryKeyColumns = []string{"volume"}
)

type (
	// VolumeCitationSlice is an alias for a slice of pointers to VolumeCitation.
	// This should generally be used opposed to []VolumeCitation.
	VolumeCitationSlice []*VolumeCitation
	// VolumeCitationHook is the signature for custom VolumeCitation hook methods
	VolumeCitationHook func(boil.Executor, *VolumeCitation) error

	volumeCitationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	volumeCitationType    = reflect.TypeOf(&VolumeCitation{})
	volumeCitationMapping = queries.MakeStructMapping(volumeCitationType)

	volumeCitationPrimaryKeyMapping, _ = queries.BindMapping(volumeCitationType, volumeCitationMapping, volumeCitationPrimaryKeyColumns)

	volumeCitationInsertCacheMut sync.RWMutex
	volumeCitationInsertCache    = make(map[string]insertCache)
	volumeCitationUpdateCacheMut sync.RWMutex
	volumeCitationUpdateCache    = make(map[string]updateCache)
	volumeCitationUpsertCacheMut sync.RWMutex
	volumeCitationUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var volumeCitationBeforeInsertHooks []VolumeCitationHook
var volumeCitationBeforeUpdateHooks []VolumeCitationHook
var volumeCitationBeforeDeleteHooks []VolumeCitationHook
var volumeCitationBeforeUpsertHooks []VolumeCitationHook

var volumeCitationAfterInsertHooks []VolumeCitationHook
var volumeCitationAfterSelectHooks []VolumeCitationHook
var volumeCitationAfterUpdateHooks []VolumeCitationHook
var volumeCitationAfterDeleteHooks []VolumeCitationHook
var volumeCitationAfterUpsertHooks []VolumeCitationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *VolumeCitation) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *VolumeCitation) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *VolumeCitation) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *VolumeCitation) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *VolumeCitation) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *VolumeCitation) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *VolumeCitation) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *VolumeCitation) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *VolumeCitation) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVolumeCitationHook registers your hook function for all future operations.
func AddVolumeCitationHook(hookPoint boil.HookPoint, volumeCitationHook VolumeCitationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		volumeCitationBeforeInsertHooks = append(volumeCitationBeforeInsertHooks, volumeCitationHook)
	case boil.BeforeUpdateHook:
		volumeCitationBeforeUpdateHooks = append(volumeCitationBeforeUpdateHooks, volumeCitationHook)
	case boil.BeforeDeleteHook:
		volumeCitationBeforeDeleteHooks = append(volumeCitationBeforeDeleteHooks, volumeCitationHook)
	case boil.BeforeUpsertHook:
		volumeCitationBeforeUpsertHooks = append(volumeCitationBeforeUpsertHooks, volumeCitationHook)
	case boil.AfterInsertHook:
		volumeCitationAfterInsertHooks = append(volumeCitationAfterInsertHooks, volumeCitationHook)
	case boil.AfterSelectHook:
		volumeCitationAfterSelectHooks = append(volumeCitationAfterSelectHooks, volumeCitationHook)
	case boil.AfterUpdateHook:
		volumeCitationAfterUpdateHooks = append(volumeCitationAfterUpdateHooks, volumeCitationHook)
	case boil.AfterDeleteHook:
		volumeCitationAfterDeleteHooks = append(volumeCitationAfterDeleteHooks, volumeCitationHook)
	case boil.AfterUpsertHook:
		volumeCitationAfterUpsertHooks = append(volumeCitationAfterUpsertHooks, volumeCitationHook)
	}
}

// OneP returns a single volumeCitation record from the query, and panics on error.
func (q volumeCitationQuery) OneP() *VolumeCitation {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single volumeCitation record from the query.
func (q volumeCitationQuery) One() (*VolumeCitation, error) {
	o := &VolumeCitation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "public: failed to execute a one query for volume_citation")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all VolumeCitation records from the query, and panics on error.
func (q volumeCitationQuery) AllP() VolumeCitationSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all VolumeCitation records from the query.
func (q volumeCitationQuery) All() (VolumeCitationSlice, error) {
	var o VolumeCitationSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "public: failed to assign all query results to VolumeCitation slice")
	}

	if len(volumeCitationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all VolumeCitation records in the query, and panics on error.
func (q volumeCitationQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all VolumeCitation records in the query.
func (q volumeCitationQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "public: failed to count volume_citation rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q volumeCitationQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q volumeCitationQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "public: failed to check if volume_citation exists")
	}

	return count > 0, nil
}

// VolumeG pointed to by the foreign key.
func (o *VolumeCitation) VolumeG(mods ...qm.QueryMod) volumeQuery {
	return o.VolumeByFk(boil.GetDB(), mods...)
}

// Volume pointed to by the foreign key.
func (o *VolumeCitation) VolumeByFk(exec boil.Executor, mods ...qm.QueryMod) volumeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.Volume),
	}

	queryMods = append(queryMods, mods...)

	query := Volumes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"volume\"")

	return query
}

// LoadVolume allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (volumeCitationL) LoadVolume(e boil.Executor, singular bool, maybeVolumeCitation interface{}) error {
	var slice []*VolumeCitation
	var object *VolumeCitation

	count := 1
	if singular {
		object = maybeVolumeCitation.(*VolumeCitation)
	} else {
		slice = *maybeVolumeCitation.(*VolumeCitationSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &volumeCitationR{}
		}
		args[0] = object.Volume
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &volumeCitationR{}
			}
			args[i] = obj.Volume
		}
	}

	query := fmt.Sprintf(
		"select * from \"volume\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Volume")
	}
	defer results.Close()

	var resultSlice []*Volume
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Volume")
	}

	if len(volumeCitationAfterSelectHooks) != 0 {
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
		object.R.Volume = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Volume == foreign.ID {
				local.R.Volume = foreign
				break
			}
		}
	}

	return nil
}

// SetVolumeG of the volume_citation to the related item.
// Sets o.R.Volume to related.
// Adds o to related.R.VolumeCitation.
// Uses the global database handle.
func (o *VolumeCitation) SetVolumeG(insert bool, related *Volume) error {
	return o.SetVolume(boil.GetDB(), insert, related)
}

// SetVolumeP of the volume_citation to the related item.
// Sets o.R.Volume to related.
// Adds o to related.R.VolumeCitation.
// Panics on error.
func (o *VolumeCitation) SetVolumeP(exec boil.Executor, insert bool, related *Volume) {
	if err := o.SetVolume(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetVolumeGP of the volume_citation to the related item.
// Sets o.R.Volume to related.
// Adds o to related.R.VolumeCitation.
// Uses the global database handle and panics on error.
func (o *VolumeCitation) SetVolumeGP(insert bool, related *Volume) {
	if err := o.SetVolume(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetVolume of the volume_citation to the related item.
// Sets o.R.Volume to related.
// Adds o to related.R.VolumeCitation.
func (o *VolumeCitation) SetVolume(exec boil.Executor, insert bool, related *Volume) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"volume_citation\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"volume"}),
		strmangle.WhereClause("\"", "\"", 2, volumeCitationPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Volume}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Volume = related.ID

	if o.R == nil {
		o.R = &volumeCitationR{
			Volume: related,
		}
	} else {
		o.R.Volume = related
	}

	if related.R == nil {
		related.R = &volumeR{
			VolumeCitation: o,
		}
	} else {
		related.R.VolumeCitation = o
	}

	return nil
}

// VolumeCitationsG retrieves all records.
func VolumeCitationsG(mods ...qm.QueryMod) volumeCitationQuery {
	return VolumeCitations(boil.GetDB(), mods...)
}

// VolumeCitations retrieves all the records using an executor.
func VolumeCitations(exec boil.Executor, mods ...qm.QueryMod) volumeCitationQuery {
	mods = append(mods, qm.From("\"volume_citation\""))
	return volumeCitationQuery{NewQuery(exec, mods...)}
}

// FindVolumeCitationG retrieves a single record by ID.
func FindVolumeCitationG(volume int, selectCols ...string) (*VolumeCitation, error) {
	return FindVolumeCitation(boil.GetDB(), volume, selectCols...)
}

// FindVolumeCitationGP retrieves a single record by ID, and panics on error.
func FindVolumeCitationGP(volume int, selectCols ...string) *VolumeCitation {
	retobj, err := FindVolumeCitation(boil.GetDB(), volume, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindVolumeCitation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVolumeCitation(exec boil.Executor, volume int, selectCols ...string) (*VolumeCitation, error) {
	volumeCitationObj := &VolumeCitation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"volume_citation\" where \"volume\"=$1", sel,
	)

	q := queries.Raw(exec, query, volume)

	err := q.Bind(volumeCitationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "public: unable to select from volume_citation")
	}

	return volumeCitationObj, nil
}

// FindVolumeCitationP retrieves a single record by ID with an executor, and panics on error.
func FindVolumeCitationP(exec boil.Executor, volume int, selectCols ...string) *VolumeCitation {
	retobj, err := FindVolumeCitation(exec, volume, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *VolumeCitation) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *VolumeCitation) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *VolumeCitation) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *VolumeCitation) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("public: no volume_citation provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(volumeCitationColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	volumeCitationInsertCacheMut.RLock()
	cache, cached := volumeCitationInsertCache[key]
	volumeCitationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			volumeCitationColumns,
			volumeCitationColumnsWithDefault,
			volumeCitationColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(volumeCitationType, volumeCitationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(volumeCitationType, volumeCitationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"volume_citation\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"volume_citation\" DEFAULT VALUES"
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
		return errors.Wrap(err, "public: unable to insert into volume_citation")
	}

	if !cached {
		volumeCitationInsertCacheMut.Lock()
		volumeCitationInsertCache[key] = cache
		volumeCitationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single VolumeCitation record. See Update for
// whitelist behavior description.
func (o *VolumeCitation) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single VolumeCitation record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *VolumeCitation) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the VolumeCitation, and panics on error.
// See Update for whitelist behavior description.
func (o *VolumeCitation) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the VolumeCitation.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *VolumeCitation) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	volumeCitationUpdateCacheMut.RLock()
	cache, cached := volumeCitationUpdateCache[key]
	volumeCitationUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(volumeCitationColumns, volumeCitationPrimaryKeyColumns, whitelist)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("public: unable to update volume_citation, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"volume_citation\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, volumeCitationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(volumeCitationType, volumeCitationMapping, append(wl, volumeCitationPrimaryKeyColumns...))
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
		return errors.Wrap(err, "public: unable to update volume_citation row")
	}

	if !cached {
		volumeCitationUpdateCacheMut.Lock()
		volumeCitationUpdateCache[key] = cache
		volumeCitationUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q volumeCitationQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q volumeCitationQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "public: unable to update all for volume_citation")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o VolumeCitationSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o VolumeCitationSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o VolumeCitationSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VolumeCitationSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("public: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), volumeCitationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"UPDATE \"volume_citation\" SET %s WHERE (\"volume\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(volumeCitationPrimaryKeyColumns), len(colNames)+1, len(volumeCitationPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "public: unable to update all in volumeCitation slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *VolumeCitation) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *VolumeCitation) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *VolumeCitation) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *VolumeCitation) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("public: no volume_citation provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(volumeCitationColumnsWithDefault, o)

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

	volumeCitationUpsertCacheMut.RLock()
	cache, cached := volumeCitationUpsertCache[key]
	volumeCitationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			volumeCitationColumns,
			volumeCitationColumnsWithDefault,
			volumeCitationColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			volumeCitationColumns,
			volumeCitationPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("public: unable to upsert volume_citation, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(volumeCitationPrimaryKeyColumns))
			copy(conflict, volumeCitationPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"volume_citation\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(volumeCitationType, volumeCitationMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(volumeCitationType, volumeCitationMapping, ret)
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
		return errors.Wrap(err, "public: unable to upsert volume_citation")
	}

	if !cached {
		volumeCitationUpsertCacheMut.Lock()
		volumeCitationUpsertCache[key] = cache
		volumeCitationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single VolumeCitation record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *VolumeCitation) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single VolumeCitation record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *VolumeCitation) DeleteG() error {
	if o == nil {
		return errors.New("public: no VolumeCitation provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single VolumeCitation record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *VolumeCitation) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single VolumeCitation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *VolumeCitation) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("public: no VolumeCitation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), volumeCitationPrimaryKeyMapping)
	query := "DELETE FROM \"volume_citation\" WHERE \"volume\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "public: unable to delete from volume_citation")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q volumeCitationQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q volumeCitationQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("public: no volumeCitationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "public: unable to delete all from volume_citation")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o VolumeCitationSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o VolumeCitationSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("public: no VolumeCitation slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o VolumeCitationSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VolumeCitationSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("public: no VolumeCitation slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(volumeCitationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), volumeCitationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"DELETE FROM \"volume_citation\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, volumeCitationPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(volumeCitationPrimaryKeyColumns), 1, len(volumeCitationPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "public: unable to delete all from volumeCitation slice")
	}

	if len(volumeCitationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *VolumeCitation) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *VolumeCitation) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *VolumeCitation) ReloadG() error {
	if o == nil {
		return errors.New("public: no VolumeCitation provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *VolumeCitation) Reload(exec boil.Executor) error {
	ret, err := FindVolumeCitation(exec, o.Volume)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *VolumeCitationSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *VolumeCitationSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VolumeCitationSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("public: empty VolumeCitationSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VolumeCitationSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	volumeCitations := VolumeCitationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), volumeCitationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"SELECT \"volume_citation\".* FROM \"volume_citation\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, volumeCitationPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(volumeCitationPrimaryKeyColumns), 1, len(volumeCitationPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, query, args...)

	err := q.Bind(&volumeCitations)
	if err != nil {
		return errors.Wrap(err, "public: unable to reload all in VolumeCitationSlice")
	}

	*o = volumeCitations

	return nil
}

// VolumeCitationExists checks if the VolumeCitation row exists.
func VolumeCitationExists(exec boil.Executor, volume int) (bool, error) {
	var exists bool

	query := "select exists(select 1 from \"volume_citation\" where \"volume\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, volume)
	}

	row := exec.QueryRow(query, volume)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "public: unable to check if volume_citation exists")
	}

	return exists, nil
}

// VolumeCitationExistsG checks if the VolumeCitation row exists.
func VolumeCitationExistsG(volume int) (bool, error) {
	return VolumeCitationExists(boil.GetDB(), volume)
}

// VolumeCitationExistsGP checks if the VolumeCitation row exists. Panics on error.
func VolumeCitationExistsGP(volume int) bool {
	e, err := VolumeCitationExists(boil.GetDB(), volume)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// VolumeCitationExistsP checks if the VolumeCitation row exists. Panics on error.
func VolumeCitationExistsP(exec boil.Executor, volume int) bool {
	e, err := VolumeCitationExists(exec, volume)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}