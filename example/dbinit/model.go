package dbinit

type TestModel struct {
	Id   *int    `db:"id"`
	Name *string `db:"name"`
}
type User struct {
	Id   *int
	Name *string
	Age  int64
}
