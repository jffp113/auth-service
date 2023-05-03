CREATE TABLE auth_users (
    id bigint,
    full_name character varying(255),
    username character varying(255),
    email character varying(255),
    hash character varying(255),
    preferences jsonb,
    created_at timestamp,
    updated_at timestamp
);

INSERT INTO auth_users (full_name, username, email, hash, preferences)
VALUES ('Jorge', 'username', 'email', 'password', '{}');