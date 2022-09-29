import {Course} from '@schkola/student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {CourseDataService} from '../../service/course-data.service';

@Component({
  selector: 'app-course-list',
  templateUrl: './course-entity-list.component.html',
  styleUrls: ['./course-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: CourseDataService}]
})

export class CourseListComponent implements OnInit {

    course: Course = new Course();

    tableHeader: Array<String> = [];

    constructor(public courseDataService: CourseDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'begin', 'end', 'first', 'last', 'name', 'start', 'end', 'dates', 'id', 'fee', 'description', 'id'];
    }
}




