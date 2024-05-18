CREATE TABLE Types
(
    id SERIAL,
    name text,
    PRIMARY KEY(id)
);

CREATE TABLE Pokemon
(
    id SERIAL,
    name TEXT,
    pokedexNumber int,
    type1Id int,
    type2Id int,
    PRIMARY KEY(id),
    CONSTRAINT fk_type1
        FOREIGN KEY(type1Id)
            REFERENCES Types(id),
    CONSTRAINT fk_type2
        FOREIGN KEY(type2Id)
            REFERENCES Types(id)
); 

INSERT INTO Types(name)
VALUES ('Normal'),
('Fire'),
('Water'),
('Electric'),
('Grass'),
('Ice'),
('Fighting'),
('Poison'),
('Ground'),
('Flying'),
('Psychic'),
('Bug'),
('Rock'),
('Ghost'),
('Dragon'),
('Dark'),
('Steel'),
('Fairy');