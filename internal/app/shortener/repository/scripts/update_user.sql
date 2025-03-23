update users
set cookie        = :cookie,
    cookie_finish = :cookie_finish
where id = :id