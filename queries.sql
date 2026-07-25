-- name: FindResultsByDescription :many
select * from results where description like :description;

-- name: AddResult :exec
insert into results(id,address,description) values(:id,:address,:description);
