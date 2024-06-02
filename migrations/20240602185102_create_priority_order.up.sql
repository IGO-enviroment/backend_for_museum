/* Таблица установки приоритетов */
BEGIN;

create table priority_orders (
    id SERIAL PRIMARY KEY,              -- Уникальный ID

    priority INTEGER,

    model_id INTEGER,
    model_type VARCHAR(50),

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

COMMIT;
