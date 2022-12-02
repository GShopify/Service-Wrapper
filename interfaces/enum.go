package interfaces

type Enum interface {
	IsValid() bool
	String() string
}
