db = db.getSiblingDB('users'); // seu database (MONGODB_USER_DB)

db.users.insertOne({
  email: "admin@crud.com",
  password: "46c17fb99b9c8c8ae8214834e1edf6b1", // senha: admin123#
  name: "Admin User",
  age: 30
});

