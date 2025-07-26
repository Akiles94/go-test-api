package models

type DomainError struct {
	Code    string
	Message string
	Details map[string]interface{}
}

func (e DomainError) Error() string {
	return e.Message
}
