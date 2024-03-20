package helper

type ResponseFor uint8

const (
	Default ResponseFor = iota
	UserRegister
	UserLogin
	UserUpdate
	UserDelete
	PhotoCreate
	PhotoGetAll
	PhotoUpdate
	PhotoDelete
	CommentCreate
	CommentGetAll
	CommentUpdate
	CommentDelete
	SocialMediaCreate
	SocialMediaGetAll
	SocialMediaUpdate
	SocialMediaDelete
	PanicRecovery
	Authentication
)

var responseMessages = map[ResponseFor]func(int) string{
	Default: func(numError int) string {
		if numError > 0 {
			return "failed to process request"
		}
		return "request processed successfully"
	},
	UserRegister: func(numError int) string {
		if numError > 0 {
			return "failed to register"
		}
		return "register success"
	},
	UserLogin: func(numError int) string {
		if numError > 0 {
			return "failed to login"
		}
		return "login success"
	},
	UserUpdate: func(numError int) string {
		if numError > 0 {
			return "failed to update user"
		}
		return "user updated successfully"
	},
	UserDelete: func(numError int) string {
		if numError > 0 {
			return "failed to delete user"
		}
		return "user deleted successfully"
	},
	PhotoCreate: func(numError int) string {
		if numError > 0 {
			return "failed to create photo"
		}
		return "photo created successfully"
	},
	PhotoGetAll: func(numError int) string {
		if numError > 0 {
			return "failed to get all photos"
		}
		return "get all photos success"
	},
	PhotoUpdate: func(numError int) string {
		if numError > 0 {
			return "failed to update photo"
		}
		return "photo updated successfully"
	},
	PhotoDelete: func(numError int) string {
		if numError > 0 {
			return "failed to delete photo"
		}
		return "photo deleted successfully"
	},
	CommentCreate: func(numError int) string {
		if numError > 0 {
			return "failed to create comment"
		}
		return "comment created successfully"
	},
	CommentGetAll: func(numError int) string {
		if numError > 0 {
			return "failed to get all comments"
		}
		return "get all comments success"
	},
	CommentUpdate: func(numError int) string {
		if numError > 0 {
			return "failed to update comment"
		}
		return "comment updated successfully"
	},
	CommentDelete: func(numError int) string {
		if numError > 0 {
			return "failed to delete comment"
		}
		return "comment deleted successfully"
	},
	SocialMediaCreate: func(numError int) string {
		if numError > 0 {
			return "failed to create social media"
		}
		return "social media created successfully"
	},
	SocialMediaGetAll: func(numError int) string {
		if numError > 0 {
			return "failed to get all social media"
		}
		return "get all social media success"
	},
	SocialMediaUpdate: func(numError int) string {
		if numError > 0 {
			return "failed to update social media"
		}
		return "social media updated successfully"
	},
	SocialMediaDelete: func(numError int) string {
		if numError > 0 {
			return "failed to delete social media"
		}
		return "social media deleted successfully"
	},
	PanicRecovery: func(numError int) string {
		return "internal server error"
	},
	Authentication: func(numError int) string {
		return "unauthenticated"
	},
}
