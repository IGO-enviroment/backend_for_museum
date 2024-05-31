/* Таблица для проверки данных */
BEGIN;

create table verifications (
    id SERIAL PRIMARY KEY,          -- Уникальный ID   

    code TEXT,                          -- Код для проверки/верификации
    until_at TIMESTAMP,                 -- Дата действительности

    expired BOOLEAN NOT NULL DEFAULT false, -- Просрочен или нет
    sended BOOLEAN NOT NULL DEFAULT false,  -- Отправлен или нет

    model_id INTEGER,                   -- ID сущности к которой привязан
    model_type VARCHAR(50),             -- Название таблицы для корректного поиска

    options JSON NOT NULL DEFAULT '{}'::jsonb,   -- Дополнительный настройки

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
