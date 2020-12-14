
-- +migrate Up
-- +migrate StatementBegin
create or replace function reset_limit()
returns int
language plpgsql
as
$$
declare
   now_quantity integer;
begin
   	select ol.quantity into now_quantity from orders_limit ol where id = 1 for update;
  	if now_quantity > 0  then
  		update orders_limit set quantity = now_quantity-1 where id = 1;
	end if;
commit;
end;
$$
language plpgsql;
-- +migrate StatementEnd

-- +migrate Down
DROP FUNCTION reset_limit;
