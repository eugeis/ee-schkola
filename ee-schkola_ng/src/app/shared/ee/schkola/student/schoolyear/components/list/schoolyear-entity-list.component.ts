import {SchoolYear} from '@schkola/student/StudentApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {SchoolYearDataService} from '../../service/schoolyear-data.service';

@Component({
  selector: 'app-schoolyear-list',
  templateUrl: './schoolyear-entity-list.component.html',
  styleUrls: ['./schoolyear-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: SchoolYearDataService}]
})

export class SchoolYearListComponent implements OnInit {

    schoolyear: SchoolYear = new SchoolYear();

    tableHeader: Array<String> = [];

    constructor(@Inject(SchoolYearDataService) public schoolyearDataService: SchoolYearDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'start', 'end', 'dates', 'id'];
    }
}




