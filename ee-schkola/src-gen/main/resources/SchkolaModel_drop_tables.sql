db1596_schkola
use ;
#User
DROP TABLE IF EXISTS user;


#Church
DROP TABLE IF EXISTS church;


#Graduation
DROP TABLE IF EXISTS graduation;


#Expense
DROP TABLE IF EXISTS expense;


#ExpenseType
DROP TABLE IF EXISTS expense_type;


#Fee
DROP TABLE IF EXISTS fee;


#FeeType
DROP TABLE IF EXISTS fee_type;


#SchoolApplication
DROP TABLE IF EXISTS school_application;


#SchoolYear
DROP TABLE IF EXISTS school_year;

#SchoolYear.dates(Date)
DROP TABLE IF EXISTS school_year_dates;


#Course
DROP TABLE IF EXISTS course;


#Group
DROP TABLE IF EXISTS group_;

#Group.students(User)
DROP TABLE IF EXISTS group__user;


#Group.courses(Course)
DROP TABLE IF EXISTS group__course;


#Grade
DROP TABLE IF EXISTS grade;


#Attendance
DROP TABLE IF EXISTS attendance;


#Book
DROP TABLE IF EXISTS book;



