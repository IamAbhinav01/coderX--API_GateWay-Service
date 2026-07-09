-- +goose Up
INSERT INTO users(name,email,password) VALUES("Abhinav","abhinavsunil@hotmail.com","iatwurn2568");
INSERT INTO users(name,email,password) VALUES("mallika","mallika75@gmail.com","godsila857");
INSERT INTO users(name,email,password) VALUES("suber","suberF857@gmail.com","fathbe8591");

-- +goose Down

DELETE FROM users
WHERE email IN (
    "abhinavsunil@hotmail.com",
    "mallika75@gmail.com",
    "suberF857@gmail.com"
);