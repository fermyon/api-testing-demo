package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/fermyon/api-testing-demo/pkg/query"
	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
	"github.com/fermyon/spin/sdk/go/v2/sqlite"
)

func getUser(w http.ResponseWriter, r *http.Request, params spinhttp.Params) {
	db := sqlite.Open("default")
	ctx := context.Background()

	userIDString := params.ByName("id")
	if userIDString == "" {
		writeParamError(w, "/id")
		return
	}

	userID, err := stringToInt(userIDString)
	if err != nil {
		// The assumption is that if the stringToInt method fails, the string isn't an int
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := query.GetUser(userID, db, ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "user not found", http.StatusNotFound)
		} else {
			writeInternalErr(w, err)
		}
		return
	}

	writeSuccess(w, http.StatusOK, user)
	return
}

func getAllUsers(w http.ResponseWriter, r *http.Request, params spinhttp.Params) {
	db := sqlite.Open("default")

	users, err := query.GetAllUsers(db)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "no users in database", http.StatusNoContent)
		} else {
			writeInternalErr(w, err)

		}
		return
	}

	writeSuccess(w, http.StatusOK, users)
	return
}

func createUser(w http.ResponseWriter, r *http.Request, params spinhttp.Params) {
	db := sqlite.Open("default")
	ctx := context.Background()

	if params.ByName("username") == "" {
		writeParamError(w, "/:username")
		return
	}

	if err := query.CreateUser(params.ByName("username"), db, ctx); err != nil {
		if isUniqueConstraintErr(err) {
			http.Error(w, "username already exists in database", http.StatusBadRequest)
			return
		} else {
			writeInternalErr(w, err)
			return
		}
	}

	writeSuccess(w, http.StatusAccepted, nil)
	return
}

// writeSuccess returns an HTTP success status, and if required, returns a JSON payload
func writeSuccess(w http.ResponseWriter, httpStatus int, data any) {
	w.WriteHeader(httpStatus)

	if data != nil {
		w.Header().Set("content-type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.Encode(data)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request, params spinhttp.Params) {
	db := sqlite.Open("default")
	ctx := context.Background()

	userIDString := params.ByName("id")
	if userIDString == "" {
		writeParamError(w, "/id")
		return
	}

	userID, err := stringToInt(userIDString)
	if err != nil {
		// The assumption is that if the stringToInt method fails, the string isn't an int
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := query.DeleteUser(userID, db, ctx); err != nil {
		writeInternalErr(w, err)
		return
	}

	writeSuccess(w, http.StatusAccepted, nil)
	return

}

// writeInternalErr returns an HTTP 500 status (internal server error)
func writeInternalErr(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// writeParamError returns a HTTP 400 status (bad request)
func writeParamError(w http.ResponseWriter, param string) {
	http.Error(w, fmt.Sprintf("you must include the %q parameter in your request URL", param), http.StatusBadRequest)
}

// isUniqueConstraintErr checks to see whether a CREATE request conflicted with an existing database entry
func isUniqueConstraintErr(err error) bool {
	return strings.Contains(err.Error(), "UNIQUE constraint failed")
}

func stringToInt(s string) (int, error) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("invalid id: unable to parse string %q to integer", s)
	}

	return int(num), nil
}
