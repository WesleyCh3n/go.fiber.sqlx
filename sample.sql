create table tbl_hwdevice (
id            serial       not null primary key,
parent_id     int          not null,
name          varchar(256) not null,
description   varchar(256) not null,
hardware_id   int          not null,
enumerator_id int          not null
);

do $$
begin
for r in 1..20 loop
  insert into tbl_hwdevice values ( default,
  r, concat('foobar', r), concat('foobar',r,'msg'), r, r);
end loop;
end;
$$;

create table tbl_fileinfo (
id             serial       not null primary key,
name           varchar(256) not null,
description    varchar(256) not null,
date           timestamp    default now(),
size           int          not null,
md5            varchar(256) not null,
catalog_id     int          not null,
vendor_id      int          not null,
ver_major      int          not null,
ver_minor      int          not null,
ver_release    int          not null,
ver_build      int          not null,
ver_revision   varchar(256) not null,
upload_user_id int          not null,
upload_date    timestamp    default now()
);

do $$
begin
for r in 1..20 loop
  insert into tbl_fileinfo values (default,
  concat('foobar', r), concat('foobar',r,'msg'),
  now()::date + r, r*10,
  concat('md5-', r), r, r, r, r, r, r,
  concat('revision', r), r,
  now()::date + r);
end loop;
end;
$$;
