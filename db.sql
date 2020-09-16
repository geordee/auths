drop table if exists orgs cascade;
create table orgs (
    id      bigint primary key,
    name    text
);
insert into orgs values(1, 'acme');

drop table if exists users cascade;
create table users (
    id      bigint primary key,
    org_id  bigint references orgs(id),
    name    text
);
insert into users values(1, 1, 'geordee');

drop table if exists roles cascade;
create table roles (
    id      bigint primary key,
    name    text
);
insert into roles values(1, 'user');

drop table if exists scopes cascade;
create table scopes (
    id      bigint primary key,
    role_id bigint references roles(id),
    name    text
);
insert into scopes values(1, 1, 'read:data');

drop table if exists users_roles cascade;
create table users_roles (
    user_id bigint references users(id),
    role_id bigint references roles(id)
);
insert into users_roles values(1, 1);

select o.name as org
     , u.name as user
     , r.name as role
     , s.name as scope
from users_roles x
  inner join users u on x.user_id = u.id
  inner join orgs  o on u.id = o.id
  inner join roles r on x.role_id = r.id
  inner join scopes s on s.role_id = r.id
where u.name = 'geordee'
;
