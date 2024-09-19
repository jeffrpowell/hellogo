CREATE TABLE IF NOT EXISTS hellogo.colorgradient (
    Id SERIAL PRIMARY KEY,
    Name VARCHAR NOT NULL,
    Description VARCHAR NULL,
    Colors VARCHAR(7)[] CHECK (array_length(colors, 1) <= 5) NOT NULL,
    UNIQUE (Name)
);

CREATE INDEX IF NOT EXISTS idx_colorgradient_name ON hellogo.colorgradient (Name);

INSERT INTO hellogo.colorgradient (Name, Colors)
VALUES 
('Foothill sunrise', ARRAY['#C2DCF0', '#A49B7F', '#716754', '#E9B786', '#CF9B79']),
('True sky blue', ARRAY['#4467A1', '#557BBA', '#699BD4', '#88B8ED', '#A5CEEF'])
ON CONFLICT (Name) DO NOTHING;