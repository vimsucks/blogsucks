package response

type (
	RequestResponse struct {
		StatusCode int
	}

	Error struct {
		ErrorCode int
		Message string
	}

	Ok struct {
		Message string
	}
)
