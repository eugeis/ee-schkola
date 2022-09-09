import {SchoolYear} from '../student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {SchoolYearDataService} from '../../services/schoolyear-data.service';

@Component({
  selector: 'app-schoolyear-list',
  templateUrl: './schoolyear-view.component.html',
  styleUrls: ['./schoolyear-view.component.scss'],
  providers: [{provide: TableDataService, useClass: SchoolYearDataService}]
})

export class SchoolYearViewComponent implements OnInit {

    schoolyear: SchoolYear = new SchoolYear();

    tableHeader: Array<String> = [];

    constructor(public schoolyearDataService: SchoolYearDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'start', 'end', 'dates', 'id'];
    }
}




