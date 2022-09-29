import {Course, SchoolYear} from '@schkola/student/StudentApiBase'
import {PersonName} from '@schkola/person/PersonApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {CourseDataService} from '../../service/course-data.service';

@Component({
  selector: 'app-course-form',
  templateUrl: './course-form.component.html',
  styleUrls: ['./course-form.component.scss']
})

export class CourseFormComponent implements OnInit {


    @Input() course: Course;

    constructor(public courseDataService: CourseDataService) {}

    ngOnInit(): void {
        if (this.course.teacher === undefined) {
            this.course.teacher = new PersonName();
        }
        if (this.course.schoolYear === undefined) {
            this.course.schoolYear = new SchoolYear();
        }
    }
}




