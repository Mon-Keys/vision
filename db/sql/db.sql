SET SYNCHRONOUS_COMMIT = 'off';

CREATE EXTENSION IF NOT EXISTS citext;

-- .__                      .__    .___
-- |  |   ____  ____   ____ |__| __| _/
-- |  | _/ __ \/  _ \ /    \|  |/ __ |
-- |  |_\  ___(  <_> )   |  \  / /_/ |
-- |____/\___  >____/|___|  /__\____ |
--           \/           \/        \/
--                     .__  .__
-- ______   ___________|  | |__| ____
-- \____ \_/ __ \_  __ \  | |  |/    \
-- |  |_> >  ___/|  | \/  |_|  |   |  \
-- |   __/ \___  >__|  |____/__|___|  /
-- |__|        \/                   \/
-- _______________   ________ ________
-- \_____  \   _  \  \_____  \\_____  \
--  /  ____/  /_\  \  /  ____/ /  ____/
-- /       \  \_/   \/       \/       \
-- \_______ \_____  /\_______ \_______ \
--         \/     \/         \/       \/
-- terminal
DROP TABLE IF EXISTS terminal;

CREATE UNLOGGED TABLE IF NOT EXISTS terminal (
    terminal_id serial NOT NULL PRIMARY KEY,
    terminal_expiration_date date NOT NULL,
    terminal_name varchar(256) NOT NULL,
    terminal_latitude real,
    terminal_longitude real
);

-- roles
DROP TABLE IF EXISTS roles;

CREATE UNLOGGED TABLE IF NOT EXISTS roles (
    role_id serial NOT NULL PRIMARY KEY,
    role_name varchar(256) NOT NULL,
    role_access_level smallint NOT NULL
);

-- user
DROP TABLE IF EXISTS users;

CREATE UNLOGGED TABLE IF NOT EXISTS users (
    user_id serial NOT NULL PRIMARY KEY,
    user_password_hash varchar(256) NOT NULL,
    user_email varchar(128) NOT NULL UNIQUE,
    user_created_date timestamptz DEFAULT now() NOT NULL
);

-- account
DROP TABLE IF EXISTS account;

CREATE UNLOGGED TABLE IF NOT EXISTS account (
    account_id serial NOT NULL PRIMARY KEY,
    account_role_id int REFERENCES roles (role_id),
    account_user_id int REFERENCES users (user_id),
    account_fullname varchar(256) NOT NULL UNIQUE
);

-- pass
DROP TABLE IF EXISTS pass;

CREATE UNLOGGED TABLE IF NOT EXISTS pass (
    pass_id serial NOT NULL PRIMARY KEY,
    pass_account_id int REFERENCES account (account_id),
    pass_dynamic_qr boolean NOT NULL,
    pass_expiration_date timestamptz NOT NULL,
    pass_issue_date timestamptz DEFAULT now() NOT NULL,
    pass_name varchar(256) NOT NULL,
    pass_secure_data varchar(1024) NOT NULL
);

-- passage
DROP TABLE IF EXISTS passage;

CREATE UNLOGGED TABLE IF NOT EXISTS passage (
    passage_id serial NOT NULL PRIMARY KEY,
    passage_terminal_id int REFERENCES terminal (terminal_id) NOT NULL,
    pass_id int REFERENCES pass (pass_id) NOT NULL,
    passage_status int NOT NULL,
    is_exit boolean NOT NULL,
    passage_datetime timestamptz DEFAULT now()
);

-- event type
DROP TABLE IF EXISTS event_type;

CREATE UNLOGGED TABLE IF NOT EXISTS event_type (
    event_type_id serial NOT NULL PRIMARY KEY,
    event_type_name varchar(256) NOT NULL
);

-- pass_declaration
DROP TABLE IF EXISTS pass_declaration;

CREATE UNLOGGED TABLE IF NOT EXISTS pass_declaration (
    pass_declaration_id serial NOT NULL PRIMARY KEY,
    pass_declaration_name citext NOT NULL
);

-- declaration events
DROP TABLE IF EXISTS declaration_events;

CREATE UNLOGGED TABLE IF NOT EXISTS declaration_events (
    declaration_events_id serial NOT NULL PRIMARY KEY,
    declaration_events_event_type_id int REFERENCES event_type (event_type_id),
    declaration_pass_declaration_id int REFERENCES pass_declaration (pass_declaration_id),
    pass_declaration_creator_id int REFERENCES account (account_id),
    pass_declaration_create_date timestamptz DEFAULT now(),
    pass_declaration_comment citext,
    pass_declaration_requests_count smallint NOT NULL CHECK (pass_declaration_requests_count > 0)
);

-- pass request
DROP TABLE IF EXISTS pass_request;

CREATE UNLOGGED TABLE IF NOT EXISTS pass_request (
    pass_request_id serial NOT NULL PRIMARY KEY,
    pass_request_account_id int REFERENCES account (account_id),
    pass_request_declaration_id int REFERENCES pass_declaration (pass_declaration_id),
    pass_request_approved boolean NOT NULL,
    pass_request_comment citext
);

