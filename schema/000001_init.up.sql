CREATE TABLE company
(
    id int not null unique,
    title varchar(255)
);

CREATE TABLE bill
(
    id serial not null unique,
    amount int,
    description varchar(255),
    email varchar(255),
    company_id int,
    time timestamp
);

CREATE TABLE users
(
    id serial not null unique,
    chat_id bigint unique,
    condition varchar(255) default 'started',
    current_bill_id int references bill (id)  on delete set null,
    company_id int references company (id) on delete set null
);





