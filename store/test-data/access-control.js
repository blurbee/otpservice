
// use admin
db.createUser(
  {
    user: "myUserAdmin",
    pwd: passwordPrompt(), // or cleartext password
    roles: [
      { role: "userAdminAnyDatabase", db: "admin" },
      { role: "readWriteAnyDatabase", db: "admin" }
    ]
  }
)


db.createUser({user: "appuser", pwd: passwordPrompt(), roles: [{ role: "readWrite", db: "test" }], mechanisms: ["SCRAM-SHA-1"]})
db.createUser({user: "rouser", pwd: passwordPrompt(), roles: [{ role: "read", db: "test" }], mechanisms: ["SCRAM-SHA-1"]})