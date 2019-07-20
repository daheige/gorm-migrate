# migrate数据迁移
    go gorm数据迁移
    通过使用GROM提供的一系列方法，根据数据模型定义好的规则，进行创建、
    删除数据表等操作，也就是数据库的DDL操作

# 数据表操作

    //根据模型自动创建数据表
    func (s *DB) AutoMigrate(values ...interface{}) *DB 
     
    //根据模型创建数据表
    func (s *DB) CreateTable(models ...interface{}) *DB
     
    //删除数据表，相当于drop table语句
    func (s *DB) DropTable(values ...interface{}) *DB 
     
    //相当于drop table if exsist 语句
    func (s *DB) DropTableIfExists(values ...interface{}) *DB 
     
    //根据模型判断数据表是否存在
    func (s *DB) HasTable(value interface{}) bool
   
# 列操作
    //删除数据表字段
    func (s *DB) DropColumn(column string) *DB
     
    //修改数据表字段的数据类型
    func (s *DB) ModifyColumn(column string, typ string) *DB

# 索引操作
    //添加外键
    func (s *DB) AddForeignKey(field string, dest string, onDelete string, onUpdate string) *DB
     
    //给数据表字段添加索引
    func (s *DB) AddIndex(indexName string, columns ...string) *DB
     
    //给数据表字段添加唯一索引
    func (s *DB) AddUniqueIndex(indexName string, columns ...string) *DB
    
# 参考文档
    http://gorm.book.jasperxu.com/database.html#m
    
    http://gorm.book.jasperxu.com/
    
    https://gorm.io/zh_CN/docs/
    
    https://github.com/jinzhu/gorm
    
