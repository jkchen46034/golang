CREATE or REPLACE function insert_userinfo(name varchar(100), dept varchar(500)) 
returns void AS $$ 
BEGIN
	INSERT into userinfo(username, department, created) values (name, dept, now());
end;
$$ LANGUAGE plpgsql
