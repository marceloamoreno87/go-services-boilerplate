package finance

import (
	"sendzap-checkout/common/core"
	"sendzap-checkout/services/checkout/modules/finance/handler"
	"sendzap-checkout/services/checkout/modules/finance/repository"
	"sendzap-checkout/services/checkout/modules/finance/usecase"
)

type CategoryModule struct{}

func (f CategoryModule) SetupCategoryHandler() handler.ICategoryHandler {
	return handler.NewCategoryHandler(
		core.POSTGRESCONN,
		core.OPTL,
		f.SetupCategoryUseCase(),
	)
}

func (f CategoryModule) SetupCategoryUseCase() usecase.ICategoryUseCase {
	return usecase.NewCategoryUseCase(
		core.OPTL,
		f.SetupCategoryRepository(),
	)
}

func (f CategoryModule) SetupCategoryRepository() repository.ICategoryRepository {
	return repository.NewCategoryRepository(
		core.POSTGRESCONN,
		core.OPTL,
	)
}
