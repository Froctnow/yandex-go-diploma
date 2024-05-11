{{define "CreateUser"}}
INSERT INTO gophermart.users (login, password) VALUES ($1, $2) ON CONFLICT (login) DO NOTHING RETURNING id;
{{end}}

{{define "GetUserForLogin"}}
SELECT id, login, password FROM gophermart.users WHERE login = $1;
{{end}}
