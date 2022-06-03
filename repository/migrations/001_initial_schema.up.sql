CREATE TABLE stories (
    id char(27) PRIMARY KEY,
    user_id TEXT NOT NULL,
    party_id TEXT NOT NULL,
    url TEXT NOT NULL,
    tagged_friends TEXT[]
);

CREATE UNIQUE INDEX user_id_idx ON stories (user_id);
CREATE UNIQUE INDEX party_id_idx ON stories (party_id);
