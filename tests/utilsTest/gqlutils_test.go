package utilsTest

import (
	"testing"

	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/internal/utils/gqlutils"
)

func TestJoinFields(t *testing.T) {

	t.Run("should panic because of the fields having conflicting key.", func(t *testing.T) {
		defer func() {
			if err := recover(); err != gqlutils.ErrFieldsHasConflictName.Error() {
				t.Error(err)
			}
		}()

		var (
			F_1 = gql.Fields{
				"duplicate": &gql.Field{},
			}
			F_2 = gql.Fields{
				"duplicate": &gql.Field{},
			}
		)

		gqlutils.JoinFields(F_1, F_2)
	})

	t.Run("should join to gql.Fields into one.", func(t *testing.T) {
		var (
			commentFields = gql.Fields{
				"commentContent": &gql.Field{
					Name:        "commentContent",
					Description: "A super important field that must not conflict with another.",
					Type:        gql.String,
				},
			}
			userObjectType = gql.NewObject(gql.ObjectConfig{
				Name:        "user",
				Description: "Just another *gql.Object example.",
				Fields: gql.Fields{
					"id": &gql.Field{
						Name:        "id",
						Description: "An unique identifier that identifies a user.",
						Type:        gql.Int,
					},
				},
			})
		)

		j := gqlutils.JoinFields(commentFields, gql.Fields{
			"user": &gql.Field{
				Name:        "user",
				Description: "The owner of the comment.",
				Type:        userObjectType,
			},
		})

		var IsNotOkay = func(ok bool) bool {
			return !ok
		}

		if _, ok := j["user"]; IsNotOkay(ok) {
			t.Error("expect gqlutils.JoinFields to concatenate fields into one.")
		}
	})

}
