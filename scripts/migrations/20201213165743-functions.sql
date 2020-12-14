
-- +migrate Up
-- +migrate StatementBegin
create or replace function get_product(id int, q int)
returns int
language plpgsql
as
$$
declare
   avail_quantity integer;
begin
   	select pq.quantity into avail_quantity from products_q pq where product_id = id for update;
  	if avail_quantity < q then
  		return 0;
  	else
  		update products_q set quantity = avail_quantity-q where product_id = id;
  		return 1;
	end if;
commit;
end;
$$
language plpgsql;
-- +migrate StatementEnd

-- +migrate Down
DROP FUNCTION get_product;
