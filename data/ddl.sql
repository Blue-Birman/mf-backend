create table customers
(
    id         int auto_increment
        primary key,
    name       varchar(255) not null,
    email      varchar(255) not null,
    password   varchar(255) not null,
    address    varchar(255) not null,
    created_at date         not null,
    updated_at date         null,
    constraint email
        unique (email)
);

create table products
(
    id         int auto_increment
        primary key,
    name       varchar(255) not null,
    image_url  varchar(255) not null,
    price      varchar(255) not null,
    created_at date         not null,
    updated_at date         null
);

create table transactions
(
    id          int auto_increment
        primary key,
    customer_id int  not null,
    date        date not null,
    created_at  date not null,
    updated_at  date null
);

create table transaction_products
(
    id             int auto_increment,
    transaction_id int null,
    product_id     int null,
    qty            int null,
    constraint transaction_products_id_uindex
        unique (id),
    constraint transaction_products_products_id_fk
        foreign key (product_id) references products (id),
    constraint transaction_products_transactions_id_fk
        foreign key (transaction_id) references transactions (id)
);

alter table transaction_products
    add primary key (id);


