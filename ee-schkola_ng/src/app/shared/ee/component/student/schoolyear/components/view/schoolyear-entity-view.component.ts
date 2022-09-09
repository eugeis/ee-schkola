import {SchoolYear} from '../student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {SchoolYearDataService} from '../../services/schoolyear-data.service';

@Component({
  selector: 'app-schoolyear-view',
  templateUrl: './schoolyear-view.component.html',
  styleUrls: ['./schoolyear-view.component.scss'],
  providers: [{provide: TableDataService, useClass: SchoolYearDataService}]
})

export class SchoolYearViewComponent implements OnInit {


    schoolyear: SchoolYear;

    constructor(public schoolyearDataService: SchoolYearDataService) {}

    ngOnInit(): void {
        this.schoolyear = this.schoolyearDataService.getFirst();
        this.schoolyearDataService.checkRoute(this.schoolyear);
    }
}




