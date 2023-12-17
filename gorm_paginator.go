package plank

import (
	"math"

	"gorm.io/gorm"
)

// # Examples
// 	countDB := sR.DB.Group("id").Model(&models.Spora{})
// 	rowsDB := sR.DB.Scopes(paginator.Scopes()).Where("user_id = ?", userID).Preload(clause.Associations)
// 	if err = paginator.SetCount(countDB); err != nil {
// 		return
// 	}
// 	if err = rowsDB.Find(&data).Error; err != nil {
// 		return
// 	}
// 	paginator.Paginate(data)
//
// Custom GPaginator
//	var customPaginator struct {
//		plank.GPaginator
//		Query string
//	}
// 	countDB := sR.DB.Group("id").Model(&models.Spora{})
// 	rowsDB := sR.DB.Scopes(customPaginator.Scopes()).Where("user_id = ?", userID).Preload(clause.Associations)
// 	if len(customPaginator.Query) > 0 {
// 		rowsDB = rowsDB.Where("name ilike ?", "customPaginator.Query)
// 		countDB = countDB.Where("name ilike ?",customPaginator.Query)
//	}
// 	if err = customPaginator.SetCount(countDB); err != nil {
// 		return
// 	}
// 	if err = rowsDB.Find(&data).Error; err != nil {
// 		return
// 	}
// 	customPaginator.Paginate(data)
//
// The GPaginator class is a struct that represents pagination parameters for a GORM query.
type GPaginator struct {
	Limit      int    `query:"limit" form:"limit, omitempty" json:"limit"` // Specifies the maximum number of rows to fetch per page.
	Page       int    `query:"page" form:"page, omitempty" json:"page"`    // Specifies the page number to fetch.
	Sort       string `query:"sort" form:"sort, omitempty" json:"sort"`    // Specifies the column to sort the results by.
	Order      string `query:"order" form:"order" json:"order"`            // Specifies the order of the sorted results (ascending or descending).
	TotalRows  int64  `json:"total_rows"`                                  // Stores the total number of rows in the query result.
	TotalPages int    `json:"total_pages"`                                 // Stores the total number of pages based on the total rows and the limit.
	Rows       any    `json:"rows"`                                        // Stores the actual rows fetched from the query.
	Offset     int    `json:"-"`                                           // Stores the offset for pagination.
}

// Scopes returns a function that applies the pagination and ordering parameters
// specified in the GPaginator instance to the provided *gorm.DB object.
//
// Parameters:
// - db: The *gorm.DB object to apply the scopes to.
//
// Returns:
// - A function that takes a *gorm.DB parameter and returns a *gorm.DB with the
//   pagination and ordering parameters applied.
func (p *GPaginator) Scopes() func(db *gorm.DB) *gorm.DB {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.Offset).Limit(p.Limit).Order(p.Order + " " + p.Sort)
	}
}

// SetCount sets the count for the GPaginator.
//
// It takes a countCondition of type *gorm.DB and returns an error.
// The countCondition is used to query the database and retrieve the total count.
// The total count is stored in the `total` variable of type int64.
// If an error occurs during the count operation, it is returned.
// Otherwise, the `total` count is passed to the SetNCount method of the GPaginator and any resulting error is returned.
func (p *GPaginator) SetCount(countCondition *gorm.DB) (err error) {
	var total int64
	if err = countCondition.Count(&total).Error; err != nil {
		return
	}
	return p.SetNCount(total)
}

// SetNCount sets the total count for pagination.
//
// total: the total count of items.
// err: an error if there was a problem setting the total count.
func (p *GPaginator) SetNCount(total int64) (err error) {

	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.Page <= 0 {
		p.Page = 1
	}

	p.Offset = (p.Page - 1) * p.Limit

	if len(p.Sort) < 1 {
		p.Sort = "asc"
	}

	if len(p.Order) < 1 {
		p.Order = "id"
	}

	p.TotalRows = total
	p.TotalPages = int(math.Ceil(float64(p.TotalRows) / float64(p.Limit)))
	return
}

// Paginate sets the rows for pagination.
//
// rows: any - The rows to be paginated.
func (p *GPaginator) Paginate(rows any) {
	p.Rows = rows
}
