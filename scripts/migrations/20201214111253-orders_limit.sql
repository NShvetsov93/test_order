
-- +migrate Up
-- +migrate StatementBegin
create or replace function check_limit()
returns int
language plpgsql
as
$$
declare
   now_quantity integer;
begin
   	select ol.quantity into now_quantity from orders_limit ol where id = 1 for update;
  	if now_quantity >= 3  then
  		return 0;
  	else
  		update orders_limit set quantity = now_quantity+1 where id = 1;
  		return 1;
	end if;
commit;
end;
$$
-- +migrate StatementEnd

-- +migrate Down
DROP FUNCTION check_limit;
