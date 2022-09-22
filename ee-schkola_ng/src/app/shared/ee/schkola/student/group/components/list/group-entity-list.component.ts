import {Group} from '@schkola/student/StudentApiBase'

import {Component, Inject, OnInit} from '@angular/core';
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

    constructor(@Inject(GroupDataService) public groupDataService: GroupDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'category', 'schoolyear', 'representative', 'students', 'courses', 'id'];
    }
}




