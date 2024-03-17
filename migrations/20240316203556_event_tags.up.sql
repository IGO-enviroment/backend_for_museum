/* Таблица связи между мероприятиями и тегами */
BEGIN;

create table event_tags (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    event_id SERIAL REFERENCES events(event_id),    -- К какому мероприятию относится
    tag_id SERIAL REFERENCES tags(tag_id),          -- С каким тегом есть связь

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
