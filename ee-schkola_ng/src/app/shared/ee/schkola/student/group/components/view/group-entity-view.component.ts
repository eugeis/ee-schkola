import {Group, GroupCategory} from '@schkola/student/StudentApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {GroupDataService} from '../../service/group-data.service';

@Component({
  selector: 'app-group-view',
  templateUrl: './group-entity-view.component.html',
  styleUrls: ['./group-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: GroupDataService}]
})

export class GroupViewComponent implements OnInit {

    groupcategoryEnum = this.groupDataService.loadEnumElement(GroupCategory);
    group: Group;

    constructor(@Inject(GroupDataService) public groupDataService: GroupDataService) {}

    ngOnInit(): void {
        this.group = this.groupDataService.getFirst();
        
        this.groupDataService.checkRoute(this.group);
    }
}




