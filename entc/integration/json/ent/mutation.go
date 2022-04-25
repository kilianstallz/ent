// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"entgo.io/ent/entc/integration/json/ent/predicate"
	"entgo.io/ent/entc/integration/json/ent/schema"
	"entgo.io/ent/entc/integration/json/ent/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	t             **schema.T
	url           **url.URL
	raw           *json.RawMessage
	dirs          *[]http.Dir
	ints          *[]int
	floats        *[]float64
	strings       *[]string
	addr          *schema.Addr
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetT sets the "t" field.
func (m *UserMutation) SetT(s *schema.T) {
	m.t = &s
}

// T returns the value of the "t" field in the mutation.
func (m *UserMutation) T() (r *schema.T, exists bool) {
	v := m.t
	if v == nil {
		return
	}
	return *v, true
}

// OldT returns the old "t" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldT(ctx context.Context) (v *schema.T, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldT is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldT requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldT: %w", err)
	}
	return oldValue.T, nil
}

// ClearT clears the value of the "t" field.
func (m *UserMutation) ClearT() {
	m.t = nil
	m.clearedFields[user.FieldT] = struct{}{}
}

// TCleared returns if the "t" field was cleared in this mutation.
func (m *UserMutation) TCleared() bool {
	_, ok := m.clearedFields[user.FieldT]
	return ok
}

// ResetT resets all changes to the "t" field.
func (m *UserMutation) ResetT() {
	m.t = nil
	delete(m.clearedFields, user.FieldT)
}

// SetURL sets the "url" field.
func (m *UserMutation) SetURL(u *url.URL) {
	m.url = &u
}

// URL returns the value of the "url" field in the mutation.
func (m *UserMutation) URL() (r *url.URL, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldURL(ctx context.Context) (v *url.URL, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ClearURL clears the value of the "url" field.
func (m *UserMutation) ClearURL() {
	m.url = nil
	m.clearedFields[user.FieldURL] = struct{}{}
}

// URLCleared returns if the "url" field was cleared in this mutation.
func (m *UserMutation) URLCleared() bool {
	_, ok := m.clearedFields[user.FieldURL]
	return ok
}

// ResetURL resets all changes to the "url" field.
func (m *UserMutation) ResetURL() {
	m.url = nil
	delete(m.clearedFields, user.FieldURL)
}

// SetRaw sets the "raw" field.
func (m *UserMutation) SetRaw(jm json.RawMessage) {
	m.raw = &jm
}

// Raw returns the value of the "raw" field in the mutation.
func (m *UserMutation) Raw() (r json.RawMessage, exists bool) {
	v := m.raw
	if v == nil {
		return
	}
	return *v, true
}

// OldRaw returns the old "raw" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRaw(ctx context.Context) (v json.RawMessage, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRaw is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRaw requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRaw: %w", err)
	}
	return oldValue.Raw, nil
}

// ClearRaw clears the value of the "raw" field.
func (m *UserMutation) ClearRaw() {
	m.raw = nil
	m.clearedFields[user.FieldRaw] = struct{}{}
}

// RawCleared returns if the "raw" field was cleared in this mutation.
func (m *UserMutation) RawCleared() bool {
	_, ok := m.clearedFields[user.FieldRaw]
	return ok
}

// ResetRaw resets all changes to the "raw" field.
func (m *UserMutation) ResetRaw() {
	m.raw = nil
	delete(m.clearedFields, user.FieldRaw)
}

// SetDirs sets the "dirs" field.
func (m *UserMutation) SetDirs(h []http.Dir) {
	m.dirs = &h
}

// Dirs returns the value of the "dirs" field in the mutation.
func (m *UserMutation) Dirs() (r []http.Dir, exists bool) {
	v := m.dirs
	if v == nil {
		return
	}
	return *v, true
}

// OldDirs returns the old "dirs" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldDirs(ctx context.Context) (v []http.Dir, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDirs is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDirs requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDirs: %w", err)
	}
	return oldValue.Dirs, nil
}

// ResetDirs resets all changes to the "dirs" field.
func (m *UserMutation) ResetDirs() {
	m.dirs = nil
}

// SetInts sets the "ints" field.
func (m *UserMutation) SetInts(i []int) {
	m.ints = &i
}

// Ints returns the value of the "ints" field in the mutation.
func (m *UserMutation) Ints() (r []int, exists bool) {
	v := m.ints
	if v == nil {
		return
	}
	return *v, true
}

// OldInts returns the old "ints" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldInts(ctx context.Context) (v []int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldInts is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldInts requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInts: %w", err)
	}
	return oldValue.Ints, nil
}

// ClearInts clears the value of the "ints" field.
func (m *UserMutation) ClearInts() {
	m.ints = nil
	m.clearedFields[user.FieldInts] = struct{}{}
}

// IntsCleared returns if the "ints" field was cleared in this mutation.
func (m *UserMutation) IntsCleared() bool {
	_, ok := m.clearedFields[user.FieldInts]
	return ok
}

// ResetInts resets all changes to the "ints" field.
func (m *UserMutation) ResetInts() {
	m.ints = nil
	delete(m.clearedFields, user.FieldInts)
}

// SetFloats sets the "floats" field.
func (m *UserMutation) SetFloats(f []float64) {
	m.floats = &f
}

// Floats returns the value of the "floats" field in the mutation.
func (m *UserMutation) Floats() (r []float64, exists bool) {
	v := m.floats
	if v == nil {
		return
	}
	return *v, true
}

// OldFloats returns the old "floats" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldFloats(ctx context.Context) (v []float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFloats is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFloats requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFloats: %w", err)
	}
	return oldValue.Floats, nil
}

// ClearFloats clears the value of the "floats" field.
func (m *UserMutation) ClearFloats() {
	m.floats = nil
	m.clearedFields[user.FieldFloats] = struct{}{}
}

// FloatsCleared returns if the "floats" field was cleared in this mutation.
func (m *UserMutation) FloatsCleared() bool {
	_, ok := m.clearedFields[user.FieldFloats]
	return ok
}

// ResetFloats resets all changes to the "floats" field.
func (m *UserMutation) ResetFloats() {
	m.floats = nil
	delete(m.clearedFields, user.FieldFloats)
}

// SetStrings sets the "strings" field.
func (m *UserMutation) SetStrings(s []string) {
	m.strings = &s
}

// Strings returns the value of the "strings" field in the mutation.
func (m *UserMutation) Strings() (r []string, exists bool) {
	v := m.strings
	if v == nil {
		return
	}
	return *v, true
}

// OldStrings returns the old "strings" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldStrings(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStrings is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStrings requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStrings: %w", err)
	}
	return oldValue.Strings, nil
}

// ClearStrings clears the value of the "strings" field.
func (m *UserMutation) ClearStrings() {
	m.strings = nil
	m.clearedFields[user.FieldStrings] = struct{}{}
}

// StringsCleared returns if the "strings" field was cleared in this mutation.
func (m *UserMutation) StringsCleared() bool {
	_, ok := m.clearedFields[user.FieldStrings]
	return ok
}

// ResetStrings resets all changes to the "strings" field.
func (m *UserMutation) ResetStrings() {
	m.strings = nil
	delete(m.clearedFields, user.FieldStrings)
}

// SetAddr sets the "addr" field.
func (m *UserMutation) SetAddr(s schema.Addr) {
	m.addr = &s
}

// Addr returns the value of the "addr" field in the mutation.
func (m *UserMutation) Addr() (r schema.Addr, exists bool) {
	v := m.addr
	if v == nil {
		return
	}
	return *v, true
}

// OldAddr returns the old "addr" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAddr(ctx context.Context) (v schema.Addr, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAddr is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAddr requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddr: %w", err)
	}
	return oldValue.Addr, nil
}

// ClearAddr clears the value of the "addr" field.
func (m *UserMutation) ClearAddr() {
	m.addr = nil
	m.clearedFields[user.FieldAddr] = struct{}{}
}

// AddrCleared returns if the "addr" field was cleared in this mutation.
func (m *UserMutation) AddrCleared() bool {
	_, ok := m.clearedFields[user.FieldAddr]
	return ok
}

// ResetAddr resets all changes to the "addr" field.
func (m *UserMutation) ResetAddr() {
	m.addr = nil
	delete(m.clearedFields, user.FieldAddr)
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.t != nil {
		fields = append(fields, user.FieldT)
	}
	if m.url != nil {
		fields = append(fields, user.FieldURL)
	}
	if m.raw != nil {
		fields = append(fields, user.FieldRaw)
	}
	if m.dirs != nil {
		fields = append(fields, user.FieldDirs)
	}
	if m.ints != nil {
		fields = append(fields, user.FieldInts)
	}
	if m.floats != nil {
		fields = append(fields, user.FieldFloats)
	}
	if m.strings != nil {
		fields = append(fields, user.FieldStrings)
	}
	if m.addr != nil {
		fields = append(fields, user.FieldAddr)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldT:
		return m.T()
	case user.FieldURL:
		return m.URL()
	case user.FieldRaw:
		return m.Raw()
	case user.FieldDirs:
		return m.Dirs()
	case user.FieldInts:
		return m.Ints()
	case user.FieldFloats:
		return m.Floats()
	case user.FieldStrings:
		return m.Strings()
	case user.FieldAddr:
		return m.Addr()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldT:
		return m.OldT(ctx)
	case user.FieldURL:
		return m.OldURL(ctx)
	case user.FieldRaw:
		return m.OldRaw(ctx)
	case user.FieldDirs:
		return m.OldDirs(ctx)
	case user.FieldInts:
		return m.OldInts(ctx)
	case user.FieldFloats:
		return m.OldFloats(ctx)
	case user.FieldStrings:
		return m.OldStrings(ctx)
	case user.FieldAddr:
		return m.OldAddr(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldT:
		v, ok := value.(*schema.T)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetT(v)
		return nil
	case user.FieldURL:
		v, ok := value.(*url.URL)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case user.FieldRaw:
		v, ok := value.(json.RawMessage)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRaw(v)
		return nil
	case user.FieldDirs:
		v, ok := value.([]http.Dir)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDirs(v)
		return nil
	case user.FieldInts:
		v, ok := value.([]int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInts(v)
		return nil
	case user.FieldFloats:
		v, ok := value.([]float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFloats(v)
		return nil
	case user.FieldStrings:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStrings(v)
		return nil
	case user.FieldAddr:
		v, ok := value.(schema.Addr)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddr(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldT) {
		fields = append(fields, user.FieldT)
	}
	if m.FieldCleared(user.FieldURL) {
		fields = append(fields, user.FieldURL)
	}
	if m.FieldCleared(user.FieldRaw) {
		fields = append(fields, user.FieldRaw)
	}
	if m.FieldCleared(user.FieldInts) {
		fields = append(fields, user.FieldInts)
	}
	if m.FieldCleared(user.FieldFloats) {
		fields = append(fields, user.FieldFloats)
	}
	if m.FieldCleared(user.FieldStrings) {
		fields = append(fields, user.FieldStrings)
	}
	if m.FieldCleared(user.FieldAddr) {
		fields = append(fields, user.FieldAddr)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldT:
		m.ClearT()
		return nil
	case user.FieldURL:
		m.ClearURL()
		return nil
	case user.FieldRaw:
		m.ClearRaw()
		return nil
	case user.FieldInts:
		m.ClearInts()
		return nil
	case user.FieldFloats:
		m.ClearFloats()
		return nil
	case user.FieldStrings:
		m.ClearStrings()
		return nil
	case user.FieldAddr:
		m.ClearAddr()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldT:
		m.ResetT()
		return nil
	case user.FieldURL:
		m.ResetURL()
		return nil
	case user.FieldRaw:
		m.ResetRaw()
		return nil
	case user.FieldDirs:
		m.ResetDirs()
		return nil
	case user.FieldInts:
		m.ResetInts()
		return nil
	case user.FieldFloats:
		m.ResetFloats()
		return nil
	case user.FieldStrings:
		m.ResetStrings()
		return nil
	case user.FieldAddr:
		m.ResetAddr()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
