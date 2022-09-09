import {Graduation, GraduationLevel} from '../person/PersonApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {GraduationDataService} from '../../services/graduation-data.service';

@Component({
  selector: 'app-graduation-view',
  templateUrl: './graduation-view.component.html',
  styleUrls: ['./graduation-view.component.scss'],
  providers: [{provide: TableDataService, useClass: GraduationDataService}]
})

export class GraduationViewComponent implements OnInit {

    graduationlevelEnum = this.graduationDataService.loadEnumElement(GraduationLevel);
    graduation: Graduation;

    constructor(public graduationDataService: GraduationDataService) {}

    ngOnInit(): void {
        this.graduation = this.graduationDataService.getFirst();
        this.graduationDataService.checkRoute(this.graduation);
    }
}




