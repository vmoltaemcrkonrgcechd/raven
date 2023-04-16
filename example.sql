DROP DATABASE IF EXISTS example;

CREATE DATABASE example;

\connect example

CREATE TABLE "user"(
                       user_id UUID DEFAULT gen_random_uuid(),
                       username VARCHAR(40) NOT NULL,
                       PRIMARY KEY (user_id)
);

CREATE TABLE article(
                        article_id UUID DEFAULT gen_random_uuid(),
                        name VARCHAR(40) NOT NULL,
                        user_id UUID NOT NULL,
                        PRIMARY KEY (article_id),
                        FOREIGN KEY (user_id)
                            REFERENCES "user" (user_id)
                            ON DELETE CASCADE
);

CREATE TABLE "like"(
                       like_id UUID DEFAULT gen_random_uuid(),
                       user_id UUID NOT NULL,
                       PRIMARY KEY (like_id),
                       FOREIGN KEY (user_id)
                           REFERENCES "user" (user_id)
                           ON DELETE CASCADE
);

CREATE TABLE article_like(
                             like_id UUID NOT NULL,
                             article_id UUID NOT NULL,
                             FOREIGN KEY (like_id)
                                 REFERENCES "like" (like_id)
                                 ON DELETE CASCADE,
                             FOREIGN KEY (article_id)
                                 REFERENCES article (article_id)
                                 ON DELETE CASCADE
);
