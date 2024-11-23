CREATE TABLE workout (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    type TEXT NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    duration INTEGER NOT NULL,
    calories_burned INTEGER NOT NULL,
    workload INTEGER NOT NULL,
    description TEXT
);

CREATE TABLE meal (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    type TEXT NOT NULL,
    calories INTEGER NOT NULL,
    protein INTEGER NOT NULL,
    carbs INTEGER NOT NULL,
    fat INTEGER NOT NULL,
    description TEXT,
    date TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL
);
