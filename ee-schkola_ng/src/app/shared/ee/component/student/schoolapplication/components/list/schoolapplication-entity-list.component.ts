import {SchoolApplication} from '../student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {SchoolApplicationDataService} from '../../services/schoolapplication-data.service';

@Component({
  selector: 'app-schoolapplication-list',
  templateUrl: './schoolapplication-view.component.html',
  styleUrls: ['./schoolapplication-view.component.scss'],
  providers: [{provide: TableDataService, useClass: SchoolApplicationDataService}]
})

export class SchoolApplicationViewComponent implements OnInit {

    schoolapplication: SchoolApplication = new SchoolApplication();

    tableHeader: Array<String> = [];

    constructor(public schoolapplicationDataService: SchoolApplicationDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'profile', 'churchcontactperson', 'churchcontact', 'churchcommitment', 'schoolyear', 'group', 'id'];
    }
}




