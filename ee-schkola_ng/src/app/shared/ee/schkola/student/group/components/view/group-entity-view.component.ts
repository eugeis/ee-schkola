import {Group} from '@schkola/student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {GroupDataService} from '../../service/group-data.service';

@Component({
  selector: 'app-group-view',
  templateUrl: './group-entity-view.component.html',
  styleUrls: ['./group-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: GroupDataService}]
})

export class GroupViewComponent implements OnInit {

    group: Group;

    constructor(public groupDataService: GroupDataService) {}

    ngOnInit(): void {
        this.group = this.groupDataService.getFirst();
        this.groupDataService.checkRoute(this.group);
    }
}




