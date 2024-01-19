CREATE TABLE  IF NOT EXISTS users (
    "id" bigserial primary key,
    "name" character varying(250) NOT NULL DEFAULT '',
    "surname" character varying(250) NOT NULL DEFAULT '',
    "patronymic" character varying(250) NOT NULL DEFAULT '',
    "age" int NOT NULL DEFAULT 0,
    "probability" numeric(17,0) NOT NULL DEFAULT 0,
    "gender" character varying(250) NOT NULL DEFAULT '',
    "country" json NOT NULL DEFAULT '{}'::json,
    "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL,
    "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL
    );