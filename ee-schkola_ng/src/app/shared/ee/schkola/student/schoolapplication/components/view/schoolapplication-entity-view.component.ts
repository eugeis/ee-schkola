import {SchoolApplication} from '@schkola/student/StudentApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {SchoolApplicationDataService} from '../../service/schoolapplication-data.service';

@Component({
  selector: 'app-schoolapplication-view',
  templateUrl: './schoolapplication-entity-view.component.html',
  styleUrls: ['./schoolapplication-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: SchoolApplicationDataService}]
})

export class SchoolApplicationViewComponent implements OnInit {

    schoolapplication: SchoolApplication;

    constructor(public schoolapplicationDataService: SchoolApplicationDataService) {}

    ngOnInit(): void {
        this.schoolapplication = this.schoolapplicationDataService.getFirst();
        this.schoolapplicationDataService.checkRoute(this.schoolapplication);
    }
}




