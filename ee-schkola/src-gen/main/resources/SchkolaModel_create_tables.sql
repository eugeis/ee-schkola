db1596_schkola
use ;
#User
CREATE TABLE IF NOT EXISTS user (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    gender VARCHAR(255) COMMENT 'gender(Gender)',
    #name(PersonName)
    name_first VARCHAR(255) COMMENT 'name.first',
    name_last VARCHAR(255) COMMENT 'name.last',
    birth_name VARCHAR(255) COMMENT 'birthName',
    birthday DATETIME COMMENT 'birthday',
    #address(Address)
    address_street VARCHAR(255) COMMENT 'address.street',
    address_suite VARCHAR(255) COMMENT 'address.suite',
    address_city VARCHAR(255) COMMENT 'address.city',
    address_code VARCHAR(255) COMMENT 'address.code',
    address_country VARCHAR(255) COMMENT 'address.country',
    #contact(Contact)
    contact_phone VARCHAR(255) COMMENT 'contact.phone',
    contact_email VARCHAR(255) COMMENT 'contact.email',
    contact_cellphone VARCHAR(255) COMMENT 'contact.cellphone',
    photo VARCHAR(255) COMMENT 'photo',
    #family(Family)
    family_marital_status VARCHAR(255) COMMENT 'family.maritalStatus(MaritalStatus)',
    family_children_count INTEGER COMMENT 'family.childrenCount',
    #family.partner(PersonName)
    partner_first VARCHAR(255) COMMENT 'family.partner.first',
    partner_last VARCHAR(255) COMMENT 'family.partner.last',
    #church(ChurchInfo)
    church_name VARCHAR(255) COMMENT 'church.name',
    church_member BOOLEAN COMMENT 'church.member',
    church_church VARCHAR(255) COMMENT 'church.church',
    #education(Education)
    #education.graduation(Graduation)
    graduation_id VARCHAR(255) COMMENT 'education.graduation.id',
    education_profession VARCHAR(255) COMMENT 'education.profession',
    PRIMARY KEY (id));

#User.gender(Gender)
ALTER TABLE user
    ADD INDEX (gender);


#Church
CREATE TABLE IF NOT EXISTS church (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    name VARCHAR(255) COMMENT 'name',
    #address(Address)
    address_street VARCHAR(255) COMMENT 'address.street',
    address_suite VARCHAR(255) COMMENT 'address.suite',
    address_city VARCHAR(255) COMMENT 'address.city',
    address_code VARCHAR(255) COMMENT 'address.code',
    address_country VARCHAR(255) COMMENT 'address.country',
    #pastor(PersonName)
    pastor_first VARCHAR(255) COMMENT 'pastor.first',
    pastor_last VARCHAR(255) COMMENT 'pastor.last',
    #contact(Contact)
    contact_phone VARCHAR(255) COMMENT 'contact.phone',
    contact_email VARCHAR(255) COMMENT 'contact.email',
    contact_cellphone VARCHAR(255) COMMENT 'contact.cellphone',
    PRIMARY KEY (id));


#Graduation
CREATE TABLE IF NOT EXISTS graduation (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    name VARCHAR(255) COMMENT 'name',
    type VARCHAR(255) COMMENT 'type(GraduationType)',
    PRIMARY KEY (id));

#Graduation.type(GraduationType)
ALTER TABLE graduation
    ADD INDEX (type);


#Expense
CREATE TABLE IF NOT EXISTS expense (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    #purpose(ExpenseType)
    expense_type_id VARCHAR(255) COMMENT 'purpose.id',
    amount FLOAT(7,4) COMMENT 'amount',
    #person(User)
    user_id VARCHAR(255) COMMENT 'person.id',
    date DATETIME COMMENT 'date',
    PRIMARY KEY (id));


#ExpenseType
CREATE TABLE IF NOT EXISTS expense_type (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    name VARCHAR(255) COMMENT 'name',
    description VARCHAR(255) COMMENT 'description',
    PRIMARY KEY (id));


#Fee
CREATE TABLE IF NOT EXISTS fee (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    #student(User)
    user_id VARCHAR(255) COMMENT 'student.id',
    amount FLOAT(7,4) COMMENT 'amount',
    #type(FeeType)
    fee_type_id VARCHAR(255) COMMENT 'type.id',
    date DATETIME COMMENT 'date',
    PRIMARY KEY (id));


#FeeType
CREATE TABLE IF NOT EXISTS fee_type (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    name VARCHAR(255) COMMENT 'name',
    amount FLOAT(7,4) COMMENT 'amount',
    description VARCHAR(255) COMMENT 'description',
    PRIMARY KEY (id));


#SchoolApplication
CREATE TABLE IF NOT EXISTS school_application (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    #person(User)
    user_id VARCHAR(255) COMMENT 'person.id',
    has_recommendation BOOLEAN COMMENT 'hasRecommendation',
    #recommendationOf(PersonName)
    recommendation_of_first VARCHAR(255) COMMENT 'recommendationOf.first',
    recommendation_of_last VARCHAR(255) COMMENT 'recommendationOf.last',
    #mentor(PersonName)
    mentor_first VARCHAR(255) COMMENT 'mentor.first',
    mentor_last VARCHAR(255) COMMENT 'mentor.last',
    #mentorContact(Contact)
    mentor_contact_phone VARCHAR(255) COMMENT 'mentorContact.phone',
    mentor_contact_email VARCHAR(255) COMMENT 'mentorContact.email',
    mentor_contact_cellphone VARCHAR(255) COMMENT 'mentorContact.cellphone',
    #schoolYear(SchoolYear)
    school_year_id VARCHAR(255) COMMENT 'schoolYear.id',
    group_ VARCHAR(255) COMMENT 'group',
    PRIMARY KEY (id));


#SchoolYear
CREATE TABLE IF NOT EXISTS school_year (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    name VARCHAR(255) COMMENT 'name',
    start DATETIME COMMENT 'start',
    end DATETIME COMMENT 'end',
    PRIMARY KEY (id));


#SchoolYear.dates(Date)
CREATE TABLE IF NOT EXISTS school_year_dates (
    school_year_id VARCHAR(255) COMMENT 'SchoolYear.id',
    dates DATETIME COMMENT 'dates',
    PRIMARY KEY (school_year_id, dates));

#Course
CREATE TABLE IF NOT EXISTS course (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    name VARCHAR(255) COMMENT 'name',
    begin DATETIME COMMENT 'begin',
    end DATETIME COMMENT 'end',
    #teacher(PersonName)
    teacher_first VARCHAR(255) COMMENT 'teacher.first',
    teacher_last VARCHAR(255) COMMENT 'teacher.last',
    #schoolYear(SchoolYear)
    school_year_id VARCHAR(255) COMMENT 'schoolYear.id',
    fee FLOAT(7,4) COMMENT 'fee',
    description LONGTEXT COMMENT 'description',
    PRIMARY KEY (id));


#Group
CREATE TABLE IF NOT EXISTS group_ (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    name VARCHAR(255) COMMENT 'name',
    type VARCHAR(255) COMMENT 'type(GroupType)',
    #schoolYear(SchoolYear)
    school_year_id VARCHAR(255) COMMENT 'schoolYear.id',
    #representative(User)
    user_id VARCHAR(255) COMMENT 'representative.id',
    PRIMARY KEY (id));

#Group.type(GroupType)
ALTER TABLE group_
    ADD INDEX (type);


#Group.students(User)
CREATE TABLE IF NOT EXISTS group__user (
    group__id VARCHAR(255) COMMENT 'Group.id',
    #students(User)
    user_id VARCHAR(255) COMMENT 'students.id',
    PRIMARY KEY (group__id, user_id));


#Group.courses(Course)
CREATE TABLE IF NOT EXISTS group__course (
    group__id VARCHAR(255) COMMENT 'Group.id',
    #courses(Course)
    course_id VARCHAR(255) COMMENT 'courses.id',
    PRIMARY KEY (group__id, course_id));

#Grade
CREATE TABLE IF NOT EXISTS grade (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    #student(User)
    user_id VARCHAR(255) COMMENT 'student.id',
    #course(Course)
    course_id VARCHAR(255) COMMENT 'course.id',
    grade FLOAT(7,4) COMMENT 'grade',
    #gradeTrace(Trace)
    grade_trace_created_at DATETIME COMMENT 'gradeTrace.createdAt',
    grade_trace_updated_at DATETIME COMMENT 'gradeTrace.updatedAt',
    grade_trace_modified_by VARCHAR(255) COMMENT 'gradeTrace.modifiedBy',
    comment VARCHAR(255) COMMENT 'comment',
    PRIMARY KEY (id));


#Attendance
CREATE TABLE IF NOT EXISTS attendance (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    #student(User)
    user_id VARCHAR(255) COMMENT 'student.id',
    date DATETIME COMMENT 'date',
    #course(Course)
    course_id VARCHAR(255) COMMENT 'course.id',
    hours INTEGER COMMENT 'hours',
    status VARCHAR(255) COMMENT 'status(AttendanceStatus)',
    #statusTrace(Trace)
    status_trace_created_at DATETIME COMMENT 'statusTrace.createdAt',
    status_trace_updated_at DATETIME COMMENT 'statusTrace.updatedAt',
    status_trace_modified_by VARCHAR(255) COMMENT 'statusTrace.modifiedBy',
    token VARCHAR(255) COMMENT 'token',
    #tokenTrace(Trace)
    token_trace_created_at DATETIME COMMENT 'tokenTrace.createdAt',
    token_trace_updated_at DATETIME COMMENT 'tokenTrace.updatedAt',
    token_trace_modified_by VARCHAR(255) COMMENT 'tokenTrace.modifiedBy',
    PRIMARY KEY (id));

#Attendance.status(AttendanceStatus)
ALTER TABLE attendance
    ADD INDEX (status);


#Book
CREATE TABLE IF NOT EXISTS book (
    id VARCHAR(255) COMMENT 'id',
    #trace(Trace)
    trace_created_at DATETIME COMMENT 'trace.createdAt',
    trace_updated_at DATETIME COMMENT 'trace.updatedAt',
    trace_modified_by VARCHAR(255) COMMENT 'trace.modifiedBy',
    title VARCHAR(255) COMMENT 'title',
    description LONGTEXT COMMENT 'description',
    language VARCHAR(255) COMMENT 'language',
    release_date DATETIME COMMENT 'releaseDate',
    edition VARCHAR(255) COMMENT 'edition',
    category VARCHAR(255) COMMENT 'category',
    #author(PersonName)
    author_first VARCHAR(255) COMMENT 'author.first',
    author_last VARCHAR(255) COMMENT 'author.last',
    #location(Location)
    location_shelf VARCHAR(255) COMMENT 'location.shelf',
    location_fold VARCHAR(255) COMMENT 'location.fold',
    PRIMARY KEY (id));






#Expense.purpose(ExpenseType)
ALTER TABLE expense
    ADD FOREIGN KEY (expense_type_id)
    REFERENCES expense_type (id);


#Expense.person(User)
ALTER TABLE expense
    ADD FOREIGN KEY (user_id)
    REFERENCES user (id);



#Fee.student(User)
ALTER TABLE fee
    ADD FOREIGN KEY (user_id)
    REFERENCES user (id);


#Fee.type(FeeType)
ALTER TABLE fee
    ADD FOREIGN KEY (fee_type_id)
    REFERENCES fee_type (id);



#SchoolApplication.person(User)
ALTER TABLE school_application
    ADD FOREIGN KEY (user_id)
    REFERENCES user (id);


#SchoolApplication.schoolYear(SchoolYear)
ALTER TABLE school_application
    ADD FOREIGN KEY (school_year_id)
    REFERENCES school_year (id);



#Course.schoolYear(SchoolYear)
ALTER TABLE course
    ADD FOREIGN KEY (school_year_id)
    REFERENCES school_year (id);


#Group.schoolYear(SchoolYear)
ALTER TABLE group_
    ADD FOREIGN KEY (school_year_id)
    REFERENCES school_year (id);


#Group.representative(User)
ALTER TABLE group_
    ADD FOREIGN KEY (user_id)
    REFERENCES user (id);


#Grade.student(User)
ALTER TABLE grade
    ADD FOREIGN KEY (user_id)
    REFERENCES user (id);


#Grade.course(Course)
ALTER TABLE grade
    ADD FOREIGN KEY (course_id)
    REFERENCES course (id);


#Attendance.student(User)
ALTER TABLE attendance
    ADD FOREIGN KEY (user_id)
    REFERENCES user (id);


#Attendance.course(Course)
ALTER TABLE attendance
    ADD FOREIGN KEY (course_id)
    REFERENCES course (id);



