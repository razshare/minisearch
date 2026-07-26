-- name: FindResultsByDescription :many
select * from results where description like :description limit :count offset :offset;

-- name: AddResult :exec
insert into results(id,address,description) values(:id,:address,:description);

-- name: CountResultsByDescription :one
select count(*) from results where description like :description;