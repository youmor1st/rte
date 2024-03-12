### Школьная система merit/demerit

Эта такая система созданная для школ, чтоб в школе была дисциплина, внутренная правила. Если школа будет давать призы за Топ 5 учеников и за Топ 3 класса, призы это в виде мотиваций для учеников. Если будет приз за топ класс это будет одной причиной обьеденений класса. 

### **Структура API**

1. **Authentication (Аутентификация)**
    
    `/admin/student-register` - Регистрация нового пользователя (доступен только админу)
    `/login` - Вход в систему (все)

1. **Управление пользователями**
    
    `/admin/users` - Получить список всех пользователей (завуч и админ)
    
    `/users/{id}/points` - Получить баллы конкретного пользователя (завуч, админ и учитель) (не доступен пока)
    
    `/users/top` - Получить список топ 5 учеников (все)  (не доступен пока)
    
    `/admin/delete-student` - Удалить пользователя (доступно только админу)
    
    `/admin/update-student` - Обновить данные пользователя (доступно только админу)
    
2. **Управление баллами**
    
    `/points` - Добавление или вычитание баллов (CRUD) (учитель и завуч) (не доступен пока)
    
3. **Управление правилами**
    
    `/rules` - Получить список правил (все) (не доступен пока)
    
    `/admin/rules/add ` - Добавление нового правила (доступно только админу) (не доступен пока)
    
    `/admin/rules/delete` - Удаление правила (доступно только админу) (не доступен пока)
    
4. **Главное меню**
       `/home` - Видят все свои функции, которые могут выполнять (не доступен пока)
   
DATA BASE

`CREATE TABLE IF NOT EXISTS classes (`                                       
id SERIAL PRIMARY KEY,                                       
class_name VARCHAR(255) NOT NULL);

`CREATE TABLE IF NOT EXISTS users (`                                     
id SERIAL PRIMARY KEY,                                     
role VARCHAR(50) NOT NULL,                                     
username VARCHAR(255) UNIQUE NOT NULL,                                     
password VARCHAR(255) NOT NULL,                                     
f_name VARCHAR(255) NOT NULL,                                     
s_name VARCHAR(255) NOT NULL,                                     
class_id INT REFERENCES classes(id),                                     
points INT DEFAULT 100 CHECK (points >= 0));

`CREATE TABLE IF NOT EXISTS points_rules (`                                           
id SERIAL PRIMARY KEY,                                            
rule_name VARCHAR(255) NOT NULL,
rule_description VARCHAR(255),                                            
rule_point INT NOT NULL,                                            
rule_type VARCHAR(10) CHECK (rule_type IN ('positive', 'negative')));

`CREATE TABLE IF NOT EXISTS points (`                                       
id SERIAL PRIMARY KEY,                                      
user_id INT REFERENCES users(id),                                      
awarded_by INT REFERENCES users(id),                                      
rule_id INT REFERENCES points_rules(id),                                      
points INT NOT NULL,                                      
reason VARCHAR(255),                                      
timestamp TIMESTAMP DEFAULT *CURRENT_TIMESTAMP*);
