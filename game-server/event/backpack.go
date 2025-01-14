package event

type ItemChangeEvent struct {
	Id           int64
	Cid          int64
	ChangeBefore int64
	ChangeAfter  int64
}
