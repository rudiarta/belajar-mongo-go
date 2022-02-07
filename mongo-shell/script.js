use belajar;
db.students.find();
db.students.createIndex({"name":"text","address":"text"});
db.students.getIndexes();
db.students.dropIndex("name_text");

db.students.find({$text:{$search:"lagi"}});