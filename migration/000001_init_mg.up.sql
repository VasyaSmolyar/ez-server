create table if not exists tasks (
		id serial primary key,
		title text not null,
		description text not null
)