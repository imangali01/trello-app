package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/Imangali2002/trello-app/view"
)

type Suit struct {
	testCase       string
	input          view.UserCreate
	expectedErr    error
	expectedResult view.UserInfo
}

func TestUserRepositoryInit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	ur := UserRepositoryInit(db)

	suits := []Suit{
		{
			testCase: "RegisterUser - create user success",
			input: view.UserCreate{
				Username: "Zhantore",
				Email:    "Zhantore@nitec.kz",
			},
			expectedErr: nil,
		},
		{
			testCase: "RegisterUser - create user fail",
			input: view.UserCreate{
				Username: "uuu",
				Email:    "uuu@mail",
			},
			expectedErr: nil,
		},
	}

	for _, tc := range suits {

		mock.ExpectQuery("INSERT INTO public.user").
			WithArgs(sqlmock.AnyArg(), tc.input.Username, tc.input.Email).
			WillReturnError(nil)

		err := ur.RegisterUser(tc.input)
		if err != tc.expectedErr {
			t.Errorf("something wrong: %v", err)
		}
	}
}
