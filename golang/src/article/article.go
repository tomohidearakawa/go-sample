package article

import (
	"database/sql"
)

type Article struct {
	Id    int
	Title string
	Body  string
}

func ReadAll(db *sql.DB) []Article {
	var articles []Article
	rows, err := db.Query("SELECT * FROM article;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		article := Article{}
		err = rows.Scan(&article.Id, &article.Title, &article.Body)
		if err != nil {
			panic(err)
		}
		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return articles
}
