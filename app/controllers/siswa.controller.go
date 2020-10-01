package controllers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	"simple-crud/app/config"
	"simple-crud/app/models"
)

const (
	table          = "siswa"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Siswa, error) {

	var siswas []models.Siswa

	db, err := config.Mysql()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var siswa models.Siswa
		var createdAt, updatedAt string

		if err = rowQuery.Scan(
			&siswa.ID,
			&siswa.NIM,
			&siswa.Name,
			&siswa.Semester,
			&createdAt,
			&updatedAt); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		siswa.CreatedAt = time.Now()
		siswa.UpdatedAt = time.Now()

		siswas = append(siswas, siswa)
	}

	return siswas, nil
}
// Insert
func Insert(ctx context.Context, sw models.Siswa) error {

	db, err := config.Mysql()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nim, name, semester, created_at, updated_at) values(%v,'%v',%v,'%v','%v')",
		table,
		sw.NIM,
		sw.Name,
		sw.Semester,
		time.Now(),
		time.Now())

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// Update
func Update(ctx context.Context, sw models.Siswa) error {

	db, err := config.Mysql()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nim = %d, name ='%s', semester = %d, updated_at = '%v' where id = '%d'",
		table,
		sw.NIM,
		sw.Name,
		sw.Semester,
		time.Now(),
		sw.ID,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// Delete
func Delete(ctx context.Context, sw models.Siswa) error {

	db, err := config.Mysql()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, sw.ID)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()

	if check == 0 {
		return errors.New("id tidak ada ")
	}

	return nil
}