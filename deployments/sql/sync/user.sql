
drop table if exists account;
CREATE TABLE IF NOT EXISTS "account"
(
    "id"          SERIAL       NOT NULL primary key,
    "user_name"   VARCHAR(140) NULL,
    "nick_name"   VARCHAR(140) NULL,
    "password"    char(32) not null,
    "avatar"      varchar(380) null,
    "score"       int          NULL,
    "create_time" timestamp    NULL,
    "update_time" timestamp    NULL,
    "delete_time" timestamp    NULL
);
insert into account
values (9, 'seven4x', 'seven4x', '2d12ed02b04144390546484a6ad73b23','', 0, null, null, null);
select * from account;
drop table if exists  register_code  ;
create table if not exists "register_code"
(
    "id"      Serial  not null primary key,
    "user_id" int     not null,
    "code"    char(8) not null unique
);
insert into register_code values (1,0,'seven4x');

create table if not exists "register_info"
(
    "id"           Serial  not null primary key,
    "code"         char(8) not null,
    create_by      int,
    used_by        int,
    used_user_name varchar(140),
    used_time      timestamp
)

