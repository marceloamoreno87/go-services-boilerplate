package mapper

import (
	"context"
	"database/sql"

	"go.opentelemetry.io/otel/trace"
	"sendzap-checkout/common/helpers"
	"sendzap-checkout/services/checkout/modules/finance/entity"
	"sendzap-checkout/services/checkout/modules/finance/response"
)

func RepoToEntityCategoryMapper(ctx context.Context, Tracer trace.Tracer, row *sql.Row) (output *entity.Category, err error) {
	_, span := Tracer.Start(ctx, "RepoToEntityCategoryMapper")
	defer span.End()
	output = &entity.Category{}
	err = row.Scan(
		&output.ID,
		&output.UserID,
		&output.Name,
		&output.Icon,
		&output.Color,
		&output.CreatedAt,
		&output.UpdatedAt,
	)
	if err != nil {
		helpers.LogError(span, helpers.MAPPER_ERROR, err.Error())
		return
	}
	return
}

func RepoToEntityCategoriesMapper(ctx context.Context, Tracer trace.Tracer, rows *sql.Rows) (output []*entity.Category, err error) {
	_, span := Tracer.Start(ctx, "RepoToEntityCategoriesMapper")
	defer span.End()
	defer rows.Close()
	output = make([]*entity.Category, 0)
	for rows.Next() {
		category := &entity.Category{}

		err = rows.Scan(
			&category.ID,
			&category.UserID,
			&category.Name,
			&category.Icon,
			&category.Color,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			helpers.LogError(span, helpers.MAPPER_ERROR, err.Error())

			return
		}

		output = append(output, category)
	}
	return
}

func EntitiesCategoryToResponseMapper(ctx context.Context, Tracer trace.Tracer, categories []*entity.Category) (output []response.FindCategoryResponse) {
	_, span := Tracer.Start(ctx, "EntitiesCategoryToResponseMapper")
	defer span.End()
	output = make([]response.FindCategoryResponse, 0)

	for _, category := range categories {
		output = append(output, response.FindCategoryResponse{
			ID:        category.GetID(),
			UserID:    category.GetUserID(),
			Name:      category.GetName(),
			Icon:      category.GetIcon(),
			Color:     category.GetColor(),
			CreatedAt: category.GetCreatedAt(),
			UpdatedAt: category.GetUpdatedAt(),
		})
	}
	return
}

func EntitiesCategoryToResponseMapperWithPagination(ctx context.Context, Tracer trace.Tracer, categories []*entity.Category) (output response.FindAllCategoriesResponse) {
	_, span := Tracer.Start(ctx, "EntitiesCategoryToResponseMapperWithPagination")
	defer span.End()

	output.Items = make([]response.FindCategoryResponse, 0)
	for _, category := range categories {
		output.Items = append(output.Items, response.FindCategoryResponse{
			ID:        category.GetID(),
			UserID:    category.GetUserID(),
			Name:      category.GetName(),
			CreatedAt: category.GetCreatedAt(),
			UpdatedAt: category.GetUpdatedAt(),
		})
	}
	return
}
