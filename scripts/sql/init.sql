CREATE TYPE ROLE AS ENUM ('staff', 'student', 'dev', 'applicant', 'unknown');
CREATE TYPE BOOKING_TYPE AS ENUM ('places', 'inventory');

CREATE TABLE campus
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE users
(
    id          BIGINT UNIQUE NOT NULL,
    nickname    VARCHAR,
    email       VARCHAR,
    campus_id   INTEGER                DEFAULT NULL REFERENCES campus (id) ON DELETE SET NULL,
    role        ROLE          NOT NULL DEFAULT 'unknown',
    handle_step INTEGER                DEFAULT 0,
    created_at  TIMESTAMP     NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP     NOT NULL DEFAULT now()
);



INSERT INTO campus(name)
VALUES ('Москва'),
       ('Казань'),
       ('Новосибирск');

CREATE TABLE Sessions
(
    id         SERIAL PRIMARY KEY,
    user_id    BIGINT    REFERENCES users (id) ON DELETE SET NULL,
    code       int       NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    end_at     TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE category
(
    id        SERIAL PRIMARY KEY,
    name      VARCHAR      NOT NULL,
    type      BOOKING_TYPE NOT NULL,
    create_at TIMESTAMP    NOT NULL DEFAULT now(),
    update_at TIMESTAMP    NOT NULL DEFAULT now()
);

INSERT INTO category(name, type)
VALUES ('Переговорные', 'places'),
       ('Кухни', 'places'),
       ('Игровые', 'places'),
       ('книги', 'inventory');

CREATE TABLE places
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR   NOT NULL,
    description VARCHAR   NOT NULL,
    campus_id   INTEGER REFERENCES campus (id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES category (id) ON DELETE CASCADE,
    floor       INTEGER   NOT NULL,
    room        INTEGER   NOT NULL,
    period      INTERVAL,
    permission  ROLE      NOT NULL DEFAULT 'student',
    create_at   TIMESTAMP NOT NULL DEFAULT now(),
    update_at   TIMESTAMP NOT NULL DEFAULT now()
);
INSERT INTO places(name, description, campus_id, category_id, floor, room, period, permission)
VALUES ('Плазма', 'Большая переговорная с телевизором и интерактивной доской', 1, 1, 1, 100, '1 hours'::interval,
        'student'),
       ('Кухня X', 'Кухня для АДМ', 1, 2, 3, 313, 'hour 1', 'student'),
       ('Игровая Люси', 'Игровая с пс4', 1, 3, 2, 213, 'hour 1', 'student');

CREATE TABLE inventory
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR   NOT NULL,
    description VARCHAR   NOT NULL,
    campus_id   INTEGER REFERENCES campus (id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES category (id) ON DELETE CASCADE,
    period      INTERVAL,
    permission  ROLE      NOT NULL DEFAULT 'student',
    create_at   TIMESTAMP NOT NULL DEFAULT now(),
    update_at   TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE bookings
(
    id           SERIAL PRIMARY KEY,
    user_id      BIGINT REFERENCES users (id) ON DELETE CASCADE,
    type         BOOKING_TYPE NOT NULL,
    inventory_id INTEGER               DEFAULT 0 REFERENCES inventory (id) ON DELETE CASCADE,
    places_id    INTEGER               DEFAULT 0 REFERENCES places (id) ON DELETE CASCADE,
    confirm      BOOLEAN      NOT NULL DEFAULT FALSE,
    status       BOOLEAN      NOT NULL DEFAULT FALSE,--staff
    start_at     TIMESTAMP    NOT NULL,
    end_at       TIMESTAMP    NOT NULL,
    create_at    TIMESTAMP    NOT NULL DEFAULT now(),
    update_at    TIMESTAMP    NOT NULL DEFAULT now()
);
