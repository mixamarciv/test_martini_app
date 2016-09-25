CREATE TABLE tuser (
    uuid          VARCHAR(36),
    type          INTEGER,      -- 0 - admin, 10 - spec, 100 - user  
    fam           VARCHAR(200),
    name          VARCHAR(200),
    otch 	  VARCHAR(200),
    mail          VARCHAR(200),
    phone         VARCHAR(20),
    pass          VARCHAR(100),
    street        VARCHAR(200),
    house         VARCHAR(20),
    flat          VARCHAR(20),
    info          VARCHAR(2000)
);

CREATE UNIQUE INDEX tuser_IDX1 ON tuser (uuid);
CREATE        INDEX tuser_IDX2 ON tuser (mail);


CREATE TABLE tpost (
    uuid_user     VARCHAR(36),
    uuid          VARCHAR(36),
    type          VARCHAR(200),
    name          VARCHAR(200),
    text 	  VARCHAR(2000),
    data          BLOB,
    postdate      VARCHAR(40)
);

CREATE UNIQUE INDEX tpost_IDX1 ON tpost (uuid);
CREATE        INDEX tpost_IDX2 ON tpost (uuid_user);


COMMIT WORK;
