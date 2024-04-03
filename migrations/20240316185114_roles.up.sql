/* Таблица ролей */
BEGIN;

create table roles (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    named VARCHAR(255),         -- Название роли

    property VARCHAR(255),      -- Описание роли и ее доступов 

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
