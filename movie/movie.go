package movie

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-mysql/config"
	"golang-mysql/models"
	"log"
	"time"
)

const (
	table          string = "movie"
	layoutDateTime string = "2006-01-02 15:04:05"
)

func GetAll(context context.Context) ([]models.Movie, error) {
	var movies []models.Movie

	db, err := config.ConnectToMySQL()

	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	query := fmt.Sprintf("SELECT * FROM %v ORDER BY created_at DESC", table)
	rowQuery, errQuery := db.QueryContext(context, query)

	if errQuery != nil {
		log.Fatal(errQuery)
	}

	for rowQuery.Next() {
		var movie models.Movie
		var createdAt, updatedAt string

		if err3 := rowQuery.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Year,
			&createdAt,
			&updatedAt,
		); err3 != nil {
			return nil, err3
		}
		//  Change format string to datetime for created_at and updated_at
		movie.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		movie.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	return movies, nil
}

func InsertMovie(context context.Context, movie models.Movie) error {
	db, err := config.ConnectToMySQL()

	if err != nil {
		log.Fatal("There is some problem with the database...", err)
	}

	query := fmt.Sprintf("INSERT INTO %v (title, year, created_at, updated_at) VALUES ('%v', '%v', NOW(), NOW())", table, movie.Title, movie.Year)

	_, err = db.ExecContext(context, query)

	if err != nil {
		return err
	}

	return nil
}

func UpdateMovieById(context context.Context, movie models.Movie, id string) error {
	db, err := config.ConnectToMySQL()

	if err != nil {
		log.Fatal("There is some problem with the database...", err)
	}

	query := fmt.Sprintf("UPDATE %v set title ='%s', year =%d, updated_at = NOW() WHERE id = %s", table, movie.Title, movie.Year, id)

	_, err = db.ExecContext(context, query)

	if err != nil {
		return err
	}

	return nil
}

func DeleteMovie(context context.Context, idMovie string) error {
	db, err := config.ConnectToMySQL()

	if err != nil {
		log.Fatal("There is some problem with the database...", err)
	}

	query := fmt.Sprintf("DELETE FROM %v WHERE id = %s", table, idMovie)

	result, err := db.ExecContext(context, query)

	if err != nil && err != sql.ErrNoRows {
		return errors.New("Id not found")
	}

	check, err := result.RowsAffected()
	fmt.Println(check)

	if check == 0 {
		return errors.New(err.Error())
	}

	return nil
}
