package redis

type IResult interface {
	Unwrap()
	Default(v string) string
}
