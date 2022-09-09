import {Course} from '../student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {CourseDataService} from '../../services/course-data.service';

@Component({
  selector: 'app-course-list',
  templateUrl: './course-view.component.html',
  styleUrls: ['./course-view.component.scss'],
  providers: [{provide: TableDataService, useClass: CourseDataService}]
})

export class CourseViewComponent implements OnInit {

    course: Course = new Course();

    tableHeader: Array<String> = [];

    constructor(public courseDataService: CourseDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'begin', 'end', 'teacher', 'schoolyear', 'fee', 'description', 'id'];
    }
}




