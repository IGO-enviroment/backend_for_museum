/* Таблица подключенных к пользоватею ролей */
BEGIN;

create table user_roles (
    id SERIAL PRIMARY KEY,                  -- Уникальный ID

    user_id SERIAL REFERENCES users(user_id), -- К какому пользователю относится
    role_id SERIAL REFERENCES roles(role_id), -- С какой ролью есть связь

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
