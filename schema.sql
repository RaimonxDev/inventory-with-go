create database if not exists `max_inventory`;

create table if not exists USERS {
    id int not null auto_increment,
    email varchar(255) not null,
    name varchar(255) not null,
    password varchar(255) not null,
    primary key (id)
}
create table if not exists PRODUCTS {
    id int not null auto_increment,
    name varchar(255) not null,
    description varchar(255) not null,
    price float not null,
    created_by int not null,
    primary key (id),
    foregin key (created_by) references USERS(id)
}

create table if not exists ROLES {
    id int not null auto_increment,
    name varchar(255) not null,
    primary key (id)
}

create table if not exists USER_ROLES {
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    primary key (id),
    foregin key (user_id) references USERS(id),
    foregin key (role_id) references ROLES(id)
}

