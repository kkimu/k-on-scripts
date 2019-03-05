create table if not exists artists (
	id varchar(36) primary key,
	name varchar(256) not null,
	kanaPrefix varchar(128) not null,
	createdAt timestamp default CURRENT_TIMESTAMP,
	updatedAt timestamp default CURRENT_TIMESTAMP
);

