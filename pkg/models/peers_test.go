// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPeers(t *testing.T) {
	t.Parallel()

	query := Peers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPeersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
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

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Peers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PeerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PeerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Peer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PeerExists to return true, but got false.")
	}
}

func testPeersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	peerFound, err := FindPeer(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if peerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPeersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Peers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPeersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Peers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPeersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	peerOne := &Peer{}
	peerTwo := &Peer{}
	if err = randomize.Struct(seed, peerOne, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}
	if err = randomize.Struct(seed, peerTwo, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = peerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = peerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Peers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPeersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	peerOne := &Peer{}
	peerTwo := &Peer{}
	if err = randomize.Struct(seed, peerOne, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}
	if err = randomize.Struct(seed, peerTwo, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = peerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = peerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func peerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func peerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Peer) error {
	*o = Peer{}
	return nil
}

func testPeersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Peer{}
	o := &Peer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, peerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Peer object: %s", err)
	}

	AddPeerHook(boil.BeforeInsertHook, peerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	peerBeforeInsertHooks = []PeerHook{}

	AddPeerHook(boil.AfterInsertHook, peerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	peerAfterInsertHooks = []PeerHook{}

	AddPeerHook(boil.AfterSelectHook, peerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	peerAfterSelectHooks = []PeerHook{}

	AddPeerHook(boil.BeforeUpdateHook, peerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	peerBeforeUpdateHooks = []PeerHook{}

	AddPeerHook(boil.AfterUpdateHook, peerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	peerAfterUpdateHooks = []PeerHook{}

	AddPeerHook(boil.BeforeDeleteHook, peerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	peerBeforeDeleteHooks = []PeerHook{}

	AddPeerHook(boil.AfterDeleteHook, peerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	peerAfterDeleteHooks = []PeerHook{}

	AddPeerHook(boil.BeforeUpsertHook, peerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	peerBeforeUpsertHooks = []PeerHook{}

	AddPeerHook(boil.AfterUpsertHook, peerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	peerAfterUpsertHooks = []PeerHook{}
}

func testPeersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPeersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(peerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPeerToManyLocalConnectionEvents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c ConnectionEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, connectionEventDBTypes, false, connectionEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, connectionEventDBTypes, false, connectionEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.LocalID = a.ID
	c.LocalID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.LocalConnectionEvents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.LocalID == b.LocalID {
			bFound = true
		}
		if v.LocalID == c.LocalID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeerSlice{&a}
	if err = a.L.LoadLocalConnectionEvents(ctx, tx, false, (*[]*Peer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LocalConnectionEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.LocalConnectionEvents = nil
	if err = a.L.LoadLocalConnectionEvents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LocalConnectionEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPeerToManyRemoteConnectionEvents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c ConnectionEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, connectionEventDBTypes, false, connectionEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, connectionEventDBTypes, false, connectionEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.RemoteID = a.ID
	c.RemoteID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.RemoteConnectionEvents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.RemoteID == b.RemoteID {
			bFound = true
		}
		if v.RemoteID == c.RemoteID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeerSlice{&a}
	if err = a.L.LoadRemoteConnectionEvents(ctx, tx, false, (*[]*Peer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RemoteConnectionEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RemoteConnectionEvents = nil
	if err = a.L.LoadRemoteConnectionEvents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RemoteConnectionEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPeerToManyClientHolePunchResults(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c HolePunchResult

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, holePunchResultDBTypes, false, holePunchResultColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, holePunchResultDBTypes, false, holePunchResultColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ClientID = a.ID
	c.ClientID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ClientHolePunchResults().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ClientID == b.ClientID {
			bFound = true
		}
		if v.ClientID == c.ClientID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeerSlice{&a}
	if err = a.L.LoadClientHolePunchResults(ctx, tx, false, (*[]*Peer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ClientHolePunchResults); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ClientHolePunchResults = nil
	if err = a.L.LoadClientHolePunchResults(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ClientHolePunchResults); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPeerToManyRemoteHolePunchResults(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c HolePunchResult

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, holePunchResultDBTypes, false, holePunchResultColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, holePunchResultDBTypes, false, holePunchResultColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.RemoteID = a.ID
	c.RemoteID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.RemoteHolePunchResults().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.RemoteID == b.RemoteID {
			bFound = true
		}
		if v.RemoteID == c.RemoteID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeerSlice{&a}
	if err = a.L.LoadRemoteHolePunchResults(ctx, tx, false, (*[]*Peer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RemoteHolePunchResults); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RemoteHolePunchResults = nil
	if err = a.L.LoadRemoteHolePunchResults(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RemoteHolePunchResults); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPeerToManyPeerLogs(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c PeerLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, peerLogDBTypes, false, peerLogColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, peerLogDBTypes, false, peerLogColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PeerID = a.ID
	c.PeerID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.PeerLogs().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PeerID == b.PeerID {
			bFound = true
		}
		if v.PeerID == c.PeerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PeerSlice{&a}
	if err = a.L.LoadPeerLogs(ctx, tx, false, (*[]*Peer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PeerLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PeerLogs = nil
	if err = a.L.LoadPeerLogs(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PeerLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPeerToManyAddOpLocalConnectionEvents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c, d, e ConnectionEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ConnectionEvent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, connectionEventDBTypes, false, strmangle.SetComplement(connectionEventPrimaryKeyColumns, connectionEventColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ConnectionEvent{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddLocalConnectionEvents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.LocalID {
			t.Error("foreign key was wrong value", a.ID, first.LocalID)
		}
		if a.ID != second.LocalID {
			t.Error("foreign key was wrong value", a.ID, second.LocalID)
		}

		if first.R.Local != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Local != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.LocalConnectionEvents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.LocalConnectionEvents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.LocalConnectionEvents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPeerToManyAddOpRemoteConnectionEvents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c, d, e ConnectionEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ConnectionEvent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, connectionEventDBTypes, false, strmangle.SetComplement(connectionEventPrimaryKeyColumns, connectionEventColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ConnectionEvent{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRemoteConnectionEvents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.RemoteID {
			t.Error("foreign key was wrong value", a.ID, first.RemoteID)
		}
		if a.ID != second.RemoteID {
			t.Error("foreign key was wrong value", a.ID, second.RemoteID)
		}

		if first.R.Remote != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Remote != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RemoteConnectionEvents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RemoteConnectionEvents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RemoteConnectionEvents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPeerToManyAddOpClientHolePunchResults(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c, d, e HolePunchResult

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*HolePunchResult{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, holePunchResultDBTypes, false, strmangle.SetComplement(holePunchResultPrimaryKeyColumns, holePunchResultColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*HolePunchResult{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddClientHolePunchResults(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ClientID {
			t.Error("foreign key was wrong value", a.ID, first.ClientID)
		}
		if a.ID != second.ClientID {
			t.Error("foreign key was wrong value", a.ID, second.ClientID)
		}

		if first.R.Client != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Client != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ClientHolePunchResults[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ClientHolePunchResults[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ClientHolePunchResults().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPeerToManyAddOpRemoteHolePunchResults(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c, d, e HolePunchResult

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*HolePunchResult{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, holePunchResultDBTypes, false, strmangle.SetComplement(holePunchResultPrimaryKeyColumns, holePunchResultColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*HolePunchResult{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRemoteHolePunchResults(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.RemoteID {
			t.Error("foreign key was wrong value", a.ID, first.RemoteID)
		}
		if a.ID != second.RemoteID {
			t.Error("foreign key was wrong value", a.ID, second.RemoteID)
		}

		if first.R.Remote != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Remote != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RemoteHolePunchResults[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RemoteHolePunchResults[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RemoteHolePunchResults().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPeerToManyAddOpPeerLogs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Peer
	var b, c, d, e PeerLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PeerLog{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, peerLogDBTypes, false, strmangle.SetComplement(peerLogPrimaryKeyColumns, peerLogColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*PeerLog{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPeerLogs(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PeerID {
			t.Error("foreign key was wrong value", a.ID, first.PeerID)
		}
		if a.ID != second.PeerID {
			t.Error("foreign key was wrong value", a.ID, second.PeerID)
		}

		if first.R.Peer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Peer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PeerLogs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PeerLogs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PeerLogs().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPeersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
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

func testPeersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PeerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPeersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Peers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	peerDBTypes = map[string]string{`ID`: `bigint`, `MultiHash`: `text`, `AgentVersion`: `text`, `Protocols`: `ARRAYtext`, `SupportsDcutr`: `boolean`, `UpdatedAt`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`}
	_           = bytes.MinRead
)

func testPeersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(peerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(peerAllColumns) == len(peerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, peerDBTypes, true, peerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPeersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(peerAllColumns) == len(peerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Peer{}
	if err = randomize.Struct(seed, o, peerDBTypes, true, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, peerDBTypes, true, peerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(peerAllColumns, peerPrimaryKeyColumns) {
		fields = peerAllColumns
	} else {
		fields = strmangle.SetComplement(
			peerAllColumns,
			peerPrimaryKeyColumns,
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

	slice := PeerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPeersUpsert(t *testing.T) {
	t.Parallel()

	if len(peerAllColumns) == len(peerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Peer{}
	if err = randomize.Struct(seed, &o, peerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Peer: %s", err)
	}

	count, err := Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, peerDBTypes, false, peerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Peer: %s", err)
	}

	count, err = Peers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
