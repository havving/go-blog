package controllers

import (
	db "github.com/revel/modules/db/app"
	"github.com/revel/revel"
	"myapp/app/models"
)

type Post struct {
	*revel.Controller
	db.Transactional
}

func (c Post) Index() revel.Result {
	var posts []models.Post
	rows, err := c.Txn.Query("select id, title, body, created_at, updated_at from posts" +
		"order by created_at desc")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		post := models.Post{}
		if err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	return c.Render(posts)
}
