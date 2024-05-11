{{define "CreateUser"}}
INSERT INTO gophermart.users (login, password) VALUES ($1, $2) ON CONFLICT (login) DO NOTHING RETURNING id;
{{end}}
