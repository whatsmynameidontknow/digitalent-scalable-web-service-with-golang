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
	PhotoGetByID
	CommentCreate
	CommentGetAll
	CommentUpdate
	CommentDelete
	CommentGetByID
	SocialMediaCreate
	SocialMediaGetAll
	SocialMediaUpdate
	SocialMediaDelete
	SocialMediaGetByID
	PanicRecovery
	Authentication
)

var responseMessages = map[ResponseFor]func(int) string{
	Default: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to process request"
		}
		return "request processed successfully"
	},
	UserRegister: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to register"
		}
		return "register success"
	},
	UserLogin: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to login"
		}
		return "login success"
	},
	UserUpdate: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to update user"
		}
		return "user updated successfully"
	},
	UserDelete: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to delete user"
		}
		return "user deleted successfully"
	},
	PhotoCreate: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to create photo"
		}
		return "photo created successfully"
	},
	PhotoGetAll: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to get all photos"
		}
		return "get all photos success"
	},
	PhotoUpdate: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to update photo"
		}
		return "photo updated successfully"
	},
	PhotoDelete: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to delete photo"
		}
		return "photo deleted successfully"
	},
	PhotoGetByID: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to get photo by id"
		}
		return "get photo by id success"
	},
	CommentCreate: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to create comment"
		}
		return "comment created successfully"
	},
	CommentGetAll: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to get all comments"
		}
		return "get all comments success"
	},
	CommentUpdate: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to update comment"
		}
		return "comment updated successfully"
	},
	CommentDelete: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to delete comment"
		}
		return "comment deleted successfully"
	},
	CommentGetByID: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to get comment by id"
		}
		return "get comment by id success"
	},
	SocialMediaCreate: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to create social media"
		}
		return "social media created successfully"
	},
	SocialMediaGetAll: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to get all social media"
		}
		return "get all social media success"
	},
	SocialMediaUpdate: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to update social media"
		}
		return "social media updated successfully"
	},
	SocialMediaDelete: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to delete social media"
		}
		return "social media deleted successfully"
	},
	SocialMediaGetByID: func(errorCount int) string {
		if errorCount > 0 {
			return "failed to get social media by id"
		}
		return "get social media by id success"
	},
	PanicRecovery: func(errorCount int) string {
		return "internal server error"
	},
	Authentication: func(errorCount int) string {
		return "unauthenticated"
	},
}
