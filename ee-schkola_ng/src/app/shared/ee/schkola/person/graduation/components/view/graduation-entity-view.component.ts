import {Graduation, GraduationLevel} from '@schkola/person/PersonApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {GraduationDataService} from '../../service/graduation-data.service';

@Component({
  selector: 'app-graduation-view',
  templateUrl: './graduation-entity-view.component.html',
  styleUrls: ['./graduation-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: GraduationDataService}]
})

export class GraduationViewComponent implements OnInit {

    graduationlevelEnum = this.graduationDataService.loadEnumElement(GraduationLevel);
    graduation: Graduation;

    constructor(@Inject(GraduationDataService) public graduationDataService: GraduationDataService) {}

    ngOnInit(): void {
        this.graduation = this.graduationDataService.getFirst();
        
        this.graduationDataService.checkRoute(this.graduation);
    }
}




