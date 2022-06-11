package bitcoins

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type WalletTestSuite struct {
	suite.Suite
	Wallet Wallet
}

func (s *WalletTestSuite) SetupTest() {
	s.Wallet = Wallet{balance: Bitcoin(0)}
}

func (s *WalletTestSuite) TestWallet() {
	s.Run("deposit", func() {
		s.Wallet.Deposit(10)
		s.assertBalance(10)
	})

	s.Run("withdraw within balance", func() {
		s.Wallet = Wallet{balance: 10}
		err := s.Wallet.Withdraw(10)
		s.NoError(err)
		s.assertBalance(0)
	})

	s.Run("withdraw with insufficient funds", func() {
		err := s.Wallet.Withdraw(25)
		s.Errorf(err, ErrInsufficientFunds.Error())
	})
}

func (s *WalletTestSuite) assertBalance(bitcoin Bitcoin) bool {
	return s.Equal(bitcoin, s.Wallet.Balance())
}

func TestWalletTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}
