CREATE TABLE IF NOT EXISTS transfers
(
    id            SERIAL PRIMARY KEY,
    timestamp     BIGINT         NOT NULL,
    from_address  TEXT           NOT NULL,
    from_owner    TEXT,
    to_address    TEXT           NOT NULL,
    to_owner      TEXT,
    amount        NUMERIC(78, 0) NOT NULL, -- large enough for big.Int
    token_address TEXT           NOT NULL,
    symbol        TEXT           NOT NULL,
    chain         TEXT           NOT NULL,
    network       TEXT           NOT NULL,
    tx_hash       TEXT           NOT NULL,
    decimals      SMALLINT       NOT NULL, -- uint8 fits in SMALLINT
    position      BIGINT         NOT NULL
);
