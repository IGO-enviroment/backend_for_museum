/* Таблица контента для постов */
BEGIN;

create table contents (
    id SERIAL PRIMARY KEY,          -- Уникальный ID

    type_value VARCHAR(50),         -- Тип контента

    data_value TEXT,                -- Данные блока с контентом: [текст, сслыка и т.п.]

    order_value INTEGER,            -- Порядок вывода

    model_id INTEGER,               -- ID поста к которому привязан
    model_type VARCHAR(50),         -- Название таблицы для корректного поиска

    options JSON NOT NULL DEFAULT '{}'::jsonb,   -- Дополнительные настройки

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
