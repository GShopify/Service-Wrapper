package main

type preset string

const (
	modeAuth     preset = "auth"
	modeDatabase preset = "database"
)

func (p preset) Valid() bool {
	switch p {
	case modeAuth, modeDatabase:
		return true
	}

	return false
}

func (p preset) String() string {
	return string(p)
}
