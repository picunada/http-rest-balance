-- create table balance (
--                          user_id uuid primary key default gen_random_uuid(),
--                          balance numeric
-- );

create table "user"
(
    id uuid primary key default gen_random_uuid(),
    name varchar(100) not null
    );

create table "account"
(
    id      uuid primary key default gen_random_uuid(),
    user_id uuid not null references "user"(id),
    balance numeric
        constraint positive_balance CHECK (balance >= 0)
);

create table "operation"
(
    id         serial primary key,
    account_id uuid not null,
    constraint account_fk foreign key (account_id) references "account"(id),
    type       varchar,
    "from"     int,
    created_at timestamp default now()
);