create table
if not exists categories
(
    id         int auto_increment
        primary key,
    name       varchar
(255) not null,
    created_at date         not null,
    updated_at date         not null,
    constraint category_id_uindex
        unique
(id)
);

alter table categories
    add primary key (id);

create table
if not exists customers
(
    id         int auto_increment
        primary key,
    name       varchar
(255) not null,
    email      varchar
(255) not null,
    password   varchar
(255) not null,
    address    varchar
(255) not null,
    created_at date         not null,
    updated_at date         null,
    constraint email
        unique
(email)
);

create table
if not exists products
(
    id          int auto_increment
        primary key,
    category_id int          not null,
    name        varchar
(255) not null,
    description varchar
(255) null,
    image_url   varchar
(255) not null,
    preview_url varchar
(255) null,
    price       varchar
(255) not null,
    created_at  date         not null,
    updated_at  date         null,
    constraint products_category_id_fk
        foreign key
(category_id) references categories
(id)
);

create table
if not exists transactions
(
    id          int auto_increment
        primary key,
    customer_id int  not null,
    date        date not null,
    created_at  date not null,
    updated_at  date null
);

create table
if not exists transaction_products
(
    id             int auto_increment,
    transaction_id int  null,
    product_id     int  null,
    qty            int  null,
    created_at     date not null,
    updated_at     date not null,
    constraint transaction_products_id_uindex
        unique
(id),
    constraint transaction_products_products_id_fk
        foreign key
(product_id) references products
(id),
    constraint transaction_products_transactions_id_fk
        foreign key
(transaction_id) references transactions
(id)
);

alter table transaction_products
    add primary key (id);

create table
if not exists carts
(
    customer_id int  not null,
    product_id int  not null,
    created_at  date not null,
    updated_at  date null,
    constraint carts_customer_id_fk
        foreign key
(customer_id) references customers
(id),
    constraint carts_product_id_fk
        foreign key
(product_id) references products
(id)
);

alter table carts
    add primary key (customer_id, product_id);

create table
if not exists wishlists
(
    customer_id int  not null,
    product_id int  not null,
    created_at  date not null,
    updated_at  date null,
    constraint  wishlists_customer_id_fk
        foreign key
(customer_id) references customers
(id),
    constraint wishlists_product_id_fk
        foreign key
(product_id) references products
(id)
);

alter table wishlists
    add primary key (customer_id, product_id);


