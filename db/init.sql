create table statuses
(
  id   serial      not null
    constraint statuses_pk
      primary key,
  name varchar(60) not null
);

alter table statuses
  owner to postgres;

create table animaltypes
(
  id   serial      not null
    constraint animaltypes_pk
      primary key,
  name varchar(60) not null
);

alter table animaltypes
  owner to postgres;

create table usertypes
(
  id   serial       not null
    constraint usertypes_pk
      primary key,
  name varchar(100) not null
);

alter table usertypes
  owner to postgres;

create table users
(
  id       serial       not null
    constraint users_pk
      primary key,
  utype_id integer      not null
    constraint users_usertypes_id_fk
      references usertypes,
  name     varchar(100) not null
);

alter table users
  owner to postgres;

create table lost
(
  id          serial                  not null
    constraint lost_pk
      primary key,
  type_id     integer                 not null
    constraint lost_animaltypes_id_fk
      references animaltypes,
  author_id   integer                 not null
    constraint lost_users_id_fk
      references users,
  sex         varchar(4)              not null,
  breed       varchar(50),
  description text,
  status_id   integer                 not null
    constraint lost_statuses_id_fk
      references statuses,
  date        timestamp default now() not null,
  place       varchar(400)            not null
);

alter table lost
  owner to postgres;

create index lost_date_index
  on lost (date desc);

create table found
(
  id          serial                  not null
    constraint found_pk
      primary key,
  type_id     integer                 not null
    constraint found_animaltypes_id_fk
      references animaltypes,
  author_id   integer                 not null
    constraint found_users_id_fk
      references users,
  sex         varchar(4)              not null,
  breed       varchar(50),
  description text,
  status_id   integer                 not null
    constraint found_statuses_id_fk
      references statuses,
  date        timestamp default now() not null,
  place       varchar(400)            not null
);

alter table found
  owner to postgres;

create index found_date_index
  on found (date desc);

