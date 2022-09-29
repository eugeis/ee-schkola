import {SchoolApplication} from '@schkola/student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {SchoolApplicationDataService} from '../../service/schoolapplication-data.service';

@Component({
  selector: 'app-schoolapplication-list',
  templateUrl: './schoolapplication-entity-list.component.html',
  styleUrls: ['./schoolapplication-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: SchoolApplicationDataService}]
})

export class SchoolApplicationListComponent implements OnInit {

    schoolapplication: SchoolApplication = new SchoolApplication();

    tableHeader: Array<String> = [];

    constructor(public schoolapplicationDataService: SchoolApplicationDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'gender', 'first', 'last', 'birthName', 'birthday', 'street', 'suite', 'city', 'code', 'country', 'phone', 'email', 'cellphone', 'photoData', 'photo', 'maritalState', 'childrenCount', 'first', 'last', 'church', 'member', 'services', 'name', 'level', 'id', 'other', 'profession', 'id', 'first', 'last', 'phone', 'email', 'cellphone', 'churchCommitment', 'name', 'start', 'end', 'dates', 'id', 'group', 'id'];
    }
}




