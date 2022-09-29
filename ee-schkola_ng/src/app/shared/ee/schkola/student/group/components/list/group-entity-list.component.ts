import {Course, Group} from '@schkola/student/StudentApiBase'
import {Profile} from '@schkola/person/PersonApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {GroupDataService} from '../../service/group-data.service';

@Component({
  selector: 'app-group-list',
  templateUrl: './group-entity-list.component.html',
  styleUrls: ['./group-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: GroupDataService}]
})

export class GroupListComponent implements OnInit {

    group: Group = new Group();

    tableHeader: Array<String> = [];

    constructor(public groupDataService: GroupDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'category', 'name', 'start', 'end', 'dates', 'id', 'gender', 'first', 'last', 'birthName', 'birthday', 'street', 'suite', 'city', 'code', 'country', 'phone', 'email', 'cellphone', 'photoData', 'photo', 'maritalState', 'childrenCount', 'first', 'last', 'church', 'member', 'services', 'name', 'level', 'id', 'other', 'profession', 'id', 'students', 'courses', 'id'];
    }
}




