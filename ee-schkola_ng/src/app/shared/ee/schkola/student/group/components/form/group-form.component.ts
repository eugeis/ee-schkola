import {Group, GroupCategory, SchoolYear} from '@schkola/student/StudentApiBase'
import {Profile} from '@schkola/person/PersonApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {GroupDataService} from '../../service/group-data.service';

@Component({
  selector: 'app-group-form',
  templateUrl: './group-form.component.html',
  styleUrls: ['./group-form.component.scss']
})

export class GroupFormComponent implements OnInit {

    groupcategoryEnum = this.groupDataService.loadEnumElement(GroupCategory);
    @Input() group: Group;

    constructor(public groupDataService: GroupDataService) {}

    ngOnInit(): void {
        if (this.group.schoolYear === undefined) {
            this.group.schoolYear = new SchoolYear();
        }
        if (this.group.representative === undefined) {
            this.group.representative = new Profile();
        }
    }
}




