CREATE TABLE stories (
    id varchar(27) PRIMARY KEY,
    user_id varchar(27) NOT NULL,
    party_id varchar(27) NOT NULL,
    url TEXT NOT NULL,
    tagged_friends TEXT[]
);

CREATE INDEX user_id_idx ON stories (user_id);
CREATE INDEX party_id_idx ON stories (party_id);
