import { NgModule } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import {StudentRoutingModules} from './student-routing.module';
import {CommonModule} from '@angular/common';
import {TemplateModule} from '../../template/template.module';
import {MaterialModule} from '../../template/material.module';

import {StudentViewComponent} from './components/view/student-module-view.component';
import {AttendanceViewComponent} from '@schkola/student/attendance/components/view/attendance-entity-view.component';
import {AttendanceListComponent} from '@schkola/student/attendance/components/list/attendance-entity-list.component';
import {AttendanceFormComponent} from '@schkola/student/attendance/components/form/attendance-form.component';
import {CourseViewComponent} from '@schkola/student/course/components/view/course-entity-view.component';
import {CourseListComponent} from '@schkola/student/course/components/list/course-entity-list.component';
import {CourseFormComponent} from '@schkola/student/course/components/form/course-form.component';
import {GradeViewComponent} from '@schkola/student/grade/components/view/grade-entity-view.component';
import {GradeListComponent} from '@schkola/student/grade/components/list/grade-entity-list.component';
import {GradeFormComponent} from '@schkola/student/grade/components/form/grade-form.component';
import {GroupViewComponent} from '@schkola/student/group/components/view/group-entity-view.component';
import {GroupListComponent} from '@schkola/student/group/components/list/group-entity-list.component';
import {GroupFormComponent} from '@schkola/student/group/components/form/group-form.component';
import {SchoolApplicationViewComponent} from '@schkola/student/schoolapplication/components/view/schoolapplication-entity-view.component';
import {SchoolApplicationListComponent} from '@schkola/student/schoolapplication/components/list/schoolapplication-entity-list.component';
import {SchoolApplicationFormComponent} from '@schkola/student/schoolapplication/components/form/schoolapplication-form.component';
import {SchoolYearViewComponent} from '@schkola/student/schoolyear/components/view/schoolyear-entity-view.component';
import {SchoolYearListComponent} from '@schkola/student/schoolyear/components/list/schoolyear-entity-list.component';
import {SchoolYearFormComponent} from '@schkola/student/schoolyear/components/form/schoolyear-form.component';
import {PersonModule} from '@schkola/person/person-model.module';


@NgModule({
    declarations: [
        StudentViewComponent,
        AttendanceViewComponent,
        AttendanceListComponent,
        AttendanceFormComponent,
        CourseViewComponent,
        CourseListComponent,
        CourseFormComponent,
        GradeViewComponent,
        GradeListComponent,
        GradeFormComponent,
        GroupViewComponent,
        GroupListComponent,
        GroupFormComponent,
        SchoolApplicationViewComponent,
        SchoolApplicationListComponent,
        SchoolApplicationFormComponent,
        SchoolYearViewComponent,
        SchoolYearListComponent,
        SchoolYearFormComponent,

    ],
    imports: [
        StudentRoutingModules,
        TemplateModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        MaterialModule,
        PersonModule
    ],
    providers: [],
    exports: [
        AttendanceFormComponent,
        CourseFormComponent,
        GradeFormComponent,
        GroupFormComponent,
        SchoolApplicationFormComponent,
        SchoolYearFormComponent,

    ]
})
export class StudentModule {}



