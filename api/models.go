package api

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v9/orm"
	"github.com/pkg/errors"
)

type operation int

const (
	deposit operation = iota + 1
	withdraw
)

type wallet struct {
	ID      int32
	Name    string
	Balance float32
}

func (w *wallet) retrieveWithLock(db orm.DB, id int32) error {
	if err := db.Model(w).Where("id = ?", id).For("UPDATE").Select(); err != nil {
		return errors.Wrap(err, "wallet doesn't exists")
	}

	return nil
}

func (w *wallet) deposit(db orm.DB, amount float32) error {
	cd := time.Now()
	trn := &transaction{
		WalletID:      w.ID,
		Amount:        amount,
		BalanceBefore: w.Balance,
		Comment:       "new deposit",
		Operation:     deposit,
		CreatedAt:     &cd,
	}

	if _, err := db.Model(trn).Insert(); err != nil {
		return errors.Wrap(err, "cannot log deposit transaction")
	}

	if _, err := db.Model(w).WherePK().Set("balance = balance + ?", amount).Update(); err != nil {
		return errors.Wrap(err, "cannot update balance")
	}

	return nil
}

func (w *wallet) withdraw(db orm.DB, amount float32) error {
	cd := time.Now()
	trn := &transaction{
		WalletID:      w.ID,
		Amount:        amount,
		BalanceBefore: w.Balance,
		Comment:       "new withdraw",
		Operation:     withdraw,
		CreatedAt:     &cd,
	}

	if _, err := db.Model(trn).Insert(); err != nil {
		return errors.Wrap(err, "cannot log withdraw transaction")
	}

	if _, err := db.Model(w).WherePK().Set("balance = balance - ?", amount).Update(); err != nil {
		return errors.Wrap(err, "cannot update balance")
	}

	return nil
}

type transaction struct {
	ID            int32
	WalletID      int32
	Operation     operation
	Amount        float32
	BalanceBefore float32
	Comment       string
	CreatedAt     *time.Time
}

func transferTransaction(db orm.DB, srcWallet, dstWallet wallet, amount float32) error {
	operationTime := time.Now()

	srcTrn := &transaction{
		WalletID:      srcWallet.ID,
		Operation:     withdraw,
		Amount:        amount,
		BalanceBefore: srcWallet.Balance,
		Comment:       fmt.Sprintf("transfer to wallet %v", dstWallet.Name),
		CreatedAt:     &operationTime,
	}

	if _, err := db.Model(srcTrn).Insert(); err != nil {
		return errors.Wrap(err, "cannot save src transaction log")
	}

	dstTrn := &transaction{
		WalletID:      dstWallet.ID,
		Operation:     deposit,
		Amount:        amount,
		BalanceBefore: dstWallet.Balance,
		Comment:       fmt.Sprintf("transfer from wallet %v", srcWallet.Name),
		CreatedAt:     &operationTime,
	}

	if _, err := db.Model(dstTrn).Insert(); err != nil {
		return errors.Wrap(err, "cannot save src transaction log")
	}

	return nil
}
