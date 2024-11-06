package handler

import (
	"database/sql"
	"net/http"

	"go.opentelemetry.io/otel/trace"
	"sendzap-checkout/common/helpers"
	"sendzap-checkout/services/checkout/modules/finance/filters"
	"sendzap-checkout/services/checkout/modules/finance/request"
	"sendzap-checkout/services/checkout/modules/finance/usecase"
)

type ICategoryHandler interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	FindCategory(w http.ResponseWriter, r *http.Request)
	FindAllCategories(w http.ResponseWriter, r *http.Request)
}

type CategoryHandler struct {
	Database        *sql.DB
	Tracer          trace.Tracer
	CategoryUseCase usecase.ICategoryUseCase
}

func NewCategoryHandler(
	Database *sql.DB,
	Tracer trace.Tracer,
	CategoryUseCase usecase.ICategoryUseCase,
) CategoryHandler {
	return CategoryHandler{
		Tracer:          Tracer,
		Database:        Database,
		CategoryUseCase: CategoryUseCase,
	}
}

func (h CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	ctx, span := h.Tracer.Start(r.Context(), "CategoryHandler.CreateCategory")
	defer span.End()

	req := request.CreateCategoryRequest{
		UserID: helpers.GetUserIDFromContext(ctx),
	}
	if err := helpers.DecodeRequestBody(r, &req); err != nil {
		helpers.LogError(span, helpers.REQUEST_BODY, err.Error())
		helpers.ReturnError(w, *helpers.NewBadRequestError(helpers.ErrUnexpected.Error()), span.SpanContext().TraceID().String())
		return
	}
	defer r.Body.Close()

	if !helpers.ValidateRequest(span, w, req) {
		return
	}

	helpers.HandleTransaction(ctx, span, w, h.Database, func(tx *sql.Tx) (interface{}, *helpers.CustomError) {
		return h.CategoryUseCase.CreateCategory(ctx, tx, req)
	})
}

func (h CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	ctx, span := h.Tracer.Start(r.Context(), "CategoryHandler.UpdateCategory")
	defer span.End()

	req := request.UpdateCategoryRequest{
		ID:     helpers.GetParam(r, "id"),
		UserID: helpers.GetUserIDFromContext(ctx),
	}

	if err := helpers.DecodeRequestBody(r, &req); err != nil {
		helpers.LogError(span, helpers.REQUEST_BODY, err.Error())
		helpers.ReturnError(w, *helpers.NewBadRequestError(helpers.ErrUnexpected.Error()), span.SpanContext().TraceID().String())
		return
	}
	defer r.Body.Close()

	if !helpers.ValidateRequest(span, w, req) {
		return
	}

	helpers.HandleTransaction(ctx, span, w, h.Database, func(tx *sql.Tx) (interface{}, *helpers.CustomError) {
		return h.CategoryUseCase.UpdateCategory(ctx, tx, req)
	})
}

func (h CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	ctx, span := h.Tracer.Start(r.Context(), "CategoryHandler.DeleteCategory")
	defer span.End()

	req := request.DeleteCategoryRequest{
		ID:     helpers.GetParam(r, "id"),
		UserID: helpers.GetUserIDFromContext(ctx),
	}

	if !helpers.ValidateRequest(span, w, req) {
		return
	}

	helpers.HandleTransaction(ctx, span, w, h.Database, func(tx *sql.Tx) (interface{}, *helpers.CustomError) {
		return h.CategoryUseCase.DeleteCategory(ctx, tx, req)
	})
}

func (h CategoryHandler) FindCategory(w http.ResponseWriter, r *http.Request) {
	ctx, span := h.Tracer.Start(r.Context(), "CategoryHandler.FindCategory")
	defer span.End()

	req := request.FindCategoryRequest{
		ID:     helpers.GetParam(r, "id"),
		UserID: helpers.GetUserIDFromContext(ctx),
	}

	if !helpers.ValidateRequest(span, w, req) {
		return
	}

	data, err := h.CategoryUseCase.FindCategory(ctx, req)
	if err != nil {
		helpers.LogError(span, helpers.USE_CASE_ERROR, err.Error())
		helpers.ReturnError(w, *err, span.SpanContext().TraceID().String())

		return
	}

	helpers.ReturnSuccess(ctx, w, data)
}

func (h CategoryHandler) FindAllCategories(w http.ResponseWriter, r *http.Request) {
	ctx, span := h.Tracer.Start(r.Context(), "CategoryHandler.FindAllCategories")
	defer span.End()

	filter := helpers.GetAllFilters(r)

	req := request.FindAllCategoriesRequest{
		UserID:          helpers.GetUserIDFromContext(ctx),
		CategoryFilters: *filters.NewCategoryFilters().SetFilters(filter),
	}

	if !helpers.ValidateRequest(span, w, req) {
		return
	}

	data, err := h.CategoryUseCase.FindAllCategories(ctx, req)
	if err != nil {
		helpers.LogError(span, helpers.USE_CASE_ERROR, err.Error())
		helpers.ReturnError(w, *err, span.SpanContext().TraceID().String())

		return
	}

	helpers.ReturnSuccess(ctx, w, data)
}
