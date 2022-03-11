/* create ui_test */
CREATE ROLE ui_test LOGIN PASSWORD 'ui_test' NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;

/*
    setup db
*/
CREATE DATABASE ui with ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8' CONNECTION LIMIT = -1 template=template0;
ALTER DATABASE ui OWNER TO ui_test;
ALTER DATABASE ui SET timezone TO 'UTC';
REVOKE USAGE ON SCHEMA public FROM PUBLIC;
REVOKE CREATE ON SCHEMA public FROM PUBLIC;
GRANT USAGE ON SCHEMA public to ui_test;
GRANT CREATE ON SCHEMA public to ui_test;


/*
    create table
*/

\connect ui;
DROP TABLE IF EXISTS users CASCADE;

create table users
(
    acct character varying(255) not null PRIMARY KEY,
    pwd character varying(1000) not null,
    fullname character varying(255) not null,
    created_at timestamp without time zone not null default current_timestamp,
    updated_at timestamp without time zone not null default current_timestamp

);

\connect ui;

ALTER TABLE users OWNER TO ui_test;


