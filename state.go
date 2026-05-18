package main

import (
	"github.com/emrebozdogan/blog-aggregator-go/internal/config"
	"github.com/emrebozdogan/blog-aggregator-go/internal/database"
)

type State struct {
	db  *database.Queries
	cfg *config.Config
}
