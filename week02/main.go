package main

import (
	"week02/service"
	"database/sql"
	"errors"
	"fmt"
	pkg_errors "github.com/pkg/errors"
)

func main() {
	user, err := service.Query()
	if err != nil {
		if errors.Is(pkg_errors.Cause(err), sql.ErrNoRows) {
			fmt.Printf("%+v\n", pkg_errors.Unwrap(err))
			return
		}
	}
	fmt.Println(user)
}