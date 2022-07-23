package main

import (
	_ "github.com/lib/pq"
	driver "github.com/yuya0729/light-clean-architecture/cmd/app/Driver"
)

func main() {
	driver.Serve()
}

// var Db *sql.DB

// func init() {
// 	var err error
// 	// これめんどい
// 	// https: //qiita.com/mabubu0203/items/5cdff1caf2b024df1d95
// 	Db, err = sql.Open("postgres", "user=postgres host=postgres port=5432 dbname=postgres password=postgres sslmode=disable")
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func GetPosts() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// post := Post{}
// 		// posts := []*Post{}

// 		// rows, err := Db.Query("select id, content, author from posts")
// 		// if err != nil {
// 		// 	return errors.Wrapf(err, "connot connect SQL")
// 		// }
// 		// defer rows.Close()

// 		// for rows.Next() {
// 		// 	if err := rows.Scan(&post.Id, &post.Content, &post.Author); err != nil {
// 		// 		return errors.Wrapf(err, "connot connect SQL")
// 		// 	}
// 		// 	posts = append(posts, &Post{Id: post.Id, Content: post.Content, Author: post.Author})
// 		// }
// 		var data = "Hello, world"

// 		return c.JSON(http.StatusOK, data)

// 	}
// }
