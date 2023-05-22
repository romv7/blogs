package grpcTest

import (
	"context"
	"errors"
	"testing"

	endpoint "github.com/romv7/blogs/endpoints/grpc"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/store"
	"gorm.io/gorm"
)

func TestGrpcNewUser(t *testing.T) {
	u := &endpoint.UserService{}

	startTestServer(u)
	defer closeTestServer()

	conn, client := createTestUserClient()
	defer conn.Close()

	for _, tcase := range globalTestCases {
		if tcase.For != UserServiceTest {
			continue
		}

		email := tcase.testValues["email"].(string)
		fullName := tcase.testValues["full_name"].(string)
		name := tcase.testValues["name"].(string)
		utype := tcase.testValues["type"].(int)

		params := &pb.UserService_New_Params{
			Email:    email,
			Name:     name,
			FullName: fullName,
			Type:     pb.User_Type(utype),
		}

		u, err := client.New(context.Background(), params)
		if err != nil {
			t.Fatal(err)
		}

		if u.Email != email || u.FullName != fullName || u.Name != name || u.Type != pb.User_Type(utype) {
			t.Errorf("UserService: didn't match the expected values from the globalTestCases.")
		}
	}

}

func TestGrpcSaveUser(t *testing.T) {
	u := &endpoint.UserService{}

	startTestServer(u)
	defer closeTestServer()

	conn, client := createTestUserClient()
	defer conn.Close()

	for _, tcase := range globalTestCases {
		if tcase.For != UserServiceTest {
			continue
		}

		params := &pb.UserService_New_Params{
			Email:    tcase.testValues["email"].(string),
			Name:     tcase.testValues["name"].(string),
			FullName: tcase.testValues["full_name"].(string),
			Type:     pb.User_Type(tcase.testValues["type"].(int)),
		}

		// Create a new user
		u, err := client.New(context.Background(), params)
		if err != nil {
			t.Fatal(err)
		}

		// Persist the newly created user to the database.
		_, err = client.Save(context.Background(), &pb.UserService_Save_Params{User: u})
		if err != nil {
			t.Fatal(err)
		}

		// Check if the client.Save function really does work by fetching
		// the saved record from the store.
		ustore := store.NewUserStore(store.SqlStore)

		defer ustore.NewUser(u).Delete()

		if _, err := ustore.GetByUuid(u.Uuid); errors.Is(err, gorm.ErrRecordNotFound) {
			t.Errorf("client.Save did not properly saved the record to the database")
		}
	}
}

func TestGrpcDeleteUser(t *testing.T) {
	u := &endpoint.UserService{}

	startTestServer(u)
	defer closeTestServer()

	conn, client := createTestUserClient()
	defer conn.Close()

	ustore := store.NewUserStore(store.SqlStore)

	for _, tcase := range globalTestCases {
		if tcase.For != UserServiceTest {
			continue
		}

		params := &pb.UserService_New_Params{
			Name:     tcase.testValues["name"].(string),
			Email:    tcase.testValues["email"].(string),
			FullName: tcase.testValues["full_name"].(string),
			Type:     pb.User_Type(tcase.testValues["type"].(int)),
		}

		pbuser, err := client.New(context.Background(), params)
		if err != nil {
			t.Error(err)
		}

		if err := ustore.NewUser(pbuser).Save(); err != nil {
			t.Error(err)
		}

		_, err = client.Delete(context.Background(), &pb.UserService_Delete_Params{User: pbuser})
		if err != nil {
			t.Error(err)
		}
	}
}
