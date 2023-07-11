db = new Mongo().getDB("patients");

db.createCollection('patients', { capped: false });

db.patients.insert([
  { "name":"小明" },
  { "name":"阿花" },
  { "name":"隨意" },
  { "name":"瑞榮" },
  { "name":"錦華" }
]);