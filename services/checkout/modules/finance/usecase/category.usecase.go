package usecase

import (
	"context"
	"database/sql"

	"go.opentelemetry.io/otel/trace"
	"sendzap-checkout/common/helpers"
	"sendzap-checkout/services/checkout/modules/finance/entity"
	"sendzap-checkout/services/checkout/modules/finance/mapper"
	"sendzap-checkout/services/checkout/modules/finance/repository"
	"sendzap-checkout/services/checkout/modules/finance/request"
	"sendzap-checkout/services/checkout/modules/finance/response"
)

type ICategoryUseCase interface {
	CreateCategory(ctx context.Context, tx *sql.Tx, input request.CreateCategoryRequest) (output response.CreateCategoryResponse, err *helpers.CustomError)
	UpdateCategory(ctx context.Context, tx *sql.Tx, input request.UpdateCategoryRequest) (output response.UpdateCategoryResponse, err *helpers.CustomError)
	DeleteCategory(ctx context.Context, tx *sql.Tx, input request.DeleteCategoryRequest) (output response.DeleteCategoryResponse, err *helpers.CustomError)
	FindCategory(ctx context.Context, input request.FindCategoryRequest) (output response.FindCategoryResponse, err *helpers.CustomError)
	FindAllCategories(ctx context.Context, input request.FindAllCategoriesRequest) (output response.FindAllCategoriesResponse, err *helpers.CustomError)
}

type CategoryUseCase struct {
	Tracer             trace.Tracer
	categoryRepository repository.ICategoryRepository
}

func NewCategoryUseCase(
	Tracer trace.Tracer,
	categoryRepository repository.ICategoryRepository,
) ICategoryUseCase {
	return CategoryUseCase{
		Tracer:             Tracer,
		categoryRepository: categoryRepository,
	}
}

func (u CategoryUseCase) CreateCategory(
	ctx context.Context,
	tx *sql.Tx,
	input request.CreateCategoryRequest,
) (
	output response.CreateCategoryResponse,
	customErr *helpers.CustomError,
) {
	ctx, span := u.Tracer.Start(ctx, "CategoryUseCase.CreateCategory")
	defer span.End()

	category := entity.NewCategory(
		&input.UserID,
		input.Name,
		input.Icon,
		input.Color,
	)

	data, err := u.categoryRepository.CreateCategory(ctx, tx, category)
	if err != nil {
		helpers.LogError(span, helpers.REPOSITORY_ERROR, err.Error())
		return
	}

	return response.CreateCategoryResponse{
		ID:        data.GetID(),
		UserID:    data.GetUserID(),
		Name:      data.GetName(),
		Icon:      data.GetIcon(),
		Color:     data.GetColor(),
		CreatedAt: data.GetCreatedAt(),
		UpdatedAt: data.GetUpdatedAt(),
	}, nil
}

func (u CategoryUseCase) UpdateCategory(ctx context.Context,
	tx *sql.Tx,
	input request.UpdateCategoryRequest,
) (
	output response.UpdateCategoryResponse,
	customErr *helpers.CustomError,
) {
	ctx, span := u.Tracer.Start(ctx, "CategoryUseCase.UpdateCategory")
	defer span.End()

	_, err := u.categoryRepository.FindCategory(ctx, input.ID, input.UserID)
	if err != nil {
		helpers.LogError(span, helpers.REPOSITORY_ERROR, err.Error())
		return
	}

	category := entity.NewCategory(
		&input.UserID,
		input.Name,
		input.Icon,
		input.Color,
	).
		SetID(input.ID)

	data, err := u.categoryRepository.UpdateCategory(ctx, tx, category)
	if err != nil {
		helpers.LogError(span, helpers.REPOSITORY_ERROR, err.Error())
		return
	}

	return response.UpdateCategoryResponse{
		ID:        data.GetID(),
		UserID:    data.GetUserID(),
		Name:      data.GetName(),
		Icon:      data.GetIcon(),
		Color:     data.GetColor(),
		CreatedAt: data.GetCreatedAt(),
		UpdatedAt: data.GetUpdatedAt(),
	}, nil
}

func (u CategoryUseCase) DeleteCategory(
	ctx context.Context,
	tx *sql.Tx,
	input request.DeleteCategoryRequest,
) (
	output response.DeleteCategoryResponse,
	customErr *helpers.CustomError,
) {
	ctx, span := u.Tracer.Start(ctx, "CategoryUseCase.DeleteCategory")
	defer span.End()

	_, err := u.categoryRepository.FindCategory(ctx, input.ID, input.UserID)
	if err != nil {
		helpers.LogError(span, helpers.REPOSITORY_ERROR, err.Error())
		return
	}

	data, err := u.categoryRepository.DeleteCategory(ctx, tx, input.ID, input.UserID)
	if err != nil {
		helpers.LogError(span, helpers.REPOSITORY_ERROR, err.Error())
		return
	}

	return response.DeleteCategoryResponse{
		ID:        data.GetID(),
		UserID:    data.GetUserID(),
		Name:      data.GetName(),
		Icon:      data.GetIcon(),
		Color:     data.GetColor(),
		CreatedAt: data.GetCreatedAt(),
		UpdatedAt: data.GetUpdatedAt(),
	}, nil
}

func (u CategoryUseCase) FindCategory(
	ctx context.Context,
	input request.FindCategoryRequest,
) (
	output response.FindCategoryResponse,
	customErr *helpers.CustomError,
) {
	ctx, span := u.Tracer.Start(ctx, "CategoryUseCase.FindCategory")
	defer span.End()

	data, err := u.categoryRepository.FindCategory(ctx, input.ID, input.UserID)
	if err != nil {
		helpers.LogError(span, helpers.REPOSITORY_ERROR, err.Error())
		return
	}

	return response.FindCategoryResponse{
		ID:        data.GetID(),
		UserID:    data.GetUserID(),
		Name:      data.GetName(),
		Icon:      data.GetIcon(),
		Color:     data.GetColor(),
		CreatedAt: data.GetCreatedAt(),
		UpdatedAt: data.GetUpdatedAt(),
	}, nil
}

func (u CategoryUseCase) FindAllCategories(
	ctx context.Context,
	input request.FindAllCategoriesRequest,
) (
	output response.FindAllCategoriesResponse,
	customErr *helpers.CustomError,
) {
	ctx, span := u.Tracer.Start(ctx, "CategoryUseCase.FindAllCategories")
	defer span.End()

	data, count, err := u.categoryRepository.FindAllCategories(ctx, input.UserID, input.CategoryFilters)
	if err != nil {
		helpers.LogError(span, helpers.REPOSITORY_ERROR, err.Error())
		return
	}

	categories := mapper.EntitiesCategoryToResponseMapper(ctx, u.Tracer, data)

	return response.FindAllCategoriesResponse{
		Items:      categories,
		Total:      count,
		TotalPages: count / input.CategoryFilters.Limit,
		Page:       input.CategoryFilters.Page,
		Limit:      input.CategoryFilters.Limit,
	}, nil
}
