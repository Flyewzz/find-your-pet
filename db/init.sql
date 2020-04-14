CREATE EXTENSION postgis;

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

create table files
(
  file_id serial       not null
    constraint files_pk
      primary key,
  name    varchar(512) not null,
  path    varchar(1024)
);

alter table files
  owner to postgres;

create table lost
(
  id          serial                  not null,
  type_id     integer                 not null
    constraint lost_animaltypes_id_fk
      references animaltypes,
  vk_id       integer                 not null,
  sex         varchar(4)              not null,
  breed       varchar(50),
  description text,
  status_id   integer                 not null
    constraint lost_statuses_id_fk
      references statuses,
  date        timestamp default now() not null,
  location    geometry                not null,
  picture_id  integer
    constraint lost_files_file_id_fk
      references files
);

alter table lost
  owner to postgres;

create index lost_date_index
  on lost (date desc);

create index lost_gist_location_index
  on lost (location);

create unique index lost_vk_id_uindex
  on lost (vk_id);

create table found
(
  id          serial                  not null,
  type_id     integer                 not null
    constraint found_animaltypes_id_fk
      references animaltypes,
  vk_id       integer                 not null,
  sex         varchar(4)              not null,
  breed       varchar(50),
  description text,
  status_id   integer                 not null
    constraint found_statuses_id_fk
      references statuses,
  date        timestamp default now() not null,
  location    geometry                not null,
  picture_id  integer                 not null
    constraint found_files_file_id_fk
      references files
);

alter table found
  owner to postgres;

create index found_date_index
  on found (date desc);

create unique index found_vk_id_uindex
  on found (vk_id);