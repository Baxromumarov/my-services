CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY NOT NULL,
    first_name varchar(255),
    last_name varchar(255),
    email varchar(255),
    bio varchar(255),
    phone_numbers text[],
    address_id text,
    typeid text,
    status text,
    createdat timestamp ,
    updatedat timestamp ,
    deletedat timestamp ,
    user_name varchar(255),
    password varchar(255),
    email_code varchar(255)
    );