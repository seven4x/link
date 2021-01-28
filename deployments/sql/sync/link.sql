CREATE TABLE IF NOT EXISTS "link"
(
    "id"          SERIAL PRIMARY KEY NOT NULL,
    "link"        VARCHAR(380)       NULL,
    "title"       VARCHAR(255)       NULL,
    "l_group"     VARCHAR(255)       NULL,
    "tags"        VARCHAR(140)       NULL,
    "l_from"      CHAR(1)            NULL,
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
