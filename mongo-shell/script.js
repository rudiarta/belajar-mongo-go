use belajar;
db.students.find();
db.students.createIndex({"name":"text","address":"text"});
db.students.getIndexes();
db.students.dropIndex("name_text");

db.students.find({$text:{$search:"lagi"}});

db.customer.find({"customer_name":{$regex:/az/i}, "customer_account_type":"partner", "customer_ref_id":{$in:[9578]}}).explain("executionStats");
// db.customer.createIndex({"customer_name":1});
// db.customer.createIndex({"customer_account_type":1});
// db.customer.createIndex({"customer_ref_id":1});
db.customer.dropIndex("idx_search");