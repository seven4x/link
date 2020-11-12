CREATE TABLE IF NOT EXISTS "topic"
(
    "id"          SERIAL       NOT NULL primary key,
    "name"        VARCHAR(140) NULL,
    "icon"        varchar(380) null,
    "scope"       smallint default 1,
    "tags"        VARCHAR(140) NULL,
    "create_by"   varchar(32)  NULL,
    "lang"        char(4)  default 'zh',
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
)