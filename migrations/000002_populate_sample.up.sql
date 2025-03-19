INSERT INTO devices (name, brand, state)
SELECT 
  'Device_' || g AS name,
  CASE WHEN random() < 0.5 THEN 'Brand_A' ELSE 'Brand_B' END AS brand,
  CASE floor(random() * 3)
    WHEN 0 THEN 'AVAILABLE'
    WHEN 1 THEN 'IN_USE'
    ELSE 'INACTIVE'
  END AS state
FROM generate_series(1, 1000000) g;
