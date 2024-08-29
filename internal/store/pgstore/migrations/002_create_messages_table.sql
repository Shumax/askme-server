CREATE TABLE IF NOT EXISTS messages (
  "id"  uuid PRIMARY KEY  NOT NULL DEfAULT gen_random_uuid(),
  "room_id" uuid NOT NULL,
  "message" varchar(255) NOT NULL,
  "reaction_count"  BIGINT NOT NULL DEfAULT 0,
  "answered"  BOOLEAN NOT NULL DEfAULT false,

  FOREIGN KEY (room_id) REFERENCES rooms(id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS messages;