package cheesegull

// SystemService is a service that gets and sets system settings,
// which are normally stored in a database.
type SystemService interface {
	// Retrieves a security key, which is a sort of API key for using restricted
	// calls of the API.
	// def is the default key, which the service must store in the database if
	// the one currently stored does not exist.
	GetSecurityKey(def string) string
}

// Logging is a service that logs in a database certain information.
type Logging interface {
	// UpdateInLast5Minutes checks whether a beatmap (i) was requested an update
	// in the previous 5 minutes, and if it wasn't it records the update being
	// requested.
	UpdateInLast5Minutes(i int) (bool, error)
}
