import {Course, Grade} from '@schkola/student/StudentApiBase'
import {Profile} from '@schkola/person/PersonApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {GradeDataService} from '../../service/grade-data.service';

@Component({
  selector: 'app-grade-form',
  templateUrl: './grade-form.component.html',
  styleUrls: ['./grade-form.component.scss']
})

export class GradeFormComponent implements OnInit {


    @Input() grade: Grade;

    constructor(public gradeDataService: GradeDataService) {}

    ngOnInit(): void {
        if (this.grade.student === undefined) {
            this.grade.student = new Profile();
        }
        if (this.grade.course === undefined) {
            this.grade.course = new Course();
        }
    }
}




