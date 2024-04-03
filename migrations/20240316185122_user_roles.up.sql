/* Таблица подключенных к пользоватею ролей */
BEGIN;

create table user_roles (
    id SERIAL PRIMARY KEY,                  -- Уникальный ID

    user_id SERIAL REFERENCES users, -- К какому пользователю относится
    role_id SERIAL REFERENCES roles, -- С какой ролью есть связь

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
