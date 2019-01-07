package nborm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

//OneToOne represent a one point on relation
type OneToOne struct {
	srcDB   string
	srcTab  string
	srcCol  string
	dstDB   string
	dstTab  string
	dstCol  string
	srcValF func() interface{}
}

//NewOneToOne create a OneToOne
func NewOneToOne(srcDB, srcTab, srcCol, dstDB, dstTab, dstCol string, srcValF func() interface{}) *OneToOne {
	return &OneToOne{srcDB, srcTab, srcCol, dstDB, dstTab, dstCol, srcValF}
}

//Query query related table by OneToOne relation
func (oto OneToOne) Query(model table) error {
	if model.DB() != oto.dstDB || model.Tab() != oto.dstTab {
		return fmt.Errorf("nborm.OneToOne.Query() error: required %s.%s supported %s.%s", oto.dstDB, oto.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	row := relationQueryRow(oto, oto.where())
	return scanRow(modAddr, tabInfo, row)
}

func (oto OneToOne) QueryInTx(tx *sql.Tx, model table) error {
	if model.DB() != oto.dstDB || model.Tab() != oto.dstTab {
		return fmt.Errorf("nborm.OneToOne.Query() error: required %s.%s supported %s.%s", oto.dstDB, oto.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	row := relationQueryRowInTx(tx, oto, oto.where())
	return scanRow(modAddr, tabInfo, row)
}

func (oto OneToOne) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", oto.dstDB, oto.dstTab, oto.srcDB, oto.srcTab, oto.srcCol, oto.dstDB,
		oto.dstTab, oto.dstCol)
}

func (oto OneToOne) where() *Where {
	return newWhere(oto.srcDB, oto.srcTab, oto.srcCol, "=", oto.srcValF())
}

func (oto OneToOne) getSrcDB() string {
	return oto.srcDB
}

func (oto OneToOne) getSrcTab() string {
	return oto.srcTab
}

func (oto OneToOne) getDstDB() string {
	return oto.dstDB
}

func (oto OneToOne) getDstTab() string {
	return oto.dstTab
}

//ForeignKey represent a one point many relation
type ForeignKey struct {
	srcDB   string
	srcTab  string
	srcCol  string
	dstDB   string
	dstTab  string
	dstCol  string
	srcValF func() interface{}
}

//NewForeignKey create a ForeignKey
func NewForeignKey(srcDB, srcTab, srcCol, dstDB, dstTab, dstCol string, srcValF func() interface{}) *ForeignKey {
	return &ForeignKey{srcDB, srcTab, srcCol, dstDB, dstTab, dstCol, srcValF}
}

//Query query related table by this relation
func (fk ForeignKey) Query(model table) error {
	if model.DB() != fk.dstDB || model.Tab() != fk.dstTab {
		return fmt.Errorf("nborm.ForeignKey.Query() error: required %s.%s supported %s.%s", fk.dstDB, fk.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where := fk.where()
	row := relationQueryRow(fk, where)
	return scanRow(modAddr, tabInfo, row)
}

func (fk ForeignKey) QueryInTx(tx *sql.Tx, model table) error {
	if model.DB() != fk.dstDB || model.Tab() != fk.dstTab {
		return fmt.Errorf("nborm.ForeignKey.Query() error: required %s.%s supported %s.%s", fk.dstDB, fk.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where := fk.where()
	row := relationQueryRowInTx(tx, fk, where)
	return scanRow(modAddr, tabInfo, row)
}

func (fk ForeignKey) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", fk.dstDB, fk.dstTab, fk.srcDB, fk.srcTab, fk.srcCol, fk.dstDB,
		fk.dstTab, fk.dstCol)
}

func (fk ForeignKey) where() *Where {
	return newWhere(fk.srcDB, fk.srcTab, fk.srcCol, "=", fk.srcValF())
}

func (fk ForeignKey) getSrcDB() string {
	return fk.srcDB
}

func (fk ForeignKey) getSrcTab() string {
	return fk.srcTab
}

func (fk ForeignKey) getDstDB() string {
	return fk.dstDB
}

func (fk ForeignKey) getDstTab() string {
	return fk.dstTab
}

//ReverseForeignKey represent many point one relation
type ReverseForeignKey struct {
	srcDB   string
	srcTab  string
	srcCol  string
	dstDB   string
	dstTab  string
	dstCol  string
	srcValF func() interface{}
}

//NewReverseForeignKey create ReverseForeignKey
func NewReverseForeignKey(srcDB, srcTab, srcCol, dstDB, dstTab, dstCol string, srcValF func() interface{}) *ReverseForeignKey {
	return &ReverseForeignKey{srcDB, srcTab, srcCol, dstDB, dstTab, dstCol, srcValF}
}

//All query all records in related table by this relation
func (rfk ReverseForeignKey) All(slice table, sorter *Sorter, pager *Pager) error {
	if slice.DB() != rfk.dstDB || slice.Tab() != rfk.dstTab {
		return fmt.Errorf("nborm.ReverseForeignKey.All() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab, slice.DB(),
			slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	where := rfk.where()
	rows, err := relationQueryRows(rfk, where, sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

func (rfk ReverseForeignKey) AllInTx(tx *sql.Tx, slice table, sorter *Sorter, pager *Pager) error {
	if slice.DB() != rfk.dstDB || slice.Tab() != rfk.dstTab {
		return fmt.Errorf("nborm.ReverseForeignKey.All() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab, slice.DB(),
			slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	where := rfk.where()
	rows, err := relationQueryRowsInTx(tx, rfk, where, sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

//AllWithFoundRows query all records in related table by this relation and the number of found rows
func (rfk ReverseForeignKey) AllWithFoundRows(slice table, sorter *Sorter, pager *Pager) (int, error) {
	if slice.DB() != rfk.dstDB || slice.Tab() != rfk.dstTab {
		return -1, fmt.Errorf("nborm.ReverseForeignKey.AllWithFoundRows() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab,
			slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, tx, err := relationQueryRowsAndFoundRows(rfk, rfk.where(), sorter, pager)
	if err != nil {
		return -1, err
	}
	err = scanRows(sliceAddr, tabInfo, rows)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return num, nil
}

func (rfk ReverseForeignKey) AllWithFoundRowsInTx(tx *sql.Tx, slice table, sorter *Sorter, pager *Pager) (int, error) {
	if slice.DB() != rfk.dstDB || slice.Tab() != rfk.dstTab {
		return -1, fmt.Errorf("nborm.ReverseForeignKey.AllWithFoundRows() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab,
			slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, tx, err := relationQueryRowsAndFoundRowsInTx(tx, rfk, rfk.where(), sorter, pager)
	if err != nil {
		return -1, err
	}
	err = scanRows(sliceAddr, tabInfo, rows)
	if err != nil {
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	return num, nil
}

//Query query related table by this relation
func (rfk ReverseForeignKey) Query(slice table, where *Where, sorter *Sorter, pager *Pager) error {
	if slice.DB() != rfk.dstDB || slice.Tab() != rfk.dstTab {
		return fmt.Errorf("nborm.ReverseForeignKey.Query() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab,
			slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	where = rfk.where().And(where)
	rows, err := relationQueryRows(rfk, where, sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

func (rfk ReverseForeignKey) QueryInTx(tx *sql.Tx, slice table, where *Where, sorter *Sorter, pager *Pager) error {
	if slice.DB() != rfk.dstDB || slice.Tab() != rfk.dstTab {
		return fmt.Errorf("nborm.ReverseForeignKey.Query() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab,
			slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	where = rfk.where().And(where)
	rows, err := relationQueryRowsInTx(tx, rfk, where, sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

//QueryWithFoundRows query related table by this realtion and number of found rows
func (rfk ReverseForeignKey) QueryWithFoundRows(slice table, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	if slice.DB() != rfk.dstDB || slice.Tab() != rfk.dstTab {
		return -1, fmt.Errorf("nborm.ReverseForeignKey.QueryWithFoundRows() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab,
			slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	where = rfk.where().And(where)
	rows, tx, err := relationQueryRowsAndFoundRows(rfk, where, sorter, pager)
	if err != nil {
		return -1, err
	}
	err = scanRows(sliceAddr, tabInfo, rows)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return num, nil
}

func (rfk ReverseForeignKey) QueryWithFoundRowsInTx(tx *sql.Tx, slice table, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	if slice.DB() != rfk.dstDB || slice.Tab() != rfk.dstTab {
		return -1, fmt.Errorf("nborm.ReverseForeignKey.QueryWithFoundRows() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab,
			slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	where = rfk.where().And(where)
	rows, tx, err := relationQueryRowsAndFoundRowsInTx(tx, rfk, where, sorter, pager)
	if err != nil {
		return -1, err
	}
	err = scanRows(sliceAddr, tabInfo, rows)
	if err != nil {
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		return -1, err
	}
	return num, nil
}

//AddOne add a related model by reverse foreign key relation
func (rfk ReverseForeignKey) AddOne(model table) error {
	if model.DB() != rfk.dstDB || model.Tab() != rfk.dstTab {
		return fmt.Errorf("nborm.ReverseForeignKey.AddOne() error: database or table not match (%s.%s), want %s.%s", model.DB(), model.Tab(),
			rfk.dstDB, rfk.dstTab)
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	dstField := getFieldByName(modAddr, rfk.dstCol, tabInfo)
	if dstField.IsValid() {
		return fmt.Errorf("nborm.ReverseForeignKey.AddOne() error: destination field already set (%s.%s.%s)", model.DB(), model.Tab(),
			dstField.columnName)
	}
	dstField.setVal(rfk.srcValF(), false)
	lid, err := insert(modAddr, tabInfo)
	if err != nil {
		return err
	}
	setInc(modAddr, tabInfo, lid)
	setSync(modAddr, tabInfo)
	return nil
}

func (rfk ReverseForeignKey) AddOneInTx(tx *sql.Tx, model table) error {
	if model.DB() != rfk.dstDB || model.Tab() != rfk.dstTab {
		return fmt.Errorf("nborm.ReverseForeignKey.AddOne() error: database or table not match (%s.%s), want %s.%s", model.DB(), model.Tab(),
			rfk.dstDB, rfk.dstTab)
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	dstField := getFieldByName(modAddr, rfk.dstCol, tabInfo)
	if dstField.IsValid() {
		return fmt.Errorf("nborm.ReverseForeignKey.AddOne() error: destination field already set (%s.%s.%s)", model.DB(), model.Tab(),
			dstField.columnName)
	}
	dstField.setVal(rfk.srcValF(), false)
	lid, err := insertInTx(tx, modAddr, tabInfo)
	if err != nil {
		return err
	}
	setInc(modAddr, tabInfo, lid)
	setSync(modAddr, tabInfo)
	return nil
}

//AddMul add many model by reverse foreign key relation
func (rfk ReverseForeignKey) AddMul(slice table) error {
	if slice.DB() != rfk.dstDB || slice.Tab() != rfk.dstTab {
		return fmt.Errorf("nborm.ReverseForeignKey.AddMul() error: database or table not match (%s.%s), want %s.%s", slice.DB(), slice.Tab(),
			rfk.dstDB, rfk.dstTab)
	}
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, modAddr uintptr) error {
		dstField := getFieldByName(modAddr, rfk.dstCol, tabInfo)
		if dstField.IsValid() {
			return fmt.Errorf("nborm.ReverseForeignKey.AddMul() error: destination field already set (%s.%s.%s)", slice.DB(), slice.Tab(),
				dstField.columnName)
		}
		dstField.setVal(rfk.srcValF(), false)
		lid, err := insertContext(ctx, modAddr, tabInfo)
		if err != nil {
			return err
		}
		setInc(modAddr, tabInfo, lid)
		setSync(modAddr, tabInfo)
		return nil
	})
}

func (rfk ReverseForeignKey) AddMulInTx(tx *sql.Tx, slice table) error {
	if slice.DB() != rfk.dstDB || slice.Tab() != rfk.dstTab {
		return fmt.Errorf("nborm.ReverseForeignKey.AddMul() error: database or table not match (%s.%s), want %s.%s", slice.DB(), slice.Tab(),
			rfk.dstDB, rfk.dstTab)
	}
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, modAddr uintptr) error {
		dstField := getFieldByName(modAddr, rfk.dstCol, tabInfo)
		if dstField.IsValid() {
			return fmt.Errorf("nborm.ReverseForeignKey.AddMul() error: destination field already set (%s.%s.%s)", slice.DB(), slice.Tab(),
				dstField.columnName)
		}
		dstField.setVal(rfk.srcValF(), false)
		lid, err := insertContextInTx(tx, ctx, modAddr, tabInfo)
		if err != nil {
			return err
		}
		setInc(modAddr, tabInfo, lid)
		setSync(modAddr, tabInfo)
		return nil
	})
}

func (rfk ReverseForeignKey) Count(where *Where) (int, error) {
	return relationCount(rfk, where)
}

func (rfk ReverseForeignKey) CountInTx(tx *sql.Tx, where *Where) (int, error) {
	return relationCountInTx(tx, rfk, where)
}

func (rfk ReverseForeignKey) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", rfk.dstDB, rfk.dstTab, rfk.srcDB, rfk.srcTab, rfk.srcCol, rfk.dstDB,
		rfk.dstTab, rfk.dstCol)
}

func (rfk ReverseForeignKey) where() *Where {
	return newWhere(rfk.srcDB, rfk.srcTab, rfk.srcCol, "=", rfk.srcValF())
}

func (rfk ReverseForeignKey) getSrcDB() string {
	return rfk.srcDB
}

func (rfk ReverseForeignKey) getSrcTab() string {
	return rfk.srcTab
}

func (rfk ReverseForeignKey) getDstDB() string {
	return rfk.dstDB
}

func (rfk ReverseForeignKey) getDstTab() string {
	return rfk.dstTab
}

//ManyToMany represent many point many relation
type ManyToMany struct {
	srcDB       string
	srcTab      string
	srcCol      string
	midDB       string
	midTab      string
	midLeftCol  string
	midRightCol string
	dstDB       string
	dstTab      string
	dstCol      string
	srcValF     func() interface{}
}

//NewManyToMany create ManyToMany
func NewManyToMany(srcDB, srcTab, srcCol, midDB, midTab, midLeftCol, midRightCol, dstDB, dstTab, dstCol string, srcValF func() interface{}) *ManyToMany {
	return &ManyToMany{srcDB, srcTab, srcCol, midDB, midTab, midLeftCol, midRightCol, dstDB, dstTab, dstCol, srcValF}
}

//All query all records in related table by this relation
func (mtm ManyToMany) All(slice table, sorter *Sorter, pager *Pager) error {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.All() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, err := relationQueryRows(mtm, mtm.where(), sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

func (mtm ManyToMany) AllInTx(tx *sql.Tx, slice table, sorter *Sorter, pager *Pager) error {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.All() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, err := relationQueryRowsInTx(tx, mtm, mtm.where(), sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

//AllWithFoundRows query all records in related table and number of found rows by this relation
func (mtm ManyToMany) AllWithFoundRows(slice table, sorter *Sorter, pager *Pager) (int, error) {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return -1, fmt.Errorf("nborm.ManyToMany.AllWithFoundRows() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(),
			slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, tx, err := relationQueryRowsAndFoundRows(mtm, mtm.where(), sorter, pager)
	if err != nil {
		return -1, err
	}
	err = scanRows(sliceAddr, tabInfo, rows)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return num, nil
}

func (mtm ManyToMany) AllWithFoundRowsInTx(tx *sql.Tx, slice table, sorter *Sorter, pager *Pager) (int, error) {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return -1, fmt.Errorf("nborm.ManyToMany.AllWithFoundRows() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(),
			slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, tx, err := relationQueryRowsAndFoundRowsInTx(tx, mtm, mtm.where(), sorter, pager)
	if err != nil {
		return -1, err
	}
	err = scanRows(sliceAddr, tabInfo, rows)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return num, nil
}

//Query query records in related table by this relation
func (mtm ManyToMany) Query(slice table, where *Where, sorter *Sorter, pager *Pager) error {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.Query() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	where = mtm.where().And(where)
	rows, err := relationQueryRows(mtm, where, sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

func (mtm ManyToMany) QueryInTx(tx *sql.Tx, slice table, where *Where, sorter *Sorter, pager *Pager) error {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.Query() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	where = mtm.where().And(where)
	rows, err := relationQueryRowsInTx(tx, mtm, where, sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

//QueryWithFoundRows query records in related table and number of found rows by this relation
func (mtm ManyToMany) QueryWithFoundRows(slice table, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return -1, fmt.Errorf("nborm.ManyToMany.QueryWithFoundRows() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab,
			slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	where = mtm.where().And(where)
	rows, tx, err := relationQueryRowsAndFoundRows(mtm, where, sorter, pager)
	if err != nil {
		return -1, err
	}
	if err := scanRows(sliceAddr, tabInfo, rows); err != nil {
		tx.Rollback()
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return num, nil
}

func (mtm ManyToMany) QueryWithFoundRowsInTx(tx *sql.Tx, slice table, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return -1, fmt.Errorf("nborm.ManyToMany.QueryWithFoundRows() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab,
			slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	where = mtm.where().And(where)
	rows, tx, err := relationQueryRowsAndFoundRowsInTx(tx, mtm, where, sorter, pager)
	if err != nil {
		return -1, err
	}
	if err := scanRows(sliceAddr, tabInfo, rows); err != nil {
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		return -1, err
	}
	return num, nil
}

//AddOne add a relation record to middle table
func (mtm ManyToMany) AddOne(model table) error {
	if model.DB() != mtm.dstDB || model.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.AddOne() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (?, ?)", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
	if _, err := getConn(mtm.midDB).Exec(stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
		return err
	}
	return nil
}

func (mtm ManyToMany) AddOneInTx(tx *sql.Tx, model table) error {
	if model.DB() != mtm.dstDB || model.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.AddOne() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (?, ?)", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
	if _, err := tx.Exec(stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
		return err
	}
	return nil
}

//AddMul add a relation record to middle table
func (mtm ManyToMany) AddMul(slice table) error {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.AddMul() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, modAddr uintptr) error {
		if !getSync(modAddr, tabInfo) {
			lastInsertID, err := insertContext(ctx, modAddr, tabInfo)
			if err != nil {
				return err
			}
			setInc(modAddr, tabInfo, lastInsertID)
			setSync(modAddr, tabInfo)
		}
		stmt := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (?, ?)", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
		if _, err := getConn(mtm.midDB).ExecContext(ctx, stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
			return err
		}
		return nil
	})
}

func (mtm ManyToMany) AddMulInTx(tx *sql.Tx, slice table) error {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.AddMul() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, modAddr uintptr) error {
		stmt := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (?, ?)", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
		if _, err := tx.ExecContext(ctx, stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
			return err
		}
		return nil
	})
}

func (mtm ManyToMany) InsertOne(model table) error {
	if model.DB() != mtm.dstDB || model.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.InsertOne() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	if getSync(modAddr, tabInfo) {
		return errors.New("nborm.ManyToMany.InsertOne() error: model already exists. If you want to add a exist model, please use ManyToMany.AddOne()")
	}
	lid, err := insert(modAddr, tabInfo)
	if err != nil {
		return err
	}
	setInc(modAddr, tabInfo, lid)
	setSync(modAddr, tabInfo)
	dstField := getFieldByName(modAddr, mtm.dstCol, tabInfo)
	if !dstField.IsValid() {
		return fmt.Errorf("nborm.ManyToMany.InsertOne() error: destination field is invalid")
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (?, ?)", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
	if _, err := getConn(mtm.midDB).Exec(stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
		return err
	}
	return nil
}

func (mtm ManyToMany) InsertOneInTx(tx *sql.Tx, model table) error {
	if model.DB() != mtm.dstDB || model.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.InsertOne() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	if getSync(modAddr, tabInfo) {
		return errors.New("nborm.ManyToMany.InsertOne() error: model already exists. If you want to add a exist model, please use ManyToMany.AddOne()")
	}
	lid, err := insertInTx(tx, modAddr, tabInfo)
	if err != nil {
		return err
	}
	setInc(modAddr, tabInfo, lid)
	setSync(modAddr, tabInfo)
	dstField := getFieldByName(modAddr, mtm.dstCol, tabInfo)
	if !dstField.IsValid() {
		return fmt.Errorf("nborm.ManyToMany.InsertOne() error: destination field is invalid")
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (?, ?)", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
	if _, err := tx.Exec(stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
		return err
	}
	return nil
}

func (mtm ManyToMany) InsertMul(slice table) error {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.InsertOne() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, modAddr uintptr) error {
		if getSync(modAddr, tabInfo) {
			return errors.New("nborm.ManyToMany.InsertMul() error: model already exists. If you want to add a exist model, please use ManyToMany.AddOne()")
		}
		lid, err := insert(modAddr, tabInfo)
		if err != nil {
			return err
		}
		setInc(modAddr, tabInfo, lid)
		setSync(modAddr, tabInfo)
		dstField := getFieldByName(modAddr, mtm.dstCol, tabInfo)
		if !dstField.IsValid() {
			return fmt.Errorf("nborm.ManyToMany.InsertOne() error: destination field is invalid")
		}
		stmt := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (?, ?)", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
		if _, err := getConn(mtm.midDB).Exec(stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
			return err
		}
		return nil
	})
}

func (mtm ManyToMany) InsertMulInTx(tx *sql.Tx, slice table) error {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.InsertOne() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, modAddr uintptr) error {
		if getSync(modAddr, tabInfo) {
			return errors.New("nborm.ManyToMany.InsertMul() error: model already exists. If you want to add a exist model, please use ManyToMany.AddOne()")
		}
		lid, err := insertInTx(tx, modAddr, tabInfo)
		if err != nil {
			return err
		}
		setInc(modAddr, tabInfo, lid)
		setSync(modAddr, tabInfo)
		dstField := getFieldByName(modAddr, mtm.dstCol, tabInfo)
		if !dstField.IsValid() {
			return fmt.Errorf("nborm.ManyToMany.InsertOne() error: destination field is invalid")
		}
		stmt := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (?, ?)", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
		if _, err := tx.Exec(stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
			return err
		}
		return nil
	})
}

//Remove remove a record from middle table
func (mtm ManyToMany) RemoveOne(model table) error {
	if model.DB() != mtm.dstDB || model.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.RemoveOne() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	stmt := fmt.Sprintf("DELETE FROM %s.%s WHERE %s = ? AND %s = ?", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
	if _, err := getConn(mtm.midDB).Exec(stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
		return err
	}
	return nil
}

func (mtm ManyToMany) RemoveOneInTx(tx *sql.Tx, model table) error {
	if model.DB() != mtm.dstDB || model.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.RemoveOne() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	stmt := fmt.Sprintf("DELETE FROM %s.%s WHERE %s = ? AND %s = ?", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
	if _, err := tx.Exec(stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
		return err
	}
	return nil
}

func (mtm ManyToMany) RemoveMul(slice table) error {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.RemoveMul() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, modAddr uintptr) error {
		stmt := fmt.Sprintf("DELETE FROM %s.%s WHERE %s = ? AND %s = ?", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
		if _, err := getConn(mtm.midDB).ExecContext(ctx, stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
			return err
		}
		return nil
	})
}

func (mtm ManyToMany) RemoveMulInTx(tx *sql.Tx, slice table) error {
	if slice.DB() != mtm.dstDB || slice.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.RemoveMul() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, slice.DB(), slice.Tab())
	}
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, modAddr uintptr) error {
		stmt := fmt.Sprintf("DELETE FROM %s.%s WHERE %s = ? AND %s = ?", mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol)
		if _, err := tx.ExecContext(ctx, stmt, mtm.srcValF(), getFieldByName(modAddr, mtm.dstCol, tabInfo).value()); err != nil {
			return err
		}
		return nil
	})
}

func (mtm ManyToMany) BulkRemove(where *Where) error {
	if where != nil && (where.db != mtm.dstDB || where.table != mtm.dstTab) {
		return fmt.Errorf("nborm.ManyToMany.BulkRemove() error: where destination table is %s.%s, want %s.%s", where.db, where.table, mtm.dstDB,
			mtm.dstTab)
	}
	where = mtm.where().And(where)
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("DELETE %s.%s FROM %s.%s %s %s", mtm.midDB, mtm.midTab, mtm.srcDB, mtm.srcTab, mtm.joinClause(), whereClause)
	_, err := getConn(mtm.midDB).Exec(stmt, whereValues...)
	return err
}

func (mtm ManyToMany) BulkRemoveInTx(tx *sql.Tx, where *Where) error {
	if where != nil && (where.db != mtm.dstDB || where.table != mtm.dstTab) {
		return fmt.Errorf("nborm.ManyToMany.BulkRemove() error: where destination table is %s.%s, want %s.%s", where.db, where.table, mtm.dstDB,
			mtm.dstTab)
	}
	where = mtm.where().And(where)
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("DELETE %s.%s FROM %s.%s %s %s", mtm.midDB, mtm.midTab, mtm.srcDB, mtm.srcTab, mtm.joinClause(), whereClause)
	_, err := tx.Exec(stmt, whereValues...)
	return err
}

func (mtm ManyToMany) Count(where *Where) (int, error) {
	return relationCount(mtm, where)
}

func (mtm ManyToMany) CountInTx(tx *sql.Tx, where *Where) (int, error) {
	return relationCountInTx(tx, mtm, where)
}

func (mtm ManyToMany) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s", mtm.midDB, mtm.midTab, mtm.srcDB, mtm.srcTab,
		mtm.srcCol, mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.dstDB, mtm.dstTab, mtm.midDB, mtm.midTab, mtm.midRightCol, mtm.dstDB,
		mtm.dstTab, mtm.dstCol)
}

func (mtm ManyToMany) where() *Where {
	return newWhere(mtm.srcDB, mtm.srcTab, mtm.srcCol, "=", mtm.srcValF())
}

func (mtm ManyToMany) getSrcDB() string {
	return mtm.srcDB
}

func (mtm ManyToMany) getSrcTab() string {
	return mtm.srcTab
}

func (mtm ManyToMany) getDstDB() string {
	return mtm.dstDB
}

func (mtm ManyToMany) getDstTab() string {
	return mtm.dstTab
}
