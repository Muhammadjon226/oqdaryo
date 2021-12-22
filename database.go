package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

func PostInfo(data CommentInfo) error {

	db, err := sql.Open("postgres", DB_CONFIG)
	if err != nil {
		
		err = errors.New("connection failed to database")
		return err
	}
	defer db.Close()
	fmt.Println(data)
	if (data.Name == "" && data.Email == "" && data.Comment == "") {
		err = errors.New("null string")
		return err
	}
	_, err = db.Query(POSTINFO,
		data.Name,
		data.Email,
		data.Comment,
	)

	if err != nil {
		fmt.Println(err)
		err = errors.New("failed to write database")
		return err
	}
	return nil

}

func GetAllInfo() ([]CommentInfo, error) {

	db, err := sql.Open("postgres", DB_CONFIG)
	if err != nil {
		err = errors.New("connection failed to database")
		return []CommentInfo{}, err
	}
	defer db.Close()

	rows, err := db.Query(GETALLINFO)

	if err == sql.ErrNoRows {
		return []CommentInfo{}, err
	}
	defer rows.Close()

	var comments []CommentInfo

	for rows.Next() {
		var comment CommentInfo

		rows.Scan(
			&comment.Name,
			&comment.Email,
			&comment.Comment,
		)
		comments = append(comments, comment)
	}
	fmt.Println(comments)

	return comments, nil

}