package main

import (
	"database/sql"
	"errors"
)

type lead struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (l *lead) getLead(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (l *lead) updateLead(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (l *lead) deleteLead(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (l *lead) createLead(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getLeads(db *sql.DB, start, count int) ([]lead, error) {
	return nil, errors.New("Not implemented")
}
