import {Grade} from '@schkola/student/StudentApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {GradeDataService} from '../../service/grade-data.service';

@Component({
  selector: 'app-grade-list',
  templateUrl: './grade-entity-list.component.html',
  styleUrls: ['./grade-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: GradeDataService}]
})

export class GradeListComponent implements OnInit {

    grade: Grade = new Grade();

    tableHeader: Array<String> = [];

    constructor(@Inject(GradeDataService) public gradeDataService: GradeDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'student', 'course', 'grade', 'comment', 'id'];
    }
}




