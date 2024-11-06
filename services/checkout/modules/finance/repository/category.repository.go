package repository

import (
	"context"
	"database/sql"

	"go.opentelemetry.io/otel/trace"
	"sendzap-checkout/common/core"
	"sendzap-checkout/common/helpers"
	"sendzap-checkout/services/checkout/modules/finance/entity"
	"sendzap-checkout/services/checkout/modules/finance/filters"
	"sendzap-checkout/services/checkout/modules/finance/mapper"
)

type ICategoryRepository interface {
	CreateCategory(ctx context.Context, tx *sql.Tx, input *entity.Category) (output *entity.Category, err error)
	UpdateCategory(ctx context.Context, tx *sql.Tx, input *entity.Category) (output *entity.Category, err error)
	DeleteCategory(ctx context.Context, tx *sql.Tx, ID int, userID int) (output *entity.Category, err error)
	FindCategory(ctx context.Context, ID int, userID int) (output *entity.Category, err error)
	FindAllCategories(ctx context.Context, userID int, filters filters.CategoryFilters) (output []*entity.Category, count int, err error)
}

type CategoryRepository struct {
	Tracer   trace.Tracer
	Database *sql.DB
}

func NewCategoryRepository(
	Database *sql.DB,
	Tracer trace.Tracer,
) ICategoryRepository {
	return CategoryRepository{
		Tracer:   Tracer,
		Database: Database,
	}
}

func (r CategoryRepository) CreateCategory(ctx context.Context, tx *sql.Tx, input *entity.Category) (output *entity.Category, err error) {
	ctx, span := r.Tracer.Start(ctx, "CategoryRepository.CreateCategory")
	defer span.End()

	query := `INSERT INTO finance.categories (user_id, name, icon, color, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`

	row := tx.QueryRowContext(
		ctx,
		query,
		input.GetUserID(),
		input.GetName(),
		input.GetIcon(),
		input.GetColor(),
		input.GetCreatedAt(),
		input.GetUpdatedAt(),
	)

	return mapper.RepoToEntityCategoryMapper(ctx, r.Tracer, row)
}

func (r CategoryRepository) UpdateCategory(ctx context.Context, tx *sql.Tx, input *entity.Category) (output *entity.Category, err error) {
	ctx, span := r.Tracer.Start(ctx, "CategoryRepository.UpdateCategory")
	defer span.End()

	query := `UPDATE finance.categories SET user_id = $1, name = $2, icon = $3, color = $4, updated_at = $5 WHERE id = $6 and user_id = $7 RETURNING *`

	row := tx.QueryRowContext(
		ctx,
		query,
		input.GetUserID(),
		input.GetName(),
		input.GetIcon(),
		input.GetColor(),
		input.GetUpdatedAt(),
		input.GetID(),
		input.GetUserID(),
	)

	return mapper.RepoToEntityCategoryMapper(ctx, r.Tracer, row)

}

func (r CategoryRepository) DeleteCategory(ctx context.Context, tx *sql.Tx, ID int, userID int) (output *entity.Category, err error) {
	ctx, span := r.Tracer.Start(ctx, "CategoryRepository.DeleteCategory")
	defer span.End()

	query := `DELETE FROM finance.categories WHERE id = $1 and user_id = $2 RETURNING *`

	row := tx.QueryRowContext(
		ctx,
		query,
		ID,
		userID,
	)

	return mapper.RepoToEntityCategoryMapper(ctx, r.Tracer, row)

}

func (r CategoryRepository) FindCategory(ctx context.Context, ID int, userID int) (output *entity.Category, err error) {
	ctx, span := r.Tracer.Start(ctx, "CategoryRepository.FindCategory")
	defer span.End()

	query := `SELECT id, user_id, name, icon, color, created_at, updated_at FROM finance.categories WHERE id = $1 AND user_id = $2`

	row := r.Database.QueryRowContext(
		ctx,
		query,
		ID,
		userID,
	)

	return mapper.RepoToEntityCategoryMapper(ctx, r.Tracer, row)
}

func (r CategoryRepository) FindAllCategories(ctx context.Context, userID int, f filters.CategoryFilters) (output []*entity.Category, count int, err error) {
	ctx, span := r.Tracer.Start(ctx, "CategoryRepository.FindAllCategories")
	defer span.End()

	// Define a query base e a query para contagem
	baseQuery := `SELECT id, user_id, name, icon, color, created_at, updated_at FROM finance.categories WHERE (user_id = $1 OR user_id IS NULL)`
	baseCountQuery := `SELECT count(*) FROM finance.categories WHERE (user_id = $1 OR user_id IS NULL)`

	filter := f.ApplyFilters(
		&core.MountedQueries{
			Query:      baseQuery,
			CountQuery: baseCountQuery,
			Args:       []interface{}{userID},
			CountArgs:  []interface{}{userID},
		},
	)

	// Executa a query com os filtros aplicados
	rows, err := r.Database.QueryContext(ctx, filter.Query, filter.Args...)
	if err != nil {
		helpers.LogError(span, helpers.REPOSITORY_ERROR, err.Error())
		return nil, 0, err
	}

	// Executa a query de contagem
	counted := r.Database.QueryRowContext(ctx, filter.CountQuery, filter.CountArgs...)
	err = counted.Scan(&count)
	if err != nil {
		helpers.LogError(span, helpers.REPOSITORY_ERROR, err.Error())
		return nil, 0, err
	}

	// Mapeia os resultados
	result, err := mapper.RepoToEntityCategoriesMapper(ctx, r.Tracer, rows)
	return result, count, err
}
