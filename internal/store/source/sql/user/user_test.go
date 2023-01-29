package user_test

import (
	"github.com/rommms07/blogs/internal/entities"
	"github.com/rommms07/blogs/internal/store/source/sql/user"
	"testing"
)

func Test_shouldBeAbleToCreateANewMockUser(t *testing.T) {
	user.InitSql()

	sqlsrc := user.UserStoreSql{}
	me := entities.NewUser("rommms", "Rom Vales Villanueva", "idream.rommms@gmail.com")
	sqlsrc.Save(me)

	//sqlsrc.DeleteByUuid(me.Uuid, me.State.CreatedAt.AsTime().Unix())
}

func Test_mustBeAbleToEditAnExistingMockUser(t *testing.T) {
	
}

func Test_ableToRemoveOrDeleteAMockUser(t *testing.T) {

}

func test_mustReturnASetOfUserBasedOnTheProvidedQuery(t *testing.T) {

}
