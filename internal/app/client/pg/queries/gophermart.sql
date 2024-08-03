{{define "CreateUser"}}
INSERT INTO gophermart.users (login, password)
VALUES ($1, $2)
ON CONFLICT (login) DO NOTHING
RETURNING id;
{{end}}

{{define "GetUserForLogin"}}
SELECT id, login, password
FROM gophermart.users
WHERE login = $1;
{{end}}

{{define "CreateOrder"}}
INSERT INTO gophermart.orders (number, user_id)
VALUES ($1, $2)
ON CONFLICT (number) DO NOTHING
RETURNING number;
{{end}}

{{define "CheckUserOrder"}}
SELECT EXISTS(SELECT 1 FROM gophermart.orders WHERE number = $1 AND user_id = $2);
{{end}}

{{define "GetOrders"}}
SELECT number, status, accrual, uploaded_at
FROM gophermart.orders
WHERE user_id = $1
ORDER BY uploaded_at DESC;
{{end}}

{{define "ExpandOrder"}}
UPDATE gophermart.orders SET status = $1, accrual = $2 WHERE number = $3
{{end}}

{{define "GetUserBalance"}}
SELECT balance
FROM gophermart.users
WHERE id = $1;
{{end}}

{{define "GetUserWithdrawn"}}
SELECT SUM(sum) as withdrawn
FROM gophermart.transactions
WHERE user_id = $1
GROUP BY user_id;
{{end}}

{{define "GetUserBalanceForUpdate"}}
SELECT balance
FROM gophermart.users
WHERE id = $1
    FOR UPDATE;
{{end}}

{{define "CreateWithdrawTransaction"}}
INSERT INTO gophermart.transactions (user_id, sum, order_number)
VALUES ($1, $2, $3);
{{end}}

{{define "UpdateUserBalance"}}
UPDATE gophermart.users SET balance = $1 WHERE id = $2;
{{end}}

{{define "IncreaseUserBalance"}}
UPDATE gophermart.users SET balance = balance + $1 WHERE id = $2;
{{end}}

{{define "GetUserWithdraws"}}
SELECT sum, order_number, processed_at
FROM gophermart.transactions
WHERE user_id = $1
ORDER BY processed_at DESC;
{{end}}