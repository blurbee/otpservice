
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


db.createUser({user: "myUserAdmin", pwd: passwordPrompt(), roles: ["dbAdmin"], mechanisms: ["SCRAM-SHA-1"]})