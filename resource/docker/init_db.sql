create table departments
(
    id integer default nextval('department_id_seq'::regclass) not null
        constraint department_pk
            primary key,
    name varchar(100),
    short_name varchar(20),
    description varchar(250),
    created_at timestamp,
    updated_at timestamp
);

alter table departments owner to lopmartyn;

create unique index department_id_uindex
    on departments (id);

create table employees
(
    uuid text not null,
    request_id text not null,
    title_id integer,
    first_name text,
    last_name text,
    gender_id integer,
    department_id integer,
    phone_no text,
    created_at timestamp,
    updated_at timestamp,
    emp_id varchar(25) not null
        constraint employee_pk
            primary key
);

comment on table employees is 'For keep the employee information';

alter table employees owner to lopmartyn;

create unique index employees_request_id_uindex
    on employees (request_id);

create unique index employees_uuid_uindex
    on employees (uuid);

create unique index employee_emp_id_uindex
    on employees (emp_id);

create table genders
(
    id integer default nextval('gender_id_seq'::regclass) not null
        constraint gender_pk
            primary key,
    name varchar(50),
    description varchar(250),
    created_at timestamp,
    updated_at timestamp
);

alter table genders owner to lopmartyn;

create table registers
(
    uuid text not null
        constraint register_pk
            primary key,
    email varchar(255),
    registered_at timestamp,
    confirm_status boolean,
    confirmed_at timestamp
);

alter table users owner to lopmartyn;

create unique index register_uuid_uindex
    on users (uuid);

