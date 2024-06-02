/* Таблица связи между мероприятиями и тегами */
BEGIN;

create table event_tags (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    event_id int REFERENCES events(id),    -- К какому мероприятию относится
    tag_id int REFERENCES tags(id),          -- С каким тегом есть связь

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

COMMIT;
