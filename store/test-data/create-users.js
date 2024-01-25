/* change database 
use projects */

db.createCollection("users");
db.users.insertOne({
    __id: "3F190455-6D4B-446E-AAAE-DAEF761B318B",
    name: "Piras Thiyagarajan",
    phone: "+1650-555-1212",
    text: "+1-650-555-1212",
    email: "test@testemail.com",
    whatsapp: "+1650-555-1212",
});

db.users.createIndex({
    name: "text",
})