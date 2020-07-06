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
		if err := srcWallet.retrieveWithLock(s.db, serializer.SrcWallet); err != nil {
			return errors.Wrap(err, "src wallet not exists")
		}

		dstWallet := &wallet{}
		if err := dstWallet.retrieveWithLock(s.db, serializer.DstWallet); err != nil {
			return errors.Wrap(err, "dst wallet doesn't exists")
		}

		// check balance
		if srcWallet.Balance < serializer.Amount {
			return errors.New("insufficient balance")
		}

		// transfer money
		if err := transferTransaction(s.db, *srcWallet, *dstWallet, serializer.Amount); err != nil {
			return errors.Wrap(err, "cannot log transactions")
		}

		// decrease src balance
		if _, err := s.db.Model(srcWallet).Set("balance = balance - ?", serializer.Amount).Where("id = ?", serializer.SrcWallet).Update(); err != nil {
			return errors.Wrap(err, "cannot decrease balance")
		}

		// increase dst balance
		if _, err := s.db.Model(dstWallet).Set("balance = balance + ?", serializer.Amount).Where("id = ?", serializer.DstWallet).Update(); err != nil {
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
