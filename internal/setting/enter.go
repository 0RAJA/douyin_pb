package setting

type group struct {
	Config    config
	Dao       mDao
	Log       log
	Maker     maker
	Snowflake sf
}

var Group = new(group)
