package messages

// TODO: apply error messages
const (
	UNSUCCESSFULL_BINDING = iota
)


func ApplyErrorMessage(errorCode int) string {
	result := ""

	switch errorCode {
	case UNSUCCESSFULL_BINDING:
		result = "An error occured! Binding was unsuccessful!"
	}


	return result
}