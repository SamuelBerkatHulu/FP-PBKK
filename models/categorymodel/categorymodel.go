package categorymodel

import (
	"fp-pbkk2/config"
	"fp-pbkk2/entities"
)

// mengambil semua categories dari database
func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}
//function untuk memasukan data ke tabel  UNTUK controller ADD
func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
	INSERT INTO categories (name, created_at, updated_at)
	VALUE (?, ?, ?)`,
	category.Name, category.CreatedAt, category.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

//GET id categories dari database UNTUK controller Edit
func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = ?`, id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Name); err !=  nil {
		panic(err.Error())
	}

	return category
}

//POST name id categories dari database UNTUK controller Edit
func Update(id int, category entities.Category) bool {
	result, err := config.DB.Exec(`UPDATE categories SET name = ?, updated_at = ? WHERE id = ?`, category.Name, category.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	// Mendapatkan jumlah baris yang terpengaruh
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	return rowsAffected > 0
}

//fucntion delete data id untuk controller Delete
func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}
