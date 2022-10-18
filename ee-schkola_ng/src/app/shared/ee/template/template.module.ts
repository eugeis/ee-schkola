import { NgModule } from '@angular/core';
import {CommonModule} from '@angular/common';

import {FormComponent} from './components/form/form.component';
import {ButtonComponent} from './components/button/button.component';
import {PageComponent} from './components/page/page.component';
import {TableComponent} from './components/table/table.component';

import {TemplateRoutingModules} from './template-routing.modules';
import {ProfileDataService} from '@schkola/person/profile/service/profile-data.service';
import {ChurchDataService} from '@schkola/person/church/service/church-data.service'
import {FormService} from './services/form.service';
import {TableDataService} from './services/data.service';
import {ButtonService} from './services/button.service';
import {MaterialModule} from './material.module';
import {GraduationDataService} from '@schkola/person/graduation/service/graduation-data.service';
import {ExpenseDataService} from '@schkola/finance/expense/service/expense-data.service';
import {ExpensePurposeDataService} from '@schkola/finance/expensepurpose/service/expensepurpose-data.service';
import {FeeDataService} from '@schkola/finance/fee/service/fee-data.service';
import {FeeKindDataService} from '@schkola/finance/feekind/service/feekind-data.service';
import {SchoolApplicationDataService} from '@schkola/student/schoolapplication/service/schoolapplication-data.service';
import {SchoolYearDataService} from '@schkola/student/schoolyear/service/schoolyear-data.service';
import {CourseDataService} from '@schkola/student/course/service/course-data.service';
import {GroupDataService} from '@schkola/student/group/service/group-data.service';
import {Grade} from '@schkola/student/StudentApiBase';
import {GradeDataService} from '@schkola/student/grade/service/grade-data.service';
import {AttendanceDataService} from '@schkola/student/attendance/service/attendance-data.service';
import {BookDataService} from '@schkola/library/book/service/book-data.service';

@NgModule({
    declarations: [FormComponent,
        ButtonComponent,
        PageComponent,
        TableComponent,
    ],
    imports: [
        TemplateRoutingModules,
        CommonModule,
        MaterialModule,
    ],
    providers: [
        ProfileDataService, ChurchDataService, GraduationDataService,
        ExpenseDataService, ExpensePurposeDataService, FeeDataService, FeeKindDataService,
        SchoolApplicationDataService, SchoolYearDataService, CourseDataService, GroupDataService, GradeDataService, AttendanceDataService,
        BookDataService,
        TableDataService, ButtonService, FormService
    ],
    exports: [
        FormComponent,
        ButtonComponent,
        PageComponent,
        TableComponent
    ]
})
export class TemplateModule { }
