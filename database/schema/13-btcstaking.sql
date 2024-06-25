CREATE TABLE btc_staking_mint
(
    id SERIAL PRIMARY KEY NOT NULL,
    tx_hash TEXT NOT NULL,
    height BIGINT NOT NULL REFERENCES block (height),
    signer TEXT NOT NULL,

    receiver_name TEXT NOT NULL,
    receiver_btc_address TEXT NOT NULL,
    receiver_eth_address TEXT NOT NULL,
    amount BIGINT NOT NULL,
    btc_staking_tx_hash TEXT NOT NULL
);

CREATE TABLE btc_staking_burn
(
    id SERIAL PRIMARY KEY NOT NULL,
    tx_hash TEXT NOT NULL,
    height BIGINT NOT NULL REFERENCES block (height),
    signer TEXT NOT NULL,

    target_btc_address TEXT NOT NULL,
    amount BIGINT NOT NULL
);