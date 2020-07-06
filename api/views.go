package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/pkg/errors"
)

func (s *Server) createWallet(c *gin.Context) {
	serializer := &walletCreateSerializer{}
	if err := c.BindJSON(serializer); err != nil {
		s.badRequest(c, err)
		return
	}

	if err := s.db.Insert(s); err != nil {
		s.internalError(c, err)
		return
	}

	c.JSON(http.StatusOK, "wallet created successfully")
}

func (s *Server) makeTransfer(c *gin.Context) {
	serializer := &moneyTransferSerializer{}
	if err := c.BindJSON(serializer); err != nil {
		s.badRequest(c, err)
		return
	}

	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		// check wallets exists
		// block row by `select for update`

		srcWallet := &wallet{}
		if err := srcWallet.retrieveWithLock(tx, serializer.SrcWallet); err != nil {
			return errors.Wrap(err, "src wallet not exists")
		}

		dstWallet := &wallet{}
		if err := dstWallet.retrieveWithLock(tx, serializer.DstWallet); err != nil {
			return errors.Wrap(err, "dst wallet doesn't exists")
		}

		// check balance
		if srcWallet.Balance < serializer.Amount {
			return errors.New("insufficient balance")
		}

		// transfer money
		if err := transferTransaction(tx, *srcWallet, *dstWallet, serializer.Amount); err != nil {
			return errors.Wrap(err, "cannot log transactions")
		}

		// decrease src balance
		if _, err := tx.Model(srcWallet).Set("balance = balance - ?", serializer.Amount).Where("id = ?", serializer.SrcWallet).Update(); err != nil {
			return errors.Wrap(err, "cannot decrease balance")
		}

		// increase dst balance
		if _, err := tx.Model(dstWallet).Set("balance = balance + ?", serializer.Amount).Where("id = ?", serializer.DstWallet).Update(); err != nil {
			return errors.Wrap(err, "cannot increase balance")
		}

		return nil
	})

	if err != nil {
		s.internalError(c, err)
		return
	}

	c.JSON(http.StatusOK, "money transferred successfully")

}

func (s *Server) deposit(c *gin.Context) {
	serializer := &depositSerializer{}
	if err := c.BindJSON(serializer); err != nil {
		s.badRequest(c, err)
		return
	}

	err := s.db.RunInTransaction(func(tx *pg.Tx) error {

		w := &wallet{}
		if err := w.retrieveWithLock(tx, serializer.Wallet); err != nil {
			return errors.Wrap(err, "cannot retrieve wallet")
		}

		if err := w.deposit(tx, serializer.Amount); err != nil {
			return errors.Wrap(err, "cannot deposit")
		}

		return nil
	})

	if err != nil {
		s.internalError(c, err)
		return
	}

	c.JSON(http.StatusOK, "deposit successfully saved")
}

func (s *Server) withdraw(c *gin.Context) {
	serializer := &withdrawSerializer{}
	if err := c.BindJSON(serializer); err != nil {
		s.badRequest(c, err)
		return
	}

	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		w := &wallet{}
		if err := w.retrieveWithLock(tx, serializer.Wallet); err != nil {
			return errors.Wrap(err, "cannot retrieve wallet")
		}

		if err := w.withdraw(tx, serializer.Amount); err != nil {
			return errors.Wrap(err, "cannot withdraw")
		}

		return nil
	})

	if err != nil {
		s.internalError(c, err)
		return
	}

	c.JSON(http.StatusOK, "withdraw successfull")
}

func (s *Server) getOperations(c *gin.Context) {
	var do, oo string
	wallet := c.Param("wallet")

	dateOrd := c.Query("date")
	operationOrd := c.Query("operation")

	if dateOrd != "" && (dateOrd == "asc" || dateOrd == "desc") {
		do = "created_at " + dateOrd
	}

	if operationOrd != "" && (operationOrd == "asc" || operationOrd == "desc") {
		oo = "operation " + operationOrd
	}

	var transactions []transaction

	if err := s.db.Model(&transactions).Where("wallet_id = ?", wallet).Order(do, oo).Select(); err != nil {
		s.internalError(c, err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}
