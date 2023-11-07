CREATE TABLE profiles (
    name varchar,
    age int,
    salary int
);

INSERT INTO profiles VALUES
('Arash', 28, 12000),
('Masoud', 22, 9000),
('Elahe', 31, 24000),
('Nima', 38, 20000),
('Behnam', 51, 27000),
('Ahmad', 40, 26000),
('Roja', 36, 19000);


SELECT (profiles.age / 10) * 10 as age_range,  min(profiles.salary) as min, max(profiles.salary) as max,  sum(profiles.salary) / count(1) as avg FROM profiles
GROUP BY profiles.age / 10
order by age_range;
