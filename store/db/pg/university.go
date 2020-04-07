package pg

// import (
// 	. "github.com/Kotyarich/find-your-pet/models"
// )

// type UniversityControllerPg struct {
// 	itemsPerPage int
// 	db           *sql.DB
// }

// func NewUniversityControllerPg(itemsPerPage int, db *sql.DB) *UniversityControllerPg {
// 	return &UniversityControllerPg{
// 		itemsPerPage: itemsPerPage,
// 		db:           db,
// 	}
// }

// func (uc *UniversityControllerPg) GetAll(page int) ([]University, error) {
// 	var rows *sql.Rows
// 	var err error
// 	if page > 0 {
// 		rows, err = uc.db.Query("SELECT university_id, name, full_name FROM universities "+
// 			"ORDER BY name ASC "+
// 			"LIMIT $1 OFFSET $2",
// 			uc.itemsPerPage, (page-1)*uc.itemsPerPage)
// 	} else if page == 0 {
// 		// All objects
// 		rows, err = uc.db.Query("SELECT university_id, name, full_name FROM universities " +
// 			"ORDER BY name ASC")
// 	} else {
// 		return nil, errs.IncorrectPageNumber
// 	}
// 	var universities []University
// 	if err != nil {
// 		// 'universities' is an empty struct
// 		return universities, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var u University
// 		rows.Scan(&u.Id, &u.Name, &u.FullName)
// 		universities = append(universities, u)
// 	}
// 	return universities, nil
// }

// func (uc *UniversityControllerPg) GetById(id int) (*University, error) {
// 	row := uc.db.QueryRow("SELECT university_id, name, full_name FROM universities "+
// 		"WHERE university_id = $1", id)
// 	var u University
// 	err := row.Scan(&u.Id, &u.Name, &u.FullName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &u, nil
// }

// func (uc *UniversityControllerPg) Add(name, fullName string) (int, error) {
// 	var idUniversity int
// 	err := uc.db.QueryRow("INSERT INTO universities (name, full_name) VALUES ($1, $2) "+
// 		"RETURNING university_id", name, fullName).Scan(&idUniversity)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return idUniversity, nil
// }

// func (uc *UniversityControllerPg) RemoveById(id int) error {
// 	result, err := uc.db.Exec("DELETE FROM universities WHERE university_id = $1", id)
// 	if err != nil {
// 		return err
// 	}
// 	affected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if affected == 0 {
// 		return errs.UniversityDoesntExist
// 	}
// 	return err
// }

// func (uc *UniversityControllerPg) RemoveAll() error {
// 	// #! Removed all the subjects too. Warning!
// 	_, err := uc.db.Exec("TRUNCATE TABLE universities CASCADE")
// 	return err
// }

// func (uc *UniversityControllerPg) Search(name string, page int) ([]models.University, error) {
// 	var rows *sql.Rows
// 	var err error
// 	if page > 0 {
// 		rows, err = uc.db.Query("SELECT university_id, name, full_name FROM universities "+
// 			"WHERE LOWER(name) LIKE '%' || $1 || '%' "+
// 			"ORDER BY name ASC LIMIT $2 OFFSET $3",
// 			strings.ToLower(name), uc.itemsPerPage, (page-1)*uc.itemsPerPage)
// 	} else if page == 0 {
// 		// All objects
// 		rows, err = uc.db.Query("SELECT university_id, name, full_name FROM universities "+
// 			"WHERE LOWER(name) LIKE '%' || $1 || '%' "+
// 			"ORDER BY name ASC", strings.ToLower(name))
// 	} else {
// 		return nil, errs.IncorrectPageNumber
// 	}
// 	defer rows.Close()
// 	var universities []models.University
// 	for rows.Next() {
// 		var u models.University
// 		err = rows.Scan(&u.Id, &u.Name, &u.FullName)
// 		if err != nil {
// 			continue
// 		}
// 		universities = append(universities, u)
// 	}
// 	return universities, nil
// }

// func (uc UniversityControllerPg) GetItemsPerPageCount() int {
// 	return uc.itemsPerPage
// }
