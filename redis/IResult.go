package redis

type IResult interface {
	Unwrap() ResultType
	Default(v string) ResultType
}

type ResultType interface {
	Output(out *interface{})
}
