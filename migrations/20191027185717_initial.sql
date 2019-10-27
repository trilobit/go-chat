-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table users
(
	id serial not null
		constraint users_pk
			primary key,
	email varchar(128) not null,
	pswd varchar(256) not null,
	token varchar(256),
	created_at timestamp without time zone NOT NULL DEFAULT now(),
	updated_at timestamp without time zone NOT NULL DEFAULT now()
);
create unique index users_email_uindex
	on users (email);

create table messages
(
	id serial not null
		constraint history_pk
			primary key,
	sender varchar(128) not null,
	receiver varchar(128) not null,
	text text not null,
	created_at timestamp without time zone NOT NULL DEFAULT now()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table messages;
drop table users;