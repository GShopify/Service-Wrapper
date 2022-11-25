package model

type PublishScope string

const (
	PublishScopeWeb    PublishScope = "web"
	PublishScopeGlobal PublishScope = "global"
)

var AllPublishScopes = []PublishScope{
	PublishScopeWeb,
	PublishScopeGlobal,
}

func (e PublishScope) IsValid() bool {
	for _, t := range AllPublishScopes {
		if t == e {
			return true
		}
	}

	return false
}

func (e PublishScope) String() string {
	return string(e)
}
