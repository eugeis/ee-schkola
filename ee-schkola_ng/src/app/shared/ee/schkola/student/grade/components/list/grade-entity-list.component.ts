import {Grade} from '@schkola/student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
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

    constructor(public gradeDataService: GradeDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'gender', 'first', 'last', 'birthName', 'birthday', 'street', 'suite', 'city', 'code', 'country', 'phone', 'email', 'cellphone', 'photoData', 'photo', 'maritalState', 'childrenCount', 'first', 'last', 'church', 'member', 'services', 'name', 'level', 'id', 'other', 'profession', 'id', 'name', 'begin', 'end', 'first', 'last', 'name', 'start', 'end', 'dates', 'id', 'fee', 'description', 'id', 'grade', 'comment', 'id'];
    }
}




