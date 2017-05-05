// Package xo_models contains the types for schema 'public'.
package xo_models

import (
	"github.com/databrary/databrary/db/models/custom_types"
	"github.com/vattle/sqlboiler/boil"
)

// GENERATED BY XO. DO NOT EDIT.

// SegmentUnion calls the stored procedure 'public.segment_union(segment) segment' on db.
func SegmentUnion(exec boil.Executor, v0 custom_types.Segment) (custom_types.Segment, error) {
	var err error

	// sql query
	const query = `SELECT public.segment_union($1)`

	// run query
	var ret custom_types.Segment

	err = exec.QueryRow(query, v0).Scan(&ret)
	if err != nil {
		return custom_types.Segment{}, err
	}

	return ret, nil
}
