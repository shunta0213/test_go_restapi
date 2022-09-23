// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrations)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsInsert)
	t.Run("SchemaMigrations", testSchemaMigrationsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
