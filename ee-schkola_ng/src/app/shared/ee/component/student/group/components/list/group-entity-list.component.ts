import {Group} from '../student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {GroupDataService} from '../../services/group-data.service';

@Component({
  selector: 'app-group-list',
  templateUrl: './group-view.component.html',
  styleUrls: ['./group-view.component.scss'],
  providers: [{provide: TableDataService, useClass: GroupDataService}]
})

export class GroupViewComponent implements OnInit {

    group: Group = new Group();

    tableHeader: Array<String> = [];

    constructor(public groupDataService: GroupDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'category', 'schoolyear', 'representative', 'students', 'courses', 'id'];
    }
}




