select id, login, password, cookie, cookie_finish
from users
where id = $1