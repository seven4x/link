CREATE TABLE IF NOT EXISTS "comment"
(
    "id"          SERIAL PRIMARY KEY NOT NULL,
    "link_id"     INTEGER            NULL,
    "context"     VARCHAR(240)       NULL,
    "score"       INTEGER DEFAULT 0  NULL,
    "agree"       INTEGER DEFAULT 0  NULL,
    "disagree"    INTEGER DEFAULT 0  NULL,
    "create_by"   INTEGER            NULL,
    "create_time" TIMESTAMP          NULL,
    "update_time" TIMESTAMP          NULL,
    "delete_time" TIMESTAMP          NULL
)
