create table if not exists users (
		id serial primary key,
		login text not null unique,
		hashed_pass text not null
)