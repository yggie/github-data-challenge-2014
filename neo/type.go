package neo

type _Type string

const (
	EVENTS       = _Type("Event")
	USERS        = _Type("User")
	REPOSITORIES = _Type("Repository")
	COMMITS      = _Type("Commit")
	LANGUAGES    = _Type("Language")
	ALL          = _Type("All")
)

type _Relation string

const (
	WRITTEN_IN = _Relation("WRITTEN_IN")
)
