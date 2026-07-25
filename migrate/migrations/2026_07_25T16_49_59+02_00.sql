-- migration: down
drop table if exists results;

-- migration: up
create table if not exists results(
    id varchar(36) primary key,
    address varchar(1024) not null default '',
    description varchar(1024) not null default ''
);
