SET SYNCHRONOUS_COMMIT = 'off';
create extension if not exists citext;


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

CREATE UNLOGGED TABLE IF NOT EXISTS terminal
(
    terminal_id serial not null primary key,
    terminal_expiration_date date not null,
    terminal_name varchar(256) not null
);


-- roles

DROP TABLE IF EXISTS roles;

CREATE UNLOGGED TABLE IF NOT EXISTS roles
(
    role_id serial not null primary key,
    role_name varchar(256) not null,
    roll_access_level smallint not null
);


-- user 

DROP TABLE IF EXISTS user;

CREATE UNLOGGED TABLE IF NOT EXISTS user
(
    user_id serial not null primary key,
    user_password_hash varchar(256) not null,
    user_login varchar(256) not null unique,
    user_email varchar(128) not null unique,
    user_created_date timestampz default now() not null,
);


-- account

DROP TABLE IF EXISTS account;

CREATE UNLOGGED TABLE IF NOT EXISTS account
(
    account_id serial not null primary key,
    account_role_id int REFERENCES roles(role_id),
    account_user_id int REFERENCES user(user_id),
    account_fullname varchar(256) not null unique,
    account_photo varbinary(2048)
);


-- pass

DROP TABLE IF EXISTS pass;

CREATE UNLOGGED TABLE IF NOT EXISTS pass
(
    pass_id serial not null primary key,
    pass_account_id int REFERENCES account(account_id),
    pass_dynamic_qr boolean not null,
    pass_expiration_date timestampz not null,
    pass_issue_date timestampz DEFAULT now() not null,
    pass_name varchar(256) not null,
    pass_secure_data varbinary(1024) not null,
);


-- passage

DROP TABLE IF EXISTS passage;

CREATE UNLOGGED TABLE IF NOT EXISTS passage 
(
    passage_id serial not null primary key,
    passage_terminal_id int REFERENCES terminal(terminal_id) not null,
    pass_id int REFERENCES pass(pass_id) not null, 
    passage_status int not null,
    is_exit boolean not null,
    passage_datetime timestampz DEFAULT now()
);

-- event type

DROP TABLE IF EXISTS event_type;

CREATE UNLOGGED TABLE IF NOT EXISTS event_type
(
    event_type_id serial not null primary key,
    event_type_name varchar(256) not null
);

-- pass_declaration

DROP TABLE IF EXISTS pass_declaration;

CREATE UNLOGGED TABLE IF NOT EXISTS pass_declaration
(
    pass_declaration_id serial not null primary key,
    pass_declaration_name citext not null,
);

-- declaration events

DROP TABLE IF EXISTS declaration_events;

CREATE UNLOGGED TABLE IF NOT EXISTS declaration_events
(
    declaration_events_id serial not null primary key,
    declaration_events_event_type_id int REFERENCES event_type(event_type_id),
    declaration_pass_declaration_id int REFERENCES pass_declaration(pass_declaration_id),
    pass_declaration_creator_id int REFERENCES account(account_id),
    pass_declaration_create_date timestampz DEFAULT now(),
    pass_declaration_comment citext,
    pass_declaration_requests_count smallint not null
    CHECK (pass_declaration_requests_count > 0)
);


-- pass request

DROP TABLE IF EXISTS pass_request;

CREATE UNLOGGED TABLE IF NOT EXISTS pass_request
(
    pass_request_id serial not null primary key,
    pass_request_account_id int REFERENCES account(account_id),
    pass_request_declaration_id int REFERENCES pass_declaration(pass_declaration_id),
    pass_request_approved boolean not null,
    pass_request_comment citext
);