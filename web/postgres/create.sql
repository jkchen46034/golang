create table userinfo
(
	uid serial not null,
	username character varying(100) not null,
	department character varying(500) not null,
	created date,
	constraint userinfo_pkey primary key(uid)
)
with (OIDS=FALSE)
