{{define "CreateUser"}}
INSERT INTO gophermart.users (login, password) VALUES ($1, $2) ON CONFLICT (login) DO NOTHING RETURNING id;
{{end}}

{{define "GetUserForLogin"}}
SELECT id, login, password FROM gophermart.users WHERE login = $1;
{{end}}

{{define "CreateOrder"}}
INSERT INTO gophermart.orders (number, user_id) VALUES ($1, $2) ON CONFLICT (number) DO NOTHING RETURNING number;
{{end}}

{{define "CheckUserOrder"}}
SELECT EXISTS(SELECT 1 FROM gophermart.orders WHERE number = $1 AND user_id = $2);
{{end}}
