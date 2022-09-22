import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

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





