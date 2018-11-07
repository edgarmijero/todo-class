UPDATE items SET task = (CASE WHEN '' = '' IS TRUE THEN task ELSE 'PUT endpoint' END), completed = false WHERE id = 'a7df8e93-2921-4626-a103-295e4f4c90fc';
