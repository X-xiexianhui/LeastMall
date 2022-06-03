create table product
(
    id           int auto_increment
        primary key,
    product_name varchar(64) not null,
    price        float(6, 2) null,
    descriptions text        null,
    cover        mediumtext  null
);

create table banner
(
    id         int auto_increment
        primary key,
    product_id int        null,
    image      mediumtext null,
    constraint banner_product_product_id_fk
        foreign key (product_id) references product (id)
);

create table images
(
    id         int auto_increment
        primary key,
    product_id int        null,
    image      mediumtext null,
    constraint pictures_product_product_id_fk
        foreign key (product_id) references product (id)
);


