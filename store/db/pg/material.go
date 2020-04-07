package pg

// import (
// 	"database/sql"
// 	"errors"
// 	"strings"

// 	"github.com/Kotyarich/find-your-pet/errs"
// 	"github.com/Kotyarich/find-your-pet/models"
// 	. "github.com/Kotyarich/find-your-pet/models"
// )

// type MaterialControllerPg struct {
// 	itemsPerPage int
// 	db           *sql.DB
// }

// func NewMaterialControllerPg(itemsPerPage int, db *sql.DB) *MaterialControllerPg {
// 	return &MaterialControllerPg{
// 		itemsPerPage: itemsPerPage,
// 		db:           db,
// 	}
// }

// func (mc *MaterialControllerPg) GetAllMaterials(subjectId, page int) ([]MaterialData, error) {
// 	itemsPerPage := mc.itemsPerPage
// 	var rows *sql.Rows
// 	var err error
// 	if page > 0 {
// 		rows, err = mc.db.Query("SELECT m.material_id, m.name, wt.name, u.email, m.date "+
// 			"FROM materials m "+
// 			"INNER JOIN worktypes wt ON m.type_id = wt.type_id "+
// 			"INNER JOIN users u ON m.author_id = u.user_id "+
// 			"WHERE m.subject_id = $1 "+
// 			"ORDER BY m.name ASC LIMIT $2 OFFSET $3", subjectId,
// 			itemsPerPage, (page-1)*itemsPerPage)
// 	} else if page == 0 {
// 		// All objects
// 		rows, err = mc.db.Query("SELECT m.material_id, m.name, wt.name, u.email, m.date "+
// 			"FROM materials m "+
// 			"INNER JOIN worktypes wt ON m.type_id = wt.type_id "+
// 			"INNER JOIN users u ON m.author_id = u.user_id "+
// 			"WHERE m.subject_id = $1 "+
// 			"ORDER BY m.name ASC", subjectId)
// 	} else {
// 		return nil, errs.IncorrectPageNumber
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	var materials []MaterialData
// 	for rows.Next() {
// 		var m MaterialData
// 		err := rows.Scan(&m.MaterialId, &m.Name, &m.TypeName, &m.UserEmail, &m.Date)
// 		if err != nil {
// 			continue
// 		}
// 		materials = append(materials, m)
// 	}
// 	return materials, nil
// }

// func (mc *MaterialControllerPg) GetById(id int) (*MaterialData, error) {
// 	row := mc.db.QueryRow("SELECT m.material_id, m.name, wt.name, u.email, m.date "+
// 		"FROM materials m "+
// 		"INNER JOIN worktypes wt ON m.type_id = wt.type_id "+
// 		"INNER JOIN users u ON m.author_id = u.user_id "+
// 		"WHERE m.material_id = $1", id)
// 	var m MaterialData
// 	err := row.Scan(&m.MaterialId, &m.Name,
// 		&m.TypeName, &m.UserEmail, &m.Date)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &m, nil
// }

// func (mc *MaterialControllerPg) Add(subjectId int, name string, typeId, authorId int) (int, error) {
// 	var idMaterial int
// 	err := mc.db.QueryRow("INSERT INTO materials (subject_id, name, type_id, author_id) "+
// 		"VALUES ($1, $2, $3, $4) RETURNING material_id",
// 		subjectId, name, typeId, authorId).Scan(&idMaterial)
// 	return idMaterial, err
// }

// func (mc *MaterialControllerPg) RemoveById(id int) error {
// 	result, err := mc.db.Exec("DELETE FROM materials WHERE material_id = $1", id)
// 	if err != nil {
// 		return err
// 	}
// 	affected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if affected == 0 {
// 		return errs.MaterialDoesntExist
// 	}
// 	return err
// }

// func (mc *MaterialControllerPg) RemoveAll(subjectId int) error {
// 	// #! Removed all the files too. Warning!
// 	_, err := mc.db.Exec("DELETE FROM materials WHERE subject_id = $1", subjectId)
// 	return err
// }

// func (mc *MaterialControllerPg) Search(subjectId int, name string, typeId, page int) ([]models.MaterialData, error) {
// 	var rows *sql.Rows
// 	var err error
// 	itemsPerPage := mc.itemsPerPage
// 	if typeId != 0 && name != "" {
// 		// Search by name and work type
// 		if page > 0 {
// 			rows, err = mc.db.Query("SELECT m.material_id, m.name, wt.name, u.email, m.date "+
// 				"FROM materials m "+
// 				"INNER JOIN worktypes wt ON m.type_id = wt.type_id AND m.type_id = $1 "+
// 				"INNER JOIN users u ON m.author_id = u.user_id "+
// 				"WHERE m.subject_id = $2 AND LOWER(m.name) LIKE '%' || $3 || '%' "+
// 				"ORDER BY m.name ASC LIMIT $4 OFFSET $5", typeId, subjectId, strings.ToLower(name),
// 				itemsPerPage, (page-1)*itemsPerPage)
// 		} else if page == 0 {
// 			// All objects
// 			rows, err = mc.db.Query("SELECT m.material_id, m.name, wt.name, u.email, m.date "+
// 				"FROM materials m "+
// 				"INNER JOIN worktypes wt ON m.type_id = wt.type_id AND m.type_id = $1 "+
// 				"INNER JOIN users u ON m.author_id = u.user_id "+
// 				"WHERE m.subject_id = $2 AND LOWER(m.name) LIKE '%' || $3 || '%' "+
// 				"ORDER BY m.name ASC", typeId, subjectId, strings.ToLower(name))
// 		} else {
// 			return nil, errs.IncorrectPageNumber
// 		}
// 	} else if typeId == 0 {
// 		// Search only by name
// 		if page > 0 {
// 			rows, err = mc.db.Query("SELECT m.material_id, m.name, wt.name, u.email, m.date "+
// 				"FROM materials m "+
// 				"INNER JOIN worktypes wt ON m.type_id = wt.type_id "+
// 				"INNER JOIN users u ON m.author_id = u.user_id "+
// 				"WHERE m.subject_id = $1 AND LOWER(m.name) LIKE '%' || $2 || '%' "+
// 				"ORDER BY m.name ASC LIMIT $3 OFFSET $4", subjectId, strings.ToLower(name),
// 				itemsPerPage, (page-1)*itemsPerPage)
// 		} else if page == 0 {
// 			// All objects
// 			rows, err = mc.db.Query("SELECT m.material_id, m.name, wt.name, u.email, m.date "+
// 				"FROM materials m "+
// 				"INNER JOIN worktypes wt ON m.type_id = wt.type_id "+
// 				"INNER JOIN users u ON m.author_id = u.user_id "+
// 				"WHERE m.subject_id = $1 AND LOWER(m.name) LIKE '%' || $2 || '%' "+
// 				"ORDER BY m.name ASC", subjectId, strings.ToLower(name))
// 		} else {
// 			return nil, errs.IncorrectPageNumber
// 		}
// 	} else if name == "" {
// 		// Search only by work type
// 		if page > 0 {
// 			rows, err = mc.db.Query("SELECT m.material_id, m.name, wt.name, u.email, m.date "+
// 				"FROM materials m "+
// 				"INNER JOIN worktypes wt ON m.type_id = wt.type_id AND m.type_id = $1 "+
// 				"INNER JOIN users u ON m.author_id = u.user_id "+
// 				"WHERE m.subject_id = $2 "+
// 				"ORDER BY m.name ASC LIMIT $3 OFFSET $4", typeId, subjectId,
// 				itemsPerPage, (page-1)*itemsPerPage)
// 		} else if page == 0 {
// 			// All objects
// 			rows, err = mc.db.Query("SELECT m.material_id, m.name, wt.name, u.email, m.date "+
// 				"FROM materials m "+
// 				"INNER JOIN worktypes wt ON m.type_id = wt.type_id AND m.type_id = $1 "+
// 				"INNER JOIN users u ON m.author_id = u.user_id "+
// 				"WHERE m.subject_id = $2 "+
// 				"ORDER BY m.name ASC", typeId, subjectId)
// 		} else {
// 			return nil, errs.IncorrectPageNumber
// 		}
// 		if err != nil {
// 			return nil, err
// 		}
// 	} else {
// 		return nil, errors.New("Incorrect data")
// 	}
// 	defer rows.Close()

// 	var materials []models.MaterialData
// 	for rows.Next() {
// 		var m models.MaterialData
// 		err := rows.Scan(&m.MaterialId, &m.Name, &m.TypeName, &m.UserEmail, &m.Date)
// 		if err != nil {
// 			continue
// 		}
// 		materials = append(materials, m)
// 	}
// 	return materials, nil
// }

// func (mc MaterialControllerPg) GetItemsPerPageCount() int {
// 	return mc.itemsPerPage
// }
