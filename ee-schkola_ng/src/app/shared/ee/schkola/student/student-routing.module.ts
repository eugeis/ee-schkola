import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

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

const routes: Routes = [
    { path: '', component: StudentViewComponent },
    { path: 'attendance', component: AttendanceListComponent },
    { path: 'attendance/new', component: AttendanceViewComponent },
    { path: 'attendance/edit/:id', component: AttendanceViewComponent },
    { path: 'course', component: CourseListComponent },
    { path: 'course/new', component: CourseViewComponent },
    { path: 'course/edit/:id', component: CourseViewComponent },
    { path: 'grade', component: GradeListComponent },
    { path: 'grade/new', component: GradeViewComponent },
    { path: 'grade/edit/:id', component: GradeViewComponent },
    { path: 'group', component: GroupListComponent },
    { path: 'group/new', component: GroupViewComponent },
    { path: 'group/edit/:id', component: GroupViewComponent },
    { path: 'schoolapplication', component: SchoolApplicationListComponent },
    { path: 'schoolapplication/new', component: SchoolApplicationViewComponent },
    { path: 'schoolapplication/edit/:id', component: SchoolApplicationViewComponent },
    { path: 'schoolyear', component: SchoolYearListComponent },
    { path: 'schoolyear/new', component: SchoolYearViewComponent },
    { path: 'schoolyear/edit/:id', component: SchoolYearViewComponent }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class StudentRoutingModules {}





