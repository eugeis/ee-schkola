import {Course} from '@schkola/student/StudentApiBase'
import {PersonName} from '@schkola/person/PersonApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {CourseDataService} from '../../service/course-data.service';

@Component({
  selector: 'app-course-view',
  templateUrl: './course-entity-view.component.html',
  styleUrls: ['./course-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: CourseDataService}]
})

export class CourseViewComponent implements OnInit {


    course: Course;

    constructor(@Inject(CourseDataService) public courseDataService: CourseDataService) {}

    ngOnInit(): void {
        this.course = this.courseDataService.getFirst();
        this.course.teacher = new PersonName();
        this.courseDataService.checkRoute(this.course);
    }
}




