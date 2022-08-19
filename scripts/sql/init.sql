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
    user_id    BIGINT REFERENCES users (id),
    code       int      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    end_at     TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE category
(
    id        SERIAL PRIMARY KEY,
    name      VARCHAR   NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT now(),
    update_at TIMESTAMP NOT NULL DEFAULT now()
);


CREATE TABLE places
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR   NOT NULL,
    description VARCHAR   NOT NULL,
    campus_id   INTEGER REFERENCES campus (id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES category (id) ON DELETE CASCADE,
    floor       INTEGER   NOT NULL,
    room        INTEGER   NOT NULL,
    create_at   TIMESTAMP NOT NULL DEFAULT now(),
    update_at   TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE inventory
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR   NOT NULL,
    description VARCHAR   NOT NULL,
    campus_id   INTEGER REFERENCES campus (id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES category (id) ON DELETE CASCADE,
    create_at   TIMESTAMP NOT NULL DEFAULT now(),
    update_at   TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE bookings
(
    id           SERIAL PRIMARY KEY,
    user_id      BIGINT REFERENCES users (id) ON DELETE CASCADE,
    type         BOOKING_TYPE NOT NULL,
    inventory_id INTEGER REFERENCES inventory (id) ON DELETE CASCADE,
    places_id    INTEGER REFERENCES places (id) ON DELETE CASCADE,
    confirm      BOOLEAN      NOT NULL DEFAULT FALSE,
    start_at     TIMESTAMP    NOT NULL,
    end_at       TIMESTAMP    NOT NULL,
    create_at    TIMESTAMP    NOT NULL DEFAULT now(),
    update_at    TIMESTAMP    NOT NULL DEFAULT now()
);
