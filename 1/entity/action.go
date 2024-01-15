package entity

type Action struct {
	// встраиваем Human в Action
	*Human
}

func NewAction(human *Human) *Action {
	return &Action{
		human,
	}
}
