package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"."
)

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/leads", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func TestGetNonExistentLead(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/lead/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Lead not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Lead not found'. Got '%s'", m["error"])
	}
}

func TestCreateLead(t *testing.T) {
	clearTable()

	payload := []byte(`{"firstname":"Test","lastname":"Lead"}`)

	req, _ := http.NewRequest("POST", "/lead", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["firstname"] != "Test" {
		t.Errorf("Expected lead first name to be 'Test'. Got '%v'", m["firstname"])
	}

	if m["lastname"] != "Lead" {
		t.Errorf("Expected lead last name to be 'Lead'. Got '%v'", m["lastname"])
	}

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected lead ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetLead(t *testing.T) {
	clearTable()
	addLeads(1)

	req, _ := http.NewRequest("GET", "/lead/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateLead(t *testing.T) {
	clearTable()
	addLeads(1)

	req, _ := http.NewRequest("GET", "/lead/1", nil)
	response := executeRequest(req)
	var originalLead map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalLead)

	payload := []byte(`{"firstname":"Test - updated name","lastname":"Lead"}`)

	req, _ = http.NewRequest("PUT", "/lead/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalLead["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalLead["id"], m["id"])
	}

	if m["firstname"] == originalLead["firstname"] {
		t.Errorf("Expected the first name to change from '%v' to '%v'. Got '%v'", originalLead["name"], m["name"], m["name"])
	}

	if m["price"] == originalLead["lastname"] {
		t.Errorf("Expected the last name to change from '%v' to '%v'. Got '%v'", originalLead["price"], m["price"], m["price"])
	}
}

func TestDeleteLead(t *testing.T) {
	clearTable()
	addLeads(1)

	req, _ := http.NewRequest("GET", "/lead/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/lead/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/lead/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func clearTable() {
	a.DB.Exec("DELETE FROM leads_test.leads")
	a.DB.Exec("ALTER SEQUENCE leads_test.leads_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS leads_test.leads
(
    id bigint NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    CONSTRAINT leads_pkey PRIMARY KEY (id)
)`

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got code %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func addLeads(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO leads_test.leads(first_name, last_name) VALUES($1, $2)", "Test", "Lead "+strconv.Itoa(i))
	}
}
