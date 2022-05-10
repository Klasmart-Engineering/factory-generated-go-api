CREATE TABLE if not exists go_template_table
(
    id         SERIAL PRIMARY KEY,
    name       character varying(255)                             NOT NULL UNIQUE,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO go_template_table (id, name)
VALUES (1, 'Example Data 1'),
       (2, 'Example Data 2');
