package server_test

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	// "math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	// "time"
	
	_ "github.com/lib/pq"

	"github.com/tarekbadrshalaan/goStuff/configuration"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/api"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/db"  
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/dto"
) 

//!+test
//go test -v
func TestBaseRentals(t *testing.T) {
	// configurations.
	c := &config{}
	err := configuration.JSON("test.json", c)
	if err != nil {
		panic(err)
	}
	// configurations.

	// database.
	if err := db.InitDB(c.DBEngine, c.DBConnectionString); err != nil {
		panic(err)
	}
	// database.

	h := api.ConfigRouter()

	tt := []struct {
		name string
		f    func(t *testing.T, h http.Handler)
	}{
		{name: "getRentals", f: getRentals},
		{name: "getAllRentals", f: getAllRentals},
		{name: "postRentals", f: postRentals},
		{name: "putRentals", f: putRentals},
		{name: "deleteRentals", f: deleteRentals},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.f(t, h)
		})
	}
}

func getRentals(t *testing.T, h http.Handler) {
	tt := []struct {
		name       string
		value      string
		expecte    string
		err        string
		statusCode int
	}{ 
		// {name: "two", value: "2", expecte: `{"rental_id":"","rental_date":"","inventory_id":"","customer_id":"","return_date":"","staff_id":"","last_update":""}`},
		{name: "missing id value", value: "", err: `<a href="http://:/rentals">Moved Permanently</a>.`, statusCode: 301},
		{name: "id not int32", value: "x", err: "Error: parameter (id) should be int32", statusCode: 400},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "http://::/rentals/"+tc.value, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()

			h.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				if res.StatusCode != tc.statusCode {
					t.Errorf("expected status code %d; got %v", tc.statusCode, res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
				return
			}

			if string(bytes.TrimSpace(b)) != tc.expecte {
				t.Fatalf("expected %v; got %s", tc.expecte, b)
			}
		})
	}
}

func getAllRentals(t *testing.T, h http.Handler) {
	tt := []struct {
		name       string
		value      string
		expecte    int
		err        string
		statusCode int
	}{
		// {name: "test by count", expecte: 200},
		{name: "wrong parameter", value: "x", err: "404 page not found", statusCode: 404},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "http://::/rentals"+tc.value, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()

			h.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				if res.StatusCode != tc.statusCode {
					t.Errorf("expected status code %d; got %v", tc.statusCode, res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
				return
			}

			rentals := []dto.RentalDTO{}
			err = json.Unmarshal(bytes.TrimSpace(b), &rentals)
			if err != nil {
				t.Fatal(err)
			}
			if len(rentals) != tc.expecte {
				t.Fatalf("expected %v; got %d", tc.expecte, len(rentals))
			}
		})
	}
}

func postRentals(t *testing.T, h http.Handler) {
	tt := []struct {
		name       string
		body       string
		expecte    string
		err        string
		statusCode int
	}{
		// {name: "duplicate key", body: `{"rental_id":"","rental_date":"","inventory_id":"","customer_id":"","return_date":"","staff_id":"","last_update":""}`, err: `pq: duplicate key value violates unique constraint "rental_pkey"`, statusCode: 500},
		{name: "wrong parameter", body: "x", err: "invalid character 'x' looking for beginning of value", statusCode: 400},
		// {name: "new rentals", body: `{"rental_id":"","rental_date":"","inventory_id":"","customer_id":"","return_date":"","staff_id":"","last_update":""}`, expecte: `{"rental_id":"","rental_date":"","inventory_id":"","customer_id":"","return_date":"","staff_id":"","last_update":""}`},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "http://::/rentals", bytes.NewBuffer([]byte(tc.body)))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()

			h.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				if res.StatusCode != tc.statusCode {
					t.Errorf("expected status code %d; got %v", tc.statusCode, res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
				return
			}

			if string(bytes.TrimSpace(b)) != tc.expecte {
				t.Fatalf("expected %v; got %s", tc.expecte, b)
			}
		})
	}
}

func putRentals(t *testing.T, h http.Handler) {
	tt := []struct {
		name       string
		body       string
		expecte    string
		err        string
		statusCode int
	}{
		// {name: "wrong key", body: `{"rental_id":"","rental_date":"","inventory_id":"","customer_id":"","return_date":"","staff_id":"","last_update":""}`, err: "record not found", statusCode: 500},
		{name: "wrong parameter", body: "x", err: "invalid character 'x' looking for beginning of value", statusCode: 400},
		// {name: "update rentals", body: `{"rental_id":"","rental_date":"","inventory_id":"","customer_id":"","return_date":"","staff_id":"","last_update":""}`, expecte: `{"rental_id":"","rental_date":"","inventory_id":"","customer_id":"","return_date":"","staff_id":"","last_update":""}`},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("PUT", "http://::/rentals", bytes.NewBuffer([]byte(tc.body)))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()

			h.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				if res.StatusCode != tc.statusCode {
					t.Errorf("expected status code %d; got %v", tc.statusCode, res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
				return
			}

			if string(bytes.TrimSpace(b)) != tc.expecte {
				t.Fatalf("expected %v; got %s", tc.expecte, b)
			}
		})
	}
}

func deleteRentals(t *testing.T, h http.Handler) {
	tt := []struct {
		name       string
		value      string
		expecte    string
		err        string
		statusCode int
	}{
		// {name: "delete one", value: "201", expecte: ""},
		{name: "missing id value", value: "", err: "404 page not found", statusCode: 404},
		{name: "id not int32", value: "x", err: "Error: parameter (id) should be int32", statusCode: 400},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "http://::/rentals/"+tc.value, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()

			h.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				if res.StatusCode != tc.statusCode {
					t.Errorf("expected status code %d; got %v", tc.statusCode, res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
				return
			}

			if string(bytes.TrimSpace(b)) != tc.expecte {
				t.Fatalf("expected %v; got %s", tc.expecte, b)
			}
		})
	}
}

//!-tests


/*
//!+bench
//go test -v  -bench=.
func BenchmarkReadRentals(b *testing.B) {
	h := api.ConfigRouter()

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	for index := 0; index < b.N; index++ {
		url := fmt.Sprintf("http://::/rentals/%d", rnd.Intn(100)+1)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			b.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			b.Errorf("expected status OK; got %v", res.Status)
			return
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			b.Fatalf("could not read response: %v", err)
		}
	}
}


func BenchmarkWriteRentals(b *testing.B) {
	h := api.ConfigRouter()

	for index := 0; index < b.N; index++ {
		// POST
		req, err := http.NewRequest("POST", "http://::/rentals", bytes.NewBuffer([]byte(`{"rental_id":"","rental_date":"","inventory_id":"","customer_id":"","return_date":"","staff_id":"","last_update":""}`)))
		if err != nil {
			b.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			b.Errorf("expected status OK; got %v", res.Status)
			return
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			b.Fatalf("could not read response: %v", err)
		}

		// PUT
		req, err = http.NewRequest("PUT", "http://::/rentals", bytes.NewBuffer([]byte(`{"rental_id":"","rental_date":"","inventory_id":"","customer_id":"","return_date":"","staff_id":"","last_update":""}`)))
		if err != nil {
			b.Fatalf("could not create request: %v", err)
		}
		rec = httptest.NewRecorder()

		h.ServeHTTP(rec, req)

		res = rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			b.Errorf("expected status OK; got %v", res.Status)
			return
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			b.Fatalf("could not read response: %v", err)
		}

		// DELETE
		req, err = http.NewRequest("DELETE", "http://::/rentals/201", nil)
		if err != nil {
			b.Fatalf("could not create request: %v", err)
		}
		rec = httptest.NewRecorder()

		h.ServeHTTP(rec, req)

		res = rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			b.Errorf("expected status OK; got %v", res.Status)
			return
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			b.Fatalf("could not read response: %v", err)
		}

	}
}
*/
//!-bench

