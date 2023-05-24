package mutations

type UserOpsPostMutationsParameter struct {
	Uuid                               string
	HeadlineText, SummaryText, Content *string
	Tags, Attachments, Refs            *[]string
	Stage                              *int32
	Status                             *int32
}

type UserOpsCommentMutationsParameter struct {
	Id                      float64
	Uuid                    *string
	CommentText, TargetUuid string
	TargetType              string
}

type ArgsCreatePost struct {
	Input *UserOpsPostMutationsParameter
}

type ArgsUpdatePost struct {
	Input *UserOpsPostMutationsParameter
}

type ArgsDeletePost struct {
	Uuid string
}

type ArgsCreateComment struct {
	Input *UserOpsCommentMutationsParameter
}
type ArgsUpdateComment struct {
	Input *UserOpsCommentMutationsParameter
}

type ArgsDeleteComment struct {
	Uuid string
}
