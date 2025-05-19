package process_transaction

import (
	"testing"
	"time"
	"github.com/PkMs7/gateway-pagamento-fullcycle/domain/entity"
	mock_repository "github.com/PkMs7/gateway-pagamento-fullcycle/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionExecuteInvalidCreditCard (t *testing.T) {

	input := TransactionDtoInput {

		ID: "1",
		AccountID: "1",
		TestCreditCardNumber: "40000000000000000",
		CreditCardName: "Patrick Marques",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear: time.Now().Year(),
		CreditCardCVV: 123,
		Amount: 300,

	}

	expectedOutput := TransactionDtoOutput{

		ID: "1",
		Status: entity.REJECTED,
		ErrorMessage: "invalid credit card number",

	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}

func TestProcessTransactionExecuteRejectedTransaction (t *testing.T) {

	input := TransactionDtoInput {

		ID: "1",
		AccountID: "1",
		TestCreditCardNumber: "4193523830170205",
		CreditCardName: "Patrick Marques",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear: time.Now().Year(),
		CreditCardCVV: 123,
		Amount: 1250,

	}

	expectedOutput := TransactionDtoOutput{

		ID: "1",
		Status: entity.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",

	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}

func TestProcessTransactionExecuteApprovedTransaction (t *testing.T) {

	input := TransactionDtoInput {

		ID: "1",
		AccountID: "1",
		TestCreditCardNumber: "4193523830170205",
		CreditCardName: "Patrick Marques",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear: time.Now().Year(),
		CreditCardCVV: 123,
		Amount: 900,

	}

	expectedOutput := TransactionDtoOutput{

		ID: "1",
		Status: entity.APPROVED,
		ErrorMessage: "",

	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}