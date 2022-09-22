import {SchoolApplication} from '@schkola/student/StudentApiBase'

import {Component, Inject, OnInit} from '@angular/core';
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

    constructor(@Inject(SchoolApplicationDataService) public schoolapplicationDataService: SchoolApplicationDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'profile', 'churchcontactperson', 'churchcontact', 'churchcommitment', 'schoolyear', 'group', 'id'];
    }
}




