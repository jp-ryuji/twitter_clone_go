// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testFollowings(t *testing.T) {
	t.Parallel()

	query := Followings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFollowingsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Followings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFollowingsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Followings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Followings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFollowingsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FollowingSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Followings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFollowingsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FollowingExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Following exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FollowingExists to return true, but got false.")
	}
}

func testFollowingsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	followingFound, err := FindFollowing(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if followingFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFollowingsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Followings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFollowingsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Followings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFollowingsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	followingOne := &Following{}
	followingTwo := &Following{}
	if err = randomize.Struct(seed, followingOne, followingDBTypes, false, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}
	if err = randomize.Struct(seed, followingTwo, followingDBTypes, false, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = followingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = followingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Followings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFollowingsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	followingOne := &Following{}
	followingTwo := &Following{}
	if err = randomize.Struct(seed, followingOne, followingDBTypes, false, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}
	if err = randomize.Struct(seed, followingTwo, followingDBTypes, false, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = followingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = followingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Followings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func followingBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Following) error {
	*o = Following{}
	return nil
}

func followingAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Following) error {
	*o = Following{}
	return nil
}

func followingAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Following) error {
	*o = Following{}
	return nil
}

func followingBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Following) error {
	*o = Following{}
	return nil
}

func followingAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Following) error {
	*o = Following{}
	return nil
}

func followingBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Following) error {
	*o = Following{}
	return nil
}

func followingAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Following) error {
	*o = Following{}
	return nil
}

func followingBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Following) error {
	*o = Following{}
	return nil
}

func followingAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Following) error {
	*o = Following{}
	return nil
}

func testFollowingsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Following{}
	o := &Following{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, followingDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Following object: %s", err)
	}

	AddFollowingHook(boil.BeforeInsertHook, followingBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	followingBeforeInsertHooks = []FollowingHook{}

	AddFollowingHook(boil.AfterInsertHook, followingAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	followingAfterInsertHooks = []FollowingHook{}

	AddFollowingHook(boil.AfterSelectHook, followingAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	followingAfterSelectHooks = []FollowingHook{}

	AddFollowingHook(boil.BeforeUpdateHook, followingBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	followingBeforeUpdateHooks = []FollowingHook{}

	AddFollowingHook(boil.AfterUpdateHook, followingAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	followingAfterUpdateHooks = []FollowingHook{}

	AddFollowingHook(boil.BeforeDeleteHook, followingBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	followingBeforeDeleteHooks = []FollowingHook{}

	AddFollowingHook(boil.AfterDeleteHook, followingAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	followingAfterDeleteHooks = []FollowingHook{}

	AddFollowingHook(boil.BeforeUpsertHook, followingBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	followingBeforeUpsertHooks = []FollowingHook{}

	AddFollowingHook(boil.AfterUpsertHook, followingAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	followingAfterUpsertHooks = []FollowingHook{}
}

func testFollowingsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Followings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFollowingsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(followingColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Followings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFollowingsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFollowingsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FollowingSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFollowingsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Followings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	followingDBTypes = map[string]string{`ID`: `bigint`, `FollowingUserID`: `integer`, `FollowerID`: `integer`}
	_                = bytes.MinRead
)

func testFollowingsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(followingPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(followingAllColumns) == len(followingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Followings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, followingDBTypes, true, followingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFollowingsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(followingAllColumns) == len(followingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Following{}
	if err = randomize.Struct(seed, o, followingDBTypes, true, followingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Followings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, followingDBTypes, true, followingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(followingAllColumns, followingPrimaryKeyColumns) {
		fields = followingAllColumns
	} else {
		fields = strmangle.SetComplement(
			followingAllColumns,
			followingPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := FollowingSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFollowingsUpsert(t *testing.T) {
	t.Parallel()

	if len(followingAllColumns) == len(followingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Following{}
	if err = randomize.Struct(seed, &o, followingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Following: %s", err)
	}

	count, err := Followings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, followingDBTypes, false, followingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Following struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Following: %s", err)
	}

	count, err = Followings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
