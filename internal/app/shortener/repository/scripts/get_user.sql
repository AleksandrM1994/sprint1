select id, login, password
from users
where login = $1
  and password = $2