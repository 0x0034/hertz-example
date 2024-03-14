namespace go user.exmple

struct User {
  1: i64 ID
  2: string Name
  3: list<User> Friends (go.tag = "gorm:\"many2many:user_friends\"")
}