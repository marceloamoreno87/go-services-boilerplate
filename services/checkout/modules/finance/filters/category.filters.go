package filters

import (
	"fmt"

	"sendzap-checkout/common/core"
)

type CategoryFilters struct {
	Name string
	core.DefaultFilters
}

func NewCategoryFilters() *CategoryFilters {
	return &CategoryFilters{}
}

func (f *CategoryFilters) SetFilters(filters map[string]string) *CategoryFilters {
	f.Name = filters["name"]
	f.DefaultFilters.SetFilters(filters)
	return f
}

// Método para construir dinamicamente a query e os parâmetros
func (f *CategoryFilters) ApplyFilters(mountedQueries *core.MountedQueries) *core.MountedQueries {

	if f.Name != "" {
		mountedQueries.Query += " AND name = $" + fmt.Sprintf("%d", len(mountedQueries.Args)+1)
		mountedQueries.CountQuery += " AND name = $" + fmt.Sprintf("%d", len(mountedQueries.Args)+1)
		mountedQueries.Args = append(mountedQueries.Args, f.Name)
		mountedQueries.CountArgs = append(mountedQueries.CountArgs, f.Name)
	}

	mountedQueries = f.DefaultFilters.ApplyFilters(mountedQueries)

	return mountedQueries
}
