import {SchoolYear} from '@schkola/student/StudentApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {SchoolYearDataService} from '../../service/schoolyear-data.service';

@Component({
  selector: 'app-schoolyear-view',
  templateUrl: './schoolyear-entity-view.component.html',
  styleUrls: ['./schoolyear-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: SchoolYearDataService}]
})

export class SchoolYearViewComponent implements OnInit {


    schoolyear: SchoolYear;

    constructor(@Inject(SchoolYearDataService) public schoolyearDataService: SchoolYearDataService) {}

    ngOnInit(): void {
        this.schoolyear = this.schoolyearDataService.getFirst();
        
        this.schoolyearDataService.checkRoute(this.schoolyear);
    }
}




