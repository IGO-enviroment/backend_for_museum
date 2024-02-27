/* Таблица заказов */
BEGIN;

create table orders (
    id serial PRIMARY KEY,              -- Уникальный ID 
    email VARCHAR(255) UNIQUE,          -- Уникальная почта
    password_digest VARCHAR(255),       -- Зашифрованный пароль
    password_reset_token VARCHAR(255),  -- Токен для восстановления пароля
    password_reset_sent_at TIMESTAMP,   -- Время установки токена восстановления пароля

    -- Таймстамп
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

COMMIT;
