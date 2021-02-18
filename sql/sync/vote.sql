
CREATE TABLE IF NOT EXISTS "user_vote"
(
    "user_id" int     not null,
    "id"      int     NOT NULL,
    "type"    char(1) not null,
    "is_like" char(1) not null,
    unique (user_id, type, id)
);