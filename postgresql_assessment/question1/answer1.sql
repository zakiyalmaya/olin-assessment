SELECT u.name AS user_name, 
    SUM(o.amount) AS total_spent
FROM users u
JOIN orders o ON u.id = o.user_id
WHERE o.created_at >= '2022-01-01'
GROUP BY u.name
HAVING SUM(o.amount) >= 1000;

-- The answer is no user has spent a minimum of 1000 on transactions
-- made on or after January 1, 2022