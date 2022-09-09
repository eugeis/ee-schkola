import {Course} from '../student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {CourseDataService} from '../../services/course-data.service';

@Component({
  selector: 'app-course-view',
  templateUrl: './course-view.component.html',
  styleUrls: ['./course-view.component.scss'],
  providers: [{provide: TableDataService, useClass: CourseDataService}]
})

export class CourseViewComponent implements OnInit {


    course: Course;

    constructor(public courseDataService: CourseDataService) {}

    ngOnInit(): void {
        this.course = this.courseDataService.getFirst();
        this.courseDataService.checkRoute(this.course);
    }
}




