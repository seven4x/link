CREATE TABLE IF NOT EXISTS "t_user"
(
    "id"          SERIAL       NOT NULL primary key,
    "name"        VARCHAR(140) NULL,
    "avatar"      varchar(380) null,
    "score"       int          NULL,
    "create_time" timestamp    NULL,
    "update_time" timestamp    NULL,
    "delete_time" timestamp    NULL
);