create table if not exists users (
		id serial primary key,
		login text not null,
		hashed_pass text not null
)