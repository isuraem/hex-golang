package ports

type DBport interface {
	CloseDbConnection() // Ensure the correct spelling here
	AddToHistory(answer int32, operation string) error
}
