import {Grade} from '../student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {GradeDataService} from '../../services/grade-data.service';

@Component({
  selector: 'app-grade-list',
  templateUrl: './grade-view.component.html',
  styleUrls: ['./grade-view.component.scss'],
  providers: [{provide: TableDataService, useClass: GradeDataService}]
})

export class GradeViewComponent implements OnInit {

    grade: Grade = new Grade();

    tableHeader: Array<String> = [];

    constructor(public gradeDataService: GradeDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'student', 'course', 'grade', 'comment', 'id'];
    }
}




