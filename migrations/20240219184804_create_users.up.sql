/* Таблица пользователей */
BEGIN;

create table users (
    id serial PRIMARY KEY,              -- Уникальный ID 

    email VARCHAR(255) UNIQUE,          -- Уникальная почта
    password_digest VARCHAR(255),       -- Зашифрованный пароль

    is_admin BOOLEAN NOT NULL DEFAULT "0", -- Админ или нет

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
