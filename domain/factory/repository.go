package factory

import "github.com/PkMs7/gateway-pagamento-fullcycle/domain/repository"

type RepositoryFactory interface {

	CreateTransactionRepository() repository.TransactionRepository
	
}