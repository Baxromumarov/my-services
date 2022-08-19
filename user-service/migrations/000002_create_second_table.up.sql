CREATE TABLE IF NOT EXISTS posts (
    id uuid PRIMARY KEY NOT NULL,
    description varchar(255) NOT NULL,
    link varchar(255) NOT NULL
);