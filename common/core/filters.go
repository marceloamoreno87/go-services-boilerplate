package core

import (
	"fmt"

	"sendzap-checkout/common/helpers"
)

type MountedQueries struct {
	Query      string
	CountQuery string
	Args       []interface{}
	CountArgs  []interface{}
}

type DefaultFilters struct {
	CreatedAtStart   string
	CreatedAtEnd     string
	OrderByField     string
	OrderByDirection string
	Page             int
	Limit            int
}

func NewDefaultFilters() *DefaultFilters {
	return &DefaultFilters{}
}

func (f *DefaultFilters) SetFilters(filters map[string]string) *DefaultFilters {
	f.CreatedAtStart = filters["created_at_start"]
	f.CreatedAtEnd = filters["created_at_end"]
	f.OrderByField = filters["order_by_field"]
	f.OrderByDirection = filters["order_by_direction"]
	f.Limit = helpers.StringToInt(filters["limit"])
	f.Page = helpers.StringToInt(filters["page"])
	return f
}

func (f *DefaultFilters) ApplyFilters(mountedQueries *MountedQueries) *MountedQueries {

	if f.CreatedAtStart != "" {
		mountedQueries.Query += " AND created_at >= $" + fmt.Sprintf("%d", len(mountedQueries.Args)+1)
		mountedQueries.CountQuery += " AND created_at >= $" + fmt.Sprintf("%d", len(mountedQueries.Args)+1)
		mountedQueries.Args = append(mountedQueries.Args, f.CreatedAtStart)
		mountedQueries.CountArgs = append(mountedQueries.CountArgs, f.CreatedAtStart)
	}

	if f.CreatedAtEnd != "" {
		mountedQueries.Query += " AND created_at <= $" + fmt.Sprintf("%d", len(mountedQueries.Args)+1)
		mountedQueries.CountQuery += " AND created_at <= $" + fmt.Sprintf("%d", len(mountedQueries.Args)+1)
		mountedQueries.Args = append(mountedQueries.Args, f.CreatedAtEnd)
		mountedQueries.CountArgs = append(mountedQueries.CountArgs, f.CreatedAtEnd)
	}

	if f.OrderByField != "" {
		orderByDirection := "ASC"
		if f.OrderByDirection != "" {
			orderByDirection = f.OrderByDirection
		}

		mountedQueries.Query += " ORDER BY " + f.OrderByField + " " + orderByDirection
	}

	mountedQueries.Query += " LIMIT $" + fmt.Sprintf("%d", len(mountedQueries.Args)+1) + " OFFSET $" + fmt.Sprintf("%d", len(mountedQueries.Args)+2)
	mountedQueries.Args = append(mountedQueries.Args, f.Limit)
	mountedQueries.Args = append(mountedQueries.Args, f.Limit*(f.Page-1))

	return mountedQueries
}
