create database wallet;
create extension "uuid-ossp";
create table wallets(id serial not null primary key, name text not null unique, balance numeric default 0);
create table transactions(id serial not null primary key, wallet_id integer not null references wallets(id), operation smallint not null, amount numeric not null, balance_before numeric not null, comment text, created_at timestamp, trn_id uuid not null);
alter table transactions add constraint wallet_trn_unq unique(wallet_id, trn_id);