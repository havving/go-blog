package controllers

import (
	db "github.com/revel/modules/db/app"
	"github.com/revel/revel"
)

type Post struct {
	*revel.Controller
	db.Transactional
}
