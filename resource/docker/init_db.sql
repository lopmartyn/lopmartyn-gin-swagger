create table if not exists departments
(
    id          integer default nextval('department_id_seq'::regclass) not null
        constraint department_pk
            primary key,
    name        varchar(100),
    short_name  varchar(20),
    description varchar(250),
    created_at  timestamp,
    updated_at  timestamp
);

alter table departments
    owner to lopmartyn;

create unique index if not exists department_id_uindex
    on departments (id);

create table if not exists genders
(
    id          integer default nextval('gender_id_seq'::regclass) not null
        constraint pk_genders
            primary key,
    name        varchar(50),
    description varchar(250),
    created_at  timestamp,
    updated_at  timestamp
);

alter table genders
    owner to lopmartyn;

create table if not exists users
(
    uuid          text not null
        constraint register_pk
            primary key,
    email         varchar(255),
    registered_at timestamp,
    status        varchar(40),
    password      text
);

alter table users
    owner to lopmartyn;

create unique index if not exists register_uuid_uindex
    on users (uuid);

create table if not exists titles
(
    id          serial
        constraint titles_pk
            primary key,
    name        varchar(70),
    description varchar(250)
);

alter table titles
    owner to lopmartyn;

create table if not exists employees
(
    uuid          text        not null,
    request_id    text        not null,
    title_id      integer
        constraint fk_employees_titles
            references titles,
    first_name    text,
    last_name     text,
    gender_id     integer
        constraint fk_employees_genders
            references genders (id),
    department_id integer
        constraint fk_employees_departments
            references departments,
    phone_no      text,
    created_at    timestamp,
    updated_at    timestamp,
    emp_id        varchar(25) not null
        constraint employee_pk
            primary key
);

comment on table employees is 'For keep the employee information';

alter table employees
    owner to lopmartyn;

create unique index if not exists employees_request_id_uindex
    on employees (request_id);

create unique index if not exists employees_uuid_uindex
    on employees (uuid);

create unique index if not exists employee_emp_id_uindex
    on employees (emp_id);

create unique index if not exists titles_id_uindex
    on titles (id);

