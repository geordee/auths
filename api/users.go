package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/geordee/auths/config"
	"github.com/geordee/auths/model"
	"github.com/geordee/auths/util"
)

type userRecord struct {
	Org   string
	User  string
	Role  string
	Scope string
}

func mapUser(records []userRecord) model.User {
	var user model.User
	for _, record := range records {
		user.ID = record.User
		user.Orgs = append(user.Orgs, record.Org)
		user.Roles = append(user.Roles, record.Role)
		user.Scopes = append(user.Scopes, record.Scope)
	}
	user.Orgs = util.Deduplicate(user.Orgs)
	user.Roles = util.Deduplicate(user.Roles)
	user.Scopes = util.Deduplicate(user.Scopes)
	return user
}

func queryUsersByID(userName string) ([]userRecord, error) {
	var userRecords []userRecord
	rows, err := config.DB.Query(`
		select u.name as user
			, o.name as org
			, r.name as role
			, s.name as scope
		from users_roles x
			inner join users u on x.user_id = u.id
			inner join orgs  o on u.id = o.id
			inner join roles r on x.role_id = r.id
			inner join scopes s on s.role_id = r.id
		where u.name = $1`, userName)
	if err != nil {
		log.Println("Database Error")
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Println("Query Error")
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		user := userRecord{}
		err = rows.Scan(
			&user.User,
			&user.Org,
			&user.Role,
			&user.Scope,
		)
		if err != nil {
			log.Println("Row Scan Error")
			log.Println(err.Error())
			return nil, err
		}
		userRecords = append(userRecords, user)
	}

	if len(userRecords) == 0 {
		return userRecords, errors.New("record_not_found")
	}

	return userRecords, nil
}

// Users API
func Users(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) > 3 {
		util.NotFound(&w, "path_not_found")
		return
	}

	user, err := queryUsersByID(paths[2])
	if err != nil {
		util.NotFound(&w, "user_not_found")
		return
	}

	response := mapUser(user)
	out, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	fmt.Fprintf(w, string(out))
}
