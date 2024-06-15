CREATE TABLE "user" (
                         "id" bigserial PRIMARY KEY,
                         "username" varchar UNIQUE NOT NULL,
                         "hashed_password" varchar NOT NULL,
                         "full_name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "widget" (
                          "id" bigserial PRIMARY KEY,
                          "name" varchar NOT NULL,
                          "user_id" bigserial NOT NULL,
                          "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "counter" (
                           "id" bigserial PRIMARY KEY,
                           "widget_id" bigserial NOT NULL,
                           "order" integer,
                           "value" integer DEFAULT 0,
                           "increment" integer DEFAULT 0
);

ALTER TABLE "widget" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "counter" ADD FOREIGN KEY ("widget_id") REFERENCES "widget" ("id");
