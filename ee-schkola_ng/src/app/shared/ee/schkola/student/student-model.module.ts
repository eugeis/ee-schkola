import { NgModule } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import {StudentRoutingModules} from './student-routing.module';
import {CommonModule} from '@angular/common';
import {TemplateModule} from '../../template/template.module';
import {MaterialModule} from '../../template/material.module';

import {StudentViewComponent} from './components/view/student-module-view.component';
import {AttendanceViewComponent} from './attendance/components/view/attendance-entity-view.component';
import {AttendanceListComponent} from './attendance/components/list/attendance-entity-list.component';
import {CourseViewComponent} from './course/components/view/course-entity-view.component';
import {CourseListComponent} from './course/components/list/course-entity-list.component';
import {GradeViewComponent} from './grade/components/view/grade-entity-view.component';
import {GradeListComponent} from './grade/components/list/grade-entity-list.component';
import {GroupViewComponent} from './group/components/view/group-entity-view.component';
import {GroupListComponent} from './group/components/list/group-entity-list.component';
import {SchoolApplicationViewComponent} from './schoolapplication/components/view/schoolapplication-entity-view.component';
import {SchoolApplicationListComponent} from './schoolapplication/components/list/schoolapplication-entity-list.component';
import {SchoolYearViewComponent} from './schoolyear/components/view/schoolyear-entity-view.component';
import {SchoolYearListComponent} from './schoolyear/components/list/schoolyear-entity-list.component';
import {PersonModule} from '@schkola/person/person-model.module';



@NgModule({
    declarations: [
        StudentViewComponent,
        AttendanceViewComponent,
        AttendanceListComponent,
        CourseViewComponent,
        CourseListComponent,
        GradeViewComponent,
        GradeListComponent,
        GroupViewComponent,
        GroupListComponent,
        SchoolApplicationViewComponent,
        SchoolApplicationListComponent,
        SchoolYearViewComponent,
        SchoolYearListComponent,

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
        AttendanceViewComponent,
        CourseViewComponent,
        GradeViewComponent,
        GroupViewComponent,
        SchoolApplicationViewComponent,
        SchoolYearViewComponent,

    ]
})
export class StudentModule { }




