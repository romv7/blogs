package mutationsTest

import "testing"

func TestUserOpsCreatePost(t *testing.T) {
	_ = `
	mutations UserOpsForCreatePost($post: UserOpsPostMutationsParameter!) {
		userOps {
			createPost(input: $post) {

			}
		}
	}
	`

	t.Error("not implemented")
}

func TestUserOpsUpdatePost(t *testing.T) {
	_ = `
	mutations UserOpsForUpdatePost($post: UserOpsPostMutationsParameter!) {
		userOps {
			updatePost(input: $post) {

			}
		}
	}
	`

	t.Error("not implemented")
}

func TestUserOpsDeletePost(t *testing.T) {
	_ = `
	mutations UserOpsForDeletePost($postUuid: String!) {
		userOps {
			deletePost(uuid: $postUuid) {

			}
		}
	}
	`

	t.Error("not implemented")
}

func TestUserOpsCreateComment(t *testing.T) {
	_ = `
	mutations UserOpsCreateComment() {
		userOps {
			createComment() {

			}
		}
	}
	`

	t.Error("not implemented")
}

func TestUserOpsUpdateComment(t *testing.T) {
	_ = `
	mutations UserOpsUpdateComment() {
		userOps {
			updateComment() {

			}
		}
	}
	`

	t.Error("not implemented")
}

func TestUserOpsDeleteComment(t *testing.T) {
	_ = `
	mutations UserOpsDeleteComment() {
		userOps {
			deleteComment() {

			}
		}
	}
	`

	t.Error("not implemented")
}
