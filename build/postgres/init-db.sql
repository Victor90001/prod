create table Section(
    id serial not null primary key,
    name varchar
);

create table Groups(
    id serial not null primary key,
    sectionID integer not null references Section(id) on delete cascade,
    name varchar
);

create table Product(
    id serial not null primary key,
    groupID integer not null references Groups(id) on delete cascade,
    name varchar,
    code int,
    price int,
    prodDate date,
    describe varchar,
    size varchar,
    country varchar,
    addParam varchar
);

create table Users(
    id serial not null primary key,
    login varchar,
    pwd varchar,
    user_type int
);

insert into users (login, pwd, user_type) values ('admin', 'admin', 1), ('user', 'user', 0);
