CREATE TABLE IF NOT EXISTS "comment"
(
    "id"          INTEGER PRIMARY KEY NOT NULL,
    "link_id"     INTEGER            NULL,
    "context"     VARCHAR(240)       NULL,
    "score"       INTEGER DEFAULT 0  NULL,
    "agree"       INTEGER DEFAULT 0  NULL,
    "disagree"    INTEGER DEFAULT 0  NULL,
    "create_by"   INTEGER            NULL,
    "create_time" TIMESTAMP          NULL,
    "update_time" TIMESTAMP          NULL,
    "delete_time" TIMESTAMP          NULL
);
CREATE TABLE IF NOT EXISTS "link"
(
    "id"          INTEGER PRIMARY KEY NOT NULL,
    "link"        VARCHAR(380)       NULL,
    "title"       VARCHAR(255)       NULL,
    "l_group"     VARCHAR(255)       NULL,
    "tags"        VARCHAR(140)       NULL,
    "l_from"      VARCHAR(1)            NULL,
    "topic_id"    INTEGER            NULL,
    "score"       INTEGER            NULL default 0,
    "agree"       INTEGER            NULL default 0,
    "comment_cnt" int                     default 0,
    "disagree"    INTEGER            NULL default 0,
    "create_time" TIMESTAMP          NULL,
    "update_time" TIMESTAMP          NULL,
    "delete_time" TIMESTAMP          NULL,
    "create_by"   int                NULL
);
CREATE TABLE IF NOT EXISTS "topic"
(
    "id"          INTEGER       NOT NULL primary key,
    "name"        VARCHAR(140) NULL,
    "short_code" varchar(32) default '',
    "icon"        varchar(380) null,
    "scope"       smallint default 1,
    "tags"        VARCHAR(140) NULL,
    "create_by"   varchar(32)  NULL,
    "lang"        VARCHAR(4)  default 'zh',
    "score"       int          NULL,
    "agree"       int          NULL,
    "disagree"    int          NULL,
    "create_time" timestamp    NULL,
    "update_time" timestamp    NULL,
    "delete_time" timestamp    NULL
);
create table if not exists "topic_rel"
(
    "aid"         int         not null,
    "bid"         int         not null,
    "position"    smallint    not null default 1,
    "create_by"   varchar(32) null,
    "predicate"   varchar(140),
    "create_time" timestamp   NULL,
    "delete_time" timestamp   NULL
);

create table hot_topic
(
    id            int,
    expire        timestamp,
    "create_time" timestamp NULL,
    "update_time" timestamp NULL,
    "delete_time" timestamp NULL
);

create table if not exists "topic_alias"
(
    "alias" varchar (32) not null unique ,
    "topic_id" integer not null
);

create view topic_shadow as
select name, id, short_code
from topic
union all
select a.alias, a.topic_id, t.short_code
from topic_alias a
         inner join topic t on a.topic_id = t.id;


drop table if exists account;
CREATE TABLE IF NOT EXISTS "account"
(
    "id"          INTEGER       NOT NULL primary key autoincrement ,
    "user_name"   VARCHAR(140) NULL,
    "nick_name"   VARCHAR(140) NULL,
    "password"    varchar(32) not null,
    "avatar"      varchar(380) null,
    "score"       int          NULL,
    "create_time" timestamp    NULL,
    "update_time" timestamp    NULL,
    "delete_time" timestamp    NULL
);
insert into account
values (9, 'seven4x', 'seven4x', '2d12ed02b04144390546484a6ad73b23','', 0, null, null, null);

drop table if exists  register_code  ;
create table if not exists "register_code"
(
    "id"      INTEGER  not null primary key,
    "user_id" int     not null,
    "code"    varchar(8) not null unique
);
insert into register_code values (1,0,'seven4x');

create table if not exists "register_info"
(
    "id"           INTEGER  not null primary key,
    "code"         varchar(8) not null,
    create_by      int,
    used_by        int,
    used_user_name varchar(140),
    used_time      timestamp
);


CREATE TABLE IF NOT EXISTS "user_vote"
(
    "user_id" int     not null,
    "id"      int     NOT NULL,
    "type"    varchar(1) not null,
    "is_like" varchar(1) not null,
    unique (user_id, type, id)
);

