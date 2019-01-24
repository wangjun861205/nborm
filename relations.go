package nborm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

//OneToOne represent a one point on relation
type OneToOne struct {
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

func (oto OneToOne) Set(model table) error {
	if model.DB() != oto.dstDB || model.Tab() != oto.dstTab {
		return fmt.Errorf("nborm.OneToOne.Set() error: required %s.%s supported %s.%s", oto.dstDB, oto.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	addr := getTabAddr(model)
	if getSync(addr, tabInfo) {
		err := insertMiddleTable(oto, addr, tabInfo)
		if err != nil {
			return err
		}
	} else {
		tx, err := Begin()
		if err != nil {
			return err
		}
		lastInsertID, err := insertInTx(tx, addr, tabInfo)
		if err != nil {
			tx.Rollback()
			return err
		}
		setInc(addr, tabInfo, lastInsertID)
		err = insertMiddleTable(oto, addr, tabInfo)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}

func (oto OneToOne) Unset(model table) error {
	if model.DB() != oto.dstDB || model.Tab() != oto.dstTab {
		return fmt.Errorf("nborm.OneToOne.Unset() error: required %s.%s supported %s.%s", oto.dstDB, oto.dstTab, model.DB(), model.Tab())
	}
	tabInfo := getTabInfo(model)
	addr := getTabAddr(model)
	return deleteMiddleTable(oto, addr, tabInfo)
}

func (oto OneToOne) where() *Where {
	return newWhere(oto.srcDB, oto.srcTab, oto.srcCol, "=", oto.srcValF())
}

func (oto OneToOne) getSrcDB() string {
	return wrap(oto.srcDB)
}

func (oto OneToOne) getRawSrcDB() string {
	return oto.srcDB
}

func (oto OneToOne) getSrcTab() string {
	return wrap(oto.srcTab)
}

func (oto OneToOne) getRawSrcTab() string {
	return oto.srcTab
}

func (oto OneToOne) getSrcCol() string {
	return wrap(oto.srcCol)
}

func (oto OneToOne) getRawSrcCol() string {
	return oto.srcCol
}

func (oto OneToOne) getMidDB() string {
	return wrap(oto.midDB)
}

func (oto OneToOne) getRawMidDB() string {
	return oto.midDB
}

func (oto OneToOne) getMidTab() string {
	return wrap(oto.midTab)
}

func (oto OneToOne) getRawMidTab() string {
	return oto.midTab
}

func (oto OneToOne) getMidLeftCol() string {
	return wrap(oto.midLeftCol)
}

func (oto OneToOne) getRawMidLeftCol() string {
	return oto.midLeftCol
}

func (oto OneToOne) getMidRightCol() string {
	return wrap(oto.midRightCol)
}

func (oto OneToOne) getRawMidRightCol() string {
	return oto.midRightCol
}

func (oto OneToOne) getDstDB() string {
	return wrap(oto.dstDB)
}

func (oto OneToOne) getRawDstDB() string {
	return oto.dstDB
}

func (oto OneToOne) getDstTab() string {
	return wrap(oto.dstTab)
}

func (oto OneToOne) getRawDstTab() string {
	return oto.dstTab
}

func (oto OneToOne) getDstCol() string {
	return wrap(oto.dstCol)
}

func (oto OneToOne) getRawDstCol() string {
	return oto.dstCol
}

func (oto OneToOne) getFullSrcTab() string {
	return fmt.Sprintf("%s.%s", oto.getSrcDB(), oto.getSrcTab())
}

func (oto OneToOne) getFullDstTab() string {
	return fmt.Sprintf("%s.%s", oto.getDstDB(), oto.getDstTab())
}

func (oto OneToOne) getFullMidTab() string {
	return fmt.Sprintf("%s.%s", oto.getMidDB(), oto.getMidTab())
}

func (oto OneToOne) getFullSrcCol() string {
	return fmt.Sprintf("%s.%s.%s", oto.getSrcDB(), oto.getSrcTab(), oto.getSrcCol())
}

func (oto OneToOne) getFullDstCol() string {
	return fmt.Sprintf("%s.%s.%s", oto.getDstDB(), oto.getDstTab(), oto.getDstCol())
}

func (oto OneToOne) getFullMidLeftCol() string {
	return fmt.Sprintf("%s.%s.%s", oto.getMidDB(), oto.getMidTab(), oto.getMidLeftCol())
}

func (oto OneToOne) getFullMidRightCol() string {
	return fmt.Sprintf("%s.%s.%s", oto.getMidDB(), oto.getMidTab(), oto.getMidRightCol())
}

func (oto OneToOne) getSrcVal() interface{} {
	return oto.srcValF()
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

func (fk ForeignKey) where() *Where {
	return newWhere(fk.srcDB, fk.srcTab, fk.srcCol, "=", fk.srcValF())
}

func (fk ForeignKey) getSrcDB() string {
	return wrap(fk.srcDB)
}

func (fk ForeignKey) getRawSrcDB() string {
	return fk.srcDB
}

func (fk ForeignKey) getSrcTab() string {
	return wrap(fk.srcTab)
}

func (fk ForeignKey) getRawSrcTab() string {
	return fk.srcTab
}

func (fk ForeignKey) getSrcCol() string {
	return wrap(fk.srcCol)
}

func (fk ForeignKey) getRawSrcCol() string {
	return fk.srcCol
}

func (fk ForeignKey) getDstDB() string {
	return wrap(fk.dstDB)
}

func (fk ForeignKey) getRawDstDB() string {
	return fk.dstDB
}

func (fk ForeignKey) getDstTab() string {
	return wrap(fk.dstTab)
}

func (fk ForeignKey) getRawDstTab() string {
	return fk.dstTab
}

func (fk ForeignKey) getDstCol() string {
	return wrap(fk.dstCol)
}

func (fk ForeignKey) getRawDstCol() string {
	return fk.dstCol
}

func (fk ForeignKey) getFullSrcTab() string {
	return fmt.Sprintf("%s.%s", fk.getSrcDB(), fk.getSrcTab())
}

func (fk ForeignKey) getFullDstTab() string {
	return fmt.Sprintf("%s.%s", fk.getDstDB(), fk.getDstTab())
}

func (fk ForeignKey) getFullSrcCol() string {
	return fmt.Sprintf("%s.%s.%s", fk.getSrcDB(), fk.getSrcTab(), fk.getSrcCol())
}

func (fk ForeignKey) getFullDstCol() string {
	return fmt.Sprintf("%s.%s.%s", fk.getDstDB(), fk.getDstTab(), fk.getDstCol())
}

func (fk ForeignKey) getSrcVal() interface{} {
	return fk.srcValF()
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
	// where = rfk.where().And(where)
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
	// where = rfk.where().And(where)
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
	// where = rfk.where().And(where)
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
	// where = rfk.where().And(where)
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
	if dstField.IsValid() && !dstField.IsNull() && dstField.value() != rfk.srcValF() {
		return fmt.Errorf("nborm.ReverseForeignKey.AddOne() error: relation field values not match src(%s.%s.%s: %v), dst(%s: %v)", wrap(rfk.srcDB),
			wrap(rfk.srcTab), wrap(rfk.srcCol), rfk.srcValF(), dstField.fullColName(), dstField.value())
	} else {
		dstField.setVal(rfk.srcValF())
	}
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
		return fmt.Errorf("nborm.ReverseForeignKey.AddOneInTx() error: database or table not match (%s.%s), want %s.%s", model.DB(), model.Tab(),
			rfk.dstDB, rfk.dstTab)
	}
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	dstField := getFieldByName(modAddr, rfk.dstCol, tabInfo)
	if dstField.IsValid() && !dstField.IsNull() && dstField.value() != rfk.srcValF() {
		return fmt.Errorf("nborm.ReverseForeignKey.AddOneInTx() error: relation field values not match src(%s.%s.%s: %v), dst(%s: %v)", wrap(rfk.srcDB),
			wrap(rfk.srcTab), wrap(rfk.srcCol), rfk.srcValF(), dstField.fullColName(), dstField.value())
	} else {
		dstField.setVal(rfk.srcValF())
	}
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
		if dstField.IsValid() && !dstField.IsNull() && dstField.value() != rfk.srcValF() {
			return fmt.Errorf("nborm.ReverseForeignKey.AddMul() error: relation field values not match src(%s.%s.%s: %v), dst(%s: %v)", wrap(rfk.srcDB),
				wrap(rfk.srcTab), wrap(rfk.srcCol), rfk.srcValF(), dstField.fullColName(), dstField.value())
		} else {
			dstField.setVal(rfk.srcValF())
		}
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
		return fmt.Errorf("nborm.ReverseForeignKey.AddMulInTx() error: database or table not match (%s.%s), want %s.%s", slice.DB(), slice.Tab(),
			rfk.dstDB, rfk.dstTab)
	}
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, modAddr uintptr) error {
		dstField := getFieldByName(modAddr, rfk.dstCol, tabInfo)
		if dstField.IsValid() && !dstField.IsNull() && dstField.value() != rfk.srcValF() {
			return fmt.Errorf("nborm.ReverseForeignKey.AddMulInTx() error: relation field values not match src(%s.%s.%s: %v), dst(%s: %v)", wrap(rfk.srcDB),
				wrap(rfk.srcTab), wrap(rfk.srcCol), rfk.srcValF(), dstField.fullColName(), dstField.value())
		} else {
			dstField.setVal(rfk.srcValF())
		}
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

// func (rfk ReverseForeignKey) joinClause() string {
// 	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", rfk.dstDB, rfk.dstTab, rfk.srcDB, rfk.srcTab, rfk.srcCol, rfk.dstDB,
// 		rfk.dstTab, rfk.dstCol)
// }

func (rfk ReverseForeignKey) where() *Where {
	return newWhere(rfk.srcDB, rfk.srcTab, rfk.srcCol, "=", rfk.srcValF())
}

func (rfk ReverseForeignKey) getSrcDB() string {
	return wrap(rfk.srcDB)
}

func (rfk ReverseForeignKey) getRawSrcDB() string {
	return rfk.srcDB
}

func (rfk ReverseForeignKey) getSrcTab() string {
	return wrap(rfk.srcTab)
}

func (rfk ReverseForeignKey) getRawSrcTab() string {
	return rfk.srcTab
}

func (rfk ReverseForeignKey) getSrcCol() string {
	return wrap(rfk.srcCol)
}

func (rfk ReverseForeignKey) getRawSrcCol() string {
	return rfk.srcCol
}

func (rfk ReverseForeignKey) getDstDB() string {
	return wrap(rfk.dstDB)
}

func (rfk ReverseForeignKey) getRawDstDB() string {
	return rfk.dstDB
}

func (rfk ReverseForeignKey) getDstTab() string {
	return wrap(rfk.dstTab)
}

func (rfk ReverseForeignKey) getRawDstTab() string {
	return rfk.dstTab
}

func (rfk ReverseForeignKey) getDstCol() string {
	return wrap(rfk.dstCol)
}

func (rfk ReverseForeignKey) getRawDstCol() string {
	return rfk.dstCol
}

func (rfk ReverseForeignKey) getFullSrcTab() string {
	return fmt.Sprintf("%s.%s", rfk.getSrcDB(), rfk.getSrcTab())
}

func (rfk ReverseForeignKey) getFullDstTab() string {
	return fmt.Sprintf("%s.%s", rfk.getDstDB(), rfk.getDstTab())
}

func (rfk ReverseForeignKey) getFullSrcCol() string {
	return fmt.Sprintf("%s.%s.%s", rfk.getSrcDB(), rfk.getSrcTab(), rfk.getSrcCol())
}

func (rfk ReverseForeignKey) getFullDstCol() string {
	return fmt.Sprintf("%s.%s.%s", rfk.getDstDB(), rfk.getDstTab(), rfk.getDstCol())
}

func (rfk ReverseForeignKey) getSrcVal() interface{} {
	return rfk.srcValF()
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
	// where = mtm.where().And(where)
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
	// where = mtm.where().And(where)
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
	// where = mtm.where().And(where)
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
	// where = mtm.where().And(where)
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
	joinClause, _ := genJoinClause(mtm)
	stmt := fmt.Sprintf("DELETE %s.%s FROM %s %s", mtm.midDB, mtm.midTab, joinClause, whereClause)
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
	joinClause, _ := genJoinClause(mtm)
	stmt := fmt.Sprintf("DELETE %s.%s FROM %s %s", mtm.midDB, mtm.midTab, joinClause, whereClause)
	_, err := tx.Exec(stmt, whereValues...)
	return err
}

func (mtm ManyToMany) Count(where *Where) (int, error) {
	return relationCount(mtm, where)
}

func (mtm ManyToMany) CountInTx(tx *sql.Tx, where *Where) (int, error) {
	return relationCountInTx(tx, mtm, where)
}

func (mtm ManyToMany) where() *Where {
	return newWhere(mtm.srcDB, mtm.srcTab, mtm.srcCol, "=", mtm.srcValF())
}

func (mtm ManyToMany) getSrcDB() string {
	return wrap(mtm.srcDB)
}

func (mtm ManyToMany) getRawSrcDB() string {
	return mtm.srcDB
}

func (mtm ManyToMany) getSrcTab() string {
	return wrap(mtm.srcTab)
}

func (mtm ManyToMany) getRawSrcTab() string {
	return mtm.srcTab
}

func (mtm ManyToMany) getSrcCol() string {
	return wrap(mtm.srcCol)
}

func (mtm ManyToMany) getRawSrcCol() string {
	return mtm.srcCol
}

func (mtm ManyToMany) getMidDB() string {
	return wrap(mtm.midDB)
}

func (mtm ManyToMany) getRawMidDB() string {
	return mtm.midDB
}

func (mtm ManyToMany) getMidTab() string {
	return wrap(mtm.midTab)
}

func (mtm ManyToMany) getRawMidTab() string {
	return mtm.midTab
}

func (mtm ManyToMany) getMidLeftCol() string {
	return wrap(mtm.midLeftCol)
}

func (mtm ManyToMany) getRawMidLeftCol() string {
	return mtm.midLeftCol
}

func (mtm ManyToMany) getMidRightCol() string {
	return wrap(mtm.midRightCol)
}

func (mtm ManyToMany) getRawMidRightCol() string {
	return mtm.midRightCol
}

func (mtm ManyToMany) getDstDB() string {
	return wrap(mtm.dstDB)
}

func (mtm ManyToMany) getRawDstDB() string {
	return mtm.dstDB
}

func (mtm ManyToMany) getDstTab() string {
	return wrap(mtm.dstTab)
}

func (mtm ManyToMany) getRawDstTab() string {
	return mtm.dstTab
}

func (mtm ManyToMany) getDstCol() string {
	return wrap(mtm.dstCol)
}

func (mtm ManyToMany) getRawDstCol() string {
	return mtm.dstCol
}

func (mtm ManyToMany) getFullSrcTab() string {
	return fmt.Sprintf("%s.%s", mtm.getSrcDB(), mtm.getSrcTab())
}

func (mtm ManyToMany) getFullDstTab() string {
	return fmt.Sprintf("%s.%s", mtm.getDstDB(), mtm.getDstTab())
}

func (mtm ManyToMany) getFullMidTab() string {
	return fmt.Sprintf("%s.%s", mtm.getMidDB(), mtm.getMidTab())
}

func (mtm ManyToMany) getFullSrcCol() string {
	return fmt.Sprintf("%s.%s.%s", mtm.getSrcDB(), mtm.getSrcTab(), mtm.getSrcCol())
}

func (mtm ManyToMany) getFullDstCol() string {
	return fmt.Sprintf("%s.%s.%s", mtm.getDstDB(), mtm.getDstTab(), mtm.getDstCol())
}

func (mtm ManyToMany) getFullMidLeftCol() string {
	return fmt.Sprintf("%s.%s.%s", mtm.getMidDB(), mtm.getMidTab(), mtm.getMidLeftCol())
}

func (mtm ManyToMany) getFullMidRightCol() string {
	return fmt.Sprintf("%s.%s.%s", mtm.getMidDB(), mtm.getMidTab(), mtm.getMidRightCol())
}

func (mtm ManyToMany) getSrcVal() interface{} {
	return mtm.srcValF()
}

// func genJoinClause(relations ...relation) (string, error) {
// 	if len(relations) == 0 {
// 		panic(errors.New("nborm.genJoinClause() error: no relation"))
// 	}
// 	var firstRel relation
// 	firstRel, relations = relations[0], relations[1:]
// 	srcTab, dstTab := firstRel.getFullSrcTab(), firstRel.getFullDstTab()
// 	chainedTables := map[string]bool{srcTab: true, dstTab: true}
// 	var currentJoinStr string
// 	switch r := firstRel.(type) {
// 	case OneToOne:
// 		currentJoinStr = fmt.Sprintf("%s JOIN %s ON %s = %s JOIN %s ON %s = %s", r.getFullSrcTab(), r.getFullMidTab(),
// 			r.getFullSrcCol(), r.getFullMidLeftCol(), r.getFullDstTab(), r.getFullMidRightCol(), r.getFullDstCol())
// 	case ManyToMany:
// 		currentJoinStr = fmt.Sprintf("%s JOIN %s ON %s = %s JOIN %s ON %s = %s", r.getFullSrcTab(), r.getFullMidTab(),
// 			r.getFullSrcCol(), r.getFullMidLeftCol(), r.getFullDstTab(), r.getFullMidRightCol(), r.getFullDstCol())
// 	case ForeignKey:
// 		currentJoinStr = fmt.Sprintf("%s JOIN %s ON %s = %s", r.getFullSrcTab(), r.getFullDstTab(),
// 			r.getFullSrcCol(), r.getFullDstCol())
// 	case ReverseForeignKey:
// 		currentJoinStr = fmt.Sprintf("%s JOIN %s ON %s = %s", r.getFullSrcTab(), r.getFullDstTab(),
// 			r.getFullSrcCol(), r.getFullDstCol())
// 	default:
// 		return "", errors.New("nborm.rGenJoinClause() error: unknown relation type")
// 	}
// 	if len(relations) > 0 {
// 		nextJoinStr, err := rGenJoinClause(chainedTables, relations...)
// 		if err != nil {
// 			return "", err
// 		}
// 		return currentJoinStr + " " + nextJoinStr, nil
// 	}
// 	return currentJoinStr, nil
// }

// func rGenJoinClause(chainedTables map[string]bool, relations ...relation) (string, error) {
// 	for i, rel := range relations {
// 		srcTab := rel.getFullSrcTab()
// 		dstTab := rel.getFullDstTab()
// 		switch {
// 		case chainedTables[srcTab]:
// 			switch r := rel.(type) {
// 			case OneToOne:
// 				currentJoinStr := fmt.Sprintf("JOIN %s ON %s = %s JOIN %s ON %s = %s", r.getFullMidTab(), r.getFullSrcCol(),
// 					r.getFullMidLeftCol(), r.getFullDstTab(), r.getFullMidRightCol(), r.getFullDstCol())
// 				chainedTables[dstTab] = true
// 				if len(relations) > 1 {
// 					relations = append(relations[:i], relations[i+1:]...)
// 					nextJoinStr, err := rGenJoinClause(chainedTables, relations...)
// 					if err != nil {
// 						return "", err
// 					}
// 					return currentJoinStr + " " + nextJoinStr, nil
// 				}
// 				return currentJoinStr, nil
// 			case ManyToMany:
// 				currentJoinStr := fmt.Sprintf("JOIN %s ON %s = %s JOIN %s ON %s = %s", r.getFullMidTab(), r.getFullSrcCol(),
// 					r.getFullMidLeftCol(), r.getFullDstTab(), r.getFullMidRightCol(), r.getFullDstCol())
// 				chainedTables[dstTab] = true
// 				if len(relations) > 1 {
// 					relations = append(relations[:i], relations[i+1:]...)
// 					nextJoinStr, err := rGenJoinClause(chainedTables, relations...)
// 					if err != nil {
// 						return "", err
// 					}
// 					return currentJoinStr + " " + nextJoinStr, nil
// 				}
// 				return currentJoinStr, nil
// 			case ForeignKey:
// 				currentJoinStr := fmt.Sprintf("JOIN %s ON %s = %s", r.getFullDstTab(), r.getFullSrcCol(), r.getFullDstCol())
// 				chainedTables[dstTab] = true
// 				if len(relations) > 1 {
// 					relations = append(relations[:i], relations[i+1:]...)
// 					nextJoinStr, err := rGenJoinClause(chainedTables, relations...)
// 					if err != nil {
// 						return "", err
// 					}
// 					return currentJoinStr + " " + nextJoinStr, nil
// 				}
// 			case ReverseForeignKey:
// 				currentJoinStr := fmt.Sprintf("JOIN %s ON %s = %s", r.getFullDstTab(), r.getFullSrcCol(), r.getFullDstCol())
// 				chainedTables[dstTab] = true
// 				if len(relations) > 1 {
// 					relations = append(relations[:i], relations[i+1:]...)
// 					nextJoinStr, err := rGenJoinClause(chainedTables, relations...)
// 					if err != nil {
// 						return "", err
// 					}
// 					return currentJoinStr + " " + nextJoinStr, nil
// 				}
// 			default:
// 				return "", errors.New("nborm.rGenJoinClause() error: unknown relation type")
// 			}
// 		case chainedTables[dstTab]:
// 			switch r := rel.(type) {
// 			case OneToOne:
// 				currentJoinStr := fmt.Sprintf("JOIN %s ON %s = %s JOIN %s ON %s = %s", r.getFullMidTab(),
// 					r.getMidRightCol(), r.getFullDstTab(), r.getFullSrcTab(),
// 					r.getFullMidLeftCol(), r.getFullSrcCol())
// 				chainedTables[srcTab] = true
// 				if len(relations) > 1 {
// 					relations = append(relations[:i], relations[i+1:]...)
// 					nextJoinStr, err := rGenJoinClause(chainedTables, relations...)
// 					if err != nil {
// 						return "", err
// 					}
// 					return currentJoinStr + " " + nextJoinStr, nil
// 				}
// 				return currentJoinStr, nil
// 			case ManyToMany:
// 				currentJoinStr := fmt.Sprintf("JOIN %s ON %s = %s JOIN %s ON %s = %s", r.getFullMidTab(),
// 					r.getMidRightCol(), r.getFullDstTab(), r.getFullSrcTab(),
// 					r.getFullMidLeftCol(), r.getFullSrcCol())
// 				chainedTables[srcTab] = true
// 				if len(relations) > 1 {
// 					relations = append(relations[:i], relations[i+1:]...)
// 					nextJoinStr, err := rGenJoinClause(chainedTables, relations...)
// 					if err != nil {
// 						return "", err
// 					}
// 					return currentJoinStr + " " + nextJoinStr, nil
// 				}
// 				return currentJoinStr, nil
// 			case ForeignKey:
// 				currentJoinStr := fmt.Sprintf("JOIN %s ON %s = %s", r.getFullSrcTab(), r.getFullDstCol(), r.getFullSrcCol())
// 				chainedTables[srcTab] = true
// 				if len(relations) > 1 {
// 					relations = append(relations[:i], relations[i+1:]...)
// 					nextJoinStr, err := rGenJoinClause(chainedTables, relations...)
// 					if err != nil {
// 						return "", err
// 					}
// 					return currentJoinStr + " " + nextJoinStr, nil
// 				}
// 				return currentJoinStr, nil
// 			case ReverseForeignKey:
// 				currentJoinStr := fmt.Sprintf("JOIN %s ON %s = %s", r.getFullSrcTab(), r.getFullDstCol(), r.getFullSrcCol())
// 				chainedTables[srcTab] = true
// 				if len(relations) > 1 {
// 					relations = append(relations[:i], relations[i+1:]...)
// 					nextJoinStr, err := rGenJoinClause(chainedTables, relations...)
// 					if err != nil {
// 						return "", err
// 					}
// 					return currentJoinStr + " " + nextJoinStr, nil
// 				}
// 				return currentJoinStr, nil
// 			default:
// 				return "", errors.New("nborm.rGenJoinClause() error: unknown relation type")
// 			}
// 		}
// 	}
// 	return "", fmt.Errorf("nborm.rGenJoinClause() error: cannot chain relations (%v)", relations)
// }

type relNode struct {
	srcTab      *tabNode
	dstTab      *tabNode
	srcCol      string
	dstCol      string
	midTab      string
	midLeftCol  string
	midRightCol string
}

type tabNode struct {
	rels []*relNode
	name string
}

func toRelNode(rel relation) *relNode {
	node := new(relNode)
	node.srcTab = &tabNode{[]*relNode{node}, rel.getFullSrcTab()}
	node.dstTab = &tabNode{[]*relNode{node}, rel.getFullDstTab()}
	node.srcCol = rel.getFullSrcCol()
	node.dstCol = rel.getFullDstCol()
	if r, ok := rel.(complexRelation); ok {
		node.midTab = r.getFullMidTab()
		node.midLeftCol = r.getFullMidLeftCol()
		node.midRightCol = r.getFullMidRightCol()
	}
	return node
}

func allTabMap(node *relNode) map[string]*tabNode {
	m := map[string]*tabNode{node.srcTab.name: node.srcTab, node.dstTab.name: node.dstTab}
	for _, subNode := range node.srcTab.rels {
		if subNode != node {
			subM := allTabMap(subNode)
			for k, v := range subM {
				m[k] = v
			}
		}
	}
	for _, subNode := range node.dstTab.rels {
		if subNode != node {
			subM := allTabMap(subNode)
			for k, v := range subM {
				m[k] = v
			}
		}
	}
	return m
}

func findLeaf(node *relNode) *tabNode {
	if len(node.srcTab.rels) == 1 {
		return node.srcTab
	}
	if len(node.dstTab.rels) == 1 {
		return node.dstTab
	}
	return findLeaf(node.srcTab.rels[1])
}

func (node *relNode) joinClause(prevTab *tabNode) string {
	switch prevTab {
	case node.srcTab:
		if node.midTab == "" {
			return fmt.Sprintf("JOIN %s ON %s = %s", node.dstTab.name, node.srcCol, node.dstCol)
		} else {
			return fmt.Sprintf("JOIN %s ON %s = %s JOIN %s ON %s = %s", node.midTab, node.srcCol, node.midLeftCol, node.dstTab.name,
				node.midRightCol, node.dstCol)
		}
	case node.dstTab:
		if node.midTab == "" {
			return fmt.Sprintf("JOIN %s ON %s = %s", node.srcTab.name, node.dstCol, node.srcCol)
		} else {
			return fmt.Sprintf("JOIN %s ON %s = %s JOIN %s ON %s = %s", node.midTab, node.dstCol, node.midRightCol, node.srcTab.name,
				node.midLeftCol, node.srcCol)
		}
	case nil:
		if node.midTab == "" {
			return fmt.Sprintf("%s JOIN %s ON %s = %s", node.srcTab.name, node.dstTab.name, node.dstCol, node.srcCol)
		} else {
			return fmt.Sprintf("%s JOIN %s ON %s = %s JOIN %s ON %s = %s", node.srcTab.name, node.midTab, node.srcCol, node.midLeftCol,
				node.dstTab.name, node.midRightCol, node.dstCol)
		}
	default:
		panic("nborm.relNode.joinClause() error: irrelative table")

	}
}

func (node *tabNode) otherRels(prevRel *relNode) []*relNode {
	l := make([]*relNode, 0, len(node.rels)-1)
	for _, rel := range node.rels {
		if rel != prevRel {
			l = append(l, rel)
		}
	}
	return l
}

func (node *relNode) otherTab(prevTab *tabNode) *tabNode {
	if prevTab == node.srcTab {
		return node.dstTab
	}
	return node.srcTab
}

func (node *relNode) toJoinClause(existRelMap map[*relNode]bool, prevTab *tabNode) string {
	if !existRelMap[node] {
		existRelMap[node] = true
		str := node.joinClause(prevTab)
		for _, rel := range node.srcTab.rels {
			nextStr := rel.toJoinClause(existRelMap, node.srcTab)
			if nextStr != "" {
				str = fmt.Sprintf("%s %s", str, nextStr)
			}
		}
		for _, rel := range node.dstTab.rels {
			nextStr := rel.toJoinClause(existRelMap, node.dstTab)
			if nextStr != "" {
				str = fmt.Sprintf("%s %s", str, nextStr)
			}
		}
		return str

	} else {
		return ""
	}
}

func genJoinClause(rels ...relation) (string, error) {
	nodeList := make([]*relNode, len(rels))
	for i, rel := range rels {
		nodeList[i] = toRelNode(rel)
	}
	root, remain := nodeList[0], nodeList[1:]
	atm := map[string]*tabNode{root.srcTab.name: root.srcTab, root.dstTab.name: root.dstTab}
	for len(remain) > 0 {
		var found bool
		for i, node := range remain {
			if existTab := atm[node.srcTab.name]; existTab != nil {
				atm[node.dstTab.name] = node.dstTab
				node.srcTab = existTab
				existTab.rels = append(existTab.rels, node)
				found = true
				remain = append(remain[:i], remain[i+1:]...)
				break
			} else if existTab := atm[node.dstTab.name]; existTab != nil {
				atm[node.srcTab.name] = node.srcTab
				node.dstTab = existTab
				existTab.rels = append(existTab.rels, node)
				found = true
				remain = append(remain[:i], remain[i+1:]...)
				break
			} else {
				continue
			}
		}
		if !found {
			return "", errors.New("nborm.genJoinClause() error: cannot chain relation")
		}
	}
	return root.toJoinClause(make(map[*relNode]bool), nil), nil
}
