package main

import "time"

type Artist struct {
	id         string
	name       string
	kanaPrefix string
	createdAt  time.Time
	updatedAt  time.Time
}
