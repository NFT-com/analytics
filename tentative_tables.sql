
-- Creation of mints table.
CREATE TABLE mints (
    id VARCHAR(64) PRIMARY KEY,
    chain_id UUID NOT NULL,
    collection_id UUID NOT NULL,
    block BIGINT NOT NULL,
    transaction_hash VARCHAR(64) NOT NULL,
    minter TEXT NOT NULL,
    emitted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Creation of transfers table.
CREATE TABLE transfers (
    id VARCHAR(64) PRIMARY KEY,
    chain_id UUID NOT NULL,
    collection_id UUID NOT NULL,
    block BIGINT NOT NULL,
    transaction_hash VARCHAR(64) NOT NULL,
    from_address TEXT NOT NULL,
    to_address TEXT NOT NULL,
    price BIGINT NOT NULL, -- FIXME: Check - what format will we be using for price?
    emitted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Creation of sales table.
-- FIXME: What is the difference for sales vs transfers?
-- Are sales events when a token is listed for a sale/auction?
CREATE TABLE sales (
    id VARCHAR(64) PRIMARY KEY,
    chain_id UUID NOT NULL,
    collection_id UUID NOT NULL,
    block BIGINT NOT NULL,
    transaction_hash VARCHAR(64) NOT NULL,
    marketplace TEXT NOT NULL,
    owner TEXT NOT NULL,
    price BIGINT NOT NULL, -- FIXME: Price format
    emitted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Creation of burns table.
CREATE TABLE burns (
    id VARCHAR(64) PRIMARY KEY,
    chain_id UUID NOT NULL,
    collection_id UUID NOT NULL,
    block BIGINT NOT NULL,
    transaction_hash VARCHAR(64) NOT NULL,
    burner TEXT NOT NULL,
    emitted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);
