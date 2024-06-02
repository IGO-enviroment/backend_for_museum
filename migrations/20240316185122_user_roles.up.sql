/* Таблица подключенных к пользоватею ролей */
BEGIN;

create table user_roles (
    id SERIAL PRIMARY KEY,                  -- Уникальный ID

    user_id int REFERENCES users(id), -- К какому пользователю относится
    role_id int REFERENCES roles(id), -- С какой ролью есть связь

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
