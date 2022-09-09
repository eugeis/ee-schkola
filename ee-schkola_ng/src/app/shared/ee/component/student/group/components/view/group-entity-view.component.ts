import {Group, GroupCategory} from '../student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {GroupDataService} from '../../services/group-data.service';

@Component({
  selector: 'app-group-view',
  templateUrl: './group-view.component.html',
  styleUrls: ['./group-view.component.scss'],
  providers: [{provide: TableDataService, useClass: GroupDataService}]
})

export class GroupViewComponent implements OnInit {

    groupcategoryEnum = this.groupDataService.loadEnumElement(GroupCategory);
    group: Group;

    constructor(public groupDataService: GroupDataService) {}

    ngOnInit(): void {
        this.group = this.groupDataService.getFirst();
        this.groupDataService.checkRoute(this.group);
    }
}




